# Download sv.exe
Write-Host "Downloading sv-cli binary..."
$svUrl = "https://github.com/rohsyl/sv-cli/releases/latest/download/sv.exe"
$svPath = "$env:ProgramFiles\sv-cli\sv.exe"

# Create installation directory if it doesn't exist
New-Item -ItemType Directory -Force -Path "$env:ProgramFiles\sv-cli" | Out-Null

# Download the binary
Invoke-WebRequest -Uri $svUrl -OutFile $svPath

# Add sv-cli to the PATH
Write-Host "Adding sv-cli to PATH..."
$envPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
if ($envPath -notlike "*$env:ProgramFiles\sv-cli*") {
    [System.Environment]::SetEnvironmentVariable("Path", "$envPath;$env:ProgramFiles\sv-cli", [System.EnvironmentVariableTarget]::Machine)
}

Write-Host "Installation complete. Restart your terminal to use sv-cli."
