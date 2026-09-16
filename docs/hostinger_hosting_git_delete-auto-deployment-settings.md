## hostinger hosting git delete-auto-deployment-settings

Delete Git auto-deployment settings

### Synopsis

Removes the Git auto-deployment settings of the website. Files already deployed stay on the
website; pushes stop deploying until settings are saved again. Succeeds also when nothing is
configured.

```
hostinger hosting git delete-auto-deployment-settings <username> <domain> [flags]
```

### Options

```
  -h, --help   help for delete-auto-deployment-settings
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

