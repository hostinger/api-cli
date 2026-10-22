## hostinger horizons websites edit

Edit website

### Synopsis

Edit an existing Hostinger Horizons website with a follow-up message.\n
Use this tool when the user wants to change, extend or fix a website that already exists.\n
This tool queues the requested changes and returns the website URL and ID.
The changes are applied asynchronously.\n
After invoking this tool, your chat reply must be EXACTLY 1 sentence summarizing
that Hostinger Horizons is now applying the requested changes and they will be ready
in a few minutes, and you should provide the website URL to the user immediately.
Do not write code.\n
If the tool call fails with an error, you should provide a clear explanation of the error
and do not generate code yourself in the chat.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --message
    [
      {
        "text": "Add a testimonials section below the menu and make the hero background green",
        "type": "text"
      }
    ]

```
hostinger horizons websites edit <website-id> [flags]
```

### Examples

```
  hostinger horizons websites edit 123e4567-e89b-12d3-a456-426614174000 --message @message.json
```

### Options

```
  -h, --help             help for edit
      --message string   (JSON, @file or -; shape in the help above)
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

* [hostinger horizons websites](hostinger_horizons_websites.md)	 - Websites commands

