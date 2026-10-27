## hostinger reach tags remove-contacts-from

Remove contacts from a tag

### Synopsis

Remove a tag from many contacts at once.

Pass `contact_uuids` to target specific contacts, or `all_contacts` to target every contact
in the profile. The work is queued, so a success response means it was accepted rather than
finished. The tag itself and the contacts are not deleted.

```
hostinger reach tags remove-contacts-from <profile-uuid> <tag-uuid> [flags]
```

### Examples

```
  hostinger reach tags remove-contacts-from 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000
```

### Options

```
      --all-contacts            Apply to every contact in the profile
      --contact-uuids strings   Contacts to apply the change to. Required unless all_contacts is true.
  -h, --help                    help for remove-contacts-from
  -y, --yes                     Skip the confirmation prompt
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

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

