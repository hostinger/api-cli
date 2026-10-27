## hostinger reach campaigns create-draft

Create a draft campaign

### Synopsis

Create a campaign in a profile.

The campaign is created as a draft, so nothing is sent and no contact is touched. It has no
audience yet either - targeting and scheduling are not part of this request, the draft is
finished and sent from the Reach interface.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --metadata
    {
      "preheader": "Our biggest deals of the year",
      "source": "api"
    }

```
hostinger reach campaigns create-draft <profile-uuid> [flags]
```

### Examples

```
  hostinger reach campaigns create-draft 550e8400-e09b-41d4-a716-400055000000 --sender-email marketing@example.com --sender-name 'Marketing Team'
```

### Options

```
  -h, --help                   help for create-draft
      --metadata string        Extra campaign fields. Any key outside the listed ones is rejected. (JSON, @file or -; shape in the help above)
      --sender-email string    From address of the campaign. Its domain has to be verified on the profile before the campaign can be sent.
      --sender-name string     From name shown to the recipients.
      --subject string         Subject line of the email.
      --template-uuid string   Template to send, as returned by the template endpoints. Can be left out and attached later, but the campaign cannot be sent without one.
      --title string           Name the campaign is listed under. Not shown to the recipients.
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

* [hostinger reach campaigns](hostinger_reach_campaigns.md)	 - Campaigns commands

