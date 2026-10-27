## hostinger ecommerce stores create

Create store

### Synopsis

Create a new store for your account.

A primary sales channel is created alongside the store.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --sales-channel
    {
      "external_id": "string",
      "type": "custom"
    }

```
hostinger ecommerce stores create [flags]
```

### Examples

```
  hostinger ecommerce stores create
```

### Options

```
      --company-email string   
      --company-name string    
      --country-code string    ISO 3166-1 alpha-2 country code.
  -h, --help                   help for create
      --language string        ISO 639-1 language code.
      --name string            
      --sales-channel string   (JSON, @file or -; shape in the help above)
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

* [hostinger ecommerce stores](hostinger_ecommerce_stores.md)	 - Stores commands

