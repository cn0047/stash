jq
-

[doc](https://stedolan.github.io/jq/manual/)
[play](https://jqplay.org/)

````sh
--compact-output / -c:
--color-output / -C and --monochrome-output / -M
--sort-keys / -S
````

````sh
.foo.bar
.["foo::bar"] # special chars
.foo? # Optional object identifier
.[10:15] # Array/string slice
, # (comma) Two filters separated by a comma

.items[]|.id
curl ... | jq '.items[]|.ImageLink' # ✅
... | jq -r '.taskArns[0]' # string without quotes

# prettify json output
echo '{"foo": "f", "bar": "b", "items": [1, 2, 3] }' | jq

echo '{"foo": "f", "bar": "b", "items": [1, 2, 3] }' | jq '.items | length' # count

cat data.json \
| jq '.gcp_price_list | del(.sustained_use_base,.sustained_use_tiers) | .[] | keys[]' \
| uniq

````
