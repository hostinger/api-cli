## hostinger hosting nodejs build

Get Node.js build details

### Synopsis

Returns one build by UUID: its state (`pending`, `running`, `completed`, `failed`), the
options it ran with and timestamps. Poll this while a build is pending or running. When it
is failed, read `Get NodeJS build logs` and `Analyse failed Node.js build` for the cause.
Returns 404 when the UUID does not belong to a build of this website.

```
hostinger hosting nodejs build <username> <domain> <uuid> [flags]
```

### Examples

```
  hostinger hosting nodejs build u123456789 mydomain.tld 123e4567-e89b-12d3-a456-426614174000
```

### Options

```
  -h, --help   help for build
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

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

