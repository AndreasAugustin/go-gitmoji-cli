# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class GoGitmojiCli < Formula
  desc "CLI for managing commits gitmoji and conventional commits format"
  homepage "https://github.com/AndreasAugustin/go-gitmoji-cli"
  version "0.5.2-alpha"
  license "MIT"

  depends_on "git"

  on_macos do
    on_intel do
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.5.2-alpha/go-gitmoji-cli_0.5.2-alpha_Darwin_x86_64.tar.gz"
      sha256 "62fab6f7cfe638d0c5cd82d9e33b925c80011c95c668f09cfebfb959bfe4b5e4"

      def install
        bin.install "go-gitmoji-cli"
        bash_completion.install "completions/go-gitmoji-cli.bash" => "go-gitmoji-cli"
        zsh_completion.install "completions/go-gitmoji-cli.zsh" => "_go-gitmoji-cli"
        fish_completion.install "completions/go-gitmoji-cli.fish"
        man1.install "manpages/go-gitmoji-cli.1.gz"
      end
    end
    on_arm do
      url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.5.2-alpha/go-gitmoji-cli_0.5.2-alpha_Darwin_arm64.tar.gz"
      sha256 "eda45f87d819e654f2ecd34250e4fbdd77d4b405361c51a0350ae80d45099928"

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
        url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.5.2-alpha/go-gitmoji-cli_0.5.2-alpha_Linux_x86_64.tar.gz"
        sha256 "de1cc0d97b3bdbcd4900eeab30110537bdf1eb739fd565b81f565be9dced1e17"

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
        url "https://github.com/AndreasAugustin/go-gitmoji-cli/releases/download/v0.5.2-alpha/go-gitmoji-cli_0.5.2-alpha_Linux_arm64.tar.gz"
        sha256 "17a0bde79f182af829674b668c62493941d32cb6a552ca2664cca752fbd7ce98"

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
