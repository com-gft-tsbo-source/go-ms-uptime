param( ${clientID}, 
       ${ip}, 
       ${port},
       ${sleep}
    )

if ( [string]::IsNullOrEmpty($clientID)) { $clientID = ${random_id}=(-join ((48..57) + (97..122) | Get-Random -Count 16 | % {[char]$_})) }
if ( [string]::IsNullOrEmpty($ip))       { $ip = "ms-uptime.test.gft.com" }
if ( [string]::IsNullOrEmpty($port))     { $port = "8080" }
if ( [string]::IsNullOrEmpty($sleep))    { $sleep = "3" }

$uri = "http://${ip}:${port}/uptime"

echo "##############################################################################"
echo "Querying '${ip}:${port}' as '$clientID'."

while ($True)
{
    sleep $sleep
    try {
        $res = Invoke-WebRequest -Headers @{'cid' = $clientID }  -Uri "$uri" -ErrorAction SilentlyContinue -WarningAction SilentlyContinue
    } catch {
        echo "Could not connect to service!"
        continue
    }
    
    $res | ConvertFrom-Json | ConvertTo-Json
}