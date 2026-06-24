resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
}

locals {
  smoke_suffix = random_string.suffix.result
}

resource "kinde_organization" "example" {
  name = "ms_${local.smoke_suffix}_organization_user_org"
}

resource "kinde_user" "example" {
  first_name = "Manual"
  last_name  = "Membership"
  identities = [
    {
      type  = "email"
      value = "ms.organization.user.basic.${local.smoke_suffix}@example.com"
    }
  ]
}

resource "kinde_organization_user" "example" {
  organization_code = kinde_organization.example.code
  user_id           = kinde_user.example.id
}
