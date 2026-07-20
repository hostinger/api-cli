## hostinger agency-hosting cron-jobs list-plan-website

List Agency Plan website cron jobs

### Synopsis

Returns a paginated list of cron jobs configured for an Agency Plan website.

Each entry includes the schedule expression and the command executed on that schedule.

```
hostinger agency-hosting cron-jobs list-plan-website <website_uid> [flags]
```

### Options

```
  -h, --help           help for list-plan-website
      --page int       Page number
      --per-page int   Number of items per page (default 25)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting cron-jobs](hostinger_agency-hosting_cron-jobs.md)	 - Cron Jobs commands

