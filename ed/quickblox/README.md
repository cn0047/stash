Quickblox
-

````
export apiHost='apiziipr.quickblox.com'
export chatHost='chatziipr.quickblox.com'
export host=$apiHost
````

#### Create new qb user

````bash
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: afba4149fd977ffd2f7b2ea6cf3535af9e785a7a" \
-d '{"user": {"login": "test_'`date +%s`'", "password": "12345", "external_user_id": "'`date +%s`'"}}' \
https://$apiHost/users.json
````

#### Get TOKEN
https://screencast.com/t/K0CqTyWht
````
php -r '
$USER_LOGIN = "";
$USER_PASSWORD = "";
$QB_APPLICATION_ID = "";
$QB_AUTH_KEY = "";
$QB_APP_SECRET = "";
$body = [
    "application_id" => $QB_APPLICATION_ID,
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
curl_setopt($curl, CURLOPT_HEADER, false);
$response = curl_exec($curl);
$responseHttpCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);
//var_dump($responseHttpCode, $response);
printf("Your token is: %s %s", json_decode($response, true)["session"]["token"], PHP_EOL);
'
````
````
# user 1 admin@ziipr
$USER_LOGIN = "";
$USER_PASSWORD = "";
export qbIdForUser1=
export tokenForUser1=''

# user 1 TechTeam
$USER_LOGIN = "";
$USER_PASSWORD = "";
export qbIdForUser1=
export tokenForUser1=''

# user 2 ggg
$USER_LOGIN = "";
$USER_PASSWORD = "";
export qbIdForUser2=
export tokenForUser2=''

# user 3 cnfxlr
$USER_LOGIN = "";
$USER_PASSWORD = "";
export qbIdForUser3=
export tokenForUser3=''

# user 2
user_130873 
export qbIdForUser2=
export tokenForUser2=''

# user 3
user_130874 
export qbIdForUser3=
export tokenForUser3=''
````

#### Chat

Chat types: `type=3` - PRIVATE chat, `type=2` - GROUP chat, `type=1` - PUBLIC_GROUP chat.

````bash
# get chats for user
curl -X GET -H "QB-Token: "$tokenForUser1 https://$host/chat/Dialog.json
curl -X GET -H "QB-Token: "$tokenForUser1 https://$host/chat/Dialog.json | grep name --color=always
curl -X GET -H "QB-Token: "$tokenForUser2 https://$host/chat/Dialog.json?type=2
curl -X GET -H "QB-Token: "$tokenForUser3 https://$host/chat/Dialog.json?name=chat+between+ziipr+admin+and+367773

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
export chatId='587cfafdf53b26701d012d41'

-----------------------------------------------------------------------------------------------------------------------

# gets unread messages count
curl -X GET -H "QB-Token: "$tokenForUser1 \
https://$host/chat/Message/unread.json?chat_dialog_ids=$chatId

-----------------------------------------------------------------------------------------------------------------------

# show messages for user from chat
curl -X GET -H "QB-Token: "$tokenForUser1 https://$host/chat/Message.json?chat_dialog_id=$chatId

-----------------------------------------------------------------------------------------------------------------------

# user 1 send message to user 2 DIRECTLY
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{
"message": "msg5",
"recipient_id": '$qbIdForUser2'
}' \
https://$host/chat/Message.json

# user 2 send message to user 1 DIRECTLY
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser2 \
-d '{
"message": "msg6",
"recipient_id": '$qbIdForUser1'
}' \
https://$host/chat/Message.json

