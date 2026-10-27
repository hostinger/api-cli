## hostinger vps firewall update-rule

Update firewall rule

### Synopsis

Update a specific firewall rule from a specified firewall.

Any virtual machine that has this firewall activated will lose sync with the firewall
and will have to be synced again manually.

Use this endpoint to modify existing firewall rules.

```
hostinger vps firewall update-rule <firewall-id> <rule-id> [flags]
```

### Examples

```
  hostinger vps firewall update-rule 9449049 8941182 --port 443 --protocol TCP --source any --source-detail 351.15.24.0/24
```

### Options

```
  -h, --help                help for update-rule
      --port string         Port or port range, ex: 1024:2048
      --protocol string     (one of: TCP, UDP, ICMP, GRE, any, ESP, AH, ICMPv6, SSH, HTTP, HTTPS, MySQL, PostgreSQL)
      --source string       (one of: any, custom)
      --source-detail any   IP range, CIDR, single IP or any
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

