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

```
hostinger hosting git update-auto-deployment-settings <username> <domain> [flags]
```

### Options

```
      --branch string                              Branch to deploy
      --directory string                           Subdirectory under the website document root to deploy into. Empty, null or omitted means
                                                   the document root.
  -h, --help                                       help for update-auto-deployment-settings
      --installation-uuid List Git installations   Active Git installation from List Git installations
      --is-enabled                                 Whether pushes to the branch deploy automatically (default true)
      --owner List Git installation repositories   Repository owner login, as returned by List Git installation repositories. GitLab group
                                                   paths use slashes.
      --repository string                          Repository name without the .git suffix
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

