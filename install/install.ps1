& {
    $ErrorActionPreference = 'Stop'
    $ProgressPreference = 'SilentlyContinue'
    [Net.ServicePointManager]::SecurityProtocol = [Net.ServicePointManager]::SecurityProtocol -bor [Net.SecurityProtocolType]::Tls12

    $binary = 'hostinger'
    $repo = 'hostinger/api-cli'
    $version = if ($env:HOSTINGER_VERSION) { $env:HOSTINGER_VERSION } else { '__INSTALL_VERSION__' }
    $installDir = if ($env:HOSTINGER_INSTALL_DIR) { $env:HOSTINGER_INSTALL_DIR } else { Join-Path $env:LOCALAPPDATA "Programs\$binary" }

    $processor = if ($env:PROCESSOR_ARCHITEW6432) { $env:PROCESSOR_ARCHITEW6432 } else { $env:PROCESSOR_ARCHITECTURE }
    $arch = switch ($processor) {
        'AMD64' { 'amd64' }
        'ARM64' { 'arm64' }
        'x86' { '386' }
        default { throw "Unsupported processor $processor." }
    }

    $tar = Join-Path $env:SystemRoot 'System32\tar.exe'
    if (-not (Test-Path $tar)) {
        throw 'tar.exe was not found. It ships with Windows 10 version 1803 and later.'
    }

    $version = $version.TrimStart('v')
    if ($version -notmatch '^\d') {
        $version = (Invoke-RestMethod "https://api.github.com/repos/$repo/releases/latest" -UseBasicParsing).tag_name.TrimStart('v')
    }

    $archive = "$binary-$version-windows-$arch.tar.gz"
    $checksums = "$binary-$version-checksums.sha256"
    $baseUrl = "https://github.com/$repo/releases/download/v$version"
    $tmp = Join-Path ([IO.Path]::GetTempPath()) ([Guid]::NewGuid().ToString())
    New-Item -ItemType Directory -Path $tmp | Out-Null

    try {
        Write-Host "Downloading $binary $version (windows/$arch)"
        $archivePath = Join-Path $tmp $archive
        $checksumsPath = Join-Path $tmp $checksums
        Invoke-WebRequest "$baseUrl/$archive" -OutFile $archivePath -UseBasicParsing
        Invoke-WebRequest "$baseUrl/$checksums" -OutFile $checksumsPath -UseBasicParsing

        $expected = foreach ($line in Get-Content $checksumsPath) {
            $hash, $name = $line -split '\s+', 2
            if ($name -eq $archive) { $hash }
        }
        if (-not $expected) {
            throw "No checksum published for $archive."
        }
        if ((Get-FileHash $archivePath -Algorithm SHA256).Hash -ne $expected) {
            throw "Checksum mismatch for $archive."
        }

        & $tar -xzf $archivePath -C $tmp "$binary.exe"
        if ($LASTEXITCODE -ne 0) {
            throw "Could not extract $binary.exe from $archive."
        }

        New-Item -ItemType Directory -Force -Path $installDir | Out-Null
        Get-ChildItem $installDir -Filter "$binary.*.old" | Remove-Item -Force -ErrorAction SilentlyContinue
        $installed = Join-Path $installDir "$binary.exe"
        if (Test-Path $installed) {
            Move-Item $installed (Join-Path $installDir "$binary.$([Guid]::NewGuid().ToString('N')).old")
        }
        Move-Item (Join-Path $tmp "$binary.exe") $installed
    }
    finally {
        Remove-Item -Recurse -Force $tmp -ErrorAction SilentlyContinue
    }

    $normalize = { param($dir) [Environment]::ExpandEnvironmentVariables($dir).TrimEnd('\') }
    $broadcastSettingChange = { [Environment]::SetEnvironmentVariable("$binary-install", $null, 'User') }
    $target = & $normalize $installDir
    $environmentKey = [Microsoft.Win32.Registry]::CurrentUser.OpenSubKey('Environment', $true)
    try {
        $userPath = $environmentKey.GetValue('Path', '', [Microsoft.Win32.RegistryValueOptions]::DoNotExpandEnvironmentNames)
        $entries = @($userPath -split ';' | Where-Object { $_ })
        if (-not ($entries | Where-Object { (& $normalize $_) -eq $target })) {
            $environmentKey.SetValue('Path', (($entries + $installDir) -join ';'), [Microsoft.Win32.RegistryValueKind]::ExpandString)
            & $broadcastSettingChange
        }
    }
    finally {
        $environmentKey.Close()
    }

    if (-not ($env:Path -split ';' | Where-Object { $_ -and (& $normalize $_) -eq $target })) {
        $env:Path = "$installDir;$env:Path"
    }

    Write-Host "Installed $binary $version to $installed"
    Write-Host ''
    Write-Host "Run '$binary --help' to get started."
}
