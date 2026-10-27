## hostinger hosting databases setup-website

Setup website database

### Synopsis

Creates a new MySQL database for the website and writes its connection details into the
website's environment variables, then restarts the application. The platform generates the
password (and the database name and user, unless supplied). The password is never returned;
the application reads it from the environment.

Written variables: `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD` and
`DATABASE_URL` (`mysql://user:password@host:port/name`, user and password percent-encoded).
Existing variables are kept. If the website already has any variable with one of these
names the call fails with 422 and nothing is created; the `Replace Node.js environment
variables` endpoint removes them.

After this call the variables are ordinary environment variables: the
`Replace Node.js environment variables` endpoint changes or removes them like any other.

A restart is enough for apps that read environment variables at process start, such as
Express or NestJS. Frameworks that bake variables into the build output (Next.js,
`NEXT_PUBLIC_*`) see the new values only after a fresh build (`Start Node.js build` endpoint).

A password in the request is ignored; the platform always generates it. The optional `name`
and `user` are identifiers, not secrets.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting databases setup-website [domain] [flags]
```

### Examples

```
  hostinger hosting databases setup-website mydomain.tld
```

### Options

```
  -h, --help               help for setup-website
      --name u123456789_   Optional database name. Generated when omitted. Letters, digits and underscores; must not start with an underscore. Up to 14 characters without the account username prefix (u123456789_), which is added automatically when missing. With the prefix the full name is 12 to 25 characters.
      --user u123456789_   Optional database user. Generated when omitted. Letters, digits and underscores; must not start with an underscore. Up to 14 characters without the account username prefix (u123456789_), which is added automatically when missing. With the prefix the full user is 12 to 25 characters.
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

