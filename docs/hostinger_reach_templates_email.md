## hostinger reach templates email

Get email template

### Synopsis

Get one email template of the profile with its full HTML in `template_content`.

A campaign's email content is its template: to read what a campaign sends, get the campaign
and pass its `template_uuid` here. The HTML can be large, so fetch it only when the content
itself is needed, never to browse or pick templates - the list email templates endpoint does
that without content. Personalization tags such as `{{name}}` come back unrendered.

```
hostinger reach templates email <profile-uuid> <template-uuid> [flags]
```

### Examples

```
  hostinger reach templates email 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000
```

### Options

```
  -h, --help   help for email
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

* [hostinger reach templates](hostinger_reach_templates.md)	 - Templates commands

