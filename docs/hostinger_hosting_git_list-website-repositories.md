## hostinger hosting git list-website-repositories

List website Git repositories

### Synopsis

Lists the Git repositories linked to directories of the website, with
`Deploy website Git repository` or in the Git section of hPanel: clone URL, branch and directory of
each one. A repository whose clone failed stays listed; deploying it again retries the clone. GitHub
and GitLab auto-deployments are not listed here; see `Get Git auto-deployment settings`.

```
hostinger hosting git list-website-repositories <username> <domain> [flags]
```

### Examples

```
  hostinger hosting git list-website-repositories u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for list-website-repositories
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

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

