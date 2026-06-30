## hostinger hosting cron-jobs create

Create account cron job

### Synopsis

Creates a cron job for the specified account from a schedule expression and a command.

Returns the created cron job, including its uid, which is required to delete the cron job or fetch its output.

```
hostinger hosting cron-jobs create <username> [flags]
```

### Options

```
      --command string   Command to execute on the schedule.
  -h, --help             help for create
      --time string      Cron schedule expression (for example "0 2 * * *" runs daily at 02:00).
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting cron-jobs](hostinger_hosting_cron-jobs.md)	 - Cron Jobs commands

