gen:
	@templ generate
init:
	@templ generate
	@go mod tidy
	@npm install --prefix ./typescript
run: 
	@templ generate
	@npm run build --prefix ./typescript
	@go run ./cmd $(ARGS)
build:
	@templ generate
	@npm run build --prefix ./typescript
	@go build -o ./bin ./cmd
