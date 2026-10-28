## hostinger reach integrations list-ecommerce-stores

List Ecommerce stores

### Synopsis

Get all of the customer's Hostinger eCommerce stores, each with its `store_id` and
`connection_status` for the profile in the path.

A store is connected to at most one profile. `connected_to_other_profile` does not name that
profile - call this for each profile to find it. The list is empty, not an error, when the
eCommerce service cannot be reached. Pass `store_id` to the connect or disconnect eCommerce
store endpoint.

```
hostinger reach integrations list-ecommerce-stores <profile-uuid> [flags]
```

### Examples

```
  hostinger reach integrations list-ecommerce-stores 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for list-ecommerce-stores
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
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

