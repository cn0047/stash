Quickblox
-

`
export host='apiziipr.quickblox.com'
`

#### Create new qb user

````bash
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: afba4149fd977ffd2f7b2ea6cf3535af9e785a7a" \
-d '{"user": {"login": "test_'`date +%s`'", "password": "LenaLena", "external_user_id": "'`date +%s`'"}}' \
https://$host/users.json
````

#### Get TOKEN

````
php -r '
$USER_LOGIN = "";
$USER_PASSWORD = "";
$QB_APPLICATION_ID = "";
$QB_AUTH_KEY = "";
$QB_APP_SECRET = "";
$body = [
    "application_id" => "3",
    "auth_key" => $QB_AUTH_KEY,
    "nonce" => time(),
    "timestamp" => time(),
    "user" => ["login" => $USER_LOGIN, "password" => $USER_PASSWORD]
];
$built_query = urldecode(http_build_query($body));
$signature = hash_hmac("sha1", $built_query , $QB_APP_SECRET);
$body["signature"] = $signature;
$post_body = http_build_query($body);
$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, "https://apiziipr.quickblox.com/session.json");
curl_setopt($curl, CURLOPT_POST, true);
curl_setopt($curl, CURLOPT_POSTFIELDS, $post_body);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($curl);
printf("Your token is: %s %s", json_decode($response, true)["session"]["token"], PHP_EOL);
'
````
````
# user 1
cnfxlr+27
user_1 p2
export qbIdForUser1=1
export tokenForUser1=''

# user 2
cnfxlr+1
user_2 p2
export qbIdForUser2=2
export tokenForUser2=''
````

#### Chat

Chat types: `type=3` - private chat, `type=2` - group char.

````bash
# get chats for user 1
curl -X GET -H "QB-Token: "$tokenForUser1 https://$host/chat/Dialog.json | grep name --color=always

# get chats for user 2
curl -X GET -H "QB-Token: "$tokenForUser2 https://$host/chat/Dialog.json | grep name --color=always

# get GROUP chats for user 1
curl -X GET -H "QB-Token: "$tokenForUser1 https://$host/chat/Dialog.json?type=2

#######################################################################################################################

# user 1 create chat with user 2
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{"type": 3, "name": "u1_to_u2", "occupants_ids": "'$qbIdForUser2'"}' \
https://$host/chat/Dialog.json

# user 2 create chat with user 1
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser2 \
-d '{"type": 3, "name": "u2_to_u1", "occupants_ids": "'$qbIdForUser1'"}' \
https://$host/chat/Dialog.json

# !!! IMPORTANT
export chatId='5758460ff53b264c440264c9'

#######################################################################################################################

# show messages for user 1
curl -X GET -H "QB-Token: "$tokenForUser1 https://$host/chat/Message.json?chat_dialog_id=$chatId

# show messages for user 2
curl -X GET -H "QB-Token: "$tokenForUser2 https://$host/chat/Message.json?chat_dialog_id=$chatId

#######################################################################################################################

# user 1 send message to user 2
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{
"chat_dialog_id": "'$chatId'",
"message": "itWkr$$$",
"recipient_id": '$qbIdForUser2'
}' \
https://$host/chat/Message.json

# user 2 send message to user 1
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser2 \
-d '{
"chat_dialog_id": "'$chatId'",
"message": "msg2",
"recipient_id": '$qbIdForUser1'
}' \
https://$host/chat/Message.json

#######################################################################################################################

# user 1 create GROUP chat with user 2
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{"type": 2, "name": "GROUP_u1_to_u2", "occupants_ids": "'$qbIdForUser2'"}' \
https://$host/chat/Dialog.json

# user 2 create GROUP chat with user 1
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser2 \
-d '{"type": 3, "name": "GROUP_u2_to_u1", "occupants_ids": "'$qbIdForUser1'"}' \
https://$host/chat/Dialog.json

# !!! IMPORTANT
export chatId='585d0f52f53b26f96c009700'

#######################################################################################################################

# user 1 send message to user 2 into GROUP chat
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{
"chat_dialog_id": "'$chatId'",
"message": "msg4",
"recipient_id": '$qbIdForUser2'
}' \
https://$host/chat/Message.json

#######################################################################################################################

# delete chats for user by chat ids
curl -X DELETE \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
https://$host/chat/Dialog/$chatId.json

````

#### Push

````
php -r '
$r = base64_encode(json_encode([
    "ios_content_available" => 1,
    "type_id" => 201105, /* Message from admin */
    "message" => "Push from cli.",
]));
printf("Your message is: %s %s", $r, PHP_EOL);
'
````
````bash
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{"event": {
    "notification_type": "push",
    "environment": "production",
    "user": {"ids": "'$qbIdForUser2'"},
    "message": "eyJpb3NfY29udGVudF9hdmFpbGFibGUiOjEsInR5cGVfaWQiOjIwMTEwNSwibWVzc2FnZSI6IlB1c2ggZnJvbSBjbGkgKiBsYXN0IG9uZS4ifQ=="
}}' \
https://$host/events.json
````
