## hostinger agency-hosting websites list-running-plan-processes

List running Agency Plan website processes

### Synopsis

Lists active and recently completed asynchronous processes for an Agency Plan website.

Each process has a unique ID (for tracking), a type, and a status (running, completed, failed).
Poll this endpoint after initiating async operations (SSL setup, backups, cloning) to track progress.

```
hostinger agency-hosting websites list-running-plan-processes <website_uid> [flags]
```

### Options

```
  -h, --help   help for list-running-plan-processes
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting websites](hostinger_agency-hosting_websites.md)	 - Websites commands

