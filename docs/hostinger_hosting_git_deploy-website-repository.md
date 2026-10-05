## hostinger hosting git deploy-website-repository

Deploy website Git repository

### Synopsis

Clones a Git repository into a directory of the website, or pulls it again. An empty or missing
directory gets a clone of the branch. A directory that already holds this repository and branch is
reset to its last commit and pulled: changes made on the server to files the repository tracks are
discarded, files it does not track stay. A directory that holds other files, including another
repository or another branch of this one, is rejected. `composer install` runs after the clone or
pull when the repository has a `composer.json`.

The call waits for the deployment and returns its log. `is_success` false means Git or composer
failed and the log says why. A second call for the same directory is rejected while the first is
still waiting for the server. If the request times out, the deployment may still finish on the
server; calling again later with the same repository and branch pulls.

Private repositories need an SSH URL and the account's Git SSH key from `Generate Git SSH key`,
added to the repository as a deploy key.

```
hostinger hosting git deploy-website-repository <username> <domain> [flags]
```

### Options

```
      --branch string                             Branch to clone and pull
      --directory List website Git repositories   Directory under the website document root, exactly as List website Git repositories returns
                                                  it for an existing repository. Empty, null or omitted means the document root.
  -h, --help                                      help for deploy-website-repository
      --repository-url string                     Clone URL of the repository on any Git host, SSH or HTTPS. Private repositories need an SSH URL
                                                  and the account's Git SSH key added to the repository as a deploy key. An HTTP or HTTPS URL with
                                                  a username or token, or any URL with a password, is rejected.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting git](hostinger_hosting_git.md)	 - Git commands

