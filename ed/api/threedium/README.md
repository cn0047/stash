Threedium
-

[docs](https://documentation.threedium.io/)

````sh
sk="secret-key: value"
pk="public-key: value"

h='https://cloud.unlimited3d.com'
j='Content-Type: application/json'

curl -sS -H $j -H $sk -H $pk $h'/api/thirdParty/projects/' | jq
curl -sS -H $j -H $sk -H $pk $h'/api/thirdParty/solutions/' | jq

````
