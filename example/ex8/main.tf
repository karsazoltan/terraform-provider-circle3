terraform {
  required_providers {
    circle3 = {
      version = "0.1"
      source  = "bmeik/tf/circle3"
    }
  }
}

provider "circle3" {
  address = "https://pumi.niif.cloud.bme.hu"
  port    = 443
  //token   = "secret" -> export CIRCLE3_TOKEN="..."
}

data "circle3_lease" "labor_lease" {
  name = "lab"
}

data "circle3_vlan" "default_vlan" {
  name = "vm"
}

locals {
  virtual_machines = {
   "vm1" = {  },
   "vm2" = {  },
  }
}

resource "circle3_disk" "ubuntu18" {
  for_each = local.virtual_machines
  name = "ubuntu18.04-${each.key}"
  url = "http://cloud-images.ubuntu.com/bionic/current/bionic-server-cloudimg-amd64.img"
}

resource "circle3_vm" "basic" {
  for_each = local.virtual_machines
  status        = "STOPPED"
  owner         = 1
  name          = "terraform-${each.key}"
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
  disks = [circle3_disk.ubuntu18[each.key].id]
  vlans = [data.circle3_vlan.default_vlan.vid]
}

output "vm-ipv4" {
  value = [for k, v in local.virtual_machines : "${circle3_vm.basic[k].ipv4}"]
}