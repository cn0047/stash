/**
 * @TODO: Add unique keys (email, sname) to user collection.
 */
/*
USERS:
email
sname
password
*/
{ "_id" : ObjectId("54b23de857fe2afb0c1182bf"), "email" : "codenamek2010+JamesBond@gmail.com",     "sname" : "James Bond",     "password" : "fee8edd194" }
{ "_id" : ObjectId("54b23e6ab8b3cf0211c5adf3"), "email" : "codenamek2010+M@gmail.com",             "sname" : "M",              "password" : "de3a50b65c" }
{ "_id" : ObjectId("54b24d47df23943812d375fc"), "email" : "codenamek2010+Q@gmail.com",             "sname" : "Q",              "password" : "91d11f0dc4" }
{ "_id" : ObjectId("54b24dae445e086e129d0feb"), "email" : "codenamek2010+VesperLynd@gmail.com",    "sname" : "Vesper Lynd",    "password" : "ae95511c7c" }
{ "_id" : ObjectId("54b24e095b9581dc127f428f"), "email" : "codenamek2010+FelixLeiter@gmail.com",   "sname" : "Felix Leiter",   "password" : "6c4e16ede8" }
{ "_id" : ObjectId("54b24e99af762f8013a30525"), "email" : "codenamek2010+ReneMathis@gmail.com",    "sname" : "Rene Mathis",    "password" : "bc9b7f38a9" }
{ "_id" : ObjectId("54b24f95061695e114e9910f"), "email" : "codenamek2010+AlexDimitrios@gmail.com", "sname" : "Alex Dimitrios", "password" : "db3815b6ca" }
{ "_id" : ObjectId("54b24fbe49b8802715620d76"), "email" : "codenamek2010+LeChiffre@gmail.com",     "sname" : "Le Chiffre",     "password" : "6464688c35" }
{ "_id" : ObjectId("54b24fe82607dbaf153f9afc"), "email" : "codenamek2010+MrWhite@gmail.com",       "sname" : "Mr White",       "password" : "83d5645f59" }

/*
CONTACTS:
owner
user
*/
{ "_id" : ObjectId("54b760acee2a81ef4a89579a"), "owner" : ObjectId("54b23de857fe2afb0c1182bf"), "user" : { "sname" : "Q" } }
{ "_id" : ObjectId("54b760adee2a81ef4a89579b"), "owner" : ObjectId("54b23de857fe2afb0c1182bf"), "user" : { "sname" : "Vesper Lynd" } }
{ "_id" : ObjectId("54b760adee2a81ef4a89579c"), "owner" : ObjectId("54b23de857fe2afb0c1182bf"), "user" : { "sname" : "Felix Leiter" } }
{ "_id" : ObjectId("54b760aeee2a81ef4a89579d"), "owner" : ObjectId("54b23de857fe2afb0c1182bf"), "user" : { "sname" : "Rene Mathis" } }
{ "_id" : ObjectId("54b76107ee2a81ef4a89579e"), "owner" : ObjectId("54b24fbe49b8802715620d76"), "user" : { "sname" : "Mr White" } }
{ "_id" : ObjectId("54b76107ee2a81ef4a89579f"), "owner" : ObjectId("54b24fbe49b8802715620d76"), "user" : { "sname" : "Alex Dimitrios" } }
{ "_id" : ObjectId("54b76107ee2a81ef4a8957a0"), "owner" : ObjectId("54b24fbe49b8802715620d76"), "user" : { "sname" : "Rene Mathis" } }

/*
CHATS:
caption
*/
{ "_id" : ObjectId("54b8b991c0eceb8b5083c1c2"), "caption" : "mi6" }
{ "_id" : ObjectId("54b8b995c0eceb8b5083c1c3"), "caption" : "organization" }
{ "_id" : ObjectId("54b8b996c0eceb8b5083c1c4"), "caption" : "Casino Royale" }
{ "_id" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "caption" : "M with Bond" }
{ "_id" : ObjectId("54dbac3e0b3ade33b8bb944a"), "caption" : "Bond with Vesper" }
{ "_id" : ObjectId("54dbac560b3ade33b8bb944b"), "caption" : "Bond with Mathis" }

