# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Incus"
  description = "The Incus plugin can be used with HashiCorp Packer to create Incus-compatible VM and container images."
  identifier = "packer/dontlaugh/incus"
  component {
    type = "builder"
    name = "Incus"
    slug = "incus"
  }
}
