## hostinger agency-hosting wordpress change-version

Change WordPress version

### Synopsis

Changes the installed WordPress core version on an Agency Plan website to one of the versions available for installation.

```
hostinger agency-hosting wordpress change-version <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting wordpress change-version zpwlGlp19 --version 6.5.5
```

### Options

```
  -h, --help             help for change-version
      --version string   Target WordPress core version to install. Must be one of the available versions.
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

* [hostinger agency-hosting wordpress](hostinger_agency-hosting_wordpress.md)	 - WordPress commands

