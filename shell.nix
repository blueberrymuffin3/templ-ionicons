{ pkgs ? import <nixpkgs> { } }:
pkgs.mkShell {
  nativeBuildInputs = with pkgs; [
    # Tools
    go_1_22
    templ

    # LSP
    gopls
    go-tools
    delve
  ];
}
