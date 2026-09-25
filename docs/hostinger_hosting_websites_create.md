## hostinger hosting websites create

Create website

### Synopsis

Create a new website for the authenticated client.

You must choose which hosting order to create this website on. Pass that
order as `order_id` together with the domain name. List orders to see
available IDs; the website is provisioned on that order's hosting plan.

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
      --order-id int             Hosting order ID to create this website on. Choose the order whose hosting plan should host the new website. List orders to find available IDs.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

