# Supernatural Frontend

React + Vite single-page application for the Supernatural platform.

## Running locally (development)

```bash
npm install
npm run dev
```

Opens a dev server (default: `http://localhost:5173`) with hot module reload.

## Building for production

```bash
npm run build
```

## Running with Docker

This directory includes a `Dockerfile` for the frontend service.
It's a multi-stage build: Node builds the static bundle, then an
`nginx:alpine` image serves it as a non-root user on port `8080`.

```bash
# Build the image
docker build -t supernatural-frontend .

# Run it (nginx listens on 8080 inside the container)
docker run -d -p 8080:8080 --name supernatural-frontend supernatural-frontend

# Open in browser
open http://localhost:8080

# Stop and remove when done
docker stop supernatural-frontend && docker rm supernatural-frontend
```
