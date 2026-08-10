## hostinger reach contacts update

Update a contact

### Synopsis

Update a contact's attributes and custom field values.

Only the properties present in the request body are changed, so a partial body is enough
to change a single attribute. Sending a property as `null` clears it.

The response carries the contact's core attributes. Read back its tags, custom field
values, source and note with `GET /api/reach/v1/profiles/{profileUuid}/contacts/{contactUuid}`.

```
hostinger reach contacts update <profile-uuid> <contact-uuid> [flags]
```

### Options

```
      --email string                 
      --fields string                Set custom field values. Omit to leave untouched, send an empty array to clear them all. (JSON)
  -h, --help                         help for update
      --name string                  
      --note string                  
      --phone string                 Phone number in E.164 format (leading "+" then 7-15 digits)
      --subscription-status string   (one of: subscribed, unsubscribed, confirmed, pending)
      --surname string               
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach contacts](hostinger_reach_contacts.md)	 - Contacts commands

