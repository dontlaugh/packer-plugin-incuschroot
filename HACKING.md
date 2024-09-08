# Developer Guide

Packer runs 3rd party plugins as subprocesses, and communicates with them over
RPC via stdin/stdout.

## Testing Your Binary

Running `make` drops a binary in the root of this repo.

The **tests/** directory contains a relative symlink to this binary, which means
you can do the following to use the HCL files there.

```
cd tests/
packer build [FILE].pkr.hcl
```

## Testing Your Own HCL Configs

There are two steps to make Packer use your development version of this plugin.

1. Copy the built binary to a suitable place on disk The easiest thing to do is
   to **move the binary to the directory where you invoke packer**.
2. Comment out any official released version from `required_plugins` in your HCL

For step 2, simply make sure there is no reference to "github.com/dontlaugh/incus"
in your `required_plugins` config.

```diff
packer {
  required_plugins {
-    incus = {
-      version = ">=0.3.1"
-      source = "github.com/dontlaugh/incus"
-    }
+    # incus = {
+    #  version = ">=0.3.1"
+    #  source = "github.com/dontlaugh/incus"
+    # }
   }
 }
```

## Debug Logging

Set `PACKER_LOG=1` in your environment for extra logging. This log output will
also indicate whether Packer has located your custom built plugin.

## HCL Config Codegen

If you make a change to a config struct, you need to regenerate some files.

```
make generate
```

## Regarding zlconf/go-cty

We _cannot_ upgrade this dependency until this issue is resolved

* https://github.com/hashicorp/packer-plugin-sdk/issues/12

