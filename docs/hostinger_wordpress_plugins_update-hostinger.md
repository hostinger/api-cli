## hostinger wordpress plugins update-hostinger

Update Hostinger WordPress plugin

### Synopsis

Update a Hostinger plugin to its latest version on a WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the update job
has been queued.

```
hostinger wordpress plugins update-hostinger <username> <software> [flags]
```

### Options

```
  -h, --help          help for update-hostinger
      --slug string   Slug of the Hostinger plugin to update to its latest version. (one of: hostinger, hostinger-ai-assistant, hostinger-affiliate-plugin, hostinger-easy-onboarding, hostinger-reach)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

