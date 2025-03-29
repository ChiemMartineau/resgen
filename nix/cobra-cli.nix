{ cobra-cli }:
cobra-cli.overrideAttrs (prev: { patches = [ ./cobra-cli.patch ]; })
