with import <nixpkgs> { };

stdenv.mkDerivation {
  name = "dev";
  buildInputs = [ go_1_13 gcc git ];
  shellHook = "";
}
