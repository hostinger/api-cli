## hostinger completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(hostinger completion bash)

To load completions for every new session, execute once:

#### Linux:

	hostinger completion bash > /etc/bash_completion.d/hostinger

#### macOS:

	hostinger completion bash > $(brew --prefix)/etc/bash_completion.d/hostinger

You will need to start a new shell for this setup to take effect.


```
hostinger completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
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

* [hostinger completion](hostinger_completion.md)	 - Generate the autocompletion script for the specified shell

