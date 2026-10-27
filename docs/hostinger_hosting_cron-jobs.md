## hostinger hosting cron-jobs

Cron Jobs commands

```
hostinger hosting cron-jobs [flags]
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

* [hostinger hosting](hostinger_hosting.md)	 - Hosting commands
* [hostinger hosting cron-jobs create](hostinger_hosting_cron-jobs_create.md)	 - Create account cron job
* [hostinger hosting cron-jobs delete](hostinger_hosting_cron-jobs_delete.md)	 - Delete account cron job
* [hostinger hosting cron-jobs list](hostinger_hosting_cron-jobs_list.md)	 - List account cron jobs
* [hostinger hosting cron-jobs output](hostinger_hosting_cron-jobs_output.md)	 - Get cron job output

