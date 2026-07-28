## hostinger mail api-tokens list

List API tokens

### Synopsis

Retrieve a paginated list of
[Hostinger Email API](https://api.mail.hostinger.com/) tokens across
all your mail orders, optionally filtered by order. Plaintext tokens
are never included; they are returned only when a token is created.

```
hostinger mail api-tokens list [flags]
```

### Options

```
  -h, --help              help for list
      --order-id string   Filter tokens by order resource ID. Single value or comma-separated list.
      --page int          Page number
      --per-page int      Number of items per page (default 25)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail api-tokens](hostinger_mail_api-tokens.md)	 - API Tokens commands

