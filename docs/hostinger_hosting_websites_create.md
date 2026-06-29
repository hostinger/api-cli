## hostinger hosting websites create

Create website

### Synopsis

Create a new website for the authenticated client.

Provide the domain name and associated order ID to create a new website.
The datacenter_code parameter is required when creating the first website
on a new hosting plan - this will set up and configure new hosting account
in the selected datacenter.

Subsequent websites will be hosted on the same datacenter automatically.

Website creation takes up to a few minutes to complete. Check the
websites list endpoint to see when your new website becomes available.

```
hostinger hosting websites create [flags]
```

### Options

```
      --datacenter-code string   Datacenter code. This parameter is required when creating the first website on a new hosting plan.
      --domain string            Domain name for the website. Cannot start with "www."
  -h, --help                     help for create
      --order-id int             ID of the associated order
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

