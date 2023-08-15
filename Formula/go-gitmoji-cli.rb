# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class GoGitmojiCli < Formula
  desc "CLI for managing commits gitmoji and conventional commits format"
  homepage "https://github.com/AndreasAugustin/go-gitmoji-cli"
  version "0.1.0-pre-alpha"
  license "MIT"

  depends_on "git"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.1.0-pre-alpha/go-gitmoji-cli_0.1.0-pre-alpha_Darwin_arm64.tar.gz"
      sha256 "4f188dc5ca5ea46940da3ce0f6c88d57e6018b70d7ac5743417baa27be3cb192"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.1.0-pre-alpha/go-gitmoji-cli_0.1.0-pre-alpha_Darwin_x86_64.tar.gz"
      sha256 "173c2ed761e56c1f5bc04fbe44891f3cffa7162f9e9bf1dafb47f7c8b783de4a"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.1.0-pre-alpha/go-gitmoji-cli_0.1.0-pre-alpha_Linux_arm64.tar.gz"
      sha256 "5d04346dd31f4ebd98929bdfffbaca550110133a0e265083b30119d29c2ef422"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.1.0-pre-alpha/go-gitmoji-cli_0.1.0-pre-alpha_Linux_x86_64.tar.gz"
      sha256 "fefd4592cc320df39bf30b98dcb6b56ae6eede5511afac6fc75330cbdefe4894"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
      end
    end
  end
end
