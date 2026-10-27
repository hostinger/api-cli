## hostinger hosting cron-jobs create

Create account cron job

### Synopsis

Creates a cron job for the specified account from a schedule expression and a command.

Returns the created cron job, including its uid, which is required to delete the cron job or fetch its output.

The hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting cron-jobs create [username] [flags]
```

### Examples

```
  hostinger hosting cron-jobs create u123456789 --command 'php /home/u123456789/cleanup.php' --time '0 2 * * *'
```

### Options

```
      --command string   Command to execute on the schedule.
  -h, --help             help for create
      --time string      Cron schedule expression (for example "0 2 * * *" runs daily at 02:00).
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

* [hostinger hosting cron-jobs](hostinger_hosting_cron-jobs.md)	 - Cron Jobs commands

