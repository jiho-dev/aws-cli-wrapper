all:
	go build -o acw
	/bin/cp ./acw ~/bin


zsh:
	./acw completion zsh > _acw
	cp _acw /Users/jiho.jung/.oh-my-zsh/completions/
