/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	"fmt"

	"github.com/keleustes/cluster-api-provider-airship/pkg/cloud/airship/actuators"
	"github.com/keleustes/cluster-api-provider-airship/pkg/cloud/airship/services/armada"
	"github.com/keleustes/cluster-api-provider-airship/pkg/cloud/airship/services/deckhand"
	"github.com/keleustes/cluster-api-provider-airship/pkg/deployer"
	"github.com/pkg/errors"
	"k8s.io/klog"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	client "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset/typed/cluster/v1alpha1"
)

// Actuator is responsible for performing cluster reconciliation
type Actuator struct {
	*deployer.Deployer

	client client.ClusterV1alpha1Interface
}

// ActuatorParams holds parameter information for Actuator
type ActuatorParams struct {
	Client client.ClusterV1alpha1Interface
}

// NewActuator creates a new Actuator
func NewActuator(params ActuatorParams) *Actuator {
	return &Actuator{
		Deployer: deployer.New(deployer.Params{ScopeGetter: actuators.DefaultScopeGetter}),
		client:   params.Client,
	}
}

// Reconcile creates or applies updates to the cluster.
func (a *Actuator) Reconcile(cluster *clusterv1.Cluster) error {
	klog.Infof("Reconciling cluster %v", cluster.Name)

	scope, err := actuators.NewScope(actuators.ScopeParams{Cluster: cluster, Client: a.client})
	if err != nil {
		return errors.Errorf("failed to create scope: %+v", err)
	}

	defer scope.Close()

	networkSvc := armada.NewService(scope)
	deckhandSvc := deckhand.NewService(scope)

	// Reconcile resource group
	_, err = deckhandSvc.CreateOrUpdateGroup(scope.ClusterConfig.ResourceGroup, scope.ClusterConfig.Location)
	if err != nil {
		return fmt.Errorf("failed to create or update resource group: %v", err)
	}

	// Reconcile network security group
	networkSGFuture, err := networkSvc.CreateOrUpdateNetworkSecurityGroup(scope.ClusterConfig.ResourceGroup, "ClusterAPINSG", scope.ClusterConfig.Location)
	if err != nil {
		return fmt.Errorf("error creating or updating network security group: %v", err)
	}
	err = networkSvc.WaitForNetworkSGsCreateOrUpdateFuture(*networkSGFuture)
	if err != nil {
		return fmt.Errorf("error waiting for network security group creation or update: %v", err)
	}

	// Reconcile virtual network
	vnetFuture, err := networkSvc.CreateOrUpdateVnet(scope.ClusterConfig.ResourceGroup, "", scope.ClusterConfig.Location)
	if err != nil {
		return fmt.Errorf("error creating or updating virtual network: %v", err)
	}
	err = networkSvc.WaitForVnetCreateOrUpdateFuture(*vnetFuture)
	if err != nil {
		return fmt.Errorf("error waiting for virtual network creation or update: %v", err)
	}
	return nil
}

// Delete deletes a cluster and is invoked by the Cluster Controller.
func (a *Actuator) Delete(cluster *clusterv1.Cluster) error {
	klog.Infof("Reconciling cluster %v", cluster.Name)

	scope, err := actuators.NewScope(actuators.ScopeParams{Cluster: cluster, Client: a.client})
	if err != nil {
		return errors.Errorf("failed to create scope: %+v", err)
	}

	defer scope.Close()

	deckhandSvc := deckhand.NewService(scope)

	resp, err := deckhandSvc.CheckGroupExistence(scope.ClusterConfig.ResourceGroup)
	if err != nil {
		return fmt.Errorf("error checking for resource group existence: %v", err)
	}
	if resp.StatusCode == 404 {
		return fmt.Errorf("resource group %v does not exist", scope.ClusterConfig.ResourceGroup)
	}

	groupsDeleteFuture, err := deckhandSvc.DeleteGroup(scope.ClusterConfig.ResourceGroup)
	if err != nil {
		return fmt.Errorf("error deleting resource group: %v", err)
	}
	err = deckhandSvc.WaitForGroupsDeleteFuture(groupsDeleteFuture)
	if err != nil {
		return fmt.Errorf("error waiting for resource group deletion: %v", err)
	}
	return nil
}
