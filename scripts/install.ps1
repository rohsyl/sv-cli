# Check if Go is installed
if (-not (Get-Command "go" -ErrorAction SilentlyContinue)) {
    Write-Host "Go is not installed. Installing Go..."

    # Determine the latest Go version
    $goUrl = "https://go.dev/dl/go1.21.1.windows-amd64.msi"
    $msiPath = "$env:TEMP\go.msi"

    # Download the MSI installer
    Invoke-WebRequest -Uri $goUrl -OutFile $msiPath

    # Install Go
    Start-Process msiexec.exe -ArgumentList "/i `"$msiPath`" /quiet /norestart" -Wait
    Remove-Item $msiPath

    Write-Host "Go has been installed."
} else {
    Write-Host "Go is already installed."
}

# Download sv.exe
Write-Host "Downloading sv-cli binary..."
$svUrl = "https://github.com/rohsyl/sv-cli/releases/download/latest/sv.exe"
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
