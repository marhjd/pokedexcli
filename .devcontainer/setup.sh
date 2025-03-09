#!/bin/sh

go install github.com/bootdotdev/bootdev@latest

# For Linux/WSL
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
# next, reload your shell configuration
source ~/.bashrc

bootdev login