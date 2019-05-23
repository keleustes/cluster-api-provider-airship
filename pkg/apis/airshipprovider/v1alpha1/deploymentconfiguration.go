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

package v1alpha1

// Armada
type Armada struct {
	GetReleasesTimeout    int    `json:"getReleasesTimeout,omitempty"`
	GetStatusTimeout      int    `json:"getStatusTimeout,omitempty"`
	Manifest              string `json:"manifest"`
	PostApplyTimeout      int    `json:"postApplyTimeout,omitempty"`
	ValidateDesignTimeout int    `json:"validateDesignTimeout,omitempty"`
}

// Kubernetes
type Kubernetes struct {
	NodeStatusInterval int `json:"nodeStatusInterval,omitempty"`
	NodeStatusTimeout  int `json:"nodeStatusTimeout,omitempty"`
}

// KubernetesProvisioner
type KubernetesProvisioner struct {
	ClearLabelsTimeout int `json:"clearLabelsTimeout,omitempty"`
	DrainGracePeriod   int `json:"drainGracePeriod,omitempty"`
	DrainTimeout       int `json:"drainTimeout,omitempty"`
	EtcdReadyTimeout   int `json:"etcdReadyTimeout,omitempty"`
	RemoveEtcdTimeout  int `json:"removeEtcdTimeout,omitempty"`
}

// PhysicalProvisioner
type PhysicalProvisioner struct {
	DeployInterval       int    `json:"deployInterval,omitempty"`
	DeployTimeout        int    `json:"deployTimeout,omitempty"`
	DeploymentStrategy   string `json:"deploymentStrategy,omitempty"`
	DestroyInterval      int    `json:"destroyInterval,omitempty"`
	DestroyTimeout       int    `json:"destroyTimeout,omitempty"`
	JoinWait             int    `json:"joinWait,omitempty"`
	PrepareNodeInterval  int    `json:"prepareNodeInterval,omitempty"`
	PrepareNodeTimeout   int    `json:"prepareNodeTimeout,omitempty"`
	PrepareSiteInterval  int    `json:"prepareSiteInterval,omitempty"`
	PrepareSiteTimeout   int    `json:"prepareSiteTimeout,omitempty"`
	RelabelNodesInterval int    `json:"relabelNodesInterval,omitempty"`
	RelabelNodesTimeout  int    `json:"relabelNodesTimeout,omitempty"`
	VerifyInterval       int    `json:"verifyInterval,omitempty"`
	VerifyTimeout        int    `json:"verifyTimeout,omitempty"`
}

// Root
type DeploymentConfigurationSpec struct {
	Armada                *Armada                `json:"armada"`
	Kubernetes            *Kubernetes            `json:"kubernetes,omitempty"`
	KubernetesProvisioner *KubernetesProvisioner `json:"kubernetesProvisioner,omitempty"`
	PhysicalProvisioner   *PhysicalProvisioner   `json:"physicalProvisioner,omitempty"`
}
