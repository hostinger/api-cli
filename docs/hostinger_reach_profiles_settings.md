## hostinger reach profiles settings

Get profile settings

### Synopsis

Get the profile's saved sender identity and brand details.

Read the default sender before creating a campaign instead of asking the customer, and
reuse the brand fields when generating email content. Any field can be `null` for a profile
that never filled its settings in. This does not check that the sending domain is
verified, use the connected sending domain endpoint for that.

```
hostinger reach profiles settings <profile-uuid> [flags]
```

### Options

```
  -h, --help   help for settings
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

