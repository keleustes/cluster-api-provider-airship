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

package armada

import (
	"github.com/kubekit99/airship-go-api/armada/services/armada"
)

// From Armada API Definition
// func (a *Client) GetHealth(params *GetHealthParams) (*GetHealthNoContent, error)
// func (a *Client) GetReleases(params *GetReleasesParams) (*GetReleasesOK, error)
// func (a *Client) GetStatus(params *GetStatusParams) (*GetStatusOK, error)
// func (a *Client) GetVersions(params *GetVersionsParams) (*GetVersionsOK, error)
// func (a *Client) PostApplyManifest(params *PostApplyManifestParams) (*PostApplyManifestOK, error)
// func (a *Client) PostRollbackReleaseName(params *PostRollbackReleaseNameParams) (*PostRollbackReleaseNameOK, error)
// func (a *Client) PostTestReleaseName(params *PostTestReleaseNameParams) (*PostTestReleaseNameOK, error)
// func (a *Client) PostTests(params *PostTestsParams) (*PostTestsOK, error)
// func (a *Client) PostValidateDesign(params *PostValidateDesignParams) (*PostValidateDesignOK, error)

const (
	// VnetDefaultName is the default name for the cluster's virtual armada.
	VnetDefaultName = "ClusterAPIVnet"
	// SubnetDefaultName is the default name for the cluster's subnet.
	SubnetDefaultName        = "ClusterAPISubnet"
	defaultPrivateSubnetCIDR = "10.0.0.0/24"
	// SecurityGroupDefaultName is the default name for the network security group of the cluster.
	SecurityGroupDefaultName = "ClusterAPINSG"
)

// CreateOrUpdateVnet creates or updates a virtual network resource.
func (s *Service) CreateOrUpdateVnet(resourceGroupName string, virtualNetworkName string, location string) (*armada.VirtualNetworksCreateOrUpdateFuture, error) {
	//JEB	if virtualNetworkName == "" {
	//JEB		virtualNetworkName = VnetDefaultName
	//JEB	}
	//JEB
	//JEB	subnets := []armada.Subnet{
	//JEB		{
	//JEB			Name: to.StringPtr(SubnetDefaultName),
	//JEB			SubnetPropertiesFormat: &armada.SubnetPropertiesFormat{
	//JEB				AddressPrefix: to.StringPtr(defaultPrivateSubnetCIDR),
	//JEB			},
	//JEB		},
	//JEB	}
	//JEB	virtualNetworkProperties := armada.VirtualNetworkPropertiesFormat{
	//JEB		AddressSpace: &armada.AddressSpace{
	//JEB			AddressPrefixes: &[]string{defaultPrivateSubnetCIDR},
	//JEB		},
	//JEB		Subnets: &subnets,
	//JEB	}
	//JEB	virtualNetwork := armada.VirtualNetwork{
	//JEB		Location:                       to.StringPtr(location),
	//JEB		VirtualNetworkPropertiesFormat: &virtualNetworkProperties,
	//JEB	}
	//JEB	sgFuture, err := s.scope.AirshipClients.VirtualNetworks.CreateOrUpdate(s.scope.Context, resourceGroupName, virtualNetworkName, virtualNetwork)
	//JEB	if err != nil {
	//JEB		return nil, err
	//JEB	}
	//JEB	return &sgFuture, nil
	return &armada.VirtualNetworksCreateOrUpdateFuture{}, nil
}

// WaitForVnetCreateOrUpdateFuture returns when the CreateOrUpdateVnet operation completes.
func (s *Service) WaitForVnetCreateOrUpdateFuture(future armada.VirtualNetworksCreateOrUpdateFuture) error {
	//JEB return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.VirtualNetworks.Client)
	return nil
}

// DeleteNetworkInterface deletes the NIC resource.
func (s *Service) DeleteNetworkInterface(resourceGroup string, networkInterfaceName string) (armada.InterfacesDeleteFuture, error) {
	//JEB return s.scope.AirshipClients.Interfaces.Delete(s.scope.Context, resourceGroup, networkInterfaceName)
	return armada.InterfacesDeleteFuture{}, nil
}

// WaitForNetworkInterfacesDeleteFuture waits for the DeleteNetworkInterface operation to complete.
func (s *Service) WaitForNetworkInterfacesDeleteFuture(future armada.InterfacesDeleteFuture) error {
	//JEB return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.Interfaces.Client)
	return nil
}

// GetPublicIPAddress retrieves the Public IP address resource.
func (s *Service) GetPublicIPAddress(resourceGroup string, IPName string) (armada.PublicIPAddress, error) {
	//JEB return s.scope.AirshipClients.PublicIPAddresses.Get(s.scope.Context, resourceGroup, IPName, "")
	return armada.PublicIPAddress{}, nil
}

// DeletePublicIPAddress deletes the Public IP address resource.
func (s *Service) DeletePublicIPAddress(resourceGroup string, IPName string) (armada.PublicIPAddressesDeleteFuture, error) {
	//JEB return s.scope.AirshipClients.PublicIPAddresses.Delete(s.scope.Context, resourceGroup, IPName)
	return armada.PublicIPAddressesDeleteFuture{}, nil
}

