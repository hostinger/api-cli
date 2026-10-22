## hostinger config

Read and write CLI settings

```
hostinger config [flags]
```

### Options

```
  -h, --help   help for config
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

* [hostinger](hostinger.md)	 - Hostinger API Command Line Interface
* [hostinger config get](hostinger_config_get.md)	 - Print the effective value of a setting
* [hostinger config list](hostinger_config_list.md)	 - List the settings and their effective values
* [hostinger config path](hostinger_config_path.md)	 - Print the config file path
* [hostinger config set](hostinger_config_set.md)	 - Set a setting in the config file
* [hostinger config unset](hostinger_config_unset.md)	 - Remove a setting from the config file

