$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $MyInvocation.MyCommand.Path
chcp 65001 | Out-Null
[Console]::OutputEncoding = [Text.UTF8Encoding]::new()
$OutputEncoding = [Text.UTF8Encoding]::new()

function Write-Section {
    param([string]$title)
    $line = "-" * 60
    Write-Host $line
    Write-Host $title
    Write-Host $line
}

function Run-InDir {
    param([string]$dir, [scriptblock]$script)
    Push-Location $dir
    try {
        & $script
    } finally {
        Pop-Location
    }
}

function Invoke-GoTest {
    param(
        [string]$dir,
        [string[]]$TestArgs = @(),
        [switch]$LegacyMode
    )

    Run-InDir $dir {
        $prev = $env:GO111MODULE
        $prevCache = $env:GOCACHE
        try {
            if ($LegacyMode) {
                $env:GO111MODULE = "off"
            }
            $cacheDir = Join-Path $env:TEMP "go_oop_gocache"
            New-Item -ItemType Directory -Path $cacheDir -Force | Out-Null
            $env:GOCACHE = $cacheDir
            & go clean -testcache

            if ($TestArgs.Count -gt 0) {
                & go test -count=1 @TestArgs
            } else {
                & go test -count=1
            }

            if ($LASTEXITCODE -ne 0) {
                throw "go test failed in $dir"
            }
        } finally {
            $env:GO111MODULE = $prev
            $env:GOCACHE = $prevCache
        }
    }
}

Write-Section "Tests: go test by tasks"

Write-Section "Root module"
Invoke-GoTest $root -TestArgs @("./0_all")
Invoke-GoTest $root -TestArgs @("./1_struct")
Invoke-GoTest $root -TestArgs @("./2_method/...")
Invoke-GoTest $root -TestArgs @("./5_interface")
Invoke-GoTest $root -TestArgs @("./6_alignment")

Write-Section "Module 3_encapsulation"
Invoke-GoTest "$root\\3_encapsulation" -TestArgs @("./...")

Write-Section "Module 4_composition"
$composeDir = Get-ChildItem -Path $root -Directory |
    Where-Object { $_.Name -like "4_*" -and (Test-Path (Join-Path $_.FullName "go.mod")) } |
    Select-Object -First 1 -ExpandProperty FullName

if (-not $composeDir) {
    throw "4_composition module directory not found."
}

Invoke-GoTest $composeDir -TestArgs @("./...")

Write-Section "Done"
Write-Host "All tests passed."
