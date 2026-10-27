## hostinger hosting nodejs replace-environment-variables

Replace Node.js environment variables

### Synopsis

Replaces the website's Node.js environment variables with the ones provided. This is a
full replace: any variable not in the request is deleted, and sending an empty `env_vars`
array deletes every variable. Saving writes the values and restarts the running Node.js
process.

A restart is enough for apps that read environment variables at process start, such as
Express or NestJS. It is not enough for frameworks that bake variables into the build.
Next.js standalone is one of those: build-time values (including `NEXT_PUBLIC_*`) need a
fresh build. After this call, use the `Start Node.js build` endpoint so those apps
pick up the new values.

The `List Node.js environment variables` endpoint returns masked values (`********`), so
never copy values from it into this request. Always send the full desired set with real
values taken from the project `.env` file or the user prompt.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --env-vars
    [
      {
        "key": "API_URL",
        "value": "https://api.example.com"
      }
    ]

```
hostinger hosting nodejs replace-environment-variables [domain] [flags]
```

### Examples

```
  hostinger hosting nodejs replace-environment-variables mydomain.tld --env-vars '[{"key":"API_URL","value":"https://api.example.com"}]'
```

### Options

```
      --env-vars string   Environment variables to set. This is the full desired set: any variable not in this list is deleted, and an empty array deletes every variable. (JSON, @file or -; shape in the help above)
  -h, --help              help for replace-environment-variables
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

