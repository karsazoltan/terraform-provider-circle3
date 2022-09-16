terraform {
  required_providers {
    circle3 = {
      version = "0.1"
      source  = "hashicorp.com/edu/circle3"
    }
  }
}

provider "circle3" {
  address = "https://cloud3.fured.cloud.bme.hu"
  port    = 443
  token   = "870d52e79fef266daebd1e6f781fe2c2422fde4a"
}

data "circle3_lease_byname" "labor_lease" {
  name = "lab"
}

data "circle3_vlan_byname" "default_vlan" {
  name = "vm"
}

resource "circle3_volume_download" "ubuntu18" {
  name = "ubuntu18.04"
  url = "http://cloud-images.ubuntu.com/bionic/current/bionic-server-cloudimg-amd64.img"
}

resource "circle3_vm" "basic" {
  status        = "RUNNING"
  owner         = 1
  name          = "terraform"
  access_method = "ssh"
  description   = "valami"
  boot_menu     = true
  lease         = data.circle3_lease_byname.labor_lease.id
  cloud_init    = true
  ci_meta_data  = file("${path.module}/meta-data.yaml")
  ci_user_data  = file("${path.module}/user-data.yaml")
  system        = "ubuntu 18.04"
  has_agent     = false
  num_cores     = 2
  ram_size      = 128
  max_ram_size  = 256
  priority      = 30
  arch          = "x86_64"
  disks = [circle3_volume_download.ubuntu18.id]
  vlans = [data.circle3_vlan_byname.default_vlan.vid]
}

