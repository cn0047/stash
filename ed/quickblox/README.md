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
php -r "
define('USER_LOGIN', '');
define('USER_PASSWORD', '.');

define('QB_APP_SECRET', '');

$body = [
    'application_id' => '3',
    'auth_key' => '',
    'nonce' => time(),
    'timestamp' => time(),
    'user' => ['login' => USER_LOGIN, 'password' => USER_PASSWORD]
];
$built_query = urldecode(http_build_query($body));
$signature = hash_hmac('sha1', $built_query , QB_APP_SECRET);
$body['signature'] = $signature;
$post_body = http_build_query($body);
$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, 'https://apiziipr.quickblox.com/session.json');
curl_setopt($curl, CURLOPT_POST, true);
curl_setopt($curl, CURLOPT_POSTFIELDS, $post_body);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($curl);
echo json_decode($response, true)['session']['token'];
"
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
# user 1
cnfxlr+24
user_126566 wr07btLkzHqiP0oLy5Ve
export qbIdForUser1=357806
export tokenForUser1='42cf66556833bdb41275bcb51a392f8a9fb82c4c'

# user 2
cnfxlr+25
user_126567 XUyBPtN4YaN60YbOKEKm
export qbIdForUser2=357807
export tokenForUser2='44c9c9a2fd5f5de49e69e7f73355bf9a10bd35f5'

#######################################################################################################################

# get chats for user 1
curl -X GET -H "QB-Token: "$tokenForUser1 https://$host/chat/Dialog.json | grep name --color=always

# get chats for user 2
curl -X GET -H "QB-Token: "$tokenForUser2 https://$host/chat/Dialog.json | grep name --color=always

#######################################################################################################################

# user 1 create chat with user 2
# where `type=3` - private chat, `type=2` - group char.
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{"type": 3, "name": "u1_to_u2", "occupants_ids": "'$qbIdForUser2'"}' \
https://$host/chat/Dialog.json

# user 2 create chat with user 1
# where `type=3` - private chat, `type=2` - group char.
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser2 \
-d '{"type": 3, "name": "u2_to_u1", "occupants_ids": "'$qbIdForUser1'"}' \
https://$host/chat/Dialog.json

# !!! IMPORTANT
export chatId='5877815bf53b26456400a704'

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
"message": "msg1",
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
# where `type=3` - private chat, `type=2` - group char.
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{"type": 2, "name": "GROUP_u1_to_u2", "occupants_ids": "'$qbIdForUser2'"}' \
https://$host/chat/Dialog.json

# user 2 create GROUP chat with user 1
# where `type=3` - private chat, `type=2` - group char.
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser2 \
-d '{"type": 3, "name": "GROUP_u2_to_u1", "occupants_ids": "'$qbIdForUser1'"}' \
https://$host/chat/Dialog.json

# !!! IMPORTANT
export chatId='587783fff53b265bcd007449'

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

````
