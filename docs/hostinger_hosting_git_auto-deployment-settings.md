## hostinger hosting git auto-deployment-settings

Get Git auto-deployment settings

### Synopsis

Returns the Git auto-deployment settings of the website: which repository and branch deploy
into which directory, and whether pushes trigger a deployment. `is_enabled` false keeps the
repository link but ignores pushes.

When the website has no auto-deployment configured every field is null. Save settings with
`Update Git auto-deployment settings`.

```
hostinger hosting git auto-deployment-settings <username> <domain> [flags]
```

### Examples

```
  hostinger hosting git auto-deployment-settings u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for auto-deployment-settings
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

