require "language/go"

class Converter < Formula
  desc "converter: Advanced SSH config - A transparent wrapper that adds regex, aliases, gateways, includes, dynamic hostnames to SSH"
  homepage "https://github.com/moul/converter"
  url "https://github.com/moul/converter/archive/v1.0.0.tar.gz"
  sha256 "c7d53c61c7ca20827979c1d20ba7fcefcb315ac0645821f69ddbe41c432e160b"

  head "https://github.com/moul/converter.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["CGO_ENABLED"] = "0"
    ENV.prepend_create_path "PATH", buildpath/"bin"

    mkdir_p buildpath/"src/github.com/moul"
    ln_s buildpath, buildpath/"src/github.com/moul/converter"
    Language::Go.stage_deps resources, buildpath/"src"

    # FIXME: update version
    system "go", "get", "github.com/BurntSushi/toml"
    system "go", "get", "github.com/Sirupsen/logrus"
    system "go", "build", "-o", "converter", "./cmd/converter"
    bin.install "converter"

    # FIXME: add autocompletion
  end

  test do
    output = shell_output(bin/"converter --version")
    assert output.include? "converter version 2"
  end
end
