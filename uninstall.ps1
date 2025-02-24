# Windows uninstall script for toto

# Define paths
$appDir = "$env:USERPROFILE\toto"
$exePath = "$appDir\toto.exe"
$dbPath = "$env:USERPROFILE\.toto.db"

# Remove the executable
if (Test-Path -Path $exePath) {
    Remove-Item -Path $exePath -Force
    Write-Host "Removed toto.exe"
} else {
    Write-Host "toto.exe not found at $exePath"
}

# Ask about removing the database
$removeDB = Read-Host "Do you want to remove all your todos (delete $dbPath)? (y/N)"
if ($removeDB -eq "y" -or $removeDB -eq "Y") {
    if (Test-Path -Path $dbPath) {
        Remove-Item -Path $dbPath -Force
        Write-Host "Todos database removed"
    } else {
        Write-Host "Database not found at $dbPath"
    }
} else {
    Write-Host "Keeping todos database at $dbPath"
}

# Remove the app directory if it's empty
if (Test-Path -Path $appDir) {
    $isEmpty = (Get-ChildItem -Path $appDir -Force | Measure-Object).Count -eq 0
    if ($isEmpty) {
        Remove-Item -Path $appDir -Force
        Write-Host "Removed empty application directory"
    } else {
        Write-Host "Not removing application directory as it still contains files"
    }
}

# Remove the app directory from PATH
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($currentPath.Contains($appDir)) {
    # Replace the app directory and any trailing semicolon if present
    $newPath = $currentPath -replace [regex]::Escape("$appDir;"), ""
    $newPath = $newPath -replace [regex]::Escape(";$appDir"), ""
    $newPath = $newPath -replace [regex]::Escape("$appDir"), ""
    
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "Removed $appDir from your PATH environment variable"
}

Write-Host "Uninstall complete! You may want to restart your terminal or Command Prompt windows" -ForegroundColor Green
