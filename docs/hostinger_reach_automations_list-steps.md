## hostinger reach automations list-steps

List automation steps

### Synopsis

Get the workflow of an automation as a flat list of steps.

The steps form a tree rather than a straight line: follow `parent_uuid` to reconstruct the
branches, and use `step_order` to order the steps that share a parent. An automation with no
steps yet returns an empty list.

```
hostinger reach automations list-steps <profile-uuid> <automation-uuid> [flags]
```

### Examples

```
  hostinger reach automations list-steps 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for list-steps
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach automations](hostinger_reach_automations.md)	 - Automations commands

