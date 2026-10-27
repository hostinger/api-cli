## hostinger reach integrations list-ecommerce-stores

List Ecommerce stores

### Synopsis

Get all of the customer's Hostinger eCommerce stores, each with its `store_id` and
`connection_status` for the profile in the path.

A store is connected to at most one profile. `connected_to_other_profile` does not name that
profile - call this for each profile to find it. The list is empty, not an error, when the
eCommerce service cannot be reached. Pass `store_id` to the disconnect eCommerce store
endpoint. Connecting an AI Website Builder site connects the store selling on it.

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
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach integrations](hostinger_reach_integrations.md)	 - Integrations commands

