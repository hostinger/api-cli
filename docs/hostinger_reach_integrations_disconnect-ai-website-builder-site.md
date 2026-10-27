## hostinger reach integrations disconnect-ai-website-builder-site

Disconnect AI Website Builder site

### Synopsis

Disconnect an AI Website Builder site from the profile; the site stops sending subscribers
to Reach.

Take `domain` from the list connected websites endpoint. Works for a site that was already
deleted in the builder. Fails with not found when the domain is not connected to this
profile. Contacts and any connected eCommerce store stay. The site stays disconnected until
it is connected again, so confirm the site and the plan with the customer before calling.

```
hostinger reach integrations disconnect-ai-website-builder-site <profile-uuid> [flags]
```

### Examples

```
  hostinger reach integrations disconnect-ai-website-builder-site 550e8400-e09b-41d4-a716-400055000000 --domain example.com
```

### Options

```
      --domain string   
  -h, --help            help for disconnect-ai-website-builder-site
  -y, --yes             Skip the confirmation prompt
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

* [hostinger reach integrations](hostinger_reach_integrations.md)	 - Integrations commands

