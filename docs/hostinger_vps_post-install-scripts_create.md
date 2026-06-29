## hostinger vps post-install-scripts create

Create post-install script

### Synopsis

Add a new post-install script to your account, which can then be used after virtual machine installation.

The script contents will be saved to the file `/post_install` with executable attribute set
and will be executed once virtual machine is installed.
The output of the script will be redirected to `/post_install.log`. Maximum script size is 48KB.

Use this endpoint to create automation scripts for VPS setup tasks.

```
hostinger vps post-install-scripts create [flags]
```

### Options

```
      --content string   Content of the script
  -h, --help             help for create
      --name string      Name of the script
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps post-install-scripts](hostinger_vps_post-install-scripts.md)	 - Post-install scripts commands

