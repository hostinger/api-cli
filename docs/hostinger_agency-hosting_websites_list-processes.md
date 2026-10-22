## hostinger agency-hosting websites list-processes

List website processes

### Synopsis

Lists active and recently completed asynchronous processes for an Agency Plan website.

Each process has a unique ID (for tracking), a type, and a status (running, completed, failed).
Poll this endpoint after initiating async operations (SSL setup, backups, cloning) to track progress.

```
hostinger agency-hosting websites list-processes <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting websites list-processes zpwlGlp19
```

### Options

```
  -h, --help   help for list-processes
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

