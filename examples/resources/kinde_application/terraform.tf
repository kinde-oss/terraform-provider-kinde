terraform {
  required_providers {
    kinde = {
      source = "kinde-oss/kinde"
    }
    random = {
      source = "hashicorp/random"
    }
  }
}

provider "kinde" {
}
