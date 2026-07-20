## hostinger agency-hosting cron-jobs create-plan-website

Create Agency Plan website cron job

### Synopsis

Creates a cron job for an Agency Plan website from a schedule expression and a command.

Returns the created cron job, including its uuid, which is required to delete the cron job.

```
hostinger agency-hosting cron-jobs create-plan-website <website_uid> [flags]
```

### Options

```
      --command string   Command to run on the schedule. Must not contain pipe (|) or redirection (<, >) characters.
  -h, --help             help for create-plan-website
      --time string      Cron schedule expression (standard 5-field crontab syntax).
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting cron-jobs](hostinger_agency-hosting_cron-jobs.md)	 - Cron Jobs commands

