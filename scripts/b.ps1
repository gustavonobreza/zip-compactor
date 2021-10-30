$filebin = "app.exe";

if (!(Test-Path -Path "go.mod")) {
   echo "You is a wrong path!!! Go to home dir."
   exit 1
}

$originalOS = go env GOHOSTOS;
$originalARCH = go env GOHOSTARCH;
$env:GOOS = $originalOS; $env:GOARCH = $originalARCH;

go install .;