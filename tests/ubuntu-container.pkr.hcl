
source "incus" "incus-ubuntu" {
  image = "images:ubuntu/22.04"
  output_image = "repacked-ubuntu"
  # publish_properties {
  #   description = "Trivial repackage with Packer"
  # }
}

build {
  sources = ["incus.incus-ubuntu"]
}
