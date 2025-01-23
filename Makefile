install:
	@npm install

css:
	@npx @tailwindcss/cli -i ./input.css -o ./web/assets/css/style.css --watch

seed:
	@go run ./cmd/seed/seed.go
