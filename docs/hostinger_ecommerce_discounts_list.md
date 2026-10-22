## hostinger ecommerce discounts list

List discounts

### Synopsis

List a store's discounts. Filter by free text over code and name, or by disabled state.
Amounts for fixed discounts are integers in the smallest currency unit; percentage
discounts carry a whole-number value between 1 and 100.

```
hostinger ecommerce discounts list <store_id> [flags]
```

### Examples

```
  hostinger ecommerce discounts list store_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
      --all                  Fetch every page and print all items
  -h, --help                 help for list
      --is-disabled string   Filter by disabled state. (one of: true, false)
      --page int             Page number
      --q string             Free-text search over discount code and name.
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

* [hostinger ecommerce discounts](hostinger_ecommerce_discounts.md)	 - Discounts commands

