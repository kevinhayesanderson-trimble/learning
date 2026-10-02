#!/usr/bin/env bash
set -e

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
cd "$DIR"

echo "==> Checking repository status..."
git fetch origin

if [ -n "$(git status --porcelain)" ]; then
    MSG="${1:-Sync learnings [$(hostname)] at $(date '+%Y-%m-%d %H:%M:%S')}"
    echo "==> Staging and committing local changes: '$MSG'..."
    git add -A
    git commit -m "$MSG"
else
    echo "==> No uncommitted local changes."
fi

echo "==> Pulling latest changes from origin/main..."
git pull --rebase --autostash origin main

echo "==> Pushing to origin/main..."
git push origin main

echo "[SUCCESS] Synchronized successfully with GitHub!"
