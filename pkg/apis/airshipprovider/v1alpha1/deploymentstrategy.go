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

// GroupsItems
type GroupsItems struct {
	Critical        bool              `json:"critical"`
	DependsOn       []string          `json:"dependsOn"`
	Name            string            `json:"name"`
	Selectors       []*SelectorsItems `json:"selectors"`
	SuccessCriteria *SuccessCriteria  `json:"successCriteria,omitempty"`
}

// Root
type DeploymentStrategySpec struct {
	Groups []*GroupsItems `json:"groups"`
}

// SelectorsItems
type SelectorsItems struct {
	NodeLabels []string `json:"nodeLabels,omitempty"`
	NodeNames  []string `json:"nodeNames,omitempty"`
	NodeTags   []string `json:"nodeTags,omitempty"`
	RackNames  []string `json:"rackNames,omitempty"`
}

// SuccessCriteria
type SuccessCriteria struct {
	MaximumFailedNodes     int `json:"maximumFailedNodes,omitempty"`
	MinimumSuccessfulNodes int `json:"minimumSuccessfulNodes,omitempty"`
	PercentSuccessfulNodes int `json:"percentSuccessfulNodes,omitempty"`
}
