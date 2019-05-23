/*
Copyright 2019 The Kubernetes Authors.

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

package deployer

import (
	"fmt"

	providerv1 "github.com/keleustes/cluster-api-provider-airship/pkg/apis/airshipprovider/v1alpha1"
	"github.com/keleustes/cluster-api-provider-airship/pkg/cloud/airship/actuators"
	"github.com/keleustes/cluster-api-provider-airship/pkg/cloud/airship/services/armada"
	"github.com/keleustes/cluster-api-provider-airship/pkg/cloud/airship/services/certificates"
	"github.com/keleustes/cluster-api-provider-airship/pkg/cloud/airship/services/deckhand"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/clientcmd"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

// Deployer satisfies the ProviderDeployer(https://github.com/kubernetes-sigs/cluster-api/blob/master/cmd/clusterctl/clusterdeployer/clusterdeployer.go) interface.
type Deployer struct {
	scopeGetter actuators.ScopeGetter
}

// Params is used to create a new deployer.
type Params struct {
	ScopeGetter actuators.ScopeGetter
}

// New returns a new Deployer.
func New(params Params) *Deployer {
	return &Deployer{
		scopeGetter: params.ScopeGetter,
	}
}

// GetIP returns the IP of a machine, but this is going away.
func (d *Deployer) GetIP(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (string, error) {
	scope, err := d.scopeGetter.GetScope(actuators.ScopeParams{Cluster: cluster})
	if err != nil {
		return "", err
	}

	// TODO: Reimplement using load balancer
	//JEB if scope.ClusterStatus != nil && scope.ClusterStatus.Network.APIServerIP.IPAddress != "" {
	//JEB		return scope.ClusterStatus.Network.APIServerIP.IPAddress, nil
	//JEB }

	pipsvc := armada.NewService(scope)

	pip, err := pipsvc.GetPublicIPAddress(scope.ClusterConfig.ResourceGroup, deckhand.GetPublicIPName(machine))
	if err != nil {
		return "", err
	}
	//JEB return *pip.IPAddress, nil
	_ = pip
	return "", nil
}

// GetKubeConfig returns the kubeconfig after the bootstrap process is complete.
func (d *Deployer) GetKubeConfig(cluster *clusterv1.Cluster, _ *clusterv1.Machine) (string, error) {

	// Load provider config.
	config, err := providerv1.ClusterConfigFromProviderSpec(cluster.Spec.ProviderSpec)
	if err != nil {
		return "", errors.Errorf("failed to load cluster provider status: %v", err)
	}

	cert, err := certificates.DecodeCertPEM(config.CACertificate)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode CA Cert")
	} else if cert == nil {
		return "", errors.New("certificate not found in config")
	}

	key, err := certificates.DecodePrivateKeyPEM(config.CAPrivateKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode private key")
	} else if key == nil {
		return "", errors.New("key not found in status")
	}

	dnsName, err := d.GetIP(cluster, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to get DNS address")
	}

	server := fmt.Sprintf("https://%s:6443", dnsName)

	cfg, err := certificates.NewKubeconfig(cluster.Name, server, cert, key)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate a kubeconfig")
	}

	yaml, err := clientcmd.Write(*cfg)
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize config to yaml")
	}

	return string(yaml), nil
}
