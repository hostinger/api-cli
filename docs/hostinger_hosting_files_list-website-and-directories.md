## hostinger hosting files list-website-and-directories

List website files and directories

### Synopsis

List files and directories under a website's document root.

Use `directory` to browse a subdirectory relative to the document root. Symlinked entries
are listed but never traversed into or resolved.

```
hostinger hosting files list-website-and-directories <username> <domain> [flags]
```

### Options

```
      --directory string     Directory path to check
      --file-types strings   Filter by entry type, e.g. file,directory. Omit for all types. (one of: file, directory, symlink, other)
  -h, --help                 help for list-website-and-directories
      --max-depth int        How many directory levels deep to recurse. (default 5)
      --max-items int        Max number of entries to return in this page. (default 1000)
      --offset int           Number of entries to skip. Page with offset + item count until reaching total_items.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting files](hostinger_hosting_files.md)	 - Files commands

