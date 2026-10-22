## hostinger

Hostinger API Command Line Interface

### Synopsis

Hostinger API Command Line Interface.

Get started:
  1. Sign in:        hostinger auth login   (or: hostinger config set api_token, or set HOSTINGER_API_TOKEN)
  2. Try a command:  hostinger agency-hosting domains list

Output is a table on a terminal and JSON when piped; pick one with --format
(table|json|yaml|tree) or filter it with --jq. Missing arguments and required
flags are prompted for in a terminal (disable with --no-input). Commands that
delete or stop things ask for confirmation (skip with --yes), and --dry-run
prints the request as curl without sending it.

Documentation: https://github.com/hostinger/api-cli/tree/main/docs

### Options

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
  -h, --help            help for hostinger
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger agency-hosting](hostinger_agency-hosting.md)	 - Agency Hosting commands
* [hostinger agent-readiness](hostinger_agent-readiness.md)	 - Agent Readiness commands
* [hostinger auth](hostinger_auth.md)	 - Sign in and check which credentials are used
* [hostinger billing](hostinger_billing.md)	 - Billing commands
* [hostinger completion](hostinger_completion.md)	 - Generate the autocompletion script for the specified shell
* [hostinger config](hostinger_config.md)	 - Read and write CLI settings
* [hostinger dns](hostinger_dns.md)	 - DNS commands
* [hostinger domains](hostinger_domains.md)	 - Domains commands
* [hostinger ecommerce](hostinger_ecommerce.md)	 - Ecommerce commands
* [hostinger horizons](hostinger_horizons.md)	 - Horizons commands
* [hostinger hosting](hostinger_hosting.md)	 - Hosting commands
* [hostinger mail](hostinger_mail.md)	 - Mail commands
* [hostinger reach](hostinger_reach.md)	 - Reach commands
* [hostinger version](hostinger_version.md)	 - Show the current version
* [hostinger vps](hostinger_vps.md)	 - VPS commands
* [hostinger wordpress](hostinger_wordpress.md)	 - WordPress commands

