## hostinger hosting php update-options

Update PHP options

### Synopsis

Updates PHP options for the website (e.g. `memory_limit`, `max_execution_time`, `upload_max_filesize`).
Only provide the options you want to change, inside the `options` object.

Values above the account plan limit are silently capped to that limit, so the request can succeed
with a smaller applied value. Call the Get PHP details endpoint afterwards to read the applied value.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --options
    {}

```
hostinger hosting php update-options <username> <domain> [flags]
```

### Examples

```
  hostinger hosting php update-options u123456789 mydomain.tld --options '{}'
```

### Options

```
  -h, --help             help for update-options
      --options string   Map of PHP options to update, keyed by option name. Only include options you want to change. (JSON, @file or -; shape in the help above)
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