// WaitForPublicIPAddressDeleteFuture waits for the DeletePublicIPAddress operation to complete.
func (s *Service) WaitForPublicIPAddressDeleteFuture(future armada.PublicIPAddressesDeleteFuture) error {
	//JEB return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.PublicIPAddresses.Client)
	return nil
}

// NetworkSGIfExists returns the nsg reference if the nsg resource exists.
func (s *Service) NetworkSGIfExists(resourceGroupName string, networkSecurityGroupName string) (*armada.SecurityGroup, error) {
	//JEB networkSG, err := s.scope.AirshipClients.SecurityGroups.Get(s.scope.Context, resourceGroupName, networkSecurityGroupName, "")
	//JEB if err != nil {
	//JEB 		if aerr, ok := err.(autorest.DetailedError); ok {
	//JEB 			if aerr.StatusCode.(int) == 404 {
	//JEB 				return nil, nil
	//JEB 			}
	//JEB 		}
	//JEB 		return nil, err
	//JEB 	}
	//JEB 	return &networkSG, nil
	return &armada.SecurityGroup{}, nil
}

// CreateOrUpdateNetworkSecurityGroup creates or updates the nsg resource.
func (s *Service) CreateOrUpdateNetworkSecurityGroup(resourceGroupName string, networkSecurityGroupName string, location string) (*armada.SecurityGroupsCreateOrUpdateFuture, error) {
	//JEB if networkSecurityGroupName == "" {
	//JEB networkSecurityGroupName = SecurityGroupDefaultName
	//JEB }
	//JEB sshInbound := armada.SecurityRule{
	//JEB Name: to.StringPtr("ClusterAPISSH"),
	//JEB SecurityRulePropertiesFormat: &armada.SecurityRulePropertiesFormat{
	//JEB Protocol:                 armada.SecurityRuleProtocolTCP,
	//JEB SourcePortRange:          to.StringPtr("*"),
	//JEB DestinationPortRange:     to.StringPtr("22"),
	//JEB SourceAddressPrefix:      to.StringPtr("*"),
	//JEB DestinationAddressPrefix: to.StringPtr("*"),
	//JEB Priority:                 to.Int32Ptr(1000),
	//JEB Direction:                armada.SecurityRuleDirectionInbound,
	//JEB Access:                   armada.SecurityRuleAccessAllow,
	//JEB },
	//JEB }
	//JEB
	//JEB kubernetesInbound := armada.SecurityRule{
	//JEB Name: to.StringPtr("KubernetesAPI"),
	//JEB SecurityRulePropertiesFormat: &armada.SecurityRulePropertiesFormat{
	//JEB Protocol:                 armada.SecurityRuleProtocolTCP,
	//JEB SourcePortRange:          to.StringPtr("*"),
	//JEB DestinationPortRange:     to.StringPtr("6443"),
	//JEB SourceAddressPrefix:      to.StringPtr("*"),
	//JEB DestinationAddressPrefix: to.StringPtr("*"),
	//JEB Priority:                 to.Int32Ptr(1001),
	//JEB Direction:                armada.SecurityRuleDirectionInbound,
	//JEB Access:                   armada.SecurityRuleAccessAllow,
	//JEB },
	//JEB }
	//JEB
	//JEB securityGroupProperties := armada.SecurityGroupPropertiesFormat{
	//JEB SecurityRules: &[]armada.SecurityRule{sshInbound, kubernetesInbound},
	//JEB }
	//JEB securityGroup := armada.SecurityGroup{
	//JEB Location:                      to.StringPtr(location),
	//JEB SecurityGroupPropertiesFormat: &securityGroupProperties,
	//JEB }
	//JEB sgFuture, err := s.scope.AirshipClients.SecurityGroups.CreateOrUpdate(s.scope.Context, resourceGroupName, networkSecurityGroupName, securityGroup)
	//JEB if err != nil {
	//JEB return nil, err
	//JEB }
	//JEB return &sgFuture, nil
	return &armada.SecurityGroupsCreateOrUpdateFuture{}, nil
}

// DeleteNetworkSecurityGroup deletes the nsg resource.
func (s *Service) DeleteNetworkSecurityGroup(resourceGroupName string, networkSecurityGroupName string) (armada.SecurityGroupsDeleteFuture, error) {
	//JEB return s.scope.AirshipClients.SecurityGroups.Delete(s.scope.Context, resourceGroupName, networkSecurityGroupName)
	return armada.SecurityGroupsDeleteFuture{}, nil
}

// WaitForNetworkSGsCreateOrUpdateFuture returns when the CreateOrUpdateNetworkSecurityGroup operation completes.
func (s *Service) WaitForNetworkSGsCreateOrUpdateFuture(future armada.SecurityGroupsCreateOrUpdateFuture) error {
	//JEB return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.SecurityGroups.Client)
	return nil
}
