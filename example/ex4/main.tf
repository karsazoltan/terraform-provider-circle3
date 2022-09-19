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

output "Superusers" {
  value = data.circle3_group.superusers
}

data "circle3_user" "karsa" {
  username = "karsa"
}

output "karsa_user" {
  value = data.circle3_user.karsa
}