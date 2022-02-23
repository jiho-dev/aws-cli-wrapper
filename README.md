
# Overview

aws-cli-warpper is a wrapper allowing shell completion by [TAB]

# install

copy `bin/$(OS)/spc` to `$HOME/bin/`
copy `config/aws-cli-wrapper.yaml` to `$HOME/.aws/`

if you want to add more aws cli under completion, add commands `ec2` section
the list item starts aws command and has parameter name

# Completion

## bash
`$ source <(spc completion bash)` when shell starts

## zsh
`$ ./spc completion zsh > "${fpath[1]}/_spc"`



