
# Overview
 
aws-cli-warpper is a wrapper allowing shell completion by [TAB]  

# install

copy `bin/$(OS)/acw` to `$HOME/bin/`  
copy `config/acw.yaml` to `$HOME/.aws/`  

if you want to add more aws cli under completion, add commands `ec2` section  
the list item starts aws command and has parameter name  

# Completion  

## bash  
`$ source <(~/bin/acw completion bash)` when shell starts  

## zsh
`$ ./acw completion zsh > "${fpath[1]}/_acw"`  

or 

`make zsh`  
