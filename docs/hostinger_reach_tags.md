## hostinger reach tags

Tags commands

```
hostinger reach tags [flags]
```

### Options

```
  -h, --help   help for tags
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
* [hostinger reach tags assign-contact-to](hostinger_reach_tags_assign-contact-to.md)	 - Assign a contact to a tag
* [hostinger reach tags assign-contacts-to](hostinger_reach_tags_assign-contacts-to.md)	 - Assign contacts to a tag
* [hostinger reach tags create-or-find](hostinger_reach_tags_create-or-find.md)	 - Create or find tags
* [hostinger reach tags delete](hostinger_reach_tags_delete.md)	 - Delete a tag
* [hostinger reach tags list-profile](hostinger_reach_tags_list-profile.md)	 - List profile tags
* [hostinger reach tags remove-contact-from](hostinger_reach_tags_remove-contact-from.md)	 - Remove a contact from a tag
* [hostinger reach tags remove-contacts-from](hostinger_reach_tags_remove-contacts-from.md)	 - Remove contacts from a tag
* [hostinger reach tags rename](hostinger_reach_tags_rename.md)	 - Rename a tag

