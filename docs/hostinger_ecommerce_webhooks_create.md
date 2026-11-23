## hostinger ecommerce webhooks create

Create an ecommerce webhook

### Synopsis

Create an enabled subscription for your store to receive ecommerce events in external automation
systems. Save the whsec_ signing secret securely; it is returned once. Up to ten endpoints, including
disabled endpoints, may exist per store. Deliveries follow Standard Webhooks: verify the webhook-id,
webhook-timestamp and webhook-signature headers against the raw body with an official Standard Webhooks
library before parsing. Events can arrive out of order or more than once; deduplicate by event id.

400: the URL is not a public destination, or webhook_sales_channel_invalid (a channel ID does not
belong to this store). 409 codes: webhook_endpoint_limit_reached, webhook_store_inactive,
webhook_endpoint_catalog_busy (a catalog change is in progress; retry). Errors return the upstream code in
`message`, as `[Ecommerce:<code>] <reason>`.

```
hostinger ecommerce webhooks create <store_id> [flags]
```

### Examples

```
  hostinger ecommerce webhooks create store_01J8Z5F8W9K8M4A7B3C2D1E0FG --events order.placed,product.updated --url https://example.com/webhooks/hostinger
```

### Options

```
      --events order.placed         Event types to subscribe to. - order.placed: a customer places an order, before any online payment is completed. Orders paid online may not have an order number yet. - `order.paid`: an order becomes fully paid: an online payment is captured, a free order is placed or the merchant marks it as paid. Partial payments don't send it. - `order.cancelled`: an unpaid order is canceled, for example by the merchant, a connected sales channel or after a delayed payment fails. Paid orders can't be canceled, so a cancellation never involves a refund. - `order.refunded`: each time a full or partial refund is recorded, through the payment provider or manually. - `order.fulfilled`: a fulfillment leaves every item in the order fulfilled. Partial fulfillments send `fulfillment.created` instead. - `fulfillment.created`: each time items in an order are fulfilled, including partial fulfillments and digital items fulfilled automatically after payment. - `fulfillment.updated`: a fulfillment's tracking details change, or a connected shipping service or sales channel reports progress such as delivery. - `fulfillment.cancelled`: a fulfillment is canceled before it ships, for example when it is canceled or its shipping label is voided. - `product.created`: a product is added to the store, including drafts and products created in bulk. - `product.updated`: a product changes, including its details, variants, prices, stock, visibility or collections. Stock changes from orders also send it, so expect frequent events. - `product.deleted`: a product is deleted, individually or in bulk. (one of: order.placed, order.paid, order.cancelled, order.refunded, order.fulfilled, fulfillment.created, fulfillment.updated, fulfillment.cancelled, product.created, product.updated, product.deleted)
  -h, --help[=true|false]           help for create
      --sales-channel-ids strings   Omit for store-wide delivery. IDs must belong to this store; empty arrays are invalid.
      --url string                  HTTPS receiver on port 443, without credentials or a fragment. Must be a public destination: localhost, single-label and local-network names (.local, .home.arpa, .internal), Hostinger service domains and loopback, private, link-local or reserved IP addresses are rejected. Merchant-owned domains hosted by Hostinger are allowed.
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --domain string           Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger ecommerce webhooks](hostinger_ecommerce_webhooks.md)	 - Webhooks commands

