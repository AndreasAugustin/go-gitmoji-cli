# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class GoGitmojiCli < Formula
  desc "CLI for managing commits gitmoji and conventional commits format"
  homepage "https://github.com/AndreasAugustin/go-gitmoji-cli"
  version "0.5.1-alpha"
  license "MIT"

  depends_on "git"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.5.1-alpha/go-gitmoji-cli_0.5.1-alpha_Darwin_x86_64.tar.gz"
      sha256 "8f7cdebdc1f3e59557852e2baceef5c3fac7b0e7e1e8893feb5fc90b7285bbdf"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
        man1.install "manpages/go-gitmoji-cli.1.gz"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.5.1-alpha/go-gitmoji-cli_0.5.1-alpha_Darwin_arm64.tar.gz"
      sha256 "357f03669a6310b88195c13eb7bdda66dfde0addb231405bee11585bc02192d4"

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
    if Hardware::CPU.intel?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.5.1-alpha/go-gitmoji-cli_0.5.1-alpha_Linux_x86_64.tar.gz"
      sha256 "d302aa0fc4b389e71d2359c1aa5394595609d2f7a5b2be31f759483f6f938bb7"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
        man1.install "manpages/go-gitmoji-cli.1.gz"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.5.1-alpha/go-gitmoji-cli_0.5.1-alpha_Linux_arm64.tar.gz"
      sha256 "82fd0719904caa7b3550171b6c179b9a87a501ffacf6796ba491bf7690eb5bcb"

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
