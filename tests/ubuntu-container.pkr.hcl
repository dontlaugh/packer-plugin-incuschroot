/*
# example config that would download the plugin. Commented out here because
# we are only using locally-built binaries with the name `packer-plugin-incus`
packer {
  required_plugins {
     incus = {
       version = ">= 0.0.1"
       source  = "github.com/dontlaugh/incus"
     }
  }
}
*/
source "incus" "incus-ubuntu" {
  image = "images:ubuntu/22.04"
  output_image = "repacked-ubuntu-${formatdate("YYYYMMDDhhmmss", timestamp())}"
}

build {
  sources = ["incus.incus-ubuntu"]
}
