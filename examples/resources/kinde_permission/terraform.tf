terraform {
  required_providers {
    kinde = {
      source = "nxt-fwd/kinde"
    }
    random = {
      source = "hashicorp/random"
    }
  }
}

provider "kinde" {
  # Configuration options can be provided here or via environment variables:
  # KINDE_DOMAIN
  # KINDE_AUDIENCE
  # KINDE_CLIENT_ID
  # KINDE_CLIENT_SECRET
}
