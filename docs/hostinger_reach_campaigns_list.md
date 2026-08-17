## hostinger reach campaigns list

List campaigns

### Synopsis

Get a paginated list of the campaigns in a profile.

Each campaign carries its headline engagement rates. Filter by status to find drafts,
scheduled, sending or sent campaigns, keeping in mind that a fully sent campaign has the
status `publish`. By default only regular campaigns are returned - pass `type` to get the
emails sent by automations or the double opt-in confirmations instead.

```
hostinger reach campaigns list <profile-uuid> [flags]
```

### Options

```
  -h, --help                 help for list
      --page int             Page number
      --per-page int         Number of items per page (default 25)
      --sort-direction asc   Order campaigns by creation date. Newest first unless set to asc. (one of: asc, desc)
      --status publish       Filter campaigns by status.
                             
                             A fully sent campaign has the status publish. There is no `sent` status, and campaigns can
                             be neither paused nor archived. (one of: draft, scheduled, sending, publish, failed)
      --type campaign        Filter campaigns by type.
                             
                             Defaults to campaign, which leaves out the emails sent by automations and the double
                             opt-in confirmations. (one of: campaign, automation, double_opt_in) (default "campaign")
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach campaigns](hostinger_reach_campaigns.md)	 - Campaigns commands

