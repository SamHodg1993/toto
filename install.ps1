# Windows installation script for toto

# Build the binary with the tool name "toto"
Write-Host "Building toto application..."
go build -o toto.exe

# Create directory for the app in user profile if it doesn't exist
$appDir = "$env:USERPROFILE\toto"
if (-not (Test-Path -Path $appDir)) {
    New-Item -ItemType Directory -Path $appDir | Out-Null
    Write-Host "Created application directory: $appDir"
}

# Move the binary to the app directory
Move-Item -Path "toto.exe" -Destination $appDir -Force
Write-Host "Moved toto.exe to $appDir"

# Check if database exists and move it, or create a new one
$dbPath = "$env:USERPROFILE\.toto.db"
if (Test-Path -Path "todo.db") {
    Move-Item -Path "todo.db" -Destination $dbPath -Force
    Write-Host "Moved existing database to $dbPath"
} else {
    # Create an empty database file
    New-Item -ItemType File -Path $dbPath -Force | Out-Null
    Write-Host "Created new database at $dbPath"
}

# Add the app directory to the user's PATH
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (-not $currentPath.Contains($appDir)) {
    [Environment]::SetEnvironmentVariable("Path", "$currentPath;$appDir", "User")
    Write-Host "Added $appDir to your PATH"
    
    # Also update the current session's PATH
    $env:Path = "$env:Path;$appDir"
} else {
    Write-Host "$appDir is already in your PATH"
}

Write-Host "Installation of the toto todo package has been completed!" -ForegroundColor Green
Write-Host "You may need to restart any open terminal windows to use the 'toto' command." -ForegroundColor Yellow
Write-Host "Enjoy!" -ForegroundColor Green
