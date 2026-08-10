## hostinger reach contacts create-in-bulk

Create contacts in bulk

### Synopsis

Create many contacts in a profile in a single call.

The contacts are imported in the background, so a success response means the import was
accepted rather than finished. Contacts whose email already exists in the profile are
left as they are. If double opt-in is enabled, new contacts start off pending and are
sent a confirmation email.

```
hostinger reach contacts create-in-bulk <profile-uuid> [flags]
```

### Options

```
      --contacts string      (JSON)
  -h, --help                help for create-in-bulk
      --note string         Note applied to every created contact
      --tag-uuids strings   Existing tags to attach to every created contact
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach contacts](hostinger_reach_contacts.md)	 - Contacts commands

