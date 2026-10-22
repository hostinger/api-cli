## hostinger vps docker create

Create new project

### Synopsis

Deploy new project from docker-compose.yaml contents or download contents from URL.

URL can be Github repository url in format https://github.com/[user]/[repo]
and it will be automatically resolved to docker-compose.yaml file in
master branch. Any other URL provided must return docker-compose.yaml
file contents.

If project with the same name already exists, existing project will be replaced.

```
hostinger vps docker create <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps docker create 1268054 --content value --project-name my-project-1
```

### Options

```
      --content string          URL pointing to docker-compose.yaml file, Github repository or raw YAML content of the compose file
      --environment string      Project environment variables
  -h, --help                    help for create
      --project-name string     Docker Compose project name using alphanumeric characters, dashes, and underscores only
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
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

* [hostinger vps docker](hostinger_vps_docker.md)	 - Docker Manager commands