# user 1 send message to user 2 into certain chat
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{
"chat_dialog_id": "'$chatId'",
"message": "itWkr$ 💎",
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

-----------------------------------------------------------------------------------------------------------------------

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

# user 1 (ADMIN) create GROUP chat (chat with admin) with user 2
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{"type": 2, "name": "chat between ziipr admin and '$qbIdForUser2'", "occupants_ids": "'$qbIdForUser2'"}' \
https://$host/chat/Dialog.json

# !!! IMPORTANT
export chatId='585d4a2bf53b2642b80d3619,585d4a17f53b26b847015277,585d4a21f53b26b8470152a7,585d4a08f53b26f96c014f3d,585d49dcf53b26827405dfd0,585d49e9f53b26e465015d14,585d49d2f53b26138f013131,585d49c0f53b26bd3e014d05,585d49c6f53b26b847015102,585d4993f53b26f8c003935c,585d497ff53b26827405de50,585d497df53b26f8c0039302,585d495ef53b26332501567e,585d493df53b26652f039238,585d4927f53b26f6130149aa,585d4912f53b2642b80d3132,585d4915f53b26138f012e1c,585d48faf53b26bd3e0149f9,585d48fdf53b26f6130148da,585d48d5f53b262047015a30,585d48bef53b26f96c0149bc,585d4889f53b26f96c0148dd,585d4880f53b26652f038f5c,585d4876f53b26827405da2b,585d4866f53b261e3f038273,585d4849f53b26138f012b0c,585d482cf53b26138f012a8d,585d4831f53b2626c5060772,585d481cf53b26138f012a51,585d481df53b26f61301452e,585d4800f53b26b32a014f40,585d47f5f53b26138f0129af,585d514af53b26f96c016fe1,585d5151f53b2633250179db,585d5134f53b26652f03b645,585d510ff53b2642b80d547b,585d5112f53b262047017fa1,585d50fcf53b26e465017de8,585d50d0f53b2626c5062e18,585d50bdf53b26f6d001788a,585d50a0f53b26f8c003b3b0,585d5086f53b26827405fdba,585d506ef53b26b32a0174a0,585d5079f53b262047017d06,585d5053f53b26138f014e23,585d505df53b26b32a017449,585d504df53b26652f03b224,585d503ef53b26e465017a61,585d5016f53b2642b80d4fd6,585d5022f53b26332501749f,585d4ff1f53b26b847016cfe,585d4fdbf53b26f6d001740a,585d4fd5f53b26f6d00173e1,585d4f82f53b26b847016aa3,585d4f84f53b26b847016aab,585d4f66f53b26e46501760c,585d4f50f53b26f96c01669c,585d4f3ef53b2642b80d4c01,585d4f32f53b262047017732,585d4f34f53b26bd3e01655a,585d4f2bf53b26e4650174f8,585d4f18f53b2620470176a1,585d4edaf53b26e465017380,585d4edcf53b26f96c0164dd,585d4ee2f53b263325016f25,585d4ea6f53b26652f03aab7,585d4eb4f53b26652f03aafa,585d4e92f53b2626c5062351,585d4e84f53b26f96c01631f,585d4e74f53b26652f03a9b2,585d4e4bf53b2620470172fa,585d4e32f53b26f6d0016cdc,585d4e36f53b26bd3e0160fb,585d4e20f53b26138f014449,585d4e28f53b26652f03a8ac,585d4e0df53b26b8470163fe,585d4dfdf53b26827405f283,585d4dd1f53b26e465016f05,585d4de4f53b26e465016f64,585d4dc7f53b26827405f189,585d4dcff53b26b8470162e6,585d4db0f53b2633250169f2,585d4d97f53b26f8c003a53f,585d4d9af53b26b32a0167bb,585d4d6bf53b261e3f039855,585d4d46f53b2642b80d43a7,585d4d55f53b26b32a0166aa,585d4d37f53b2642b80d435a,585d4d1ef53b26f613015be0,585d4d0ff53b26652f03a3dd,585d4d01f53b2626c5061c57,585d4d03f53b2642b80d4286,585d4cf3f53b26827405edd6,585d4cd6f53b26b32a0164a7,585d4cd8f53b26f8c003a24e,585d4cbbf53b26b32a016435,585d4cb2f53b26bd3e015af1,585d4c9bf53b26b32a01638f,585d4c81f53b261e3f03949e,585d4c79f53b262047016b05'

-----------------------------------------------------------------------------------------------------------------------

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

-----------------------------------------------------------------------------------------------------------------------

# delete chats for user by chat ids
curl -X DELETE \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser2 \
https://$host/chat/Dialog/$chatId.json?force=0

#######################################################################################################################

# user 1 create GROUP chat with user 2 & user 3
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{"type": 2, "name": "GROUP_CHAT_FOR_TT_AND_GGG_AND_FL", "occupants_ids": "'$qbIdForUser2,$qbIdForUser3'"}' \
https://$host/chat/Dialog.json

export chatId='589abcfff53b26ee88308e2b'

-----------------------------------------------------------------------------------------------------------------------

# user 1 send message into GROUP chat
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
-d '{
"chat_dialog_id": "'$chatId'",
"message": "I have inited this chat for us."
}' \
https://$host/chat/Message.json

# user 2 send message into GROUP chat
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser2 \
-d '{
"chat_dialog_id": "'$chatId'",
"message": "ty"
}' \
https://$host/chat/Message.json

# user 3 send message into GROUP chat
curl -X POST \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser3 \
-d '{
"chat_dialog_id": "'$chatId'",
"message": "me two"
}' \
https://$host/chat/Message.json

-----------------------------------------------------------------------------------------------------------------------

# users get own chats
curl -X GET -H "QB-Token: "$tokenForUser1 https://$host/chat/Dialog.json?name=GROUP_CHAT_FOR_TT_AND_GGG_AND_FL
curl -X GET -H "QB-Token: "$tokenForUser2 https://$host/chat/Dialog.json?name=GROUP_CHAT_FOR_TT_AND_GGG_AND_FL
curl -X GET -H "QB-Token: "$tokenForUser3 https://$host/chat/Dialog.json?name=GROUP_CHAT_FOR_TT_AND_GGG_AND_FL

-----------------------------------------------------------------------------------------------------------------------

# user 1 delete chat for all (FORCE)
curl -X DELETE \
-H "Content-Type: application/json" \
-H "QB-Token: "$tokenForUser1 \
https://$host/chat/Dialog/$chatId.json?force=1

-----------------------------------------------------------------------------------------------------------------------

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
