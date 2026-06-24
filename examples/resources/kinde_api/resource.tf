resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
}

locals {
  smoke_suffix = random_string.suffix.result
}

resource "kinde_api" "example" {
  name     = "ms_${local.smoke_suffix}_api"
  audience = "https://ms-${local.smoke_suffix}.example.kinde.api"
}
