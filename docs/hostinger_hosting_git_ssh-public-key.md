## hostinger hosting git ssh-public-key

Get Git SSH public key

### Synopsis

Returns the public SSH key of the hosting account. `Deploy website Git repository` uses this key to
clone and pull over SSH, so a private repository works once the key is added to it as a deploy key
on the Git host. `public_key` is null when the account has no key yet.

```
hostinger hosting git ssh-public-key <username> [flags]
```

### Examples

```
  hostinger hosting git ssh-public-key u123456789
```

### Options

```
  -h, --help   help for ssh-public-key
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

