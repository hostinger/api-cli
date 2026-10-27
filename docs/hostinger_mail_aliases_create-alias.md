## hostinger mail aliases create-alias

Create alias

### Synopsis

Create an alias for the given mailbox. The alias address is formed
from the given local part and the domain of the mailbox. Messages
sent to the alias are delivered to the mailbox.

```
hostinger mail aliases create-alias <mailbox-id> [flags]
```

### Examples

```
  hostinger mail aliases create-alias AC1a2b3c4d5e6f7g --local-part info
```

### Options

```
  -h, --help                help for create-alias
      --local-part string   Local part of the alias address (the part before the @). The domain is taken from the mailbox. Case-insensitive and stored lowercase; must start and end with a letter or digit; single dots, underscores and hyphens are allowed in between.
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

* [hostinger mail aliases](hostinger_mail_aliases.md)	 - Aliases commands

