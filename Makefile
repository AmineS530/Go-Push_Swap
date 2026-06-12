NAME = push-swap

all: $(NAME) checker

$(NAME): 
	@echo "Building $(NAME)..."
	@go build -o $(NAME) exec/push-swap.cmd/Main.go

checker:
	@echo "Building checker..."
	@go build -o checker exec/checker.cmd/Main.go

clean:
	@echo "Cleaning up..."
	@rm -f $(NAME) checker

re: clean $(NAME) checker

.phony: $(NAME) checker clean