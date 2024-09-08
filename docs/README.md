The Incus plugin allows building images for Incus, by starting an Incus instance (container or VM),
running provisioners within this container, taking a snapshot of the instance, and
publishing it as an Incus image.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then 
run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    lxd = {
      source  = "github.com/dontlaugh/incus"
      version = "~> 1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/dontlaugh/incus"
```

### Components

#### Builders

TODO: fix this link

* [incus](/packer/integrations/hashicorp/lxd/latest/components/builder/lxd) - The Incus builder builds images with LXD 
  by starting an instance, provisioning it, and exporting it as a tar.gz archive of the root file system.
