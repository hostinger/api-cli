## hostinger agency-hosting cron-jobs delete-plan-website

Delete Agency Plan website cron job

### Synopsis

Permanently deletes the cron job identified by its uuid from an Agency Plan website.

The operation is idempotent: deleting a cron job that does not exist succeeds without error.

```
hostinger agency-hosting cron-jobs delete-plan-website <website_uid> <uuid> [flags]
```

### Options

```
  -h, --help   help for delete-plan-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting cron-jobs](hostinger_agency-hosting_cron-jobs.md)	 - Cron Jobs commands

