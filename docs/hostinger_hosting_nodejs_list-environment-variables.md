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

### Options

```
  -h, --help   help for list-environment-variables
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

