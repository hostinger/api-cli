## hostinger agent-readiness scans get

Get agent readiness scan

### Synopsis

Get the latest agent readiness scan of a website.\n
Scans started with the `Run agent readiness scan` operation are asynchronous. While the status is
`pending`, the result fields are null; wait a few seconds and call this operation again until the status is
`completed` or `failed`. A `failed` scan is not retried; run the scan again.\n
A completed scan has a 0-100 score, a letter grade, a score per category, and the result of every
check with the evidence found and how to fix it. Check names are listed by the `List checks` operation.\n
Returns 404 when the domain has never been scanned; use the `Run agent readiness scan` operation first.

```
hostinger agent-readiness scans get <domain> [flags]
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agent-readiness scans](hostinger_agent-readiness_scans.md)	 - Scans commands

