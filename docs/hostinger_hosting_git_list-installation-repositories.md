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

### Examples

```
  hostinger hosting git list-installation-repositories 018f5e2a-1234-7890-abcd-1234567890ab
```

### Options

```
  -h, --help   help for list-installation-repositories
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

