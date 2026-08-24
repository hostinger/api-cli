## hostinger reach profiles connected-sending-domain

Get connected sending domain

### Synopsis

Get the sending domain connected to the profile, its verification status and any suspended
sender addresses.

Campaigns only go out once a domain is connected and active, so this is the cheapest way to
check that precondition before building one. A profile with no domain connected returns the
same shape with every field set to `null`. For the individual MX, SPF, DKIM and DMARC records
behind the status, use the DNS status endpoint.

```
hostinger reach profiles connected-sending-domain <profile-uuid> [flags]
```

### Options

```
  -h, --help   help for connected-sending-domain
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

