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

// Filesystem
type Filesystem2 struct {
	FsLabel      string `json:"fsLabel,omitempty"`
	FsUuid       string `json:"fsUuid,omitempty"`
	Fstype       string `json:"fstype,omitempty"`
	MountOptions string `json:"mountOptions,omitempty"`
	Mountpoint   string `json:"mountpoint,omitempty"`
}

// InterfacesItem
type InterfacesItem2 struct {
	DeviceLink string   `json:"deviceLink,omitempty"`
	Networks   []string `json:"networks,omitempty"`
	Slaves     []string `json:"slaves,omitempty"`
	Sriov      *Sriov   `json:"sriov,omitempty"`
}

// KernelParams
type KernelParams2 struct {
	AdditionalProperties map[string]string `json:"-,omitempty"`
}

// LogicalVolumesItems
type LogicalVolumesItems2 struct {
	Filesystem *Filesystem2 `json:"filesystem,omitempty"`
	LvUuid     string       `json:"lvUuid,omitempty"`
	Name       string       `json:"name,omitempty"`
	Size       string       `json:"size,omitempty"`
}

// Metadata
type Metadata2 struct {
	BootMac   string            `json:"bootMac,omitempty"`
	OwnerData map[string]string `json:"ownerData,omitempty"`
	Rack      string            `json:"rack,omitempty"`
	Tags      []string          `json:"tags,omitempty"`
}

// Oob
type Oob2 struct {
	Account string `json:"account,omitempty"`
	// AdditionalProperties map[string]interface{} `json:"-,omitempty"`
	Credetial string `json:"credetial,omitempty"`
	Network   string `json:"network,omitempty"`
	Type      string `json:"type,omitempty"`
}

// PartitionsItems
type PartitionsItems2 struct {
	Bootable    bool              `json:"bootable,omitempty"`
	Filesystem  *Filesystem2      `json:"filesystem,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Name        string            `json:"name,omitempty"`
	PartUuid    string            `json:"partUuid,omitempty"`
	Size        string            `json:"size,omitempty"`
	VolumeGroup string            `json:"volumeGroup,omitempty"`
}

// PhysicalDevicesItem
type PhysicalDevicesItem2 struct {
	Labels      map[string]string   `json:"labels,omitempty"`
	Partitions  []*PartitionsItems2 `json:"partitions,omitempty"`
	VolumeGroup string              `json:"volumeGroup,omitempty"`
}

// Platform
type Platform2 struct {
	Image        string         `json:"image,omitempty"`
	Kernel       string         `json:"kernel,omitempty"`
	KernelParams *KernelParams2 `json:"kernelParams,omitempty"`
}

// Root
type HostProfileSpec struct {
	HardwareProfile string                      `json:"hardwareProfile,omitempty"`
	HostProfile     string                      `json:"hostProfile,omitempty"`
	Interfaces      map[string]*InterfacesItem2 `json:"interfaces,omitempty"`
	Metadata        *Metadata2                  `json:"metadata,omitempty"`
	Oob             *Oob2                       `json:"oob,omitempty"`
	Platform        *Platform2                  `json:"platform,omitempty"`
	PrimaryNetwork  string                      `json:"primaryNetwork,omitempty"`
	Storage         *Storage2                   `json:"storage,omitempty"`
}

// Sriov
type Sriov struct {
	Trustmode bool    `json:"trustmode,omitempty"`
	VfCount   float64 `json:"vfCount,omitempty"`
}

// Storage
type Storage2 struct {
	PhysicalDevices map[string]*PhysicalDevicesItem2 `json:"physicalDevices,omitempty"`
	VolumeGroups    map[string]*VolumeGroupsItem2    `json:"volumeGroups,omitempty"`
}

// VolumeGroupsItem
type VolumeGroupsItem2 struct {
	LogicalVolumes []*LogicalVolumesItems2 `json:"logicalVolumes,omitempty"`
	VgUuid         string                  `json:"vgUuid,omitempty"`
}
