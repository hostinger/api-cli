## hostinger hosting git auto-deployment-settings

Get Git auto-deployment settings

### Synopsis

Returns the Git auto-deployment settings of the website: which repository and branch deploy
into which directory, and whether pushes trigger a deployment. `is_enabled` false keeps the
repository link but ignores pushes.

When the website has no auto-deployment configured every field is null. Save settings with
`Update Git auto-deployment settings`.

```
hostinger hosting git auto-deployment-settings <username> <domain> [flags]
```

### Options

```
  -h, --help   help for auto-deployment-settings
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

