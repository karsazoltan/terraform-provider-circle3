terraform {
  required_providers {
    circle3 = {
      version = "0.1"
      source  = "bmeik/tf/circle3"
    }
  }
}
provider "circle3" {
  address = "http://proxy.fured.cloud.bme.hu"
  port    = 6973
  authtype = "Bearer"
  // export CIRCLE3_TOKEN="secret-key"
}

resource "circle3_lbvm" "demo" {
  name = "loadbalancing"
  from_template = "ubuntu"
  username = "admin"
}

