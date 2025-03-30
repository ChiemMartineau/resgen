let pkgs = import <nixpkgs> { };
in pkgs.mkShell {
  packages = with pkgs;
    [
      go
      gotools
      gopls
      (callPackage ./cobra-cli.nix { })
      (callPackage ./pkl.nix { })
      (callPackage ./pkl-gen-go.nix { })
    ] ++ (pkgs.lib.optionals pkgs.stdenv.hostPlatform.isLinux
      (with pkgs; [ chromium ]));
}
