## hostinger hosting websites list-setups

List website setups

### Synopsis

Returns the website setups started in the last 24 hours for the hosting accounts
accessible to the authenticated client, newest first. Narrow the list with the
`order_id`, `subscription_id` or `domain` filters.

Meant for polling right after creating a website or starting a website setup: the
website shows up in the websites list before its server-side setup has finished, and
while the setup is `running` endpoints that operate on that website may respond with
`404` or `409`. Poll this endpoint with the `domain` filter every 10 to 15 seconds and
wait for `status: completed` before uploading files, deploying or creating databases.
`failed` means the setup stopped before finishing or has not reported progress for
over an hour. Setups older than 24 hours are not listed.

`type` is the website type the setup was started with (`wordpress`, `headless_empty`,
`headless_wordpress`, `headless_ecommerce`, `headless_pocketbase`), or `null` for an empty
website.

```
hostinger hosting websites list-setups [flags]
```

### Examples

```
  hostinger hosting websites list-setups
```

### Options

```
      --domain string            Filter by domain name (exact match)
  -h, --help                     help for list-setups
      --order-id int             Order ID
      --subscription-id string   Filter by hosting order subscription ID
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

