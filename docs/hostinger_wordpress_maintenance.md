## hostinger wordpress maintenance

Maintenance commands

```
hostinger wordpress maintenance [flags]
```

### Options

```
  -h, --help   help for maintenance
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

* [hostinger wordpress](hostinger_wordpress.md)	 - WordPress commands
* [hostinger wordpress maintenance show-status](hostinger_wordpress_maintenance_show-status.md)	 - Show maintenance status
* [hostinger wordpress maintenance toggle](hostinger_wordpress_maintenance_toggle.md)	 - Toggle maintenance mode

