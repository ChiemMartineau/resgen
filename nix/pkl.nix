{ stdenv, lib, fetchurl, autoPatchelfHook, zlib }:
let
  platforms = {
    "x86_64-linux" = {
      name = "linux-amd64";
      sha256 = "0fdzaqqf2z57fiy89xrd3a8mxg6il0p91jg29k11rnjz5bnl7rzx";
    };
    "aarch64-linux" = {
      name = "linux-aarch64";
      sha256 = "01hgw89jkvz6kqkakyia16p2ravqfxb2xhmszw8h0lh4xw4lrmr0";
    };
    "x86_64-darwin" = {
      name = "macos-amd64";
      sha256 = "0m5js5rj7b6ib5imbfdi18sb32w6rmc4mhyh5sjiflxf1mhvs2av";
    };
    "aarch64-darwin" = {
      name = "macos-aarch64";
      sha256 = "1q0gn53qbfa2xvz1n8sf9gv9cfj67zkqk77rdz4vazh6lzddx4yk";
    };
  };
in stdenv.mkDerivation rec {
  pname = "pkl";
  version = "0.28.1";

  src = fetchurl (with platforms.${stdenv.hostPlatform.system}; {
    url =
      "https://github.com/apple/pkl/releases/download/${version}/pkl-${name}";
    inherit sha256;
  });

  dontUnpack = true;

  nativeBuildInputs = [ autoPatchelfHook ];

  buildInputs = [ zlib ];

  installPhase = ''
    runHook preInstall
    install -D $src $out/bin/pkl
    runHook postInstall
  '';

  meta = with lib; {
    homepage = "https://pkl-lang.org";
    description =
      "A configuration as code language with rich validation and tooling. ";
    platforms = builtins.attrNames platforms;
    maintainers = with maintainers; [ samuel-martineau ];
    license = licenses.asl20;
  };
}
