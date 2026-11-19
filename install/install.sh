#!/bin/sh
set -eu

binary="hostinger"
repo="hostinger/api-cli"
version="${HOSTINGER_VERSION:-__INSTALL_VERSION__}"
install_dir="${HOSTINGER_INSTALL_DIR:-$HOME/.local/bin}"

fail() {
	printf 'error: %s\n' "$1" >&2
	exit 1
}

download() {
	if command -v curl >/dev/null 2>&1; then
		curl -fsSL --retry 3 -o "$2" "$1" || fail "could not download $1"
	elif command -v wget >/dev/null 2>&1; then
		wget -q -O "$2" "$1" || fail "could not download $1"
	else
		fail "curl or wget is required"
	fi
}

main() {
	if command -v sha256sum >/dev/null 2>&1; then
		sha256_cmd="sha256sum"
	elif command -v shasum >/dev/null 2>&1; then
		sha256_cmd="shasum -a 256"
	else
		fail "sha256sum or shasum is required"
	fi

	case "$(uname -s)" in
	Linux) os=linux ;;
	Darwin) os=darwin ;;
	*) fail "unsupported operating system $(uname -s). On Windows, run install.ps1 in PowerShell" ;;
	esac

	case "$(uname -m)" in
	x86_64 | amd64) arch=amd64 ;;
	arm64 | aarch64) arch=arm64 ;;
	i386 | i686) arch=386 ;;
	*) fail "unsupported processor $(uname -m)" ;;
	esac

	tmp="$(mktemp -d)"
	trap 'rm -rf "$tmp"' EXIT

	version="${version#v}"
	case "$version" in
	[0-9]*) ;;
	*)
		download "https://api.github.com/repos/$repo/releases/latest" "$tmp/latest.json"
		version="$(sed -n 's/.*"tag_name": *"v\{0,1\}\([^"]*\)".*/\1/p' "$tmp/latest.json" | head -n 1)"
		[ -n "$version" ] || fail "could not find the latest $binary release"
		;;
	esac

	archive="$binary-$version-$os-$arch.tar.gz"
	checksums="$binary-$version-checksums.sha256"
	base_url="https://github.com/$repo/releases/download/v$version"

	printf 'Downloading %s %s (%s/%s)\n' "$binary" "$version" "$os" "$arch"
	download "$base_url/$archive" "$tmp/$archive"
	download "$base_url/$checksums" "$tmp/$checksums"

	expected="$(awk -v name="$archive" '$2 == name { print $1 }' "$tmp/$checksums")"
	[ -n "$expected" ] || fail "no checksum published for $archive"
	actual="$($sha256_cmd "$tmp/$archive" | cut -d ' ' -f 1)"
	[ "$actual" = "$expected" ] || fail "checksum mismatch for $archive"

	tar -xzf "$tmp/$archive" -C "$tmp" "$binary"
	mkdir -p "$install_dir"
	cp "$tmp/$binary" "$install_dir/.$binary.new"
	chmod 755 "$install_dir/.$binary.new"
	mv -f "$install_dir/.$binary.new" "$install_dir/$binary"

	printf 'Installed %s %s to %s\n' "$binary" "$version" "$install_dir/$binary"

	case ":$PATH:" in
	*":$install_dir:"*) ;;
	*)
		printf '\n%s is not on your PATH. Add this line to your shell profile, then open a new terminal:\n' "$install_dir"
		printf "  export PATH=\"%s:\$PATH\"\n" "$install_dir"
		;;
	esac

	printf "\nRun '%s --help' to get started.\n" "$binary"
}

main "$@"
