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
Nevertheless, I was searching a nice project to get a bit more into golang so this cli was born.
There are some feature differences between both projects.

Per default the format is [conventional-commits][conventional-commits] with emoji
`<type>[optional scope]: :smile: <description>`

## Installation

:warning: Currently the package is in state `pre-alpha` please be aware that some functionality may change and some bugs may happen.

### brew

```bash
brew tap AndreasAugustin/go-gitmoji-cli https://github.com/AndreasAugustin/go-gitmoji-cli
brew install AndreasAugustin/go-gitmoji-cli/go-gitmoji-cli
```

### Aur

The package is located [here][go-gitmoji-cli-aur]

```bash
pamac update -a
pamac install go-gitmoji-cli-bin
```

### Docker

There are 2 docker registries available. Just pull the image to have a local test setup available.
To use it, mount the current dir into the working dir.

- [dockerhub][go-gitmoji-cli-docker-hub]
- [ghcr][go-gitmoji-cli-docker-ghcr]

```bash
# available <tag> are the git tags

# docker hub
# pull the image
docker pull andyaugustin/go-gitmoji-cli:<tag>
# run the container -> will open a zsh session
# note: the local volume mount is missing in the sample command
docker run -it andyaugustin/go-gitmoji-cli:<tag>

# if you prefer ghcr instead of docker hub
# docker pull ghcr.io/andreasaugustin/go-gitmoji-cli:<tag>
# docker run -it ghcr.io/andreasaugustin/go-gitmoji-cli:<tag>
```

### Go

```bash
 go install github.com/AndreasAugustin/go-gitmoji-cli@latest
```

### Manual

Download the related release [here][go-gitmoji-cli-releases] and unpack the related binary into your path

## Configuration

It is possible to configure the cli either with a `.go-gitmoji-cli.json` file within the repo directory
or with command line flags.
Environment variables are supported (case insensitive). The key is the same like the parameter with a prefix **GO_GITMOJI_CLI_**.
All parameters are able to be modified with flags.

| **parameter**    | **description**                                                                                 | **default**                        |
|------------------|-------------------------------------------------------------------------------------------------|------------------------------------|
| auto_add         | perform automatically a `git add .`                                                             | `false`                            |
| auto_sign        | automatically sign commits (can also be configured with git `git config -g commit.gpgsign=true` | `false`                            |
| emoji_format     | format of emojis `code/emoji`                                                                   | `code`                             |
| scope_prompt     | Prompt for adding the commit scope                                                              | `false`                            |
| body_prompt      | Prompt for adding the commit message body                                                       | `false`                            |
| capitalize_title | If set to true the commit title description will be capitalized                                 | `false`                            |
| gitmojis_url     | The URL of the gitmojis database                                                                | `https://gitmoji.dev/api/gitmojis` |
| debug            | enable debug mode                                                                               | false                              |

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
- command flags

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
go-gitmoji-cli list gitmojis
```

![list](docs/assets/list_gitmojis.gif)

```bash
# list the available commit types
go-gitmoji-cli list commit-types
```

![list](docs/assets/list_commit_types.gif)

### Commit

There are 2 ways making commits with the tool

- hooks `go-gitmoji-cli hooks --help`. With that command it is possible to install a commit hook. To use the hook, just do a `git commit <options>`.
This will trigger the cli.
- commit `go-gitmoji-cli commit --help`. It is not possible to use this command when you have installed a hook.

```bash
# doing a commit with dry-run
go-gitmoji-cli commit --dry-run
```

![commit-dry-run](docs/assets/commit.gif)

Some arguments and flags you know from git will be reused.
The first message will be parsed and the single parts will be reused.
E.g. `git commit -S -m "feat(api)!: :smile: also just parts of the message will be reused" -m "this is a message body"`
This is also true when the `go-gitmoji-cli commit -S -m "..." -m "..."` is used.

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
[go-gitmoji-cli-aur]: https://aur.archlinux.org/packages/go-gitmoji-cli-bin

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
     <td align="center" valign="top" width="14.28%"><a href="https://github.com/AndreasAugustin"><img src="https://avatars0.githubusercontent.com/u/8027933?v=4?s=100" width="100px;" alt="andy Augustin"/><br /><sub><b>andy Augustin</b></sub></a><br /><a href="https://github.com/AndreasAugustin/actions-template-sync/commits?author=AndreasAugustin" title="Documentation">📖</a> <a href="https://github.com/AndreasAugustin/actions-template-sync/commits?author=AndreasAugustin" title="Code">💻</a> <a href="https://github.com/AndreasAugustin/actions-template-sync/pulls?q=is%3Apr+reviewed-by%3AAndreasAugustin" title="Reviewed Pull Requests">👀</a> <a href="#security-AndreasAugustin" title="Security">🛡️</a> <a href="#ideas-AndreasAugustin" title="Ideas, Planning, & Feedback">🤔</a> <a href="#example-AndreasAugustin" title="Examples">💡</a> <a href="#content-AndreasAugustin" title="Content">🖋</a> </td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
