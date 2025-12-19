#!/usr/bin/env bash
set -euo pipefail

if ! command -v task >/dev/null; then
    echo "Installing Taskfile CLI helper..." >&2
    curl -1sLf 'https://dl.cloudsmith.io/public/task/task/setup.deb.sh' | sudo -E bash
    sudo apt install task
fi

sudo chown -R ${UID}:${UID} /home/vscode/.codex

if [ -f /tmp/gitconfig ]; then
    cp /tmp/gitconfig /home/vscode/.gitconfig
    sudo chown -R ${UID}:${UID} /home/vscode/.gitconfig
    # sudo chmod 644 /home/vscode/.gitconfig
fi

if [ -f /tmp/git-credential ]; then
    cp /tmp/git-credential /home/vscode/.git-credential
    sudo chown -R ${UID}:${UID} /home/vscode/.git-credential
    # sudo find /home/vscode/.git-credential -type d -exec chmod 755 {} +
    # sudo find /home/vscode/.git-credential -type f -exec chmod 644 {} +
fi

# Fix project file permissions to standard privileges
sudo chown -R ${UID}:${UID} ${PWD}
sudo find ${PWD} -type d -exec chmod 755 {} +
sudo find ${PWD} -type f -exec chmod 644 {} +

git config --global --add safe.directory ${PWD}