/*
POSTS:
chat
user
date
text
*/
{ "_id" : ObjectId("54b99b7dbde7d834a6ec8b40"), "chat" : DBRef("chat", ObjectId("54b8c04bc0eceb8b5083c1cc")), "user" : "M",          "date" : ISODate("2006-11-14T10:00:35Z"), "text" : "You are now agent with two zeros!" }
{ "_id" : ObjectId("54b99b7dbde7d834a6ec8b41"), "chat" : DBRef("chat", ObjectId("54b8c04bc0eceb8b5083c1cc")), "user" : "James Bond", "date" : ISODate("2006-11-14T10:00:47Z"), "text" : "I will not fail!" }
{ "_id" : ObjectId("54b99b7dbde7d834a6ec8b42"), "chat" : DBRef("chat", ObjectId("54b8c04bc0eceb8b5083c1cc")), "user" : "M",          "date" : ISODate("2006-11-14T10:00:53Z"), "text" : "That's the cards in your hands..." }

/*
USERS_IN_CHAT:
chat
user
*/
{ "_id" : ObjectId("54bf3992ea0f73ea669c2df7"), "chat" : DBRef("chat", ObjectId("54b8c04bc0eceb8b5083c1cc")), "user" : ObjectId("54b23e6ab8b3cf0211c5adf3") }
{ "_id" : ObjectId("54bf3996ea0f73ea669c2df8"), "chat" : DBRef("chat", ObjectId("54b8c04bc0eceb8b5083c1cc")), "user" : ObjectId("54b23de857fe2afb0c1182bf") }
{ "_id" : ObjectId("54dbadec0b3ade33b8bb944c"), "chat" : DBRef("chat", ObjectId("54dbac3e0b3ade33b8bb944a")), "user" : ObjectId("54b23de857fe2afb0c1182bf") }
{ "_id" : ObjectId("54dbaded0b3ade33b8bb944d"), "chat" : DBRef("chat", ObjectId("54dbac3e0b3ade33b8bb944a")), "user" : ObjectId("54b24dae445e086e129d0feb") }
{ "_id" : ObjectId("54dbae590b3ade33b8bb944e"), "chat" : DBRef("chat", ObjectId("54dbac560b3ade33b8bb944b")), "user" : ObjectId("54b23de857fe2afb0c1182bf") }
{ "_id" : ObjectId("54dbae590b3ade33b8bb944f"), "chat" : DBRef("chat", ObjectId("54dbac560b3ade33b8bb944b")), "user" : ObjectId("54b24e99af762f8013a30525") }



// All users.
db.user.find();
// All contacts.
db.contact.find();
// All contacts by owner.
db.contact.find({"owner" : "James Bond"});
// All chats.
db.chat.find();
db.chat.insert({"caption" : "Bond with Mathis"});
// All posts.
db.post.find();
db.post.insert({
    chat: {
        "$ref" : "chat",
        "$id" : ObjectId("54b8c04bc0eceb8b5083c1cc"),
        "$db" : "skipe"
    },
    user: "James Bond",
    date: ISODate("2006-11-14T10:01:53Z"),
    text: "Ok. Got it!"
});
db.post.insert({
    chat: {
        "$ref" : "chat",
        "$id" : ObjectId("54dbac3e0b3ade33b8bb944a"),
        "$db" : "skipe"
    },
    user: "James Bond",
    date: ISODate("2008-11-14T11:01:53Z"),
    text: "Hi, Darling!"
});
db.post.insert({
    chat: {
        "$ref" : "chat",
        "$id" : ObjectId("54dbac3e0b3ade33b8bb944a"),
        "$db" : "skipe"
    },
    user: "Vesper Lynd",
    date: ISODate("2008-11-14T11:01:59Z"),
    text: "Don't hi, agent. I don't give you money..."
});
// usersInChat.
db.usersInChat.find();
// usersInChat for user.
db.usersInChat.find({"user" : "James Bond"});
db.usersInChat.find({"chat.$id" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "user": {$ne: ObjectId("54b23de857fe2afb0c1182bf")}});
db.usersInChat.update(
    {"user" : "James Bond"},
    {$set: {"user" : ObjectId("54b23de857fe2afb0c1182bf")}},
    {}
);
db.usersInChat.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54dbac560b3ade33b8bb944b"), "$db" : "skipe"}, user: ObjectId("54b23de857fe2afb0c1182bf")});
db.usersInChat.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54dbac560b3ade33b8bb944b"), "$db" : "skipe"}, user: ObjectId("54b24e99af762f8013a30525")});
