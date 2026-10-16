## hostinger ecommerce products update

Update a product

### Synopsis

Update a product's name, description, status, SEO settings or collections. Set status to
published to make it buyable, draft to hide it, or archived to retire it. SEO settings (title,
description, slug) left out keep their values. collection_changes adds or removes collections
without touching the others. Variants, prices and inventory are managed through the variant
endpoints, not here. Returns the updated product summary.

```
hostinger ecommerce products update <store_id> <product_id> [flags]
```

### Options

```
      --collection-changes string   Adds the product to or removes it from collections without touching its other collections. Send at least one collection in add or remove. (JSON)
      --description string          The product description.
  -h, --help                        help for update
      --name string                 The product name.
      --seo-settings string         Search-engine settings to change. Send at least one field; omitted fields keep their current values. (JSON)
      --status string               Set "published" to make the product buyable, "draft" to hide it, or "archived" to retire it. (one of: draft, published, archived)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce products](hostinger_ecommerce_products.md)	 - Products commands

