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

package actuators

import (
	"github.com/kubekit99/airship-go-api/armada/services/armada"
	"github.com/kubekit99/airship-go-api/autorest"
	"github.com/kubekit99/airship-go-api/deckhand/services/deckhand"
	"github.com/kubekit99/airship-go-api/drydock/services/drydock"
	// "github.com/kubekit99/airship-go-api/promenade/services/promenade"
	// "github.com/kubekit99/airship-go-api/shipyard/services/shipyard"
	providerv1 "github.com/keleustes/cluster-api-provider-airship/pkg/apis/airshipprovider/v1alpha1"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

// AirshipClients contains all the Airship clients used by the scopes.
type AirshipClients struct {
	Drydock   AirshipDrydockClient
	Armada    AirshipArmadaClient
	Deckhand  AirshipDeckhandClient
	Shipyard  AirshipShipyardClient
	Promenade AirshipPromenadeClient

	// Drydock
	VM    drydock.VirtualMachinesClient
	Disks drydock.DisksClient

	// Network
	VirtualNetworks   armada.VirtualNetworksClient
	SecurityGroups    armada.SecurityGroupsClient
	Interfaces        armada.InterfacesClient
	LB                armada.LoadBalancersClient
	PublicIPAddresses armada.PublicIPAddressesClient

	// Resources
	Groups      deckhand.GroupsClient
	Deployments deckhand.DeploymentsClient
	Tags        deckhand.TagsClient
}

// AirshipDrydockClient defines the operations that will interact with the Airship Drydock API
type AirshipDrydockClient interface {
	// Virtual Machines Operations
	RunCommand(resoureGroup string, name string, cmd string) (drydock.VirtualMachinesRunCommandFuture, error)
	VMIfExists(resourceGroup string, name string) (*drydock.VirtualMachine, error)
	DeleteVM(resourceGroup string, name string) (drydock.VirtualMachinesDeleteFuture, error)
	WaitForVMRunCommandFuture(future drydock.VirtualMachinesRunCommandFuture) error
	WaitForVMDeletionFuture(future drydock.VirtualMachinesDeleteFuture) error

	// Disk Operations
	DeleteManagedDisk(resourceGroup string, name string) (drydock.DisksDeleteFuture, error)
	WaitForDisksDeleteFuture(future drydock.DisksDeleteFuture) error
}

// AirshipArmadaClient defines the operations that will interact with the Airship Network API
type AirshipArmadaClient interface {
	// Network Interfaces Operations
	DeleteNetworkInterface(resourceGroupName string, networkInterfaceName string) (armada.InterfacesDeleteFuture, error)
	WaitForNetworkInterfacesDeleteFuture(future armada.InterfacesDeleteFuture) error

	// Network Security Groups Operations
	CreateOrUpdateNetworkSecurityGroup(resourceGroupName string, networkSecurityGroupName string, location string) (*armada.SecurityGroupsCreateOrUpdateFuture, error)
	NetworkSGIfExists(resourceGroupName string, networkSecurityGroupName string) (*armada.SecurityGroup, error)
	WaitForNetworkSGsCreateOrUpdateFuture(future armada.SecurityGroupsCreateOrUpdateFuture) error

	// Public Ip Address Operations
	GetPublicIPAddress(resourceGroupName string, IPName string) (armada.PublicIPAddress, error)
	DeletePublicIPAddress(resourceGroup string, IPName string) (armada.PublicIPAddressesDeleteFuture, error)
	WaitForPublicIPAddressDeleteFuture(future armada.PublicIPAddressesDeleteFuture) error

	// Virtual Networks Operations
	CreateOrUpdateVnet(resourceGroupName string, virtualNetworkName string, location string) (*armada.VirtualNetworksCreateOrUpdateFuture, error)
	WaitForVnetCreateOrUpdateFuture(future armada.VirtualNetworksCreateOrUpdateFuture) error
}

// AirshipDeckhandClient defines the operations that will interact with the Airship Resources API
type AirshipDeckhandClient interface {
	// Resource Groups Operations
	CreateOrUpdateGroup(resourceGroupName string, location string) (deckhand.Group, error)
	DeleteGroup(resourceGroupName string) (deckhand.GroupsDeleteFuture, error)
	CheckGroupExistence(rgName string) (autorest.Response, error)
	WaitForGroupsDeleteFuture(future deckhand.GroupsDeleteFuture) error

	// Deployment Operations
	CreateOrUpdateDeployment(machine *clusterv1.Machine, clusterConfig *providerv1.AirshipClusterProviderSpec, machineConfig *providerv1.AirshipMachineProviderSpec) (*deckhand.DeploymentsCreateOrUpdateFuture, error)
	GetDeploymentResult(future deckhand.DeploymentsCreateOrUpdateFuture) (de deckhand.DeploymentExtended, err error)
	ValidateDeployment(machine *clusterv1.Machine, clusterConfig *providerv1.AirshipClusterProviderSpec, machineConfig *providerv1.AirshipMachineProviderSpec) error
	WaitForDeploymentsCreateOrUpdateFuture(future deckhand.DeploymentsCreateOrUpdateFuture) error
}

// AirshipShipyardClient defines the operations that will interact with the Airship Resources API
type AirshipShipyardClient interface {
	WaitForSomething(thething string) (string, error)
}

// AirshipPromenadeClient defines the operations that will interact with the Airship Resources API
type AirshipPromenadeClient interface {
	WaitForSomething(thething string) (string, error)
}
