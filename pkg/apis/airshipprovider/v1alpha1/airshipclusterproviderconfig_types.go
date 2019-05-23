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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AirshipClusterProviderSpec is the Schema for the airshipclusterproviderspecs API
// +k8s:openapi-gen=true
type AirshipClusterProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ResourceGroup string `json:"resourceGroup"`
	Location      string `json:"location"`

	// Shiyard DeploymentConfiguration
	DeploymentConfiguration DeploymentConfigurationSpec `json:"deploymentConfiguration"`
	DeploymentStrategy      DeploymentStrategySpec      `json:"deploymentStrategy"`

	// Drydock Configuration
	// JEB: This is common configuration for the hardware. Force to add it to the cluster
	// for right now.
	HardwareProfiles []HardwareProfileSpec `json:"hardwareProfiles"`

	// JEB: This is network configuration. The generic cluster definition account s for
	// CNI/calico setup up, but does describe the underlying network.
	Networks []NetworkSpec `json:"networks"`

	// CACertificate is a PEM encoded CA Certificate for the control plane nodes.
	CACertificate []byte `json:"caCertificate,omitempty"`

	// CAPrivateKey is a PEM encoded PKCS1 CA PrivateKey for the control plane nodes.
	CAPrivateKey []byte `json:"caKey,omitempty"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&AirshipClusterProviderSpec{})
}
