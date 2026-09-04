## hostinger hosting nodejs runtime-logs

Get Node.js runtime logs

### Synopsis

Returns the Node.js application's runtime console log entries, oldest first, each with
timestamp, level and message. On the first call send `period` (`1h`, `1d`, `1w` or `1m`)
and optionally `levels` and `limit` (1-5000, default 1000); when more entries match than
`limit`, the newest are kept.

To poll for new entries send `total_lines + 1` from the previous response as `from_line`
and omit `period`; `period` and `from_line` cannot be combined. Lines that are not JSON
with a timestamp, level and message are skipped, so `logs` may hold fewer than `limit`
entries while `total_lines` counts every raw line. Entries with a timestamp before
`last_deployed_at` belong to the previous deployment. Returns an empty `logs` list when
the application has not written a log file yet.

```
hostinger hosting nodejs runtime-logs <username> <domain> [flags]
```

### Options

```
      --from-line total_lines + 1   1-based line of the log file to start from. For polling send total_lines + 1 from the
                                    previous response. Cannot be combined with `period`.
  -h, --help                        help for runtime-logs
      --levels strings              Return only entries with these log levels, sent as a comma-separated list, e.g. ERROR,WARN.
                                    Matching runs on the raw log line, so entries written with numeric levels (for example by
                                    pino) are excluded while this filter is set. (one of: LOG, ERROR, WARN, INFO, DEBUG, TRACE)
      --limit int                   Maximum number of log entries to return. When more entries match, the newest are kept. (default 1000)
      --period from_line            Time window for the first fetch. Required when from_line is not sent. (one of: 1h, 1d, 1w, 1m)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

