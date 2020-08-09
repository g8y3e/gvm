# gvm
## Installation
 - Download version for your architecture:
    - amd64: https://github.com/g8y3e/gvm/releases/download/1.0.0-amd64/gvm.exe
    - 386: https://github.com/g8y3e/gvm/releases/download/1.0.0-386/gvm.exe
 - Copy it to destination folder
 - Add this folder in PATH environment variable
 
## Usage
 - Set GVM root folder running command `gvm root <path>` - this folder will be used for your Go versions 
 - Chose version of Go that you want to install from list - with command `gvm list -g` 
 - Install a Go version with `gvm install <version>`
 - And use it with `gvm use <version>`