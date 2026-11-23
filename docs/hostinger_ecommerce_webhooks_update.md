## hostinger ecommerce webhooks update

Update an ecommerce webhook

### Synopsis

Change the receiver, events, channel filter or enabled state. Omitted fields are preserved; a null
channel filter restores store-wide delivery. Endpoints disabled automatically after a 410 Gone
(receiver_gone) or persistent delivery failures (receiver_failing) resume with enabled: true; events
from the disabled period are not sent. Lifecycle-revoked endpoints cannot be enabled again: delete
them and create fresh subscriptions after recovery.

400 codes: webhook_destination_not_allowed (the new URL, or the saved URL when enabling, is not a public
destination), webhook_sales_channel_invalid. 409 codes: webhook_endpoint_revoked, webhook_store_inactive,
webhook_endpoint_catalog_busy (retry). Errors return the upstream code in `message`, as
`[Ecommerce:<code>] <reason>`.

```
hostinger ecommerce webhooks update <store_id> <endpoint_id> [flags]
```

### Examples

```
  hostinger ecommerce webhooks update store_01J8Z5F8W9K8M4A7B3C2D1E0FG whep_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
      --enabled[=true|false]        Merchant-disabled endpoints can be enabled again; lifecycle-revoked endpoints cannot.
      --events order.placed         Replacement event types. Omit to preserve. - order.placed: a customer places an order, before any online payment is completed. Orders paid online may not have an order number yet. - `order.paid`: an order becomes fully paid: an online payment is captured, a free order is placed or the merchant marks it as paid. Partial payments don't send it. - `order.cancelled`: an unpaid order is canceled, for example by the merchant, a connected sales channel or after a delayed payment fails. Paid orders can't be canceled, so a cancellation never involves a refund. - `order.refunded`: each time a full or partial refund is recorded, through the payment provider or manually. - `order.fulfilled`: a fulfillment leaves every item in the order fulfilled. Partial fulfillments send `fulfillment.created` instead. - `fulfillment.created`: each time items in an order are fulfilled, including partial fulfillments and digital items fulfilled automatically after payment. - `fulfillment.updated`: a fulfillment's tracking details change, or a connected shipping service or sales channel reports progress such as delivery. - `fulfillment.cancelled`: a fulfillment is canceled before it ships, for example when it is canceled or its shipping label is voided. - `product.created`: a product is added to the store, including drafts and products created in bulk. - `product.updated`: a product changes, including its details, variants, prices, stock, visibility or collections. Stock changes from orders also send it, so expect frequent events. - `product.deleted`: a product is deleted, individually or in bulk. (one of: order.placed, order.paid, order.cancelled, order.refunded, order.fulfilled, fulfillment.created, fulfillment.updated, fulfillment.cancelled, product.created, product.updated, product.deleted)
  -h, --help[=true|false]           help for update
      --sales-channel-ids strings   Omit to preserve; null restores store-wide delivery. Empty or foreign-store IDs are invalid.
      --url string                  HTTPS receiver on port 443, without credentials or a fragment. Omit to preserve. A changed URL must be a public destination under the same rules as creation; resending the saved URL unchanged is accepted.
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

