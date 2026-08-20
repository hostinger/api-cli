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

### Options

```
  -h, --help                  help for list
      --include strings       Opt-in heavy data: "variants" embeds each product's variants; "media" embeds its media. (one of: variants, media)
      --page int              Page number
      --product-ids strings   Restrict to these product ids. Doubles as a single-product lookup. Up to 200 ids.
      --q string              Free-text search over product title and SKU.
      --status strings        Product statuses to include. (one of: draft, proposed, published, rejected, archived)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce products](hostinger_ecommerce_products.md)	 - Products commands

