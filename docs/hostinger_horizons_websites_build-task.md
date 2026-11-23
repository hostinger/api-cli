## hostinger horizons websites build-task

Get website build task

### Synopsis

Get the state and outcome of a website build started by the `Create website` or `Edit website` tool.\n
Use this tool with the `task.id` those tools return to find out whether the build completed or failed.
Poll it every 30 seconds or so while `state` is `TASK_STATE_SUBMITTED` or `TASK_STATE_WORKING`.\n
When `state` is `TASK_STATE_FAILED` or `TASK_STATE_REJECTED`, the website does not contain the requested
changes: tell the user, using `message` and `reason_code`, and do not publish it.
When `state` is `TASK_STATE_CANCELED`, the user stopped the build and changes made before that may be kept:
publish only if the user asks to.
`reason_code` `insufficient_credits` means the user must get more credits before building again.\n
When `state` is `TASK_STATE_COMPLETED`, the website can be published. A `reason_code` of
`iteration_limit` on a completed build means it stopped early and kept the work done so far.

```
hostinger horizons websites build-task <website-id> <task-id> [flags]
```

### Examples

```
  hostinger horizons websites build-task 123e4567-e89b-12d3-a456-426614174000 223e4567-e89b-12d3-a456-426614174000
```

### Options

```
  -h, --help[=true|false]   help for build-task
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --domain string           Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger horizons websites](hostinger_horizons_websites.md)	 - Websites commands

