let
  pkgs = import <nixpkgs> { };
in
pkgs.mkShell {
  packages = with pkgs; [
    go
    gotools
    gopls
    (callPackage ./nix/cobra-cli.nix { })
    (callPackage ./nix/pkl.nix { })
    (callPackage ./nix/pkl-gen-go.nix { })
  ];
}
