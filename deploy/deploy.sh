#!/usr/bin/env bash
set -euo pipefail

REPO_DIR="${REPO_DIR:-$(pwd)}"
BRANCH="${BRANCH:-main}"

cd "$REPO_DIR"
git config --global --add safe.directory "$REPO_DIR" || true
git fetch --all --prune
git reset --hard "origin/$BRANCH"

if [ -d "$REPO_DIR/frontend" ]; then
  cd "$REPO_DIR/frontend"
  if command -v npm >/dev/null 2>&1; then
    npm ci
    npm run build
  else
    echo "npm not found" >&2
    exit 1
  fi
fi

cd "$REPO_DIR"
if command -v docker >/dev/null 2>&1; then
  if docker compose version >/dev/null 2>&1; then
    docker compose up -d --build
  elif command -v docker-compose >/dev/null 2>&1; then
    docker-compose up -d --build
  else
    echo "docker compose not available" >&2
    exit 1
  fi
else
  echo "docker not found" >&2
  exit 1
fi

echo "deploy done"

