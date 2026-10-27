## hostinger hosting git update-auto-deployment-settings

Update Git auto-deployment settings

### Synopsis

Creates or replaces the Git auto-deployment settings of the website: repository, branch, the
directory under the document root to deploy into, and `is_enabled`. Send the full set;
`is_enabled` defaults to true and `directory` to the document root. `installation_uuid` must
be an installation from `List Git installations` that belongs to the same customer as the
website.

For PHP and static websites, saving with `is_enabled` true deploys the branch right away and
every later push to that branch deploys again. For Node.js and Website Builder websites saving
does not clone anything. On a Node.js website start the first deploy with
`Start Node.js build` using `source_type` `git`; pushes then trigger new builds with the build
settings stored for the website.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting git update-auto-deployment-settings [domain] [flags]
```

### Examples

```
  hostinger hosting git update-auto-deployment-settings mydomain.tld --branch main --installation-uuid 018f5e2a-1234-7890-abcd-1234567890ab --owner octocat --repository my-site
```

### Options

```
      --branch string                              Branch to deploy
      --directory string                           Subdirectory under the website document root to deploy into. Empty, null or omitted means the document root.
  -h, --help                                       help for update-auto-deployment-settings
      --installation-uuid List Git installations   Active Git installation from List Git installations
      --is-enabled                                 Whether pushes to the branch deploy automatically (default true)
      --owner List Git installation repositories   Repository owner login, as returned by List Git installation repositories. GitLab group paths use slashes.
      --repository string                          Repository name without the .git suffix
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

