# Development Usage

## Commands

```Makefile
start: 
	@node ./esbuild.config.js
	@go build -tags '!dev' -o bin/build
	@ENVIRONMENT=PRODUCTION ./bin/build

build:
	@node ./esbuild.config.js
	@go build -tags 'dev' -o bin/build

dev:
	@wgo \
    -exit \
    -file=.go \
    -file=.html \
	-file=.css \
	-xdir=public \
	go build -tags 'dev' -o bin/build . \
    :: ENVIRONMENT=DEVELOPMENT ./bin/build \
    :: wgo -xdir=bin -xdir=node_modules -xdir=public node ./esbuild.config.js \
	:: wgo -dir=node_modules npx livereload -w 800 -ee go .
```

These are the default development commands which will be pre-configured for you.

**These are specific to Linux only.**

### For Windows

Please use git bash instead of command prompt or powershell and use the same `Makefile` above.

### If Auto Browser Reload Feels Slow

Change the delay time (ms) of the command, default will be 800ms.

```sh
wgo -dir=node_modules npx livereload -w 800 -ee go .
```

## Environment Variables

All the configurations will be done for you.

Load env vars either by:
- Creating a `.env` file.
- Using runtime injected ones.

Please visit your `config/env.go` to configure further.

# Docs

- [godotenv](https://github.com/joho/godotenv#godotenv--)
