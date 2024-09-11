
# Packer Plugin Incus Chroot


Goal: build disk images in a chroot on a pre-existing VM.

_Note: this is a brand new (Sep 2024) fork of the LXD plugin. You probably want
[bketelsen/packer-plugin-incus](https://github.com/bketelsen/packer-plugin-incus)._

## Installation

### Using pre-built releases

#### Using the `packer init` command

<details>
<summary> packer init not yet supported </summary>

Starting from version 1.7, Packer supports a new `packer init` command allowing
automatic installation of Packer plugins. Read the
[Packer documentation](https://www.packer.io/docs/commands/init) for more information.

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    incus = {
      version = ">= 1.0.0"
      source  = "github.com/dontlaugh/incuschroot"
    }
  }
}
```

</details>

## Development

Get the sdc tool from Hashicorp.

```
make install-packer-sdc
```

Then see [HACKING.md](./HACKING.md)
