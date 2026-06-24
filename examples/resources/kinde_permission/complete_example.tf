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

output "permission_ids" {
  value = [
    kinde_permission.view_users.id,
    kinde_permission.manage_users.id,
  ]
}
