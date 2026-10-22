## hostinger hosting php get

Get PHP details

### Synopsis

Returns the full PHP configuration for the website: current version, available versions
(supported and unsupported), enabled/disabled extensions, options with their current value,
default, type and the plan limit (`max`), and conflicting extension groups.

Use it to check the current PHP setup before updating the version, extensions or options.

```
hostinger hosting php get <username> <domain> [flags]
```

### Examples

```
  hostinger hosting php get u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for get
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

