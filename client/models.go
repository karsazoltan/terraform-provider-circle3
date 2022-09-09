package client

import (
	"time"
)

type Lease struct {
	ID                       int    `json:"id"`
	Name                     string `json:"name"`
	Suspend_interval_seconds int    `json:"suspend_interval_seconds"`
	Delete_interval_seconds  int    `json:"delete_interval_seconds"`
}

type VM struct {
	AccessMethod string   `json:"access_method"`
	Arch         string   `json:"arch"`
	BootMenu     bool     `json:"boot_menu"`
	CiMetaData   string   `json:"ci_meta_data"`
	CiUserData   string   `json:"ci_user_data"`
	CloudInit    bool     `json:"cloud_init"`
	Description  string   `json:"description"`
	Disks        []int    `json:"disks"`
	HasAgent     bool     `json:"has_agent"`
	ID           int      `json:"id"`
	Ipv4Addr     string   `json:"ipv4addr"`
	Ipv6Addr     string   `json:"ipv6addr"`
	IsBase       bool     `json:"is_base"`
	Lease        int      `json:"lease"`
	MaxRamSize   int      `json:"max_ram_size"`
	Name         string   `json:"name"`
	Node         string   `json:"node"`
	NumCores     int      `json:"num_cores"`
	Owner        int      `json:"owner"`
	Priority     int      `json:"priority"`
	Pw           string   `json:"pw"`
	RamSize      int      `json:"ram_size"`
	RawData      string   `json:"raw_data"`
	ReqTraits    []string `json:"req_traits"`
	Status       string   `json:"status"`
	System       string   `json:"system"`
}

type Template struct {
	AccessMethod string `json:"access_method"`
	Arch         string `json:"arch"`
	BootMenu     bool   `json:"boot_menu"`
	CiMetaData   string `json:"ci_meta_data"`
	CiUserData   string `json:"ci_user_data"`
	CloudInit    bool   `json:"cloud_init"`
	Description  string `json:"description"`
	Disks        []int  `json:"disks"`
	HasAgent     bool   `json:"has_agent"`
	ID           int    `json:"id"`
	Lease        int    `json:"lease"`
	MaxRamSize   int    `json:"max_ram_size"`
	Name         string `json:"name"`
	NumCores     int    `json:"num_cores"`
	Owner        int    `json:"owner"`
	Parent       int    `json:"parent"`
	Priority     int    `json:"priority"`
	RamSize      int    `json:"ram_size"`
	RawData      string `json:"raw_data"`
	System       string `json:"system"`
}

type Disk struct {
	Base      int       `json:"base"`
	Bus       string    `json:"bus"`
	CiDisk    bool      `json:"ci_disk"`
	Datastore int       `json:"datastore"`
	Destroyed time.Time `json:"destroyed"`
	DevNum    string    `json:"dev_num"`
	Filename  string    `json:"filename"`
	ID        int       `json:"id"`
	IsReady   bool      `json:"is_ready"`
	Name      string    `json:"name"`
	Size      int       `json:"size"`
	Type      string    `json:"type"`
}

type Interface struct {
	Host     int    `json:"host"`
	ID       int    `json:"id"`
	Instance int    `json:"instance"`
	Model    string `json:"model"`
	Vlan     int    `json:"vlan"`
}
