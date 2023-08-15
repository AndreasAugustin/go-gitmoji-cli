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
      sha256 "19f15249003b17ba27e488e8a1ad371b2b553485be934c06d934c01dfdb706cf"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.1.0-pre-alpha/go-gitmoji-cli_0.1.0-pre-alpha_Darwin_x86_64.tar.gz"
      sha256 "d5590292d52600c975ee29dd769fd828831006f49fc2c1c72cf1d63e34350000"

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
      sha256 "ca268a94ae148b17e64470f23705ddc058347ca014f71d994dc44b9a743ed55f"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.1.0-pre-alpha/go-gitmoji-cli_0.1.0-pre-alpha_Linux_x86_64.tar.gz"
      sha256 "750914e097a36eb57c66f886992c4721cdf2c0d24ef49d7d390f2f544c64483e"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
      end
    end
  end
end
