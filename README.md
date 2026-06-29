# Hostinger API CLI

`hostinger` is the command line interface for the [Hostinger API](https://developers.hostinger.com).

> **This repository is generated.** Do not open pull requests here — every file
> is overwritten by the next generation run. The CLI is generated from the
> OpenAPI specification by
> [hostinger/public-api-generator](https://github.com/hostinger/public-api-generator)
> (see the `cli/` directory there). Report issues or contribute in that repository.

## Installation

### Homebrew (macOS & Linux)

```sh
brew install hostinger/tap/hostinger
```

Upgrade with `brew upgrade hostinger`. Shell tab-completion (bash/zsh/fish) is installed
automatically.

### Binary download

Download the binary for your platform from the
[releases page](https://github.com/hostinger/api-cli/releases).

## Configuration

Create `$HOME/.hostinger.yaml`:

```yaml
api_token: <your API token>
```

or set the `HOSTINGER_API_TOKEN` environment variable. Generate a token at
[hPanel → API](https://hpanel.hostinger.com/profile/api).

## Usage

```
hostinger <group> <subgroup> <verb> [args] [flags]
```

Examples:

```
hostinger billing catalog list
hostinger vps virtual-machines list
hostinger vps vm start 123          # vm is an alias for virtual-machines
hostinger dns records list example.com
```

Output format defaults to a table; use `--format json|table|tree` to change it.
Command documentation lives in [docs/](docs/), and `manifest.json` maps every
API operation to its command. Shell autocompletion: see
[AUTOCOMPLETE.md](AUTOCOMPLETE.md).
