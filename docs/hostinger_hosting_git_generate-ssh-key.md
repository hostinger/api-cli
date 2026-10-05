## hostinger hosting git generate-ssh-key

Generate Git SSH key

### Synopsis

Creates the SSH key pair of the hosting account and returns the public key. When the account already
has a key, returns that key unchanged. One key serves every website of the account; add the public
key to a private repository as a deploy key before deploying it.

```
hostinger hosting git generate-ssh-key <username> [flags]
```

### Options

```
  -h, --help   help for generate-ssh-key
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

