$filebin = "zip-compactor";
$OSs = @('windows', 'linux', 'darwin')
$basePath = ".\bin\"
$plataforms = @('amd64','arm64')
$originalOS = go env GOOS
$originalARCH = go env GOARCH

Write-Output "Building your app..."
Write-Output "ATTENTION: if you stop before you finish you will have problems."
Write-Output "      SO, if that happens you need to restart the shell"

if (!(Test-Path -Path "go.mod")) {
   echo "You is a wrong path!!! Go to home dir."
   exit 1
}

if (Test-Path -Path $basePath) {
   Remove-Item -Force -Recurse -Confirm:$false $basePath
}

foreach ($os in $OSs) {
    foreach ($plataform in $plataforms) {
        $to = $basePath + $os + "-" + $plataform + ".zip"
		$binpath = $filebin

		if ($os -eq "windows") {
			$binpath += '.exe'
		}

		Write-Output "- Building $os/$plataform"

        # Build
		$env:GOOS = $os; $env:GOARCH = $plataform; go build -o $binpath .;
        # Zip to bin folder
        .\app.exe -from $binpath -to $to;
        # Delete file
        Remove-Item -Force -Recurse -Confirm:$false $binpath;
	}
}

# Restore initial values
$env:GOOS = $originalOS;
$env:GOARCH = $originalARCH;


Write-Output "Finished!";