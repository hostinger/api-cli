## hostinger reach contact-fields

Contact Fields commands

```
hostinger reach contact-fields [flags]
```

### Options

```
  -h, --help   help for contact-fields
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

* [hostinger reach](hostinger_reach.md)	 - Reach commands
* [hostinger reach contact-fields create](hostinger_reach_contact-fields_create.md)	 - Create a contact field
* [hostinger reach contact-fields delete](hostinger_reach_contact-fields_delete.md)	 - Delete a contact field
* [hostinger reach contact-fields list](hostinger_reach_contact-fields_list.md)	 - List contact fields
* [hostinger reach contact-fields update](hostinger_reach_contact-fields_update.md)	 - Update a contact field

