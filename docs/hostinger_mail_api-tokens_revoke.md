## hostinger mail api-tokens revoke

Revoke API token

### Synopsis

Revoke an API token. The token immediately loses access to the
Hostinger Email API (https://api.mail.hostinger.com/). This action
cannot be undone.

```
hostinger mail api-tokens revoke <token-id> [flags]
```

### Examples

```
  hostinger mail api-tokens revoke 019683f8-1234-7abc-8def-0123456789ab
```

### Options

```
  -h, --help   help for revoke
  -y, --yes    Skip the confirmation prompt
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

