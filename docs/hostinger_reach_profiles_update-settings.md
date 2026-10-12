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

```
hostinger reach profiles update-settings <profile-uuid> [flags]
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
      --sender-email @               Part of the default sender address before the @, for example `hello`. Never send a
                                     full address: Reach adds the profile's connected sending domain itself.
      --sender-name string           Default sender name. Letters, numbers and spaces only.
      --social-links strings         
      --socials string               Social profile links keyed by platform. (JSON)
      --text-color string            
      --topics strings               
      --website-url string           
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

