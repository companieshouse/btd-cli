# btd-cli

A simple command-line tool for displaying business transaction data (BTD) in a human-readable format.

![btd-cli](btd-cli.gif)

## Installation

Using [Homebrew](https://brew.sh):

```shell
$ brew tap companieshouse/homebrew-ch
$ brew install btd-cli
```

Alternatively:

```shell
$ go install github.com/companieshouse/btd-cli@latest
```

## Usage

`btd-cli` commands use the following structure:

```shell
$ btd-cli <command> [subcommand] [flags and arguments]
```

To display global, command, or subcommand usage information, use one of the following:

```shell
$ btd-cli help
$ btd-cli help <command>
$ btd-cli help <command> <subcommand>
```

`btd-cli` also support the `--help` flag (and its shortened form `-h`) as an alternative to the `help` command: 

```shell
$ btd-cli --help
$ btd-cli <command> --help
$ btd-cli <command> <subcommand> --help
```

### Parsing Data

The `parse` command supports multiple subcommands for parsing business transaction data from different sources. These include `string` and `file` subcommands, which are described in more detail below.

#### Parsing data strings

Use the `string` subcommand to parse business transaction data from a command-line argument string (quotes are required):

```shell
btd-cli parse string '...'
```

#### Parsing data files

Use the `file` subcommand to parse business transaction data from an input file. Each non-empty line of the input file is assumed to contain a complete transaction and is parsed and output independently of any other:

```shell
btd-cli parse file <path>
```

## Global Flags

`btd-cli` supports the following global flags:

| Flag              | Description                                  | Default               |
|-------------------|----------------------------------------------|-----------------------|
| `-c`, `--config`  | Config file path; see [Configuration File](#configuration-file) | `$HOME/.btd-cli.toml` |
| `-t`, `--tag-map` | Path to the tag map file                     | `tagmap.dat`          |

## Configuration File

`btd-cli` will read its settings from a [TOML](https://toml.io/en/) format configuration file at `$HOME/.btd-cli.toml` if one exists. Configuration file settings always take precedence over built-in defaults, and command-line flags always take precedence over both configuration file settings and built-in defaults. The configuration file path can be changed using the `--config` flag (or its shortened form `-c`); see [Global Flags](#global-flags).

The following configuration file settings are supported:

| Name      | Description                                                                 |
|-----------|-----------------------------------------------------------------------------|
| `tag-map` | Path to the tag map file (`$var` and `${var}` style environment variables will be expanded) |

For example, to set a default path for the tag map in the configuration file:

```toml
tag-map = '$HOME/projects/chl-tuxedo/chtuxgw/config/tagmap.dat'
```

## Updating the example gif image

To update the example `btd-cli.gif` image used in this `README.md` file using [VHS](https://github.com/charmbracelet/vhs):

- Modify the `btd-cli.tape` configuration file if needed
- Update the `btd-cli.gif` image by running:

```shell
vhs < btd-cli.tape
```

## License

This project is subject to the terms of the [MIT License](/LICENSE).
