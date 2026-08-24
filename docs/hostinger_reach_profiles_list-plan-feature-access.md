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

### Options

```
  -h, --help   help for list-plan-feature-access
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

