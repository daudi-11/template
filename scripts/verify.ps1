$ErrorActionPreference = 'Stop'
$repoRoot = Split-Path -Parent $PSScriptRoot
$backendPath = Join-Path $repoRoot 'backend'
$frontendPath = Join-Path $repoRoot 'frontend'
$originalLocation = Get-Location

try {
    Set-Location $backendPath
    Write-Host 'Running Go tests...'
    go test ./...
    if ($LASTEXITCODE -ne 0) {
        throw "Go tests failed with exit code $LASTEXITCODE."
    }

    Set-Location $frontendPath
    Write-Host 'Running frontend lint...'
    npm run lint
    if ($LASTEXITCODE -ne 0) {
        throw "Frontend lint failed with exit code $LASTEXITCODE."
    }

    Write-Host 'Building frontend...'
    npm run build
    if ($LASTEXITCODE -ne 0) {
        throw "Frontend build failed with exit code $LASTEXITCODE."
    }

    Write-Host 'All requested verification steps passed.'
}
finally {
    Set-Location $originalLocation
}
