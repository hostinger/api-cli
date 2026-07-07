## hostinger wordpress ai-tools set-option-status

Set AI option status

### Synopsis

Enable or disable an AI option for the Hostinger Tools plugin on the specified
WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress ai-tools set-option-status <username> <software> [flags]
```

### Options

```
      --enable          Enable (true) or disable (false) the AI option.
  -h, --help            help for set-option-status
      --option string   AI option name (one of: llmstxt, web2agent)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress ai-tools](hostinger_wordpress_ai-tools.md)	 - AI Tools commands

