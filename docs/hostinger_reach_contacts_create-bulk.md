## hostinger reach contacts create-bulk

Create new contacts

### Synopsis

Create a new contact in the email marketing system.

This endpoint allows you to create a new contact with basic information like name, email, and surname.

If double opt-in is enabled, the contact will be created with a pending status
and a confirmation email will be sent.

```
hostinger reach contacts create-bulk <profile-uuid> [flags]
```

### Examples

```
  hostinger reach contacts create-bulk 550e8400-e09b-41d4-a716-400055000000 --email john.doe@example.com
```

### Options

```
      --email string        
  -h, --help                help for create-bulk
      --name string         
      --note string         
      --phone string        Phone number in E.164 format (leading "+" then 7-15 digits)
      --surname string      
      --tag-uuids strings   Existing tags to attach to the created contact
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

