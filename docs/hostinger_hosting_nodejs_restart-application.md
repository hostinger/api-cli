## hostinger hosting nodejs restart-application

Restart Node.js application

### Synopsis

Restarts the Node.js server process for the website. Does not rebuild or redeploy the
application. Use it to apply environment or configuration changes, or to recover a hung
application.

Only applicable to server-side applications (Express, Next.js, NestJS, etc.). Static
front-end apps (React, Vue, Vite) have no persistent server process, so restarting them
has no effect. Returns success even when the website has no server process to restart.

```
hostinger hosting nodejs restart-application <username> <domain> [flags]
```

### Examples

```
  hostinger hosting nodejs restart-application u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for restart-application
  -y, --yes    Skip the confirmation prompt
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

