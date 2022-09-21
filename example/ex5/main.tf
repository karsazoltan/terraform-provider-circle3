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

data "circle3_template" "basetemplate" {
  name = "ubuntu v1"
}

resource "circle3_vm" "from_template_tf" {
  name = "from template"
  from_template = data.circle3_template.basetemplate.id
}

resource "circle3_port" "openport8080" {
  port = 8080
  vlan = circle3_vm.from_template_tf.vlans[0]
  vm = circle3_vm.from_template_tf.id
  type = "tcp"
}