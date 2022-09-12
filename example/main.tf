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
output "labor_output" {
  value = data.circle3_lease_byname.labor_lease
}
resource "circle3_vm" "basic" {
  owner         = 1
  name          = "terraform"
  access_method = "ssh"
  description   = "valami"
  boot_menu     = true
  lease         = data.circle3_lease_byname.labor_lease.id
  cloud_init    = true
  ci_meta_data  = "valami"
  ci_user_data  = "msk valami"
  system        = "ubuntu 18.04"
  has_agent     = false
  num_cores     = 2
  ram_size      = 128
  max_ram_size  = 256
  priority      = 30
  arch          = "x86_64"
}
output "vm_create" {
  value = circle3_vm.basic
}

//data "circle3_leases" "all" {}

//output "all_leases" {
//  value = data.circle3_leases.all.leases
//}
