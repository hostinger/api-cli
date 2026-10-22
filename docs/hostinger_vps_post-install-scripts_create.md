## hostinger vps post-install-scripts create

Create post-install script

### Synopsis

Add a new post-install script to your account, which can then be used after virtual machine installation.

The script contents will be saved to the file `/post_install` with executable attribute set
and will be executed once virtual machine is installed.
The output of the script will be redirected to `/post_install.log`. Maximum script size is 48KB.

Use this endpoint to create automation scripts for VPS setup tasks.

```
hostinger vps post-install-scripts create [flags]
```

### Examples

```
  hostinger vps post-install-scripts create --content '#!/bin/bash

echo '\''Hello, World!'\''' --name 'My Script'
```

### Options

```
      --content string   Content of the script
  -h, --help             help for create
      --name string      Name of the script
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger vps post-install-scripts](hostinger_vps_post-install-scripts.md)	 - Post-install scripts commands

