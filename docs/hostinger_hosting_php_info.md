## hostinger hosting php info

Get PHP info

### Synopsis

Returns the full phpinfo page (HTML) for the website.

Use it to debug PHP issues or inspect the complete PHP environment of the website.

```
hostinger hosting php info <username> <domain> [flags]
```

### Examples

```
  hostinger hosting php info u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for info
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

* [hostinger hosting php](hostinger_hosting_php.md)	 - PHP commands

