## hostinger hosting git delete-auto-deployment-settings

Delete Git auto-deployment settings

### Synopsis

Removes the Git auto-deployment settings of the website. Files already deployed stay on the
website; pushes stop deploying until settings are saved again. Succeeds also when nothing is
configured.

```
hostinger hosting git delete-auto-deployment-settings <username> <domain> [flags]
```

### Examples

```
  hostinger hosting git delete-auto-deployment-settings u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for delete-auto-deployment-settings
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

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

