## hostinger reach integrations disconnect-wordpress-site

Disconnect WordPress site

### Synopsis

Disconnect a WordPress site from the profile.

The Reach plugin on the site stops sending subscribers and store data, and the site's
products and coupons are removed from Reach; contacts and past orders stay. Take `domain`
from the list connected websites endpoint. Fails with not found when the domain is not
connected to this profile. This cannot be undone through the API - the customer reconnects
from the WordPress plugin - so confirm the site and the plan with the customer before calling.

```
hostinger reach integrations disconnect-wordpress-site <profile-uuid> [flags]
```

### Examples

```
  hostinger reach integrations disconnect-wordpress-site 550e8400-e09b-41d4-a716-400055000000 --domain example.com
```

### Options

```
      --domain string   
  -h, --help            help for disconnect-wordpress-site
  -y, --yes             Skip the confirmation prompt
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

* [hostinger reach integrations](hostinger_reach_integrations.md)	 - Integrations commands

