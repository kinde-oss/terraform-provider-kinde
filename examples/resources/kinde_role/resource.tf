resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
}

locals {
  smoke_suffix = random_string.suffix.result
}

resource "kinde_role" "example" {
  name        = "ms_${local.smoke_suffix}_basic_role"
  key         = "ms_${local.smoke_suffix}_basic_role"
  description = "Basic role example for manual Terraform smoke testing"
}
