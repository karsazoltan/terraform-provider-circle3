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
	Disks        []int    `json:"disks,omitempty"`
	HasAgent     bool     `json:"has_agent"`
	ID           int      `json:"id"`
	Ipv4Addr     string   `json:"ipv4addr"`
	Ipv6Addr     string   `json:"ipv6addr"`
	IsBase       bool     `json:"is_base"`
	Lease        int      `json:"lease"`
	MaxRamSize   int      `json:"max_ram_size"`
	Name         string   `json:"name"`
	Node         int      `json:"node,omitempty"`
	NumCores     int      `json:"num_cores"`
	Owner        int      `json:"owner"`
	Priority     int      `json:"priority"`
	Pw           string   `json:"pw"`
	RamSize      int      `json:"ram_size"`
	RawData      string   `json:"raw_data"`
	ReqTraits    []string `json:"req_traits,omitempty"`
	Status       string   `json:"status"`
	System       string   `json:"system"`
	Vlans        []int    `json:"vlans,omitempty"`
}

type VMResource struct {
	Lease      int `json:"lease"`
	MaxRamSize int `json:"max_ram_size"`
	NumCores   int `json:"num_cores"`
	Priority   int `json:"priority"`
	RamSize    int `json:"ram_size"`
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

type Vlan struct {
	Comment     string `json:"comment"`
	Description string `json:"description"`
	Domain      int    `json:"domain"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Vid         int    `json:"vid"`
}

type DDisk struct {
	Instance int    `json:"instance,omitempty"`
	Url      string `json:"url"`
	Name     string `json:"name"`
}

type CDisk struct {
	Instance int    `json:"instance,omitempty"`
	Size     string `json:"size"`
	Name     string `json:"name"`
}

type InstanceActivities struct {
	ActivityCode  string    `json:"activity_code"`
	Created       time.Time `json:"created"`
	Finished      time.Time `json:"finished"`
	GetPercentage int       `json:"get_percentage"`
	ID            int       `json:"id"`
	Instance      int       `json:"instance"`
	Interruptible bool      `json:"interruptible"`
	Modified      time.Time `json:"modified"`
	Parent        int       `json:"parent,omitempty"`
	ResultData    struct {
		AdminTextTemplate string `json:"admin_text_template"`
		Params            struct {
			Checksum   string `json:"checksum,omitempty"`
			DiskID     int    `json:"disk_id,omitempty"`
			DiskSize   int    `json:"disk_size,omitempty"`
			URL        string `json:"url,omitempty"`
			TemplateID int    `json:"template_id,omitempty"`
		} `json:"params,omitempty"`
		UserTextTemplate string `json:"user_text_template"`
	} `json:"result_data,omitempty"`
	ResultantState string    `json:"resultant_state"`
	Started        time.Time `json:"started"`
	Succeeded      *bool     `json:"succeeded"`
	TaskUuid       string    `json:"task_uuid"`
	User           int       `json:"user"`
}

type StorageActivity struct {
	Created    time.Time `json:"created"`
	Disk       int       `json:"disk"`
	Finished   time.Time `json:"finished"`
	ID         int       `json:"id"`
	Modified   time.Time `json:"modified"`
	Parent     int       `json:"parent"`
	ResultData string    `json:"result_data"`
	Started    time.Time `json:"started"`
	Succeeded  *bool     `json:"succeeded"`
	TaskUuid   string    `json:"task_uuid"`
	User       int       `json:"user"`
}

type User struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	Groups      []int  `json:"groups"`
	ID          int    `json:"id"`
	IsStaff     bool   `json:"is_staff"`
	IsSuperuser bool   `json:"is_superuser"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
}

type Group struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	UserSet []int  `json:"user_set"`
}

type TemplateCreating struct {
	Admin      int    `json:"admin,omitempty"`
	Amount     int    `json:"amount,omitempty"`
	MaxRamSize int    `json:"max_ram_size,omitempty"`
	Name       string `json:"name"`
	NumCores   int    `json:"num_cores,omitempty"`
	Operator   int    `json:"operator,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	RamSize    int    `json:"ram_size,omitempty"`
	Template   int    `json:"template"`
	Users      []int  `json:"users,omitempty"`
}

type PortsReq struct {
	Vlan     int `json:"vlan,omitempty"`
	Instance int `json:"instance,omitempty"`
}

type OpenPort struct {
	Forwarding      bool   `json:"forwarding,omitempty"`
	DestinationPort int    `json:"port_destination,omitempty"`
	SourcePort      int    `json:"port_source,omitempty"`
	Type            string `json:"type,omitempty"`
}

type NetRule struct {
	Action          string    `json:"action"`
	CreatedAt       time.Time `json:"created_at"`
	Description     string    `json:"description"`
	Direction       string    `json:"direction"`
	Dport           int       `json:"dport"`
	Extra           string    `json:"extra"`
	Firewall        int       `json:"firewall"`
	ForeignNetwork  int       `json:"foreign_network"`
	Host            int       `json:"host"`
	Hostgroup       int       `json:"hostgroup"`
	ID              int       `json:"id"`
	ModifiedAt      time.Time `json:"modified_at"`
	Nat             bool      `json:"nat"`
	NatExternalIpv4 string    `json:"nat_external_ipv4,omitempty"`
	NatExternalPort int       `json:"nat_external_port"`
	Owner           int       `json:"owner"`
	Proto           string    `json:"proto"`
	Sport           int       `json:"sport,omitempty"`
	Vlan            int       `json:"vlan,omitempty"`
	Vlangroup       int       `json:"vlangroup,omitempty"`
	Weight          int       `json:"weight"`
}

type Metric struct {
	Target     string      `json:"target,omitempty"`
	Datapoints [][]float64 `json:"datapoints,omitempty"`
}
