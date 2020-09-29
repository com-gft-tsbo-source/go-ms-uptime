param( ${clientID}, 
       ${ip}, 
       ${port}
    )

if ( [string]::IsNullOrEmpty($clientID)) { $clientID = ${random_id}=(-join ((48..57) + (97..122) | Get-Random -Count 16 | % {[char]$_})) }
if ( [string]::IsNullOrEmpty($ip))       { $ip = "ms-uptime" }
if ( [string]::IsNullOrEmpty($port))     { $port = "8080" }

$uri = "http://${ip}:${port}/uptime"

try {
    $res = Invoke-WebRequest -Headers @{'cid' = $clientID }  "$uri" -ErrorAction SilentlyContinue -WarningAction SilentlyContinue
} catch {
    echo "Could not connect to service!"
    exit 0
}

$res | ConvertFrom-Json | ConvertTo-Json
