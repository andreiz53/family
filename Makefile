install:
	@npm install

css:
	@npx tailwindcss -i ./input.css -o ./web/assets/css/style.css --watch