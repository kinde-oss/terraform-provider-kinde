resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
}

locals {
  smoke_suffix = random_string.suffix.result
}

resource "kinde_application" "example" {
  name          = "ms_${local.smoke_suffix}_application"
  type          = "reg"
  login_uri     = "https://${local.smoke_suffix}.example.com/oauth/login"
  homepage_uri  = "https://${local.smoke_suffix}.example.com"
  logout_uris   = ["https://${local.smoke_suffix}.example.com/oauth/logout"]
  redirect_uris = ["https://${local.smoke_suffix}.example.com/oauth/callback"]
}
