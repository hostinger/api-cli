#!/usr/bin/env bash
set -euo pipefail

if ! command -v task >/dev/null; then
    echo "Installing Taskfile CLI helper..." >&2
    curl -1sLf 'https://dl.cloudsmith.io/public/task/task/setup.deb.sh' | sudo -E bash
    sudo apt install -y task
fi

sudo chown -R ${UID}:${UID} ${HOME}/.codex

# Fix project file permissions to standard privileges
sudo chown -R ${UID}:${UID} ${PWD}
sudo find ${PWD} -type d -exec chmod 755 {} +
sudo find ${PWD} -type f ! -name "*.sh" -exec chmod 644 {} +
sudo find ${PWD} -type f -name "*.sh" -exec chmod 755 {} +

# Setup git files configurations
sudo chown -R ${UID}:${UID} ${HOME}/.gitconfig-volume

touch ${HOME}/.gitconfig-volume/config
ln -fs ${HOME}/.gitconfig-volume/config ${HOME}/.gitconfig
sudo chown -R ${UID}:${UID} ${HOME}/.gitconfig

mkdir -p ${HOME}/.gitconfig-volume/.git-credentials
ln -fs ${HOME}/.gitconfig-volume/.git-credentials ${HOME}/.git-credentials
sudo chown -R ${UID}:${UID} ${HOME}/.git-credentials

git config --global --add safe.directory ${PWD}
