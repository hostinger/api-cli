## hostinger ecommerce products list

List products

### Synopsis

List a store's products newest first as lean summaries (name, status, thumbnail, variant
count and price range). Prices are integers in the smallest currency unit and live on
variants. Filter by status, free text or a set of product ids. Use include=variants to
embed each product's variants with prices and inventory, and include=media to embed its media.

```
hostinger ecommerce products list <store_id> [flags]
```

### Examples

```
  hostinger ecommerce products list store_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
      --all                   Fetch every page and print all items
  -h, --help                  help for list
      --include strings       Opt-in heavy data: "variants" embeds each product's variants; "media" embeds its media. (one of: variants, media)
      --page int              Page number
      --product-ids strings   Restrict to these product ids. Doubles as a single-product lookup. Up to 200 ids.
      --q string              Free-text search over product title and SKU.
      --status strings        Product statuses to include. (one of: draft, proposed, published, rejected, archived)
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

* [hostinger ecommerce products](hostinger_ecommerce_products.md)	 - Products commands

