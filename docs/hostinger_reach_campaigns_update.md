## hostinger reach campaigns update

Update campaign

### Synopsis

Update a campaign's content, sender, template and audience.

Only `draft`, `scheduled` and `failed` campaigns can be updated; a campaign that is sending or
already sent is rejected. Updating a `failed` campaign moves it back to `draft`. Audience changes
replace rather than merge: `all_contacts: true` targets every contact of the profile, while
`segment_uuids` replaces the segments currently targeted. Leave both out to keep the audience as
it is. This does not send the campaign - use the send endpoint for that.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --metadata
    {
      "preheader": "Our biggest deals of the year",
      "source": "api"
    }

```
hostinger reach campaigns update <profile-uuid> <campaign-uuid> [flags]
```

### Examples

```
  hostinger reach campaigns update 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
      --all-contacts[=true|false]                   true targets every contact of the profile, `false` targets only `segment_uuids`. Required when `segment_uuids` or `excluded_segment_uuids` is sent.
      --excluded-segment-uuids all_contacts: true   Segments to leave out. Only allowed together with all_contacts: true.
  -h, --help[=true|false]                           help for update
      --metadata string                             Extra campaign fields. Any key outside the listed ones is rejected. (JSON, @file or -; shape in the help above)
      --segment-uuids strings                       Segments to target. Replaces the segments currently targeted.
      --sender-email string                         From address of the campaign. Its domain has to be verified on the profile before the campaign can be sent.
      --sender-name string                          From name shown to the recipients.
      --subject string                              Subject line of the email.
      --template-uuid string                        Template to send, as returned by the template endpoints.
      --title string                                Name the campaign is listed under. Not shown to the recipients.
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

* [hostinger reach campaigns](hostinger_reach_campaigns.md)	 - Campaigns commands

