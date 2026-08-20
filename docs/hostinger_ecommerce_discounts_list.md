## hostinger ecommerce discounts list

List discounts

### Synopsis

List a store's discounts. Filter by free text over code and name, or by disabled state.
Amounts for fixed discounts are integers in the smallest currency unit; percentage
discounts carry a whole-number value between 1 and 100.

```
hostinger ecommerce discounts list <store_id> [flags]
```

### Options

```
  -h, --help                 help for list
      --is-disabled string   Filter by disabled state. (one of: true, false)
      --page int             Page number
      --q string             Free-text search over discount code and name.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce discounts](hostinger_ecommerce_discounts.md)	 - Discounts commands

