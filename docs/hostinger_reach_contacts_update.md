## hostinger reach contacts update

Update a contact

### Synopsis

Update a contact's attributes and custom field values.

Only the properties present in the request body are changed, so a partial body is enough
to change a single attribute. Sending a property as `null` clears it.

The response carries the contact's core attributes. Read back its tags, custom field
values, source and note with `GET /api/reach/v1/profiles/{profileUuid}/contacts/{contactUuid}`.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --fields
    [
      {
        "selected_option_uuids": [
          "string"
        ],
        "uuid": "string",
        "value": "string"
      }
    ]

```
hostinger reach contacts update <profile-uuid> <contact-uuid> [flags]
```

### Examples

```
  hostinger reach contacts update 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000
```

### Options

```
      --email string                 
      --fields string                Set custom field values. Omit to leave untouched, send an empty array to clear them all. (JSON, @file or -; shape in the help above)
  -h, --help                         help for update
      --name string                  
      --note string                  
      --phone string                 Phone number in E.164 format (leading "+" then 7-15 digits)
      --subscription-status string   (one of: subscribed, unsubscribed, confirmed, pending)
      --surname string               
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

