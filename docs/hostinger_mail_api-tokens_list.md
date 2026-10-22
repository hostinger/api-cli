## hostinger mail api-tokens list

List API tokens

### Synopsis

Retrieve a paginated list of
Hostinger Email API (https://api.mail.hostinger.com/) tokens across
all your mail orders, optionally filtered by order. Plaintext tokens
are never included; they are returned only when a token is created.

```
hostinger mail api-tokens list [flags]
```

### Examples

```
  hostinger mail api-tokens list
```

### Options

```
      --all               Fetch every page and print all items
  -h, --help              help for list
      --order-id string   Filter tokens by order resource ID. Single value or comma-separated list.
      --page int          Page number
      --per-page int      Number of items per page (default 25)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger mail api-tokens](hostinger_mail_api-tokens.md)	 - API Tokens commands

