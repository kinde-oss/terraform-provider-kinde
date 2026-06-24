resource "kinde_role" "member_role" {
  name        = "ms_${local.smoke_suffix}_membership_role"
  key         = "ms_${local.smoke_suffix}_membership_role"
  description = "Role used by the organization membership complete example"
}

resource "kinde_user" "role_member" {
  first_name = "Manual"
  last_name  = "RoleMember"
  identities = [
    {
      type  = "email"
      value = "ms.organization.user.roles.${local.smoke_suffix}@example.com"
    }
  ]
}

resource "kinde_organization_user" "with_roles" {
  organization_code = kinde_organization.example.code
  user_id           = kinde_user.role_member.id
  roles             = [kinde_role.member_role.id]
}

output "organization_code" {
  value = kinde_organization.example.code
}

output "membership_ids" {
  value = [
    kinde_organization_user.example.id,
    kinde_organization_user.with_roles.id,
  ]
}
