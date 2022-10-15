terraform {
  required_providers {
    circle3 = {
      version = "0.1"
      source  = "bmeik/tf/circle3"
    }
  }
}
provider "circle3" {
  address = "https://wombat.ik.bme.hu"
  port    = 443
  // export CIRCLE3_TOKEN="secret-key"
}
data "circle3_template" "basetemplate" {
  name = "staticnet v1"
}
resource "circle3_vm" "from_template_tf" {
  name = "cloud3 init"
  from_template = data.circle3_template.basetemplate.id
  connection {
    type = "ssh"
    user = "cloud"
    password = self.pw
    host = self.hostipv4
    port = self.sshportipv4
  }
  provisioner "remote-exec" {
    inline = [
      "sudo apt update",
      "sudo apt -y  install python3 python3-dev git python3-pip git",
      "git clone https://git.ik.bme.hu/CIRCLE3/salt.git"
    ]
  }
}
output "vm-ip" {
  value = circle3_vm.from_template_tf.ipv4
}