## hostinger reach profiles list-plan-feature-access

List plan feature access

### Synopsis

List which plan features the profile can use.

This is the feature lock matrix, not a usage quota. `available` means the feature can be
used right now and `locked` means it is not part of the base plan, so an upgrade is needed.
For remaining emails, recipients and AI credits use the limits endpoint instead.

Worth checking before building something that cannot be activated afterwards, such as an
automation on a plan without automation activation.

```
hostinger reach profiles list-plan-feature-access <profile-uuid> [flags]
```

### Examples

```
  hostinger reach profiles list-plan-feature-access 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for list-plan-feature-access
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

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

