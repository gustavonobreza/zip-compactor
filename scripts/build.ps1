$filebin = "app.exe";

if (!(Test-Path -Path "go.mod")) {
   echo "You is a wrong path!!! Go to home dir."
   exit 1
}

go build -o $filebin .