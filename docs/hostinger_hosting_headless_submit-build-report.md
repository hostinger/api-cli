## hostinger hosting headless submit-build-report

Submit headless build report

### Synopsis

Submits the build report the headless skill writes at the end of a run (including early
stops), so Hostinger can learn where builds get stuck. Send one report per call, with the
report JSON as the request body, unchanged. Only send a report when the user agreed to it.

The report must match report.schema.json v1; an invalid report returns 422 and is not
stored. The report is linked to the authenticated account, so do not add account details
to it.

```
hostinger hosting headless submit-build-report [flags]
```

### Options

```
      --consent string          Upload status of the report, as written by the headless skill (JSON)
      --generated-at string     When the report was written
  -h, --help                    help for submit-build-report
      --info string             Website, tool and outcome of the run (JSON)
      --issues string           Problems hit during the run (JSON)
      --logging-gaps string     Where the run log is missing or unsure (JSON)
      --previous-runs string    Earlier runs for the same website (JSON)
      --run string              The run the report is for (JSON)
      --schema-version string   Report schema version (one of: 1)
      --skill-gaps string       Places where the skill instructions were missing or wrong (JSON)
      --timing string           Run start and end times and per-stage timings (JSON)
      --tokens string           Token usage, only when the tool reports its own usage (JSON)
      --user-feedback string    Feedback the user gave during the run (JSON)
      --what-went-well string   Things that worked well during the run (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting headless](hostinger_hosting_headless.md)	 - Headless commands

