## hostinger gpu catalog pool-groups

Get pool groups

### Synopsis

Retrieve GPU pool groups, ordered by name.

Each group bundles the pools of one GPU model and lists its currently available pools,
whether any of them is available and the lowest hourly price among its instance types.
The price is returned even when the group is out of stock.

Use this endpoint to show which GPU models can be ordered and from what price.

```
hostinger gpu catalog pool-groups [flags]
```

### Examples

```
  hostinger gpu catalog pool-groups
```

### Options

```
      --currency USD   Currency of the returned prices, USD by default (one of: USD, ARS, AMD, AUD, AZN, BDT, BRL, BTC, BYN, CAD, CHF, CLP, CNY, COP, CZK, DKK, EGP, EUR, GBP, GEL, HKD, HRK, HUF, IDR, ILS, INR, JPY, KGS, KRW, KZT, MAD, MXN, MYR, NGN, NOK, NZD, PEN, PHP, PKR, PLN, SEK, SGD, THB, TJS, TRY, TWD, UAH, VND, ZAR)
  -h, --help           help for pool-groups
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

* [hostinger gpu catalog](hostinger_gpu_catalog.md)	 - Catalog commands

