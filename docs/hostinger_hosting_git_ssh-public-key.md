## hostinger hosting git ssh-public-key

Get Git SSH public key

### Synopsis

Returns the public SSH key of the hosting account. `Deploy website Git repository` uses this key to
clone and pull over SSH, so a private repository works once the key is added to it as a deploy key
on the Git host. `public_key` is null when the account has no key yet.

```
hostinger hosting git ssh-public-key <username> [flags]
```

### Options

```
  -h, --help   help for ssh-public-key
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

