# Seperate Client Workflows

**Examples of CI Pipeline workflows for seperate client approach.**

## Basic Github Action
```yaml
name: CI Build

on:
  push:
    tags:
      - "v*.*.*"

jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    strategy:
      matrix:
        node-version: [22]

    permissions:
      contents: read

    steps:
      - name: Checkout
        uses: actions/checkout@v5

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: ${{ matrix.node-version }}

      - name: Install and Build Web Assets
        working-directory: ./web
        run: |
          npm ci
          npm run build

      - name: Build Go Binary
        run: |
          go build -tags 'dev' -o bin/build

    # The binary is in `./bin/`. Do the rest...
```

**Note: This assumes the following :-**
- Triggers on tag push.
- Your frontend/client project is inside web directory.

## Docker + Github Actions
```yaml
name: CI Build

on:
  push:
    tags:
      - "v*.*.*"

jobs:
  build:
    name: Docker Build & Push
    runs-on: ubuntu-latest

    permissions:
      contents: read
      packages: write

    steps:
      - name: Checkout
        uses: actions/checkout@v5

      - name: Extract tag name without "v"
        id: tag
        run: |
          TAG=${GITHUB_REF##*/}
          TAG_WITHOUT_V=$(echo $TAG | sed 's/^v//')
          echo "TAG=$TAG_WITHOUT_V" >> $GITHUB_OUTPUT

      - name: Log in to GHCR
        uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Build and Push Docker Image
        run: |
          IMAGE_NAME=ghcr.io/${{ github.repository }}:${{ steps.tag.outputs.TAG }}
          docker build -t $IMAGE_NAME .
          docker push $IMAGE_NAME
```

**Note: This assumes the following :-**
- Triggers on tag push.
- You're using GHCR

## Dockerfile Example
```dockerfile
FROM node:22-alpine AS bundler

WORKDIR /app

# Copy entire project
COPY . .

# Go inside web dir and build client
WORKDIR /app/web
RUN npm ci && npm run build && rm -rf node_modules

FROM golang:1.25-alpine AS builder
ENV GO111MODULE=on

WORKDIR /app

# Copy project from bundler
COPY --from=bundler /app . 

RUN go mod download

RUN go build -tags '!dev' -o bin/build

FROM scratch

WORKDIR /app

# Copy the binary
COPY --from=builder --chown=1000:1000 /app/bin /app/bin

# Switch to non-root user
USER 1000:1000

# Environment variables
ENV ENVIRONMENT="PRODUCTION"
ENV PORT="3000"

EXPOSE $PORT

CMD ["./bin/build"]
```

**Note: This assumes the following :-**
- Triggers on tag push.
- Your frontend/client project is in web directory.

> **DO NOT forget to keep .dockerignore**