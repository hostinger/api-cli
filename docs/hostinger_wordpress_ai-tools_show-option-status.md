## hostinger wordpress ai-tools show-option-status

Show AI option status

### Synopsis

Show the current AI option status for the Hostinger Tools plugin on the
specified WordPress installation. Filter by `option` to return a single
option, or omit it to return all options.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress ai-tools show-option-status <username> <software> [flags]
```

### Options

```
  -h, --help            help for show-option-status
      --option string   Filter the status by a single AI option. (one of: llmstxt, web2agent)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress ai-tools](hostinger_wordpress_ai-tools.md)	 - AI Tools commands

