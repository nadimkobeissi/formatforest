# This file was generated by GoReleaser. DO NOT EDIT.
class Formatforest < Formula
  desc "Simple and elegant blogging engine written in Go."
  homepage "https://formatforest.com"
  version "0.0.3"
  bottle :unneeded

  if OS.mac?
    url "https://gitlab.com/nadimk/formatforest/uploads//formatforest_0.0.3_macos_amd64.zip"
    sha256 "d75b6a6d98efdfdd03b151e7411f19e89eb06f5dca7b16871dd5b5ca2e16ca0b"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://gitlab.com/nadimk/formatforest/uploads//formatforest_0.0.3_linux_amd64.zip"
      sha256 "f7c02302cee0b344680bd5bba28c52515cca7a04bc63065a5a0aeb5484392423"
    end
    if Hardware::CPU.arm?
      if Hardware::CPU.is_64_bit?
        url "https://gitlab.com/nadimk/formatforest/uploads//formatforest_0.0.3_linux_arm64.zip"
        sha256 "1b915d8251bbf2fed506aadf05030658c1391049d4a19587ea6b7891b7955b37"
      else
        url "https://gitlab.com/nadimk/formatforest/uploads//formatforest_0.0.3_linux_armv6.zip"
        sha256 "5760e24b3aef4d00f1eaf847ea218f55b80689dcfc602c4b4e94b768068f3172"
      end
    end
  end

  def install
    bin.install "formatforest"
  end
end
