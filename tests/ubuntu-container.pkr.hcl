
source "incus" "incus-ubuntu" {
  image = "images:ubuntu/22.04"
  output_image = "repacked-ubuntu-${formatdate("YYYYMMDDhhmmss", timestamp())}"
}

build {
  sources = ["incus.incus-ubuntu"]
}
