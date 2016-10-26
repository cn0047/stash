Quickblox
-

`
export host='apiziipr.quickblox.com'
`

### Create new qb user

````bash
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: afba4149fd977ffd2f7b2ea6cf3535af9e785a7a" \
-d '{"user": {"login": "test_'`date +%s`'", "password": "LenaLena", "external_user_id": "'`date +%s`'"}}' \
https://$host/users.json
````

#### Push

````bash
curl -X POST \
-H "Content-Type: application/json" \
-H "QuickBlox-REST-API-Version: 0.1.1" \
-H "QB-Token: ad7b695025ad11a068fdb3459f9e5ce1c5a7ce60" \
-d '{"event": {
    "notification_type": "push",
    "environment": "production",
    "user": { "ids": "177914"},
    "message": "MDA3IHBpbmcK"
}}' \
https://$host/events.json
````

#### Chat

````bash
4 qb:165178(user_4:8HoLePxFJqN4l2ZyRaca)   a3c05a224ebd06208523b04cedaafaa5655a94d7
17891 qb:167555(user_17891:tJmi7hLtknPFKibngp5C)   ea29f867fa09452de973f84fe5a733275999ea14
admin                                              09db68bab40795cf700f37e32d96d9ed9f60521a

# get dialogs for user 17891
curl -X GET -H "QB-Token: ea29f867fa09452de973f84fe5a733275999ea14" https://$host/chat/Dialog

# get messages from dialog
curl -X GET -H "QB-Token: ea29f867fa09452de973f84fe5a733275999ea14" \
https://$host/chat/Message?chat_dialog_id=56e03c6bf53b265e7d001aa1

# send message from 17891 to 4
curl https://$host/chat/Message.json \
-H "QB-Token: ea29f867fa09452de973f84fe5a733275999ea14" \
-H "QuickBlox-REST-API-Version: 0.1.0" \
-H "Content-Type: application/json" \
-XPOST -H 'application/json' -d '{
"chat_dialog_id": "56e03c6bf53b265e7d001aa1",
"message": "cli test",
"recipient_id": "165178"
}'

# user 4 send reply to user 17891
curl https://$host/chat/Message.json \
-H "QB-Token: a3c05a224ebd06208523b04cedaafaa5655a94d7" \
-H "QuickBlox-REST-API-Version: 0.1.0" \
-H "Content-Type: application/json" \
-XPOST -H 'application/json' -d '{
"chat_dialog_id": "56e03c6bf53b265e7d001aa1",
"message": "reply to your cli test",
"recipient_id": "167555"
}'

# create direct chat
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: 5829415b1d81dd70d9432e3deee889d2431cde92" \
-d '{"type": 3, "name": "Chat_to_203379", "occupants_ids": "203379"}' \
https://$host/chat/Dialog

# user 18039 send direct message to 17891
curl https://$host/chat/Message.json \
-H "QB-Token: a7353d479ca0c496de1a2fb9b2933ea7742ec8fb" \
-H "QuickBlox-REST-API-Version: 0.1.0" \
-H "Content-Type: application/json" \
-XPOST -H 'application/json' -d '{
"chat_dialog_id": "56e04642f53b265e7d001b83",
"message": "BOLD cli test",
"recipient_id": "167555"
}'
````
