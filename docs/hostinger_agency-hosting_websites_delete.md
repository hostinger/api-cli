## hostinger agency-hosting websites delete

Delete website

### Synopsis

Permanently deletes an Agency Plan website. Deletion is processed asynchronously: the
website is immediately transitioned to a deleting state and the underlying server
resources are removed in the background.

```
hostinger agency-hosting websites delete <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting websites delete zpwlGlp19
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

* [hostinger agency-hosting websites](hostinger_agency-hosting_websites.md)	 - Websites commands

