all:
	go build -o spc


zsh:
	./spc completion zsh > _spc
	cp _spc /Users/jiho.jung/.oh-my-zsh/completions/
