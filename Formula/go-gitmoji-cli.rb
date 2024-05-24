# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class GoGitmojiCli < Formula
  desc "CLI for managing commits gitmoji and conventional commits format"
  homepage "https://github.com/AndreasAugustin/go-gitmoji-cli"
  version "0.6.0-alpha"
  license "MIT"

  depends_on "git"

  on_macos do
    on_intel do
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.6.0-alpha/go-gitmoji-cli_0.6.0-alpha_Darwin_x86_64.tar.gz"
      sha256 "d5cc5751bbd0cd1a8befddf07a79e8b3221449367be01c0138e30719c25e1ae2"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
        man1.install "manpages/go-gitmoji-cli.1.gz"
      end
    end
    on_arm do
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.6.0-alpha/go-gitmoji-cli_0.6.0-alpha_Darwin_arm64.tar.gz"
      sha256 "25a9a86cc7f3c8f9e36a54a80b85dfdc6a2ca53d5a0b3eeecf9a6307a92dfe54"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
        man1.install "manpages/go-gitmoji-cli.1.gz"
      end
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.6.0-alpha/go-gitmoji-cli_0.6.0-alpha_Linux_x86_64.tar.gz"
        sha256 "f90bfef05338df47545a1ee1cff8ff55e53a80abf6b911f47291e190c44649c4"

        def install
          bin.install "go-gitmoji-cli"
          bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
          zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
          fish_completion.install "completions/go-gitmoji-cli.fish"
          man1.install "manpages/go-gitmoji-cli.1.gz"
        end
      end
    end
    on_arm do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.6.0-alpha/go-gitmoji-cli_0.6.0-alpha_Linux_arm64.tar.gz"
        sha256 "e0ce4e74b0bd5e30147efd5623c666cf5aa226c973be87911e4066e0b0c76694"

        def install
          bin.install "go-gitmoji-cli"
          bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
          zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
          fish_completion.install "completions/go-gitmoji-cli.fish"
          man1.install "manpages/go-gitmoji-cli.1.gz"
        end
      end
    end
  end
end
