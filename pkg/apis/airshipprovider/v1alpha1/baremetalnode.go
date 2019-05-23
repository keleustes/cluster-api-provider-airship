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

// AddressingItems
type AddressingItems struct {
	Address string `json:"address,omitempty"`
	Network string `json:"network,omitempty"`
}

// Filesystem
type Filesystem struct {
	FsLabel      string `json:"fsLabel,omitempty"`
	FsUuid       string `json:"fsUuid,omitempty"`
	Fstype       string `json:"fstype,omitempty"`
	MountOptions string `json:"mountOptions,omitempty"`
	Mountpoint   string `json:"mountpoint,omitempty"`
}

// InterfacesItem
type InterfacesItem struct {
	DeviceLink string   `json:"deviceLink,omitempty"`
	Networks   []string `json:"networks,omitempty"`
	Slaves     []string `json:"slaves,omitempty"`
}

// KernelParams
type KernelParams struct {
	// AdditionalProperties map[string]interface{} `json:"-,omitempty"`
	AdditionalProperties map[string]string `json:"-,omitempty"`
}

// LogicalVolumesItems
type LogicalVolumesItems struct {
	Filesystem *Filesystem `json:"filesystem,omitempty"`
	LvUuid     string      `json:"lvUuid,omitempty"`
	Name       string      `json:"name,omitempty"`
	Size       string      `json:"size,omitempty"`
}

// Metadata
type Metadata struct {
	BootMac   string            `json:"bootMac,omitempty"`
	OwnerData map[string]string `json:"ownerData,omitempty"`
	Rack      string            `json:"rack,omitempty"`
	Tags      []string          `json:"tags,omitempty"`
}

// Oob
type Oob struct {
	Account string `json:"account,omitempty"`
	// AdditionalProperties map[string]interface{} `json:"-,omitempty"`
	Credetial string `json:"credetial,omitempty"`
	Network   string `json:"network,omitempty"`
	Type      string `json:"type,omitempty"`
}

// PartitionsItems
type PartitionsItems struct {
	Bootable    bool              `json:"bootable,omitempty"`
	Filesystem  *Filesystem       `json:"filesystem,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Name        string            `json:"name,omitempty"`
	PartUuid    string            `json:"partUuid,omitempty"`
	Size        string            `json:"size,omitempty"`
	VolumeGroup string            `json:"volumeGroup,omitempty"`
}

// PhysicalDevicesItem
type PhysicalDevicesItem struct {
	Labels      map[string]string  `json:"labels,omitempty"`
	Partitions  []*PartitionsItems `json:"partitions,omitempty"`
	VolumeGroup string             `json:"volumeGroup,omitempty"`
}

// Platform
type Platform struct {
	Image        string        `json:"image,omitempty"`
	Kernel       string        `json:"kernel,omitempty"`
	KernelParams *KernelParams `json:"kernelParams,omitempty"`
}

// Root
type BaremetalSpec struct {
	Addressing      []*AddressingItems         `json:"addressing,omitempty"`
	HardwareProfile string                     `json:"hardwareProfile,omitempty"`
	HostProfile     string                     `json:"hostProfile,omitempty"`
	Interfaces      map[string]*InterfacesItem `json:"interfaces,omitempty"`
	Metadata        *Metadata                  `json:"metadata,omitempty"`
	Oob             *Oob                       `json:"oob,omitempty"`
	Platform        *Platform                  `json:"platform,omitempty"`
	PrimaryNetwork  string                     `json:"primaryNetwork,omitempty"`
	Storage         *Storage                   `json:"storage,omitempty"`
}

// Storage
type Storage struct {
	PhysicalDevices map[string]*PhysicalDevicesItem `json:"physicalDevices,omitempty"`
	VolumeGroups    map[string]*VolumeGroupsItem    `json:"volumeGroups,omitempty"`
}

// VolumeGroupsItem
type VolumeGroupsItem struct {
	LogicalVolumes []*LogicalVolumesItems `json:"logicalVolumes,omitempty"`
	VgUuid         string                 `json:"vgUuid,omitempty"`
}
