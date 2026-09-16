## hostinger hosting git list-installation-repositories

List Git installation repositories

### Synopsis

Lists the repositories the Git installation can access, read live from the provider. Works
for github and gitlab installations. Use an active installation: a suspended or pending one
is still queried and the call fails with whatever the provider answers. The list is cut at
the first 500 repositories in the order the provider returns them; when the account has
more, name the repository directly instead of searching this list.

`owner`, `name` and a branch (`default_branch` or another one) go into `source_options` of
`Start Node.js build` or into `Update Git auto-deployment settings`. Returns 404 when the
installation does not belong to the customer. Limited to 10 calls per minute per API client
(429 above that).

```
hostinger hosting git list-installation-repositories <uuid> [flags]
```

### Options

```
  -h, --help   help for list-installation-repositories
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

