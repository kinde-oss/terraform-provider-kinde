resource "kinde_role" "readonly_role" {
  name        = "ms_${local.smoke_suffix}_readonly_role"
  key         = "ms_${local.smoke_suffix}_readonly_role"
  description = "Additional role used by the complete user role example"
}

resource "kinde_user_role" "readonly_assignment" {
  organization_code = kinde_organization.example.code
  user_id           = kinde_user.example_user.id
  role_id           = kinde_role.readonly_role.id

  depends_on = [kinde_organization_user.example_membership]
}

output "user_role_assignment_ids" {
  value = [
    kinde_user_role.example_assignment.id,
    kinde_user_role.readonly_assignment.id,
  ]
}
