jq
-

````sh
# prettify json output
echo '{"foo": "f", "bar": "b", "items": [1, 2, 3] }' | jq

echo '{"foo": "f", "bar": "b", "items": [1, 2, 3] }' | jq '.items | length'
````
