## hostinger hosting nodejs list-environment-variables

List Node.js environment variables

### Synopsis

Lists the Node.js environment variables currently set for the website. Values are always
masked as `********` and cannot be read back through this API. Use this endpoint to see
which keys are configured or to verify a change, not to read values.

To change variables, use the `Replace Node.js environment variables` endpoint. It replaces
the whole set, so never copy the masked values from this response into that request; send
the full desired set with real values taken from the project `.env` file or the user
prompt instead.

```
hostinger hosting nodejs list-environment-variables <username> <domain> [flags]
```

### Examples

```
  hostinger hosting nodejs list-environment-variables u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for list-environment-variables
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

