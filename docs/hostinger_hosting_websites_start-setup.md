## hostinger hosting websites start-setup

Start website setup

### Synopsis

Starts a website setup on a Web or Cloud hosting order and returns the created setup
right away; the website itself is provisioned asynchronously. Poll the list website
setups endpoint (`hosting_websites_list-setups`) with the `domain` filter every 10 to 15
seconds and wait for `status: completed` before uploading files, deploying or creating
databases.

Omit `type` for an empty website. `type: wordpress` installs WordPress in the website
root with the admin user, email and password from `wordpress`, the domain as the site
title, and `en_US` when `wordpress.language` is omitted. The headless types
(`headless_empty`, `headless_wordpress`, `headless_ecommerce`, `headless_pocketbase`)
create a headless website; `headless_wordpress` additionally installs WordPress into the
`cms` directory of the website root with generated credentials. `headless_empty` installs
nothing: WordPress can be added later through
`POST /api/hosting/v1/accounts/{username}/wordpress/installations`
(`wordpress_installations_install`) with `type: headless`.

Omit `domain` to set the website up on a generated temporary free subdomain. To know the
free subdomain before the call, so a timed-out call can be looked up and retried on the
same domain, generate one through `POST /api/hosting/v1/domains/free-subdomains`
(`hosting_domains_generate-free-subdomain`) and send it as `domain`. It is accepted for
24 hours, only from the client that generated it.

When the order has no hosting account yet, the setup creates one first, in the datacenter
given by `datacenter_code` or the order's default datacenter when it is omitted.
Returns 404 when the order does not exist or is not accessible to the authenticated
client, and 409 with a `Retry-After` header while a setup for the same domain is still
running.

The order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --wordpress
    {
      "admin": {
        "email": "admin@example.com",
        "password": "S3curePass123",
        "user": "site_admin"
      },
      "is_ai_builder": false,
      "language": "en_US"
    }

```
hostinger hosting websites start-setup [order-id] [flags]
```

### Examples

```
  hostinger hosting websites start-setup 12345
```

### Options

```
      --datacenter-code string                                Datacenter for the hosting account when the order does not have one yet. Omit to use the order's default datacenter. Ignored when the order already has a hosting account. List available datacenters to find valid codes.
      --domain POST /api/hosting/v1/domains/free-subdomains   Customer-owned domain, or a free subdomain you generated through POST /api/hosting/v1/domains/free-subdomains (`hosting_domains_generate-free-subdomain`) in the last 24 hours. Cannot start with "www." and can have at most one subdomain level, e.g. `blog.example.com`. Omit or `null` to set the website up on a generated temporary free subdomain.
  -h, --help                                                  help for start-setup
      --type null                                             Website type. Omit or null for an empty website. `wordpress` installs WordPress in the website root and requires `wordpress`. The headless types (`headless_empty`, `headless_wordpress`, `headless_ecommerce`, `headless_pocketbase`) create a headless website; `headless_wordpress` additionally installs WordPress into the `cms` directory with generated credentials. `headless_empty` installs nothing, so WordPress can be added later through `POST /api/hosting/v1/accounts/{username}/wordpress/installations` (`wordpress_installations_install`) with `type: headless`. (one of: wordpress, headless_empty, headless_wordpress, headless_ecommerce, headless_pocketbase)
      --wordpress type                                        WordPress install settings. Required when type is `wordpress`, not allowed otherwise. The site title is the domain. (JSON, @file or -; shape in the help above)
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

