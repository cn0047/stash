jq
-

https://jqplay.org/

````sh
# prettify json output
echo '{"foo": "f", "bar": "b", "items": [1, 2, 3] }' | jq

echo '{"foo": "f", "bar": "b", "items": [1, 2, 3] }' | jq '.items | length'

cat data.json \
| jq '.gcp_price_list | del(.sustained_use_base,.sustained_use_tiers) | .[] | keys[]' \
| uniq

````
