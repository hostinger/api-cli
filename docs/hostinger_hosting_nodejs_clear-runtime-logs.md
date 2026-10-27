## hostinger hosting nodejs clear-runtime-logs

Clear Node.js runtime logs

### Synopsis

Empties the Node.js application's runtime log file. This cannot be undone, so confirm with
the user before calling it. Returns success even when no log file exists yet.

Use it before reproducing a problem so the next `Get Node.js runtime logs` call returns
only fresh entries; start that call with `period` again instead of reusing a `from_line`
from before the clear.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting nodejs clear-runtime-logs [domain] [flags]
```

### Examples

```
  hostinger hosting nodejs clear-runtime-logs mydomain.tld
```

### Options

```
  -h, --help   help for clear-runtime-logs
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

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

