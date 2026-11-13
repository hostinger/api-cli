## hostinger gpu exposed-services create

Create exposed service

### Synopsis

Expose a service running on a specified GPU instance to the internet.

The service is exposed asynchronously: a successful request returns immediately, and the service
appears in the exposed services list once it gets its public hostname and external port.
Fails when the instance already has the maximum of 4 exposed services.

Use this endpoint to make an application on your GPU instance publicly reachable.

```
hostinger gpu exposed-services create <instance-id> [flags]
```

### Examples

```
  hostinger gpu exposed-services create 6325 --internal-port 8000 --name Ollama --protocol tcp --type http
```

### Options

```
  -h, --help[=true|false]   help for create
      --internal-port int   Port the service listens on inside the instance, unique within the instance
      --name string         Service name, unique within the instance
      --protocol string     (one of: tcp, udp)
      --type string         (one of: http, https)
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

* [hostinger gpu exposed-services](hostinger_gpu_exposed-services.md)	 - Exposed services commands

