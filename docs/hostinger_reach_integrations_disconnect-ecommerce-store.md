## hostinger reach integrations disconnect-ecommerce-store

Disconnect Ecommerce store

### Synopsis

Disconnect an eCommerce store from the profile.

Reach stops receiving the store's customers and orders, and its products and coupons are
removed from Reach; contacts and past orders stay. Take `store_id` from the list eCommerce
stores endpoint; for a store that is `connected_to_other_profile`, call this with that
profile. Also succeeds when the store is not connected to this profile, so success is not
proof that anything was disconnected. Confirm the store and the plan with the customer
before calling.

```
hostinger reach integrations disconnect-ecommerce-store <profile-uuid> <store-id> [flags]
```

### Examples

```
  hostinger reach integrations disconnect-ecommerce-store 550e8400-e09b-41d4-a716-400055000000 ec_123
```

### Options

```
  -h, --help   help for disconnect-ecommerce-store
  -y, --yes    Skip the confirmation prompt
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

