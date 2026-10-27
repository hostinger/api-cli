## hostinger hosting git list-website-repositories

List website Git repositories

### Synopsis

Lists the Git repositories linked to directories of the website, with
`Deploy website Git repository` or in the Git section of hPanel: clone URL, branch and directory of
each one. A repository whose clone failed stays listed; deploying it again retries the clone. GitHub
and GitLab auto-deployments are not listed here; see `Get Git auto-deployment settings`.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting git list-website-repositories [domain] [flags]
```

### Examples

```
  hostinger hosting git list-website-repositories mydomain.tld
```

### Options

```
  -h, --help   help for list-website-repositories
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

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

