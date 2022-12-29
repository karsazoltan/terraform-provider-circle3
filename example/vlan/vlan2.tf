variable "users" {
  type = list(string)
  default = [ "Alice", "Bob" ]
}
data "circle3_template" "meres" {
  name = "meres-sablon"
}
data "circle3_user" "user" {
  for_each = toset(var.users)
  name = "${each.key}"
}
resource "circle3_vm" "for_each_users" {
  for_each = toset(var.users)
  name = "vm-${each.key}"
  from_template = data.circle3_template.meres.id
  owner = data.circle3_user.user[each.key].id   
}



