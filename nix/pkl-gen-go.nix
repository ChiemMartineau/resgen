{ buildGoModule, fetchFromGitHub, lib }:
buildGoModule rec {
  pname = "pkl-gen-go";
  version = "0.9.0";

  src = fetchFromGitHub {
    owner = "apple";
    repo = "pkl-go";
    rev = "v${version}";
    hash = "sha256-ON5rW8Z813lPY6vVcATiGEbaYprlGqQxbp31gm+/A3k=";
  };

  vendorHash = "sha256-1+5Pe+fo7rfr5MeNIUNHln4XIRTkiWZXBeKb1HntND8=";

  subPackages = [ "cmd/pkl-gen-go" ];

  ldflags = [ "-X 'main.Version=${version}'" ];

  meta = with lib; {
    homepage = "https://github.com/apple/pkl-go";
    license = licenses.asl20;
    maintainers = with maintainers; [ samuel-martineau ];
  };
}
