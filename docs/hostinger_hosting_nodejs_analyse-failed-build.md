## hostinger hosting nodejs analyse-failed-build

Analyse failed Node.js build

### Synopsis

Returns an AI analysis of why a build failed and how to fix it, based on the build logs,
the project file list and package.json. Only builds in the `failed` state can be analysed;
any other state returns 422. When no analysis could be produced both `analysis` and
`solution` are null, in which case read `Get NodeJS build logs` instead.

Each call runs the analysis again, so call it once per failed build and keep the result.
Limited to 5 calls per minute per API client (429 above that).

```
hostinger hosting nodejs analyse-failed-build <username> <domain> <uuid> [flags]
```

### Examples

```
  hostinger hosting nodejs analyse-failed-build u123456789 mydomain.tld 123e4567-e89b-12d3-a456-426614174000
```

### Options

```
  -h, --help   help for analyse-failed-build
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

