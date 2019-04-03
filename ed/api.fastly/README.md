Fastly
-

#### API

````sh
export key={key}
export serviceID={sID}
export version={vID}

# get versions
curl -X GET https://api.fastly.com/service/$serviceID/version \
  -H 'Fastly-Key: '$key | jq

# new dictionary
curl -X POST https://api.fastly.com/service/$serviceID/version/$version/dictionary \
  -H 'Fastly-Key: '$key
  -d 'name=referer_blacklist'

{
  "name": "referer_blacklist",
  "service_id": "5U9RuBXC0nn3iDVIjHeYyr",
  "version": 2,
  "deleted_at": null,
  "created_at": "2018-04-26T11:11:23Z",
  "write_only": false,
  "updated_at": "2018-04-26T11:11:23Z",
  "id": "4JZNSfCULIB2J1hIt5XTTk"
}

export dictionaryID={dID}

# get dictionaries
curl -X GET https://api.fastly.com/service/$serviceID/version/$version/dictionary \
  -H 'Fastly-Key: '$key | jq

# get dictionary items
curl -X GET https://api.fastly.com/service/$serviceID/dictionary/$dictionaryID/items \
  -H 'Fastly-Key: '$key | jq

# create new dictionary item (redirect rule)
curl -X POST https://api.fastly.com/service/$serviceID/dictionary/$dictionaryID/item \
  -H 'Fastly-Key: '$key
  -H 'Content-Type: application/json' \
  -d '{"item_key": "cn007b.tumblr.com/fr-test", "item_value": "https://cn007b.tumblr.com/fr"}'

{
  "dictionary_id": "4JZNSfCULIB2J1hIt5XTTk",
  "service_id": "5U9RuBXC0nn3iDVIjHeYyr",
  "item_key": "cn007b.tumblr.com/fr-test",
  "item_value": "https://cn007b.tumblr.com/fr",
  "created_at": "2018-04-26T11:32:35Z",
  "deleted_at": null,
  "updated_at": "2018-04-26T11:32:35Z"
}
````
