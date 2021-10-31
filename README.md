<img align="center" src="./docs/zip.webp" alt="Zip image"/>

<br>

## Usage

### Install (CLI USAGE)
```bash
go install github.com/gustavonobreza/zip-compactor
```
### Run in CLI (Can select just one file to be zipped) 
```bash
zip-compactor -from <File to be zipped> -to <Target path>
```

### Run in GUI (Can select many files)
```bash
git clone https://github.com/gustavonobreza/zip-compactor
cd zip-compactor
go install github.com/akavel/rsrc
rsrc -ico ".\docs\zip.ico"
go build -ldflags -H=windowsgui .
```
open the folder **zip-compactor** and copy executable file to your Desktop for example, or add in Start Menu (Windows)



