## hostinger mail aliases create-alias

Create alias

### Synopsis

Create an alias for the given mailbox. The alias address is formed
from the given local part and the domain of the mailbox. Messages
sent to the alias are delivered to the mailbox.

```
hostinger mail aliases create-alias <mailbox-id> [flags]
```

### Options

```
  -h, --help                help for create-alias
      --local-part string   Local part of the alias address (the part before the @). The domain is taken from the mailbox. Case-insensitive and stored lowercase; must start and end with a letter or digit; single dots, underscores and hyphens are allowed in between.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail aliases](hostinger_mail_aliases.md)	 - Aliases commands

