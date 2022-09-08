package client

type Lease struct {
	Name                     string `json:"name"`
	Suspend_interval_seconds int    `json:"suspend_interval_seconds"`
	Delete_interval_seconds  int    `json:"delete_interval_seconds"`
}
