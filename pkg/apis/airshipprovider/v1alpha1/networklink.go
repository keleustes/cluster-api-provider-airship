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

// Bonding
type Bonding struct {
	DownDelay float64 `json:"downDelay,omitempty"`
	Hash      string  `json:"hash,omitempty"`
	Mode      string  `json:"mode,omitempty"`
	MonRate   float64 `json:"monRate,omitempty"`
	PeerRate  string  `json:"peerRate,omitempty"`
	UpDelay   float64 `json:"upDelay,omitempty"`
}

// Labels
type Labels2 struct {
	AdditionalProperties map[string]string `json:"-,omitempty"`
}

// Root
type NetworkLinkSpec struct {
	AllowedNetworks []string  `json:"allowedNetworks,omitempty"`
	Bonding         *Bonding  `json:"bonding,omitempty"`
	Labels          *Labels2  `json:"labels,omitempty"`
	Linkspeed       string    `json:"linkspeed,omitempty"`
	Mtu             float64   `json:"mtu,omitempty"`
	Trunking        *Trunking `json:"trunking,omitempty"`
}

// Trunking
type Trunking struct {
	DefaultNetwork string `json:"defaultNetwork,omitempty"`
	Mode           string `json:"mode,omitempty"`
}
