terraform {
  required_providers {
    circle3 = {
      version = "0.1"
      source  = "bmeik/tf/circle3"
    }
  }
}
provider "circle3" {
  address = "https://cloud3.fured.cloud.bme.hu"
  port    = 443
  // export CIRCLE3_TOKEN="secret-key"
}

data "circle3_lease" "labor_lease" {
  name = "lab"
}

resource "circle3_disk" "empty_disk" {
  size_format = "3Gi"
  name = "empty_disk"
}

resource "circle3_disk" "tiny_linux" {
  name = "tinycore-linux"
  url = "http://tinycorelinux.net/13.x/x86/release/TinyCore-current.iso"
}

resource "circle3_vm" "basic" {
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
  ram_size      = 128
  max_ram_size  = 256
  priority      = 30
  arch          = "x86_64"
  disks         = [ 
    circle3_disk.empty_disk.id, 
    circle3_disk.tiny_linux.id 
  ]
}

//output "vm_create" {
//  value = circle3_vm.basic
//}

//data "circle3_leases" "all" {}

//output "all_leases" {
//  value = data.circle3_leases.all.leases
//}
