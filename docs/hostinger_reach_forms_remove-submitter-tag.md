## hostinger reach forms remove-submitter-tag

Remove form submitter tag

### Synopsis

Stop a form from applying a tag to its future submitters.

Succeeds also when the form did not have the tag. Contacts who already received the tag keep it, and
the tag itself is not deleted. The automatic `form:{name}` tag cannot be removed and is rejected with
a conflict.

```
hostinger reach forms remove-submitter-tag <profile-uuid> <form-uuid> [flags]
```

### Examples

```
  hostinger reach forms remove-submitter-tag 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000 --tag-uuid 550e8400-e29b-41d4-a716-446655440000
```

### Options

```
  -h, --help              help for remove-submitter-tag
      --tag-uuid string   
  -y, --yes               Skip the confirmation prompt
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

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

