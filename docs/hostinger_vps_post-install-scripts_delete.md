## hostinger vps post-install-scripts delete

Delete post-install script

### Synopsis

Delete a post-install script from your account.

Use this endpoint to remove unused automation scripts.

```
hostinger vps post-install-scripts delete <post-install-script-id> [flags]
```

### Examples

```
  hostinger vps post-install-scripts delete 9568314
```

### Options

```
  -h, --help   help for delete
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

* [hostinger vps post-install-scripts](hostinger_vps_post-install-scripts.md)	 - Post-install scripts commands

