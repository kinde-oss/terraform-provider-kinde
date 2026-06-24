resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
}

locals {
  smoke_suffix = random_string.suffix.result
}

resource "kinde_organization" "example" {
  name = "ms_${local.smoke_suffix}_user_role_org"
}

resource "kinde_user" "example_user" {
  first_name = "Manual"
  last_name  = "Assignment"
  identities = [
    {
      type  = "email"
      value = "ms.user.role.${local.smoke_suffix}@example.com"
    }
  ]
}

resource "kinde_organization_user" "example_membership" {
  organization_code = kinde_organization.example.code
  user_id           = kinde_user.example_user.id

  lifecycle {
    ignore_changes = [roles]
  }
}

resource "kinde_role" "example_role" {
  name        = "ms_${local.smoke_suffix}_basic_role"
  key         = "ms_${local.smoke_suffix}_basic_role"
  description = "Role used by the basic user role example"
}

resource "kinde_user_role" "example_assignment" {
  organization_code = kinde_organization.example.code
  user_id           = kinde_user.example_user.id
  role_id           = kinde_role.example_role.id

  depends_on = [kinde_organization_user.example_membership]
}
