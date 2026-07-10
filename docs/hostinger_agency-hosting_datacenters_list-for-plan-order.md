## hostinger agency-hosting datacenters list-for-plan-order

List available datacenters for an Agency Plan order

### Synopsis

Lists the datacenters available for provisioning a new website on the given Agency Plan
hosting order.

Each datacenter includes a `pinger_url` you can ping from the client to measure round-trip
latency; comparing the results across datacenters lets you pick the nearest one (lowest
ping) before choosing its `code` as the `datacenter_code` when creating a website setup.

```
hostinger agency-hosting datacenters list-for-plan-order <order_id> [flags]
```

### Options

```
  -h, --help   help for list-for-plan-order
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting datacenters](hostinger_agency-hosting_datacenters.md)	 - Datacenters commands

