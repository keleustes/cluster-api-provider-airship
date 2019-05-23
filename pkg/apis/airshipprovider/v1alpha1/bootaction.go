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

// AssetsItems
type AssetsItems struct {
	Data             string   `json:"data,omitempty"`
	DataPipeline     []string `json:"dataPipeline,omitempty"`
	Location         string   `json:"location,omitempty"`
	LocationPipeline []string `json:"locationPipeline,omitempty"`
	Path             string   `json:"path,omitempty"`
	Permissions      string   `json:"permissions,omitempty"`
	Type             string   `json:"type"`
}

// FilterSetItems
type FilterSetItems struct {
	FilterType string      `json:"filterType,omitempty"`
	NodeLabels *NodeLabels `json:"nodeLabels,omitempty"`
	NodeNames  []string    `json:"nodeNames,omitempty"`
	NodeTags   []string    `json:"nodeTags,omitempty"`
	RackLabels *RackLabels `json:"rackLabels,omitempty"`
	RackNames  []string    `json:"rackNames,omitempty"`
}

// NodeFilter
type NodeFilter struct {
	FilterSet     []*FilterSetItems `json:"filterSet,omitempty"`
	FilterSetType string            `json:"filterSetType,omitempty"`
}

// NodeLabels
type NodeLabels struct {
	AdditionalProperties map[string]string `json:"-,omitempty"`
}

// RackLabels
type RackLabels struct {
	AdditionalProperties map[string]string `json:"-,omitempty"`
}

// Root
type BootactionSpec struct {
	Assets     []*AssetsItems `json:"assets,omitempty"`
	NodeFilter *NodeFilter    `json:"nodeFilter,omitempty"`
	Signaling  bool           `json:"signaling,omitempty"`
}
