# go-gitmoji-cli
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
[![CodeQL](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/github-code-scanning/codeql)

[![ci_go](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/ci_go.yml/badge.svg)](https://github.com/AndreasAugustin/go-gitmoji-cli/actions/workflows/ci_go.yml)

![Lint](https://github.com/AndreasAugustin/go-gitmoji-cli/workflows/Lint/badge.svg)

## Abstract

[Gitmoji][gitmoji] is an emoji guide for GitHub commit messages. Aims to be a standardization cheatsheet - guide for using emojis on GitHub's commit messages.
 is a nice way to standardize commit messages with emojis.

There is already a nice [gitmoji-cli][gitmoji-cli] command line interface available.
Because I was searching for a nice project to get more into golang this project was created.

It is possible to use different commit message formats.
Per default the format is [conventional-commits[conventional-commits] with emoji
`<type>[optional scope]: :smile: <description>`

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

This will create a file within the local directory or within the OS related config directory (when `-g` flag is enabled).
Reading will follow the following order:

- default values
- global config if exists
- local config if exists
- environment variables

## Debugging

There is a flag `--debug` enabling **verbose** logging

## DEV

The development environment targets are located in the [Makefile](Makefile)

```bash
make help
```

[gitmoji]: https://gitmoji.dev/
[gitmoji-cli]: https://github.com/carloscuesta/gitmoji-cli
[conventional-commits]: https://www.conventionalcommits.org/en/v1.0.0/

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
