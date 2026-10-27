## hostinger reach integrations connect-ai-website-builder-site

Connect AI Website Builder site

### Synopsis

Connect an AI Website Builder site the customer built manually to the profile, so its
subscribers flow into Reach.

Fails with not found when the customer owns no AI Website Builder site with that domain.
Fails with `DOMAIN_ALREADY_EXISTS` when the domain is connected to any profile - to move a
site to another plan, disconnect it from the old profile first - and with
`WEBSITE_CONNECTION_ALREADY_EXISTS` when this profile already has an AI Website Builder site.
The eCommerce store selling on the site, if it is not connected anywhere, is connected too.
WordPress sites and sites made with the agentic AI Website Builder cannot be connected here.

```
hostinger reach integrations connect-ai-website-builder-site <profile-uuid> [flags]
```

### Examples

```
  hostinger reach integrations connect-ai-website-builder-site 550e8400-e09b-41d4-a716-400055000000 --domain example.com
```

### Options

```
      --domain string   
  -h, --help            help for connect-ai-website-builder-site
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

