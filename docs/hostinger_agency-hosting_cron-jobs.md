## hostinger agency-hosting cron-jobs

Cron Jobs commands

```
hostinger agency-hosting cron-jobs [flags]
```

### Options

```
  -h, --help   help for cron-jobs
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

* [hostinger agency-hosting](hostinger_agency-hosting.md)	 - Agency Hosting commands
* [hostinger agency-hosting cron-jobs create-website](hostinger_agency-hosting_cron-jobs_create-website.md)	 - Create website cron job
* [hostinger agency-hosting cron-jobs delete-website](hostinger_agency-hosting_cron-jobs_delete-website.md)	 - Delete website cron job
* [hostinger agency-hosting cron-jobs list-website](hostinger_agency-hosting_cron-jobs_list-website.md)	 - List website cron jobs

