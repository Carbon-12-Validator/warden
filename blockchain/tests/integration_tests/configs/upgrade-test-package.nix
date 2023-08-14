let
  pkgs = import ../../../nix { };
  fetchEthermint = rev: builtins.fetchTarball "https://gitlab.qredo.com/qrdochain/fusionchain/archive/${rev}.tar.gz";
  released = pkgs.buildGo118Module rec {
    name = "fusiond";
    src = fetchEthermint "d29cdad6e667f6089dfecbedd36bb8d3a2a7d025";
    subPackages = [ "cmd/fusiond" ];
    vendorSha256 = "sha256-cQAol54b6hNzsA4Q3MP9mTqFWM1MvR5uMPrYpaoj3SY=";
    doCheck = false;
  };
  current = pkgs.callPackage ../../../. { };
in
pkgs.linkFarm "upgrade-test-package" [
  { name = "genesis"; path = released; }
  { name = "integration-test-upgrade"; path = current; }
]