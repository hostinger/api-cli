## hostinger mail catchalls create-catch-all

Create catch-all

### Synopsis

Create a catch-all that routes all messages sent to unknown addresses
of the domain to the given mailbox. The mailbox address receives a
confirmation email and the catch-all becomes active only after it is
confirmed. A domain can have only one catch-all.

```
hostinger mail catchalls create-catch-all <mailbox-id> [flags]
```

### Examples

```
  hostinger mail catchalls create-catch-all AC1a2b3c4d5e6f7g
```

### Options

```
  -h, --help   help for create-catch-all
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

* [hostinger mail catchalls](hostinger_mail_catchalls.md)	 - Catchalls commands

