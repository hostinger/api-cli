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

### Options

```
      --content string        URL pointing to docker-compose.yaml file, Github repository or raw YAML content of the compose file
      --environment string    Project environment variables
  -h, --help                  help for create
      --project-name string   Docker Compose project name using alphanumeric characters, dashes, and underscores only
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps docker](hostinger_vps_docker.md)	 - Docker Manager commands

