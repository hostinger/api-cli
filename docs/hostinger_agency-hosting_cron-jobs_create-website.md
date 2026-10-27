## hostinger agency-hosting cron-jobs create-website

Create website cron job

### Synopsis

Creates a cron job for an Agency Plan website from a schedule expression and a command.

Returns the created cron job, including its uuid, which is required to delete the cron job.

```
hostinger agency-hosting cron-jobs create-website <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting cron-jobs create-website zpwlGlp19 --command 'php artisan schedule:run' --time '*/30 * * * *'
```

### Options

```
      --command string   Command to run on the schedule. Must not contain pipe (|) or redirection (<, >) characters.
  -h, --help             help for create-website
      --time string      Cron schedule expression (standard 5-field crontab syntax).
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

