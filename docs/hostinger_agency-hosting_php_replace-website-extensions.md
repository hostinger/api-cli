## hostinger agency-hosting php replace-website-extensions

Replace website PHP extensions

### Synopsis

Replaces the set of PHP extensions enabled on an Agency Plan website with the ones provided. Any toggleable extension not in the request is disabled, so call the extensions endpoint first and send the full desired set. Extensions compiled into PHP, reported with the "built-in" state, are always active and are unaffected.

```
hostinger agency-hosting php replace-website-extensions <website_uid> [flags]
```

### Options

```
      --extensions strings   Extension names, exactly as returned by the extensions endpoint.
  -h, --help                 help for replace-website-extensions
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting php](hostinger_agency-hosting_php.md)	 - PHP commands

