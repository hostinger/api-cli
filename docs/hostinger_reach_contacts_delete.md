## hostinger reach contacts delete

Delete a contact

### Synopsis

Delete a contact with the specified UUID.

This endpoint permanently removes a contact from the email marketing system.

Deprecated. This endpoint cannot target a profile, so it always falls back to the
client's default profile and cannot delete contacts of any other profile. Use
`DELETE /api/reach/v1/profiles/{profileUuid}/contacts/{contactUuid}` instead.

```
hostinger reach contacts delete <uuid> [flags]
```

### Examples

```
  hostinger reach contacts delete 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d
```

### Options

```
  -h, --help   help for delete
  -y, --yes    Skip the confirmation prompt
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

* [hostinger reach contacts](hostinger_reach_contacts.md)	 - Contacts commands

