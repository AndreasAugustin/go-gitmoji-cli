# go-gitmoji-cli
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
[![CodeQL](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/github-code-scanning/codeql)

[![ci_go](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/ci_go.yml/badge.svg)](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/ci_go.yml)

[![Lint](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/lint.yml/badge.svg)](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/lint.yml)

[![goreleaser](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/release_go.yml/badge.svg)](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/release_go.yml)

```bash
  ____ ____    ____ _ ___ _  _ ____  _ _    ____ _    _
  | __ |  | __ | __ |  |  |\/| |  |  | | __ |    |    |
  |__] |__|    |__] |  |  |  | |__| _| |    |___ |___ |
```

![commit](docs/assets/commit.gif)

## Abstract

[Gitmoji][gitmoji] is an emoji guide for GitHub commit messages. Aims to be a standardization cheatsheet - guide for using emojis on GitHub's commit messages.
 is a nice way to standardize commit messages with emojis.

There is already a nice [gitmoji-cli][gitmoji-cli] command line interface available.
Because I was searching for a nice project to get more into golang this project was created.

It is possible to use different commit message formats.
Per default the format is [conventional-commits][conventional-commits] with emoji
`<type>[optional scope]: :smile: <description>`

## Installation

:warning: not all installation methods are tested (yet).
Currently all is experimental. Will be updated when the first release has been done.

### brew

```bash
brew install https://github.com/AndreasAugustin/go-gitmoji-cli/Formula
```

### Aur

```bash
pamac install go-gitmoji-cli
```

### Docker

There are 2 docker registries available. Just pull the image to have a local test setup available.
To use it, mount the current dir into the working dir.

- [dockerhub][go-gitmoji-cli-docker-hub]
- [ghcr][go-gitmoji-cli-docker-ghcr]

### Manual

Download the related release [here][go-gitmoji-cli-releases] and unpack the related binary into your path

## Configuration

It is possible to configure the cli either with a `.go-gitmoji-cli.json` file within the repo directory
or with command line flags.
Environment variables are supported (case insensitive). The key is the same like the parameter with a prefix **GO_GITMOJI_CLI_**.

| **parameter**    | **description**                                                                                 | **default**                        |
|------------------|-------------------------------------------------------------------------------------------------|------------------------------------|
| auto_add         | perform automatically a `git add .`                                                             | `false`                            |
| auto_sign        | automatically sign commits (can also be configured with git `git config -g commit.gpgsign=true` | `false`                            |
| emoji_format     | format of emojis `code/emoji`                                                                   | `code`                             |
| scope_prompt     | Prompt for adding the commit scope                                                              | `false`                            |
| body_prompt      | Prompt for adding the commit message body                                                       | `false`                            |
| capitalize_title | If set to true the commit title description will be capitalized                                 | `false`                            |
| gitmojis_url     | The URL of the gitmojis database                                                                | `https://gitmoji.dev/api/gitmojis` |

The configuration values can be changed with

```bash
go-gitmoji-cli config [-g]
```

![config](docs/assets/config.gif)

This will create a file within the local directory or within the OS related config directory (when `-g` flag is enabled).
Reading will follow the following order:

- default values
- global config if exists
- local config if exists
- environment variables

## Usage

### basic commands

```bash
# show available commands
go-gitmoji-cli --help
```

![help](docs/assets/help.gif)

```bash
# show the version
go-gitmoji-cli --version
```

![version](docs/assets/version.gif)

```bash
# list the available gitmojis
go-gitmoji-cli list
```

![list](docs/assets/list.gif)

### Commit

There are 2 ways making commits with the tool

- hooks `go-gitmoji-cli hooks --help`. This will install a commit hook.
- commit `go-gitmoji-cli commit --help`

```bash
# doing a commit with dry-run
go-gitmoji-cli commit --dry-run
```

![commit-dry-run](docs/assets/commit.gif)

## Debugging

There is a flag `--debug` enabling **verbose** logging

## DEV

The development environment targets are located in the [Makefile](Makefile)

```bash
make help
```

## Used libraries and tools

Special thanks to [gitmoji][gitmoji] and [gitmoji-cli][gitmoji-cli]

- [cobra][cobra]
- [viper][viper]
- [logrus][logrus]
- [bubbletea][bubbletea]
- [bubbles][bubbles]
- [lipgloss][lipgloss]
- [go-figure][go-figure]
- [vhs][vhs]
- [goreleaser][goreleaser]

[gitmoji]: https://gitmoji.dev/
[gitmoji-cli]: https://github.com/carloscuesta/gitmoji-cli
[conventional-commits]: https://www.conventionalcommits.org/en/v1.0.0/
[cobra]: https://github.com/spf13/cobra
[viper]: https://github.com/spf13/viper
[logrus]: https://github.com/sirupsen/logrus
[bubbletea]: https://github.com/charmbracelet/bubbletea
[bubbles]: https://github.com/charmbracelet/bubbles
[lipgloss]: https://github.com/charmbracelet/lipgloss
[go-figure]: https://github.com/common-nighthawk/go-figure
[vhs]: https://github.com/charmbracelet/vhs
[goreleaser]: https://goreleaser.com/
[go-gitmoji-cli-releases]: https://github.com/AndreasAugustin/go-gitmoji-cli/releases
[go-gitmoji-cli-docker-hub]: https://hub.docker.com/repository/docker/andyaugustin/go-gitmoji-cli/general
[go-gitmoji-cli-docker-ghcr]: https://github.com/AndreasAugustin/go-gitmoji-cli/pkgs/container/go-gitmoji-cli


## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/AndreasAugustin"><img src="https://avatars0.githubusercontent.com/u/8027933?v=4" width="100px;" alt=""/><br /><sub><b>andy Augustin</b></sub></a><br /><a href="https://github.com/AndreasAugustin/template/commits?author=AndreasAugustin" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
