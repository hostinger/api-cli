## hostinger hosting git list-installations

List Git installations

### Synopsis

Lists the Git provider accounts the customer has connected. Only installations with status
`active` are returned unless the `status` filter says otherwise.

An empty list means the customer has no active installation. Check `status=suspended` and
`status=pending` as well. If there is none at all, GitHub has to be connected once in hPanel
(Websites, Manage, Advanced, Git, Connect GitHub; or Add Website, Node.js Web App, Import Git
Repository, Continue with GitHub); this endpoint then lists the new installation.

Use `uuid` as the path parameter of `List Git installation repositories`.

```
hostinger hosting git list-installations [flags]
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
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

