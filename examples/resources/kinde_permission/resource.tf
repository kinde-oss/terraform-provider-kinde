resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
}

locals {
  smoke_suffix = random_string.suffix.result
}

resource "kinde_permission" "example" {
  name        = "ms_${local.smoke_suffix}_basic_permission"
  key         = "ms_${local.smoke_suffix}_basic_permission"
  description = "Basic permission example for manual Terraform smoke testing"
}
