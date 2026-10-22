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

### Examples

```
  hostinger wordpress ai-tools set-option-status u123456789 1232456789 --enable=true --option llmstxt
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
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger wordpress ai-tools](hostinger_wordpress_ai-tools.md)	 - AI Tools commands

