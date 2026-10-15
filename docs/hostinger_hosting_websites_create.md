## hostinger hosting websites create

Create website

### Synopsis

Deprecated: use `POST /api/hosting/v1/orders/{order_id}/onboardings` (start website setup)
instead. It takes the same `domain` and `datacenter_code`, can also install WordPress, and
returns the created setup.

Create a new website for the authenticated client.

You must choose which hosting order to create this website on. Pass that
order as `order_id` together with the domain name. List orders to see
available IDs; the website is provisioned on that order's hosting plan.

Only Web and Cloud hosting orders are accepted. To create a website on an Agency
Plan order, use `POST /api/agency-hosting/v1/orders/{order_id}/websites/setups`.

When the order has no hosting account yet, a new hosting account is set up first, in the
datacenter given by `datacenter_code` or the order's default datacenter when it is omitted.
Subsequent websites will be hosted on the same datacenter automatically.

Website creation is asynchronous and takes up to a few minutes. Poll the list website
setups endpoint with the `domain` filter every 10 to 15 seconds and wait for `status:
completed` before uploading files, deploying or creating databases. While the setup is
`running`, endpoints that operate on the website may respond with `404` or `409`.
`is_enabled` on the websites list reflects suspension, not readiness.

```
hostinger hosting websites create [flags]
```

### Options

```
      --datacenter-code string    Datacenter for the hosting account when the order does not have one yet. Omit to use the order's default datacenter. Ignored when the order already has a hosting account. Must be an available datacenter; list available datacenters to find valid codes.
      --domain blog.example.com   Domain name for the website. Cannot start with "www." and can have at most one subdomain level, e.g. blog.example.com.
  -h, --help                      help for create
      --order-id int              Hosting order ID to create this website on. Choose the order whose hosting plan should host the new website. List orders to find available IDs.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

