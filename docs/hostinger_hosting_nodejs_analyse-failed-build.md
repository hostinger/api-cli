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

### Options

```
  -h, --help   help for analyse-failed-build
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

