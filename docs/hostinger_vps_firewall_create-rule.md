## hostinger vps firewall create-rule

Create firewall rule

### Synopsis

Create new firewall rule for a specified firewall.

By default, the firewall drops all incoming traffic,
which means you must add accept rules for all ports you want to use.

Any virtual machine that has this firewall activated will lose sync with the firewall
and will have to be synced again manually.

Use this endpoint to add new security rules to firewalls.

```
hostinger vps firewall create-rule <firewall-id> [flags]
```

### Examples

```
  hostinger vps firewall create-rule 9449049 --port 443 --protocol TCP --source any --source-detail 351.15.24.0/24
```

### Options

```
  -h, --help                help for create-rule
      --port string         Port or port range, ex: 1024:2048
      --protocol string     (one of: TCP, UDP, ICMP, GRE, any, ESP, AH, ICMPv6, SSH, HTTP, HTTPS, MySQL, PostgreSQL)
      --source string       (one of: any, custom)
      --source-detail any   IP range, CIDR, single IP or any
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

* [hostinger vps firewall](hostinger_vps_firewall.md)	 - Firewall commands

