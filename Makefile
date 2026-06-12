NAME = push-swap

all:
	@echo "Building $(NAME)..."
	@go build -o $(NAME) exec/Main.go
	
clean:
	@echo "Cleaning up..."
	@rm -f $(NAME)