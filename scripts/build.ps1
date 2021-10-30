$filebin = "app.exe";


if (!(Test-Path -Path "go.mod")) {
   echo "You is a wrong path!!! Go to home dir."
   exit 1
}

go install 'github.com/akavel/rsrc'

rsrc -icon .\zip.ico

go build -o $filebin .