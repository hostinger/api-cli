## hostinger hosting nodejs build-settings-from-repository

Get Node.js build settings from repository

### Synopsis

Determine the build settings for a Node.js application from the package.json file of a branch
in a Git repository. Use before starting a Git-source build to find out the auto-detected
settings instead of guessing them.

The repository is read through a Git installation: take `installation_uuid` from
`List Git installations` and `owner` and `repository` from `List Git installation repositories`.

Returns 422 when the installation is not found, the repository or branch is not found, or
the branch has no package.json.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting nodejs build-settings-from-repository [domain] [flags]
```

### Examples

```
  hostinger hosting nodejs build-settings-from-repository mydomain.tld --owner value --repository value --branch value --installation-uuid 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d
```

### Options

```
      --branch string                              Branch whose package.json is inspected
  -h, --help                                       help for build-settings-from-repository
      --installation-uuid List Git installations   Git installation used to access the repository, taken from List Git installations
      --owner string                               Repository owner login. GitLab group paths use slashes.
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

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

