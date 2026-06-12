## hapi vps firewall create-rule

Create firewall rule

### Synopsis

Create new firewall rule for a specified firewall.

By default, the firewall drops all incoming traffic,
which means you must add accept rules for all ports you want to use.

Any virtual machine that has this firewall activated will lose sync with the firewall
and will have to be synced again manually.

Use this endpoint to add new security rules to firewalls.

```
hapi vps firewall create-rule <firewall-id> [flags]
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
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps firewall](hapi_vps_firewall.md)	 - Firewall commands

