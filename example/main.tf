terraform {
  required_providers {
    circle3 = {
      version = "0.1"
      source = "hashicorp.com/edu/circle3"
    }
  }
}

provider circle3 {
  address = "cloud3.fured.cloud.bme.hu"
  port = 443
  token = "secret"
}

data "circle3_leases" "all" {}

output "all_leases" {
  value = data.circle3_leases.all.leases
}