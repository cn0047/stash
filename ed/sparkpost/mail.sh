curl -XPOST \
https://api.sparkpost.com/api/v1/transmissions \
-H "Authorization: <YOUR API KEY>" \
-H "Content-Type: application/json" \
-d '{
"content": {"from": "", "subject": "", "text":""},
"recipients": [{"address": ""}]
}'
