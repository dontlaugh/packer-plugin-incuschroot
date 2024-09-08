_Note: this is a brand new (Sep 2024) fork of the LXD plugin_

# Packer Plugin Incus

Frienship ended with LXD. Incus is my new best friend.


## Installation

### Using pre-built releases

#### Using the `packer init` command

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
      source  = "github.com/dontlaugh/incus"
    }
  }
}
```

## Development

```
make install-packer-sdc
```
