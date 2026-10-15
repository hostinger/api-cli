## hostinger hosting websites start-setup

Start website setup

### Synopsis

Starts a website setup on a Web or Cloud hosting order and returns the created setup
right away; the website itself is provisioned asynchronously. Poll the list website
setups endpoint with the `domain` filter every 10 to 15 seconds and wait for
`status: completed` before uploading files, deploying or creating databases.

Omit `type` for an empty website. `type: wordpress` installs WordPress in the website
root with the admin user, email and password from `wordpress`, the domain as the site
title, and `en_US` when `wordpress.language` is omitted. The headless types
(`headless_wordpress`, `headless_ecommerce`, `headless_pocketbase`) create a headless
website; `headless_wordpress` additionally installs WordPress into the `cms` directory
of the website root with generated credentials.

Omit `domain` to set the website up on a generated temporary free subdomain.

When the order has no hosting account yet, the setup creates one first, in the datacenter
given by `datacenter_code` or the order's default datacenter when it is omitted.
Returns 404 when the order does not exist or is not accessible to the authenticated
client, and 409 with a `Retry-After` header while a setup for the same domain is still
running.

```
hostinger hosting websites start-setup <order_id> [flags]
```

### Options

```
      --datacenter-code string    Datacenter for the hosting account when the order does not have one yet. Omit to use the order's default datacenter. Ignored when the order already has a hosting account. List available datacenters to find valid codes.
      --domain blog.example.com   Customer-owned domain. Cannot start with "www." and can have at most one subdomain level, e.g. blog.example.com. Omit or `null` to set the website up on a generated temporary free subdomain.
  -h, --help                      help for start-setup
      --type null                 Website type. Omit or null for an empty website. `wordpress` installs WordPress in the website root and requires `wordpress`. The headless types (`headless_wordpress`, `headless_ecommerce`, `headless_pocketbase`) create a headless website; `headless_wordpress` additionally installs WordPress into the `cms` directory with generated credentials. (one of: wordpress, headless_wordpress, headless_ecommerce, headless_pocketbase)
      --wordpress type            WordPress install settings. Required when type is `wordpress`, not allowed otherwise. The site title is the domain. (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

