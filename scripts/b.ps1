$filebin = "app.exe";
$env:ZIP_COMPRESSOR_VERSION = git describe --abbrev=0 --tags

if (!(Test-Path -Path "go.mod")) {
   echo "You is a wrong path!!! Go to home dir."
   exit 1
}

$originalOS = go env GOHOSTOS;
$originalARCH = go env GOHOSTARCH;
$env:GOOS = $originalOS; $env:GOARCH = $originalARCH;

go install -ldflags "-X main.Version=$env:ZIP_COMPRESSOR_VERSION" .;

Write-Output $env:ZIP_COMPRESSOR_VERSION