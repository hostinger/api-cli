## hostinger hosting nodejs update-build-settings

Update Node.js build settings

### Synopsis

Replaces the build settings stored for the website. Send the full set: `node_version` is
required and every nullable field you omit is stored as null. Creates the settings when
none exist yet.

This does not start a build. Stored settings drive Git auto-deployment builds; a build
started through the API uses the values sent in that request, so to rebuild with corrected
settings call `Start Node.js build` with the same values. Typical fixes: a wrong `app_type`
after auto-detection, or a missing `entry_file` for express, fastify, nest, nuxt and hono
apps.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting nodejs update-build-settings [domain] [flags]
```

### Examples

```
  hostinger hosting nodejs update-build-settings mydomain.tld --node-version 20
```

### Options

```
      --app-type string           Node.js application framework. Set it explicitly when auto-detection picked the wrong one. (one of: create-react-app, gatsby, vite, angular, react, vue, parcel, next, nuxt, nest, express, fastify, astro, svelte, svelte-kit, hono, react-router, nitro, other)
      --build-script string       The package.json script that builds the application
      --entry-file string         The main entry point file for the application (required for express, fastify, nest, nuxt and hono app types)
  -h, --help                      help for update-build-settings
      --node-version int          Node.js major version (one of: 18, 20, 22, 24)
      --output-directory string   Build output directory relative to the root directory
      --package-manager string    Package manager used to install dependencies (one of: npm, yarn, pnpm)
      --root-directory string     Application root directory (where package.json is located) relative to public_html. Omit it, or send ".", for public_html itself.
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

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

