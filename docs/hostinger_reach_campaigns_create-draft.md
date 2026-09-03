## hostinger reach campaigns create-draft

Create a draft campaign

### Synopsis

Create a campaign in a profile.

The campaign is created as a draft, so nothing is sent and no contact is touched. It has no
audience yet either - targeting and scheduling are not part of this request, the draft is
finished and sent from the Reach interface.

```
hostinger reach campaigns create-draft <profile-uuid> [flags]
```

### Options

```
  -h, --help                   help for create-draft
      --metadata string        Extra campaign fields. Any key outside the listed ones is rejected. (JSON)
      --sender-email string    From address of the campaign. Its domain has to be verified on the profile before
                               the campaign can be sent.
      --sender-name string     From name shown to the recipients.
      --subject string         Subject line of the email.
      --template-uuid string   Template to send, as returned by the template endpoints. Can be left out and
                               attached later, but the campaign cannot be sent without one.
      --title string           Name the campaign is listed under. Not shown to the recipients.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach campaigns](hostinger_reach_campaigns.md)	 - Campaigns commands

