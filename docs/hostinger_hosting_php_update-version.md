## hostinger hosting php update-version

Update PHP version

### Synopsis

Changes the PHP version of the website.

Use the Get PHP details endpoint to see the versions available for the website.

```
hostinger hosting php update-version <username> <domain> [flags]
```

### Examples

```
  hostinger hosting php update-version u123456789 mydomain.tld --version 8.1
```

### Options

```
  -h, --help             help for update-version
      --version string   PHP version to switch the website to.
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

