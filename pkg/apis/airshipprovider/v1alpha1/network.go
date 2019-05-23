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

// DhcpRelay
type DhcpRelay struct {
	SelfIp         string `json:"selfIp,omitempty"`
	UpstreamTarget string `json:"upstreamTarget,omitempty"`
}

// Dns
type Dns struct {
	Domain  string `json:"domain,omitempty"`
	Servers string `json:"servers,omitempty"`
}

// Labels
type Labels struct {
	AdditionalProperties map[string]string `json:"-,omitempty"`
}

// RangesItems
type RangesItems struct {
	End   string `json:"end,omitempty"`
	Start string `json:"start,omitempty"`
	Type  string `json:"type,omitempty"`
}

// Root
type NetworkSpec struct {
	Cidr        string         `json:"cidr,omitempty"`
	DhcpRelay   *DhcpRelay     `json:"dhcpRelay,omitempty"`
	Dns         *Dns           `json:"dns,omitempty"`
	Labels      *Labels        `json:"labels,omitempty"`
	Mtu         float64        `json:"mtu,omitempty"`
	Ranges      []*RangesItems `json:"ranges,omitempty"`
	Routedomain string         `json:"routedomain,omitempty"`
	Routes      []*RoutesItems `json:"routes,omitempty"`
	Vlan        string         `json:"vlan,omitempty"`
}

// RoutesItems
type RoutesItems struct {
	Gateway     string  `json:"gateway,omitempty"`
	Metric      float64 `json:"metric,omitempty"`
	Routedomain string  `json:"routedomain,omitempty"`
	Subnet      string  `json:"subnet,omitempty"`
}
