#!/bin/sh

go install github.com/bootdotdev/bootdev@latest

# For Linux/WSL
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc
# next, reload your shell configuration
source ~/.zshrc

bootdev login
