## hostinger agency-hosting website-setups provision-plan

Provision a new Agency Plan website

### Synopsis

Provisions a new website on one of your Agency Plan hosting orders.

Choose the datacenter, stack (`flavor`), and PHP version for the site. Optionally attach
your own `domain` — omit it, set it to `null`, or leave it unavailable and a free
`*.hostingersite.com` subdomain is generated instead — and/or install WordPress by
supplying the `wordpress` details (admin account, site title, and language).

Common setups:
- **Plain PHP site**: `flavor` set to `php-fpm`, with `settings.php.version`; omit
  `wordpress` and `type`.
- **WordPress site**: `flavor` set to the desired WordPress version (e.g. `wp-7.0`), plus
  the `wordpress` block (admin account, title, language).
- **Static/Node.js frontend app**: `flavor` set to `php-fpm` and `type` set to
  `node-static`.

Provisioning runs in the background, so the response returns immediately with a setup UUID
that identifies the job. The new website becomes reachable once provisioning finishes.

```
hostinger agency-hosting website-setups provision-plan <order_id> [flags]
```

### Options

```
      --clone string                Clone the new website from an existing website (JSON)
      --datacenter-code string      Datacenter code where the website should be provisioned. Available codes depend on live capacity and are not a fixed set.
      --derive-domain string        Derive the domain from an existing vhost (JSON)
      --domain string               Primary domain to attach to the website. Omit or set to null to get a free auto-generated *.hostingersite.com subdomain instead.
      --flavor wp-<major>.<minor>   Setup flavor: a specific WordPress version in the format wp-<major>.<minor> or `wp-<major>.<minor>.<patch>` (e.g. `wp-6.8.2`), or `php-fpm` for a plain PHP stack. Generic versions like `wp-latest` are not allowed.
  -h, --help                        help for provision-plan
      --settings string             Website settings (JSON)
      --type string                 Website type (one of: horizons, node-static)
      --wordpress string            WordPress installation options (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting website-setups](hostinger_agency-hosting_website-setups.md)	 - Website Setups commands

