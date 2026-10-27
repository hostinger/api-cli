## hostinger vps firewall replace-all-rules-in-group

Replace all firewall rules in group

### Synopsis

Replaces all firewall rules within a specified firewall group with the provided set of rules
in a single atomic operation, instead of creating or deleting rules one by one.

Any virtual machine using this firewall group will need to be synchronized after replacing rules;
pass the "sync" parameter to trigger synchronization immediately.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --rules
    [
      {
        "port": "443",
        "protocol": "TCP",
        "source": "any",
        "source_detail": "351.15.24.0/24"
      }
    ]

```
hostinger vps firewall replace-all-rules-in-group <firewall-id> [flags]
```

### Examples

```
  hostinger vps firewall replace-all-rules-in-group 9449049 --rules @rules.json
```

### Options

```
  -h, --help           help for replace-all-rules-in-group
      --rules string   The complete set of firewall rules that atomically replaces all existing rules in the group (JSON, @file or -; shape in the help above)
      --sync           Synchronize the firewall group to all its virtual machines after replacing the rules
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

* [hostinger vps firewall](hostinger_vps_firewall.md)	 - Firewall commands

