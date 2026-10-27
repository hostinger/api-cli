## hostinger ecommerce products create-physical

Create physical product

### Synopsis

Create a published physical product with a single variant priced in the store currency.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --seo-settings
    {
      "description": "Soft combed cotton, unisex fit.",
      "slug": "blue-t-shirt",
      "title": "Blue T-Shirt | My Shop"
    }

```
hostinger ecommerce products create-physical <store_id> [flags]
```

### Examples

```
  hostinger ecommerce products create-physical store_01J8Z5F8W9K8M4A7B3C2D1E0FG --name 'My Product' --price 1000
```

### Options

```
      --currency string       ISO 4217 currency code. Defaults to the store's default currency when omitted.
      --description string    The product description.
  -h, --help                  help for create-physical
      --name string           The product name.
      --price int             Price in the smallest currency unit (e.g. cents). Must be positive.
      --seo-settings string   Search-engine settings. Omitted fields are generated from the name and description; a slug already in use gets a numeric suffix. Send at least one field. (JSON, @file or -; shape in the help above)
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

