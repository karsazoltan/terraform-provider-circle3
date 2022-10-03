terraform {
  required_providers {
    circle3 = {
      version = "0.1"
      source  = "bmeik/tf/circle3"
    }
  }
}

provider "circle3" {
  address = "https://axolotl.niif.cloud.bme.hu"
  port    = 443
  // export CIRCLE3_TOKEN="secret-key"
}

data "circle3_lease" "labor_lease" {
  name = "lab"
}

data "circle3_vlan" "default_vlan" {
  name = "vm"
}

resource "circle3_disk" "ubuntu18" {
  name = "ubuntu18.04"
  url = "http://cloud-images.ubuntu.com/bionic/current/bionic-server-cloudimg-amd64.img"
}

resource "circle3_vm" "basic" {
  status        = "STOPPED"
  owner         = 1
  name          = "terraform"
  access_method = "ssh"
  description   = "valami"
  boot_menu     = true
  lease         = data.circle3_lease.labor_lease.id
  cloud_init    = true
  ci_meta_data  = file("${path.module}/meta-data.yaml")
  ci_user_data  = file("${path.module}/user-data.yaml")
  system        = "ubuntu 18.04"
  has_agent     = false
  num_cores     = 2
  ram_size      = 256
  max_ram_size  = 256
  priority      = 80
  arch          = "x86_64"
  disks = [circle3_disk.ubuntu18.id]
  vlans = [data.circle3_vlan.default_vlan.vid]
}