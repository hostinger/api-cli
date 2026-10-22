## hostinger agency-hosting php replace-website-extensions

Replace website PHP extensions

### Synopsis

Replaces the set of PHP extensions enabled on an Agency Plan website with the ones provided. Any toggleable extension not in the request is disabled, so call the extensions endpoint first and send the full desired set. Extensions compiled into PHP, reported with the "built-in" state, are always active and are unaffected.

```
hostinger agency-hosting php replace-website-extensions <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting php replace-website-extensions zpwlGlp19 --extensions bz2,zip
```

### Options

```
      --extensions strings   Extension names, exactly as returned by the extensions endpoint.
  -h, --help                 help for replace-website-extensions
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

* [hostinger agency-hosting php](hostinger_agency-hosting_php.md)	 - PHP commands

