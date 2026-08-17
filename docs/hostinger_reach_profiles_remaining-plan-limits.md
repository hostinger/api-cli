## hostinger reach profiles remaining-plan-limits

Get remaining plan limits

### Synopsis

Get how much of the plan is left for the current period.

Two things to keep in mind before you build alerting on this. The period is a calendar month
rather than a billing anniversary, so the counters reset on the 1st no matter when the
subscription started. And usage is tracked per order, so every profile on the same order shares
one pool and reports the same numbers here. Only the current period is available, past usage is
not kept.

```
hostinger reach profiles remaining-plan-limits <profile-uuid> [flags]
```

### Options

```
  -h, --help   help for remaining-plan-limits
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

