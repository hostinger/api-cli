## hostinger hosting websites list-setups

List website setups

### Synopsis

Returns the website setups started in the last 24 hours for the hosting accounts
accessible to the authenticated client, newest first.

Meant for polling right after creating a website: the website shows up in the
websites list before its server-side setup has finished, and while the setup is
`running` endpoints that operate on that website may respond with `404` or `409`.
Poll this endpoint with the `domain` filter every 10 to 15 seconds and wait for
`status: completed` before uploading files, deploying or creating databases.
`failed` means the setup stopped before finishing or has not reported progress for
over an hour. Setups older than 24 hours are not listed.

```
hostinger hosting websites list-setups [flags]
```

### Options

```
      --domain string   Filter by domain name (exact match)
  -h, --help            help for list-setups
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

