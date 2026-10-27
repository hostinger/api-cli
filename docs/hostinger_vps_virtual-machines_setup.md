## hostinger vps virtual-machines setup

Setup purchased virtual machine

### Synopsis

Setup newly purchased virtual machine with `initial` state.

Use this endpoint to configure and initialize purchased VPS instances.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --public-key
    {
      "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2X...",
      "name": "my-key"
    }

```
hostinger vps virtual-machines setup <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines setup 1268054 --data-center-id 19 --template-id 1130
```

### Options

```
      --data-center-id int           Data center ID
      --enable-backups               Enable weekly backup schedule (default true)
  -h, --help                         help for setup
      --hostname string              Override default hostname of the virtual machine
      --install-monarx               Install Monarx malware scanner (if supported)
      --ns1 string                   Name server 1
      --ns2 string                   Name server 2
      --password string              Password for the virtual machine. If not provided, random password will be generated. Password will not be shown in the response.
      --post-install-script-id int   Post-install script ID
      --public-key string            Use SSH key (JSON, @file or -; shape in the help above)
      --template-id int              Template ID
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

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

