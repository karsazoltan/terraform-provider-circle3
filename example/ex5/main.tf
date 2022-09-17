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
  // export CIRCLE3_TOKEN="secret-key"
}

data "circle3_template" "basetemplate" {
  
}

resource "circle3_vm" "from template terraform" {
  name = "from template"
  from_template = data.circle3_template.basetemplate.id
}