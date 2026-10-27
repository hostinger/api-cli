## hostinger hosting headless submit-build-report

Submit headless build report

### Synopsis

Submits the build report the headless skill writes at the end of a run (including early
stops), so Hostinger can learn where builds get stuck. Send one report per call, with the
report JSON as the request body, unchanged. Only send a report when the user agreed to it.

The report must match report.schema.json v1; an invalid report returns 422 and is not
stored. The report is linked to the authenticated account, so do not add account details
to it.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --consent
    {
      "note": "string",
      "sent": false
    }

  --info
    {
      "domain": "example.com",
      "outcome": "live",
      "tool": "claude-code"
    }

  --issues
    [
      {}
    ]

  --logging-gaps
    [
      {}
    ]

  --previous-runs
    [
      {}
    ]

  --run
    {
      "follow_ups": [
        {}
      ],
      "id": "2026-10-14T10-00-00Z",
      "kind": "build",
      "skill": "string"
    }

  --skill-gaps
    [
      {}
    ]

  --timing
    {
      "ended_at": "string",
      "minutes_total": 0,
      "stages": [
        {}
      ],
      "started_at": "string"
    }

  --tokens
    {}

  --user-feedback
    [
      {}
    ]

  --what-went-well
    [
      {}
    ]

```
hostinger hosting headless submit-build-report [flags]
```

### Examples

```
  hostinger hosting headless submit-build-report --consent '{"note":"string","sent":false}' --generated-at 2026-10-14T10:30:00Z --info @info.json --issues '[{}]' --previous-runs '[{}]' --run @run.json --schema-version 1 --skill-gaps '[{}]' --timing @timing.json --tokens '{}' --user-feedback '[{}]' --what-went-well '[{}]'
```

### Options

```
      --consent string          Upload status of the report, as written by the headless skill (JSON, @file or -; shape in the help above)
      --generated-at string     When the report was written
  -h, --help                    help for submit-build-report
      --info string             Website, tool and outcome of the run (JSON, @file or -; shape in the help above)
      --issues string           Problems hit during the run (JSON, @file or -; shape in the help above)
      --logging-gaps string     Where the run log is missing or unsure (JSON, @file or -; shape in the help above)
      --previous-runs string    Earlier runs for the same website (JSON, @file or -; shape in the help above)
      --run string              The run the report is for (JSON, @file or -; shape in the help above)
      --schema-version string   Report schema version (one of: 1)
      --skill-gaps string       Places where the skill instructions were missing or wrong (JSON, @file or -; shape in the help above)
      --timing string           Run start and end times and per-stage timings (JSON, @file or -; shape in the help above)
      --tokens string           Token usage, only when the tool reports its own usage (JSON, @file or -; shape in the help above)
      --user-feedback string    Feedback the user gave during the run (JSON, @file or -; shape in the help above)
      --what-went-well string   Things that worked well during the run (JSON, @file or -; shape in the help above)
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting headless](hostinger_hosting_headless.md)	 - Headless commands

