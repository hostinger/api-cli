## hostinger agency-hosting website-setups create

Create a new website

### Synopsis

Provisions a new website on one of your Agency Plan hosting orders.

Choose the datacenter, stack (`flavor`), and PHP version for the site. Optionally attach
your own `domain` — omit it, set it to `null`, or leave it unavailable and a free
`*.hostingersite.com` subdomain is generated instead — and/or install WordPress by
supplying the `wordpress` details (admin account, site title, and language).

Common setups:
- Plain PHP site: `flavor` set to `php-fpm`, with `settings.php.version`; omit
  `wordpress` and `type`.
- WordPress site: `flavor` set to the desired WordPress version (e.g. `wp-7.0`), plus
  the `wordpress` block (admin account, title, language).
- Static/Node.js frontend app: `flavor` set to `php-fpm` and `type` set to
  `node-static`.

Provisioning runs in the background, so the response returns immediately with a setup UUID
that identifies the job. The new website becomes reachable once provisioning finishes.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --settings
    {
      "php": {
        "version": "8.3"
      }
    }

  --wordpress
    {
      "admin": {
        "email": "admin@example.com",
        "password": "S3cureP@ssw0rd",
        "user": "admin"
      },
      "language": "en_US",
      "title": "My Blog"
    }

```
hostinger agency-hosting website-setups create <order_id> [flags]
```

### Examples

```
  hostinger agency-hosting website-setups create 123456 --datacenter-code us-east --flavor wp-6.8.2 --settings '{"php":{"version":"8.3"}}'
```

### Options

```
      --datacenter-code string      Datacenter code where the website should be provisioned. Available codes depend on live capacity and are not a fixed set.
      --domain string               Primary domain to attach to the website. Omit or set to null to get a free auto-generated *.hostingersite.com subdomain instead.
      --flavor wp-<major>.<minor>   Setup flavor: a specific WordPress version in the format wp-<major>.<minor> or `wp-<major>.<minor>.<patch>` (e.g. `wp-6.8.2`), or `php-fpm` for a plain PHP stack. Generic versions like `wp-latest` are not allowed.
  -h, --help                        help for create
      --settings string             Website settings (JSON, @file or -; shape in the help above)
      --type string                 Website type (one of: node-static)
      --wordpress string            WordPress installation options (JSON, @file or -; shape in the help above)
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

* [hostinger agency-hosting website-setups](hostinger_agency-hosting_website-setups.md)	 - Website Setups commands

