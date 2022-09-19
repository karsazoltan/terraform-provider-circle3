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

data "circle3_group" "superusers" {
  name = "Superusers"
}

data "circle3_template" "basetemplate" {
  name = "ubuntu v1"
}

resource "circle3_vmpool" "pool_users" {
  name = "vm pool"
  from_template = data.circle3_template.basetemplate.id
  users = data.circle3_group.superusers.users
}