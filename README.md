# gvm
## Installation
 - Download version for your architecture:
    - amd64: https://github.com/g8y3e/gvm/releases/download/1.0.0-amd64/gvm.exe
    - 386: https://github.com/g8y3e/gvm/releases/download/1.0.0-386/gvm.exe
 - Copy it to destination folder
 - Add this folder in PATH environment variable
 
## Usage
 - Set GVM root folder running command `gvm root <path>` - this folder will be used for installing new versions of Go
 - Chose version of Go that you want to install for list - wtih command `gvm list -g` 
 - Install Go version with `gvm install <version>`
 - And use it with `gvm use <version>`