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

### Options

```
  -h, --help   help for restart-application
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

