## hostinger hosting nodejs build-settings-from-repository

Get Node.js build settings from repository

### Synopsis

Determine the build settings for a Node.js application from the package.json file of a branch
in a Git repository. Use before starting a Git-source build to find out the auto-detected
settings instead of guessing them.

The repository is read through a Git installation: take `installation_uuid` from
`List Git installations` and `owner` and `repository` from `List Git installation repositories`.

Returns 422 when the installation is not found, the repository or branch is not found, or
the branch has no package.json.

```
hostinger hosting nodejs build-settings-from-repository <username> <domain> [flags]
```

### Examples

```
  hostinger hosting nodejs build-settings-from-repository u123456789 mydomain.tld --owner value --repository value --branch value --installation-uuid 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d
```

### Options

```
      --branch string                              Branch whose package.json is inspected
  -h, --help                                       help for build-settings-from-repository
      --installation-uuid List Git installations   Git installation used to access the repository, taken from List Git installations
      --owner string                               Repository owner login. GitLab group paths use slashes.
      --repository string                          Repository name without the .git suffix
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

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

