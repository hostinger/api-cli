## hapi vps virtual-machines metrics

Get metrics

### Synopsis

Retrieve historical metrics for a specified virtual machine.

It includes the following metrics: 
- CPU usage
- Memory usage
- Disk usage
- Network usage
- Uptime

Use this endpoint to monitor VPS performance and resource utilization over time.

```
hapi vps virtual-machines metrics <virtual-machine-id> [flags]
```

### Options

```
      --date-from string   
      --date-to string     
  -h, --help               help for metrics
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps virtual-machines](hapi_vps_virtual-machines.md)	 - Virtual machine commands

