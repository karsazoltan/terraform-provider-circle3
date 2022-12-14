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

resource "circle3_vm" "basic" {
  owner         = 1
  name          = "terraform"
  access_method = "ssh"
  description   = "valami"
  boot_menu     = true
  lease         = data.circle3_lease.labor_lease.id
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

