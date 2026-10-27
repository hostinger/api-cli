## hostinger vps virtual-machines

Virtual machine commands

```
hostinger vps virtual-machines [flags]
```

### Options

```
  -h, --help   help for virtual-machines
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

* [hostinger vps](hostinger_vps.md)	 - VPS commands
* [hostinger vps virtual-machines attached-public-keys](hostinger_vps_virtual-machines_attached-public-keys.md)	 - Get attached public keys
* [hostinger vps virtual-machines get](hostinger_vps_virtual-machines_get.md)	 - Get virtual machine details
* [hostinger vps virtual-machines list](hostinger_vps_virtual-machines_list.md)	 - Get virtual machines
* [hostinger vps virtual-machines metrics](hostinger_vps_virtual-machines_metrics.md)	 - Get metrics
* [hostinger vps virtual-machines purchase](hostinger_vps_virtual-machines_purchase.md)	 - Purchase new virtual machine
* [hostinger vps virtual-machines recreate](hostinger_vps_virtual-machines_recreate.md)	 - Recreate virtual machine
* [hostinger vps virtual-machines reset-hostname](hostinger_vps_virtual-machines_reset-hostname.md)	 - Reset hostname
* [hostinger vps virtual-machines restart](hostinger_vps_virtual-machines_restart.md)	 - Restart virtual machine
* [hostinger vps virtual-machines set-hostname](hostinger_vps_virtual-machines_set-hostname.md)	 - Set hostname
* [hostinger vps virtual-machines set-nameservers](hostinger_vps_virtual-machines_set-nameservers.md)	 - Set nameservers
* [hostinger vps virtual-machines set-panel-password](hostinger_vps_virtual-machines_set-panel-password.md)	 - Set panel password
* [hostinger vps virtual-machines set-root-password](hostinger_vps_virtual-machines_set-root-password.md)	 - Set root password
* [hostinger vps virtual-machines setup](hostinger_vps_virtual-machines_setup.md)	 - Setup purchased virtual machine
* [hostinger vps virtual-machines start](hostinger_vps_virtual-machines_start.md)	 - Start virtual machine
* [hostinger vps virtual-machines stop](hostinger_vps_virtual-machines_stop.md)	 - Stop virtual machine

