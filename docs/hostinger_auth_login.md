## hostinger auth login

Sign in through the browser

### Synopsis

Sign in through the browser (OAuth). The tokens are stored locally and
refreshed automatically; commands also sign in on first use when no API token
is configured.

```
hostinger auth login [flags]
```

### Options

```
  -h, --help   help for login
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

* [hostinger auth](hostinger_auth.md)	 - Sign in and check which credentials are used

