## hostinger ecommerce products update

Update a product

### Synopsis

Update a product's name, description, status, SEO settings or collections. Set status to
published to make it buyable, draft to hide it, or archived to retire it. SEO settings (title,
description, slug) left out keep their values. collection_changes adds or removes collections
without touching the others. Variants, prices and inventory are managed through the variant
endpoints, not here. Returns the updated product summary.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --collection-changes
    {
      "add": [
        {
          "id": "pcol_01J8Z5F8W9K8M4A7B3C2D1E0FG"
        }
      ],
      "remove": [
        {
          "id": "pcol_01J8Z5F8W9K8M4A7B3C2D1E0FG"
        }
      ]
    }

  --seo-settings
    {
      "description": "Soft combed cotton, unisex fit.",
      "slug": "blue-t-shirt",
      "title": "Blue T-Shirt | My Shop"
    }

```
hostinger ecommerce products update <store_id> <product_id> [flags]
```

### Examples

```
  hostinger ecommerce products update store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
      --collection-changes string   Adds the product to or removes it from collections without touching its other collections. Send at least one collection in add or remove. (JSON, @file or -; shape in the help above)
      --description string          The product description.
  -h, --help                        help for update
      --name string                 The product name.
      --seo-settings string         Search-engine settings to change. Send at least one field; omitted fields keep their current values. (JSON, @file or -; shape in the help above)
      --status string               Set "published" to make the product buyable, "draft" to hide it, or "archived" to retire it. (one of: draft, published, archived)
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

* [hostinger ecommerce products](hostinger_ecommerce_products.md)	 - Products commands

