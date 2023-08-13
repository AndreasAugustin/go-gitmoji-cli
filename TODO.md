# TODO

- [x] edit the long descriptions for the cmds
- [x] git hook wrong command name
- [ ] lots of tests missing
- [x] local cache for list (use viper with config file)
- [x] add check for hook and disable commit if hook is created
- [ ] commands
  - [x] list -> use list bubbles prompt and local cache in homedir
  - [x] search -> remove (merge with list)
  - [ ] update
    - [x] gitmojis cache
    - [ ] cli
  - [x] info -> get infos like cache dir
  - [x] check all config parameters implemented
  - [ ] commit
    - [x] add dryrun
    - [x] gitmoji
    - [x] scope
    - [x] message
    - [ ] add 2nd message optional
    - [ ] add flags for message, scope,.. and prefill
    - [ ] config option for commit message formats
  - [x] config -> write config as local config file (with prompt)
    - [x] possibility to use env variables
    - [x] add -g --global flag for global configuration
  - [x] init -> creates git hook -> change name
  - [x] remove -> remove git hook -> change name
- [x] bubbletea and bubbles as prompt -> remove current search, use list instead (bubbles). remove spinner, use bubbles
- [x] remove query cmd -> included into list
- [ ] docker for dev  and release a example image
- [x] ci/cd -> releaser action [gorealeaser][goreleaser]
- [x] bug in config -> gitmoji url is not shown properly

[goreleaser]: https://goreleaser.com/ci/actions/
