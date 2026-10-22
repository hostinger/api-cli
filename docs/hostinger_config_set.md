## hostinger config set

Set a setting in the config file

### Synopsis

Set a setting in the config file. For api_token the value may be left out:
it is then read from a hidden prompt so it stays out of the shell history.

```
hostinger config set <key> [value] [flags]
```

### Examples

```
  config set format json
  config set api_token
```

### Options

```
  -h, --help   help for set
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

* [hostinger config](hostinger_config.md)	 - Read and write CLI settings

