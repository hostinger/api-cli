## hostinger vps post-install-scripts update

Update post-install script

### Synopsis

Update a specific post-install script.

Use this endpoint to modify existing automation scripts.

```
hostinger vps post-install-scripts update <post-install-script-id> [flags]
```

### Examples

```
  hostinger vps post-install-scripts update 9568314 --content '#!/bin/bash

echo '\''Hello, World!'\''' --name 'My Script'
```

### Options

```
      --content string   Content of the script
  -h, --help             help for update
      --name string      Name of the script
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger vps post-install-scripts](hostinger_vps_post-install-scripts.md)	 - Post-install scripts commands

