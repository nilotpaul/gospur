start: 
	@pnpm tailwindcss -i ./web/styles/globals.css -o ./public/styles.css --minify
	@node ./esbuild.config.js
	@go build -tags '!dev' -o bin/build
	@ENVIRONMENT=PRODUCTION ./bin/build

dev: build
	@ENVIRONMENT=DEVELOPMENT ./bin/build

build:
	@go build -tags 'dev' -o bin/build

watch:
	@wgo -debounce 30ms \
    -exit \
    -file=.go \
    -file=.html \
    node ./esbuild.config.js \
    :: pnpm tailwindcss -i ./web/styles/globals.css -o ./public/styles.css \
    :: ENVIRONMENT=DEVELOPMENT go build -tags 'dev' -o bin/build main.go \
    :: ENVIRONMENT=DEVELOPMENT ./bin/build \
    :: wgo -dir=public pnpm livereload public