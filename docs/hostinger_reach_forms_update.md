## hostinger reach forms update

Update form

### Synopsis

Rename a form and/or change its status.

Setting a `draft` or `paused` form to `active` takes it live: its latest saved design is published
and visitors can submit it. `paused` and `draft` stop the form from accepting submissions. A form is
not reachable by visitors until it is taken live. Design changes saved while the form is already
`active` stay unpublished until the form is paused and set to `active` again. Renaming does not rename
the automatic `form:{name}` tag.

```
hostinger reach forms update <profile-uuid> <form-uuid> [flags]
```

### Examples

```
  hostinger reach forms update 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help            help for update
      --name string     Form name. Letters, digits, spaces and - . ' _ only.
      --status active   active takes the form live, `paused` and `draft` stop it accepting submissions. (one of: active, paused, draft)
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

