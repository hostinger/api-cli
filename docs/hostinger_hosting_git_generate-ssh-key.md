## hostinger hosting git generate-ssh-key

Generate Git SSH key

### Synopsis

Creates the SSH key pair of the hosting account and returns the public key. When the account already
has a key, returns that key unchanged. One key serves every website of the account; add the public
key to a private repository as a deploy key before deploying it.

```
hostinger hosting git generate-ssh-key <username> [flags]
```

### Examples

```
  hostinger hosting git generate-ssh-key u123456789
```

### Options

```
  -h, --help   help for generate-ssh-key
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

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

