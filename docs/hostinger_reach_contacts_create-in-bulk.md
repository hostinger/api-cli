## hostinger reach contacts create-in-bulk

Create contacts in bulk

### Synopsis

Create many contacts in a profile in a single call.

The contacts are imported in the background, so a success response means the import was
accepted rather than finished. Contacts whose email already exists in the profile are
left as they are. If double opt-in is enabled, new contacts start off pending and are
sent a confirmation email.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --contacts
    [
      {
        "email": "john.doe@example.com",
        "name": "John",
        "phone": "+14155552671",
        "surname": "Doe"
      }
    ]

```
hostinger reach contacts create-in-bulk <profile-uuid> [flags]
```

### Examples

```
  hostinger reach contacts create-in-bulk 550e8400-e09b-41d4-a716-400055000000 --contacts @contacts.json
```

### Options

```
      --contacts string     (JSON, @file or -; shape in the help above)
  -h, --help                help for create-in-bulk
      --note string         Note applied to every created contact
      --tag-uuids strings   Existing tags to attach to every created contact
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach contacts](hostinger_reach_contacts.md)	 - Contacts commands

