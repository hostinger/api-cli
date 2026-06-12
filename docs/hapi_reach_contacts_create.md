## hapi reach contacts create

Create a new contact

### Synopsis

Create a new contact in the email marketing system.

This endpoint allows you to create a new contact with basic information like name, email, and surname.

If double opt-in is enabled,
the contact will be created with a pending status and a confirmation email will be sent.

```
hapi reach contacts create [flags]
```

### Options

```
      --email string     
  -h, --help             help for create
      --name string      
      --note string      
      --phone string     Phone number in E.164 format (leading "+" then 7-15 digits)
      --surname string   
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi reach contacts](hapi_reach_contacts.md)	 - Contacts commands

