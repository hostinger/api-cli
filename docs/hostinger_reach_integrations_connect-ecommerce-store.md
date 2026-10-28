## hostinger reach integrations connect-ecommerce-store

Connect Ecommerce store

### Synopsis

Connect one of the customer's eCommerce stores to the profile, so Reach receives its
customers and orders.

Take `store_id` from the list eCommerce stores endpoint. Fails with not found
(`ECOMMERCE_STORE_NOT_FOUND`) when the store is not one of the customer's own stores, also
when the eCommerce service cannot be reached. Fails with a conflict
(`ECOMMERCE_CONNECTION_ALREADY_EXISTS`) when the store is connected to any profile or this
profile already has a store - to move a store to another plan, disconnect it from the
profile that holds it first.

```
hostinger reach integrations connect-ecommerce-store <profile-uuid> <store-id> [flags]
```

### Examples

```
  hostinger reach integrations connect-ecommerce-store 550e8400-e09b-41d4-a716-400055000000 ec_123
```

### Options

```
  -h, --help   help for connect-ecommerce-store
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

