## hostinger agency-hosting cron-jobs delete-website

Delete website cron job

### Synopsis

Permanently deletes the cron job identified by its uuid from an Agency Plan website.

The operation is idempotent: deleting a cron job that does not exist succeeds without error.

```
hostinger agency-hosting cron-jobs delete-website <website_uid> <uuid> [flags]
```

### Examples

```
  hostinger agency-hosting cron-jobs delete-website zpwlGlp19 01931d6f-68f5-7b72-8d9e-09c6e1e6aa0e
```

### Options

```
  -h, --help   help for delete-website
  -y, --yes    Skip the confirmation prompt
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

* [hostinger agency-hosting cron-jobs](hostinger_agency-hosting_cron-jobs.md)	 - Cron Jobs commands

