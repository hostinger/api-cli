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

### Options

```
  -h, --help   help for list-website-repositories
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

