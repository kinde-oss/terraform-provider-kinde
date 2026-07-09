resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
}

locals {
  smoke_suffix = random_string.suffix.result
}

resource "kinde_permission" "view_users" {
  name        = "ms_${local.smoke_suffix}_view_users"
  key         = "ms_${local.smoke_suffix}_view_users"
  description = "Manual smoke-test permission for viewing users"
}

resource "kinde_permission" "manage_users" {
  name        = "ms_${local.smoke_suffix}_manage_users"
  key         = "ms_${local.smoke_suffix}_manage_users"
  description = "Manual smoke-test permission for managing users"
}

resource "kinde_role" "full_example" {
  name        = "ms_${local.smoke_suffix}_full_role"
  key         = "ms_${local.smoke_suffix}_full_role"
  description = "Complete role example for manual Terraform smoke testing"
  permissions = [
    kinde_permission.view_users.id,
    kinde_permission.manage_users.id,
  ]
}

output "full_example_role_id" {
  value = kinde_role.full_example.id
}

output "permission_ids" {
  value = [
    kinde_permission.view_users.id,
    kinde_permission.manage_users.id,
  ]
}
