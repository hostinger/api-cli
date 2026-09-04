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

### Options

```
  -h, --help   help for build
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

