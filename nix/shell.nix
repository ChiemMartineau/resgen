let pkgs = import <nixpkgs> { };
in pkgs.mkShell {
  packages = with pkgs; [
    go
    gotools
    gopls
    (callPackage ./cobra-cli.nix { })
    (callPackage ./pkl.nix { })
    (callPackage ./pkl-gen-go.nix { })
  ];
}
