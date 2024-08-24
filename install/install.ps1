$ProjectDir = Get-Location
$OutputBinary = "mankka.exe"

Set-Location $ProjectDir

Write-Host "Building Go project..."
go build -o $OutputBinary

if ($LASTEXITCODE -eq 0) {
    Write-Host "Build succeeded. Binary created: $OutputBinary"
} else {
    Write-Host "Build failed."
    exit 1
}
