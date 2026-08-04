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
}
