# Role assignment with dependencies
resource "kinde_user" "example_user" {
  first_name = "John"
  last_name  = "Doe"
  email      = "john.doe@example.com"
}

resource "kinde_role" "example_role" {
  name = "Example Role"
  key  = "example_role"
}

# Basic role assignment referencing the resources above
resource "kinde_user_role" "basic_assignment" {
  user_id           = kinde_user.example_user.id
  role_id           = kinde_role.example_role.id
  organization_code = "org_123" # Replace with your organization code
}

# Supporting resources for multiple assignments
resource "kinde_user" "admin_user" {
  first_name = "Ada"
  last_name  = "Admin"
  email      = "ada.admin@example.com"
}

resource "kinde_role" "admin" {
  name = "Administrator"
  key  = "administrator"
}

resource "kinde_role" "readonly" {
  name = "Read Only"
  key  = "read_only"
}

# Multiple role assignments for a user
resource "kinde_user_role" "admin_assignment" {
  user_id           = kinde_user.admin_user.id
  role_id           = kinde_role.admin.id
  organization_code = "org_123" # Replace with your organization code
}

resource "kinde_user_role" "readonly_assignment" {
  user_id           = kinde_user.admin_user.id
  role_id           = kinde_role.readonly.id
  organization_code = "org_123" # Replace with your organization code
}

resource "kinde_user_role" "example_assignment" {
  user_id           = kinde_user.example_user.id
  role_id           = kinde_role.example_role.id
  organization_code = "org_123" # Replace with your organization code
}
