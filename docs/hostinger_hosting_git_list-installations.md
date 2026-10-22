## hostinger hosting git list-installations

List Git installations

### Synopsis

Lists the Git provider accounts the customer has connected. Only installations with status
`active` are returned unless the `status` filter says otherwise.

An empty list means the customer has no active installation. Check `status=suspended` and
`status=pending` as well. If there is none at all, a Git provider (GitHub or GitLab) has to be
connected once in hPanel (Websites, Manage, Advanced, Git; or Add Website, Node.js Web App,
Import Git Repository); this endpoint then lists the new installation.

Use `uuid` as the path parameter of `List Git installation repositories`, and as
`installation_uuid` in `Start Node.js build` with `source_type` `git` and in
`Update Git auto-deployment settings`.

```
hostinger hosting git list-installations [flags]
```

### Examples

```
  hostinger hosting git list-installations
```

### Options

```
  -h, --help              help for list-installations
      --provider string   Filter by Git provider (one of: github, gitlab, bitbucket)
      --status string     Filter by installation status (one of: pending, active, suspended) (default "active")
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

