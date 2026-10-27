## hostinger reach profiles update-settings

Update profile settings

### Synopsis

Change the profile's default sender and brand details.

Send only the fields to change: every omitted field keeps its value, and `null` or an empty
string clears a text field. The default sender is what campaigns created afterwards start
from; existing campaigns keep their own sender. `sender_email` is only the part before the
`@`, because Reach combines it with the profile's connected sending domain, so an address
outside that domain cannot be set. Connecting a sending domain is a separate endpoint and
a connected domain cannot be changed here. Returns the settings after the change.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --socials
    {}

```
hostinger reach profiles update-settings <profile-uuid> [flags]
```

### Examples

```
  hostinger reach profiles update-settings 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
      --about string                 
      --accent-color string          
      --background-color string      
      --brand-name string            
      --business-address string      
      --button-color string          
      --button-corners int           
      --button-text string           Button text colour.
      --font-family string           
      --heading-font-family string   
  -h, --help                         help for update-settings
      --keywords strings             
      --links-color string           
      --logo-url string              Public image URL. Logo uploads are not supported.
      --privacy-policy-url string    
      --selling-points strings       
      --sender-email @               Part of the default sender address before the @, for example `hello`. Never send a full address: Reach adds the profile's connected sending domain itself.
      --sender-name string           Default sender name. Letters, numbers and spaces only.
      --social-links strings         
      --socials string               Social profile links keyed by platform. (JSON, @file or -; shape in the help above)
      --text-color string            
      --topics strings               
      --website-url string           
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

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

