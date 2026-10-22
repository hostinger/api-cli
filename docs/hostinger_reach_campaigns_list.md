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

### Examples

```
  hostinger reach campaigns list 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
      --all                  Fetch every page and print all items
  -h, --help                 help for list
      --page int             Page number
      --per-page int         Number of items per page (default 25)
      --sort-direction asc   Order campaigns by creation date. Newest first unless set to asc. (one of: asc, desc)
      --status publish       Filter campaigns by status. A fully sent campaign has the status publish. There is no `sent` status, and campaigns can be neither paused nor archived. (one of: draft, scheduled, sending, publish, failed)
      --type campaign        Filter campaigns by type. Defaults to campaign, which leaves out the emails sent by automations and the double opt-in confirmations. (one of: campaign, automation, double_opt_in) (default "campaign")
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

* [hostinger reach campaigns](hostinger_reach_campaigns.md)	 - Campaigns commands

