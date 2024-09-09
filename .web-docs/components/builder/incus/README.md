Type: `incus`
Artifact BuilderId: `incus`

The `incus` Packer builder builds containers for Incus. The builder starts an Incus
container, runs provisioners within this container, then saves the container as
an Incus image. The Incus builder requires the `incus-client` package.

## Basic Example

Below is a fully functioning example.

**HCL**

```hcl
source "incus" "incus-ubuntu" {
  image = "images:ubuntu/22.04"
  output_image = "repacked-ubuntu"
}

build {
  sources = ["incus.incus-ubuntu"]
}
```

**JSON**

```json
{}
```

## Configuration Reference

### Required:

<!-- Code generated from the comments of the Config struct in builder/incus/config.go; DO NOT EDIT MANUALLY -->

- `image` (string) - The source image to use when creating the build
  container. This can be a (local or remote) image (name or fingerprint).

<!-- End of code generated from the comments of the Config struct in builder/incus/config.go; -->


### Optional:

<!-- Code generated from the comments of the Config struct in builder/incus/config.go; DO NOT EDIT MANUALLY -->

- `output_image` (string) - The name of the output artifact. Defaults to name.

- `container_name` (string) - Container Name

- `publish_remote_name` (string) - The (optional) name of the Incus remote on which to publish the
  container image.

- `command_wrapper` (string) - Lets you prefix all builder commands, such as with ssh for a
  remote build host. Defaults to `{{.Command}}`; i.e. no wrapper.

- `profile` (string) - Profile

- `init_sleep` (string) - The number of seconds to sleep between launching
  the LXD instance and provisioning it; defaults to 3 seconds.

- `publish_properties` (map[string]string) - Pass key values to the publish step to be set as properties on
  the output image. This is most helpful to set the description, but can be
  used to set anything needed. See https://stgraber.org/2016/03/30/lxd-2-0-image-management-512/
  for more properties.

- `launch_config` (map[string]string) - List of key/value pairs you wish to
  pass to incus launch via --config. Defaults to empty.

- `virtual_machine` (bool) - Create Incus virtual-machine image. Defaults to false for container image

- `skip_publish` (bool) - Skip execute `incus publish`; defaults to false

<!-- End of code generated from the comments of the Config struct in builder/incus/config.go; -->
