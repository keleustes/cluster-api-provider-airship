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

// DeviceAliases
type DeviceAliases struct {
	AdditionalProperties map[string]string `json:"-,omitempty"`
}

// HugepagesItem
type HugepagesItem struct {
}

// Root
type HardwareProfileSpec struct {
	BiosVersion       string                    `json:"biosVersion,omitempty"`
	BootMode          string                    `json:"bootMode,omitempty"`
	BootstrapProtocol string                    `json:"bootstrapProtocol,omitempty"`
	CpuSets           map[string]string         `json:"cpuSets,omitempty"`
	DeviceAliases     *DeviceAliases            `json:"deviceAliases,omitempty"`
	Generation        string                    `json:"generation,omitempty"`
	Hugepages         map[string]*HugepagesItem `json:"hugepages,omitempty"`
	HwVersion         string                    `json:"hwVersion,omitempty"`
	PxeInterface      float64                   `json:"pxeInterface,omitempty"`
	Vendor            string                    `json:"vendor,omitempty"`
}
