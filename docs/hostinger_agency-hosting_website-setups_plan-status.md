## hostinger agency-hosting website-setups plan-status

Get Agency Plan website setup status

### Synopsis

Returns the current status of an Agency Plan website setup started via the setups
endpoint.

Poll this endpoint using the `setup_uuid` returned from the provisioning request until
`status` becomes `completed`, at which point `website_uid` identifies the new website.

```
hostinger agency-hosting website-setups plan-status <order_id> <setup_uuid> [flags]
```

### Options

```
  -h, --help   help for plan-status
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting website-setups](hostinger_agency-hosting_website-setups.md)	 - Website Setups commands

