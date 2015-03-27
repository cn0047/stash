/**
 * @TODO: Add unique keys (email, sname) to user collection.
 */
/*
USER
*/
db.user.remove({});
db.user.find();
{ "_id" : ObjectId("54b23de857fe2afb0c1182bf"), "email" : "codenamek2010+JamesBond@gmail.com", "sname" : "James Bond", "password" : "fee8edd194" }
{ "_id" : ObjectId("54b23e6ab8b3cf0211c5adf3"), "email" : "codenamek2010+M@gmail.com", "sname" : "M", "password" : "de3a50b65c" }
{ "_id" : ObjectId("54b24d47df23943812d375fc"), "email" : "codenamek2010+Q@gmail.com", "sname" : "Q", "password" : "91d11f0dc4" }
{ "_id" : ObjectId("54b24dae445e086e129d0feb"), "email" : "codenamek2010+VesperLynd@gmail.com", "sname" : "Vesper Lynd", "password" : "ae95511c7c" }
{ "_id" : ObjectId("54b24e095b9581dc127f428f"), "email" : "codenamek2010+FelixLeiter@gmail.com", "sname" : "Felix Leiter", "password" : "6c4e16ede8" }
{ "_id" : ObjectId("54b24e99af762f8013a30525"), "email" : "codenamek2010+ReneMathis@gmail.com", "sname" : "Rene Mathis", "password" : "bc9b7f38a9" }
{ "_id" : ObjectId("54b24f95061695e114e9910f"), "email" : "codenamek2010+AlexDimitrios@gmail.com", "sname" : "Alex Dimitrios", "password" : "db3815b6ca" }
{ "_id" : ObjectId("54b24fbe49b8802715620d76"), "email" : "codenamek2010+LeChiffre@gmail.com", "sname" : "Le Chiffre", "password" : "6464688c35" }
{ "_id" : ObjectId("54b24fe82607dbaf153f9afc"), "email" : "codenamek2010+MrWhite@gmail.com", "sname" : "Mr White", "password" : "83d5645f59" }
{ "_id" : ObjectId("54c091de3c6a314b63e514d3"), "email" : "x@com.com", "sname" : "x", "password" : "b27813df19" }

/*
CONTACT
*/
db.contact.remove({});
// db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {"$ref" : "user", "$id" : ObjectId("54b23e6ab8b3cf0211c5adf3"), "$db" : "skipe"}});
// db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {"$ref" : "user", "$id" : ObjectId("54b24d47df23943812d375fc"), "$db" : "skipe"}});
// db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {"$ref" : "user", "$id" : ObjectId("54b24dae445e086e129d0feb"), "$db" : "skipe"}});
// db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {"$ref" : "user", "$id" : ObjectId("54b24e095b9581dc127f428f"), "$db" : "skipe"}});
// db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {"$ref" : "user", "$id" : ObjectId("54b24e99af762f8013a30525"), "$db" : "skipe"}});
// db.contact.insert({owner : ObjectId("54b24fbe49b8802715620d76"), user: {"$ref" : "user", "$id" : ObjectId("54b24fe82607dbaf153f9afc"), "$db" : "skipe"}});
// db.contact.insert({owner : ObjectId("54b24fbe49b8802715620d76"), user: {"$ref" : "user", "$id" : ObjectId("54b24f95061695e114e9910f"), "$db" : "skipe"}});
// db.contact.insert({owner : ObjectId("54b24fbe49b8802715620d76"), user: {"$ref" : "user", "$id" : ObjectId("54b24e99af762f8013a30525"), "$db" : "skipe"}});
db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {_id: ObjectId("54b23e6ab8b3cf0211c5adf3"), sname: "M"}});
db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {_id: ObjectId("54b24d47df23943812d375fc"), sname: "Q"}});
db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {_id: ObjectId("54b24dae445e086e129d0feb"), sname: "Vesper Lynd"}});
db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {_id: ObjectId("54b24e095b9581dc127f428f"), sname: "Felix Leiter"}});
db.contact.insert({owner : ObjectId("54b23de857fe2afb0c1182bf"), user: {_id: ObjectId("54b24e99af762f8013a30525"), sname: "Rene Mathis"}});
db.contact.insert({owner : ObjectId("54b24fbe49b8802715620d76"), user: {_id: ObjectId("54b24fe82607dbaf153f9afc"), sname: "Mr White"}});
db.contact.insert({owner : ObjectId("54b24fbe49b8802715620d76"), user: {_id: ObjectId("54b24f95061695e114e9910f"), sname: "Alex Dimitrios"}});
db.contact.insert({owner : ObjectId("54b24fbe49b8802715620d76"), user: {_id: ObjectId("54b24e99af762f8013a30525"), sname: "Rene Mathis"}});
db.contact.find();
{ "_id" : ObjectId("54f9463ff2a273e199df07fb"), "owner" : ObjectId("54b23de857fe2afb0c1182bf"), "user" : { "_id" : ObjectId("54b23e6ab8b3cf0211c5adf3"), "sname" : "M" } }
{ "_id" : ObjectId("54f94643f2a273e199df07fc"), "owner" : ObjectId("54b23de857fe2afb0c1182bf"), "user" : { "_id" : ObjectId("54b24d47df23943812d375fc"), "sname" : "Q" } }
{ "_id" : ObjectId("54f94643f2a273e199df07fd"), "owner" : ObjectId("54b23de857fe2afb0c1182bf"), "user" : { "_id" : ObjectId("54b24dae445e086e129d0feb"), "sname" : "Vesper Lynd" } }
{ "_id" : ObjectId("54f94644f2a273e199df07fe"), "owner" : ObjectId("54b23de857fe2afb0c1182bf"), "user" : { "_id" : ObjectId("54b24e095b9581dc127f428f"), "sname" : "Felix Leiter" } }
{ "_id" : ObjectId("54f94644f2a273e199df07ff"), "owner" : ObjectId("54b23de857fe2afb0c1182bf"), "user" : { "_id" : ObjectId("54b24e99af762f8013a30525"), "sname" : "Rene Mathis" } }
{ "_id" : ObjectId("54f94644f2a273e199df0800"), "owner" : ObjectId("54b24fbe49b8802715620d76"), "user" : { "_id" : ObjectId("54b24fe82607dbaf153f9afc"), "sname" : "Mr White" } }
{ "_id" : ObjectId("54f94645f2a273e199df0801"), "owner" : ObjectId("54b24fbe49b8802715620d76"), "user" : { "_id" : ObjectId("54b24f95061695e114e9910f"), "sname" : "Alex Dimitrios" } }
{ "_id" : ObjectId("54f94645f2a273e199df0802"), "owner" : ObjectId("54b24fbe49b8802715620d76"), "user" : { "_id" : ObjectId("54b24e99af762f8013a30525"), "sname" : "Rene Mathis" } }
db.contact.find({"owner" : ObjectId("54b23de857fe2afb0c1182bf")});

/*
CHAT
*/
db.chat.insert({ "_id" : ObjectId("54b8b991c0eceb8b5083c1c2"), "caption" : "mi6" });
db.chat.insert({ "_id" : ObjectId("54b8b995c0eceb8b5083c1c3"), "caption" : "organization" });
db.chat.insert({ "_id" : ObjectId("54b8b996c0eceb8b5083c1c4"), "caption" : "Casino Royale" });
db.chat.insert({ "_id" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "caption" : "M with Bond" });
db.chat.insert({ "_id" : ObjectId("54dbac3e0b3ade33b8bb944a"), "caption" : "Bond with Vesper" });
db.chat.insert({ "_id" : ObjectId("54dbac560b3ade33b8bb944b"), "caption" : "Bond with Mathis" });
db.chat.find();
db.chat.count();
{ "_id" : ObjectId("54b8b991c0eceb8b5083c1c2"), "caption" : "mi6" }
{ "_id" : ObjectId("54b8b995c0eceb8b5083c1c3"), "caption" : "organization" }
{ "_id" : ObjectId("54b8b996c0eceb8b5083c1c4"), "caption" : "Casino Royale" }
{ "_id" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "caption" : "M with Bond" }
{ "_id" : ObjectId("54dbac3e0b3ade33b8bb944a"), "caption" : "Bond with Vesper" }
{ "_id" : ObjectId("54dbac560b3ade33b8bb944b"), "caption" : "Bond with Mathis" }

/*
POST
*/
db.post.remove({});
// db.post.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "$db" : "skipe"}, user: "James Bond", date: ISODate("2006-11-14T10:01:53Z"), text: "Ok. Got it!"});
// db.post.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54dbac3e0b3ade33b8bb944a"), "$db" : "skipe"}, user: "James Bond", date: ISODate("2008-11-14T11:01:53Z"), text: "Hi, Darling!"});
// db.post.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54dbac3e0b3ade33b8bb944a"), "$db" : "skipe"}, user: "Vesper Lynd", date: ISODate("2008-11-14T11:01:59Z"), text: "Don't hi, agent. I don't give you money..."});

// db.post.insert({chat: {_id: ObjectId("54b8c04bc0eceb8b5083c1cc"), caption: "M with Bond"}, user : "M",           date : "2006-11-14T10:00:35Z",     text : "You are now agent with two zeros!" });
// db.post.insert({chat: {_id: ObjectId("54b8c04bc0eceb8b5083c1cc"), caption: "M with Bond"}, user : "James Bond",  date : "2006-11-14T10:00:47Z",     text : "I will not fail!" });
// db.post.insert({chat: {_id: ObjectId("54b8c04bc0eceb8b5083c1cc"), caption: "M with Bond"}, user : "M",           date : "2006-11-14T10:00:53Z",     text : "That's the cards in your hands..." });
// db.post.insert({chat: {_id: ObjectId("54b8c04bc0eceb8b5083c1cc"), caption: "M with Bond"}, user : "James Bond",  date : "2006-11-14T10:01:53Z",     text : "Ok. Got it!" });
// db.post.insert({chat: {_id: ObjectId("54dbac3e0b3ade33b8bb944a"), caption: "Bond with Vesper"}, user : "James Bond",  date : "2008-11-14T11:01:53Z",     text : "Hi, Darling!" });
// db.post.insert({chat: {_id: ObjectId("54dbac3e0b3ade33b8bb944a"), caption: "Bond with Vesper"}, user : "Vesper Lynd", date : "2008-11-14T11:01:59Z",     text : "Don't hi, agent. I don't give you money..." });
// db.post.insert({chat: {_id: ObjectId("54dbac3e0b3ade33b8bb944a"), caption: "Bond with Vesper"}, user : "James Bond",  date : "2015-02-21T22:17:57.308Z", text : "why you so cold to me?" });
// db.post.insert({chat: {_id: ObjectId("54dbac3e0b3ade33b8bb944a"), caption: "Bond with Vesper"}, user : "James Bond",  date : "2015-02-21T22:18:26.838Z", text : "?" });
// db.post.insert({chat: {_id: ObjectId("54dbac3e0b3ade33b8bb944a"), caption: "Bond with Vesper"}, user : "James Bond",  date : "2015-02-21T22:22:23.566Z", text : "??" });
// db.post.insert({chat: {_id: ObjectId("54dbac3e0b3ade33b8bb944a"), caption: "Bond with Vesper"}, user : "James Bond",  date : "2015-02-21T23:11:57.027Z", text : "&" });
// db.post.insert({chat: {_id: ObjectId("54dbac3e0b3ade33b8bb944a"), caption: "Bond with Vesper"}, user : "James Bond",  date : "2015-02-21T23:12:39.039Z", text : "???" });
// db.post.insert({chat: {_id: ObjectId("54dbac560b3ade33b8bb944b"), caption: "Bond with Mathis"}, user : "James Bond",  date : "2015-02-21T23:01:32.405Z", text : "Hi!" });
// db.post.insert({chat: {_id: ObjectId("54dbac560b3ade33b8bb944b"), caption: "Bond with Mathis"}, user : "James Bond",  date : "22.02.2015, 1:21:23",      text : "how are you?" });
db.post.insert({chat: ObjectId("54b8c04bc0eceb8b5083c1cc"), user : "M",           date : "2006-11-14T10:00:35Z",     text : "You are now agent with two zeros!" });
db.post.insert({chat: ObjectId("54b8c04bc0eceb8b5083c1cc"), user : "James Bond",  date : "2006-11-14T10:00:47Z",     text : "I will not fail!" });
db.post.insert({chat: ObjectId("54b8c04bc0eceb8b5083c1cc"), user : "M",           date : "2006-11-14T10:00:53Z",     text : "That's the cards in your hands..." });
db.post.insert({chat: ObjectId("54b8c04bc0eceb8b5083c1cc"), user : "James Bond",  date : "2006-11-14T10:01:53Z",     text : "Ok. Got it!" });
db.post.insert({chat: ObjectId("54dbac3e0b3ade33b8bb944a"), user : "James Bond",  date : "2008-11-14T11:01:53Z",     text : "Hi, Darling!" });
db.post.insert({chat: ObjectId("54dbac3e0b3ade33b8bb944a"), user : "Vesper Lynd", date : "2008-11-14T11:01:59Z",     text : "Don't hi, agent. I don't give you money..." });
db.post.insert({chat: ObjectId("54dbac3e0b3ade33b8bb944a"), user : "James Bond",  date : "2015-02-21T22:17:57.308Z", text : "why you so cold to me?" });
db.post.insert({chat: ObjectId("54dbac3e0b3ade33b8bb944a"), user : "James Bond",  date : "2015-02-21T22:18:26.838Z", text : "?" });
db.post.insert({chat: ObjectId("54dbac3e0b3ade33b8bb944a"), user : "James Bond",  date : "2015-02-21T22:22:23.566Z", text : "??" });
db.post.insert({chat: ObjectId("54dbac3e0b3ade33b8bb944a"), user : "James Bond",  date : "2015-02-21T23:11:57.027Z", text : "&" });
db.post.insert({chat: ObjectId("54dbac3e0b3ade33b8bb944a"), user : "James Bond",  date : "2015-02-21T23:12:39.039Z", text : "???" });
db.post.insert({chat: ObjectId("54dbac560b3ade33b8bb944b"), user : "James Bond",  date : "2015-02-21T23:01:32.405Z", text : "Hi!" });
db.post.insert({chat: ObjectId("54dbac560b3ade33b8bb944b"), user : "James Bond",  date : "22.02.2015, 1:21:23",      text : "how are you?" });
db.post.find();
{ "_id" : ObjectId("54f948bcf2a273e199df0811"), "chat" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "user" : "M", "date" : "2006-11-14T10:00:35Z", "text" : "You are now agent with two zeros!" }
{ "_id" : ObjectId("54f948bcf2a273e199df0812"), "chat" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "user" : "James Bond", "date" : "2006-11-14T10:00:47Z", "text" : "I will not fail!" }
{ "_id" : ObjectId("54f948bdf2a273e199df0813"), "chat" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "user" : "M", "date" : "2006-11-14T10:00:53Z", "text" : "That's the cards in your hands..." }
{ "_id" : ObjectId("54f948bdf2a273e199df0814"), "chat" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "user" : "James Bond", "date" : "2006-11-14T10:01:53Z", "text" : "Ok. Got it!" }
{ "_id" : ObjectId("54f948bdf2a273e199df0815"), "chat" : ObjectId("54dbac3e0b3ade33b8bb944a"), "user" : "James Bond", "date" : "2008-11-14T11:01:53Z", "text" : "Hi, Darling!" }
{ "_id" : ObjectId("54f948bdf2a273e199df0816"), "chat" : ObjectId("54dbac3e0b3ade33b8bb944a"), "user" : "Vesper Lynd", "date" : "2008-11-14T11:01:59Z", "text" : "Don't hi, agent. I don't give you money..." }
{ "_id" : ObjectId("54f948bef2a273e199df0817"), "chat" : ObjectId("54dbac3e0b3ade33b8bb944a"), "user" : "James Bond", "date" : "2015-02-21T22:17:57.308Z", "text" : "why you so cold to me?" }
{ "_id" : ObjectId("54f948bef2a273e199df0818"), "chat" : ObjectId("54dbac3e0b3ade33b8bb944a"), "user" : "James Bond", "date" : "2015-02-21T22:18:26.838Z", "text" : "?" }
{ "_id" : ObjectId("54f948bef2a273e199df0819"), "chat" : ObjectId("54dbac3e0b3ade33b8bb944a"), "user" : "James Bond", "date" : "2015-02-21T22:22:23.566Z", "text" : "??" }
{ "_id" : ObjectId("54f948bff2a273e199df081a"), "chat" : ObjectId("54dbac3e0b3ade33b8bb944a"), "user" : "James Bond", "date" : "2015-02-21T23:11:57.027Z", "text" : "&" }
{ "_id" : ObjectId("54f948c0f2a273e199df081b"), "chat" : ObjectId("54dbac3e0b3ade33b8bb944a"), "user" : "James Bond", "date" : "2015-02-21T23:12:39.039Z", "text" : "???" }
{ "_id" : ObjectId("54f948c0f2a273e199df081c"), "chat" : ObjectId("54dbac560b3ade33b8bb944b"), "user" : "James Bond", "date" : "2015-02-21T23:01:32.405Z", "text" : "Hi!" }
{ "_id" : ObjectId("54f948c1f2a273e199df081d"), "chat" : ObjectId("54dbac560b3ade33b8bb944b"), "user" : "James Bond", "date" : "22.02.2015, 1:21:23", "text" : "how are you?" }

/*
USERS_IN_CHAT
*/
db.usersInChat.remove({});
// db.usersInChat.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54dbac560b3ade33b8bb944b"), "$db" : "skipe"}, user: {"$ref" : "user", "$id" : ObjectId("54b23de857fe2afb0c1182bf"), "$db" : "skipe"}});
// db.usersInChat.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54dbac560b3ade33b8bb944b"), "$db" : "skipe"}, user: {"$ref" : "user", "$id" : ObjectId("54b24e99af762f8013a30525"), "$db" : "skipe"}});
// db.usersInChat.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54dbac3e0b3ade33b8bb944a"), "$db" : "skipe"}, user: {"$ref" : "user", "$id" : ObjectId("54b24dae445e086e129d0feb"), "$db" : "skipe"}});
// db.usersInChat.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54dbac3e0b3ade33b8bb944a"), "$db" : "skipe"}, user: {"$ref" : "user", "$id" : ObjectId("54b23de857fe2afb0c1182bf"), "$db" : "skipe"}});
// db.usersInChat.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "$db" : "skipe"}, user: {"$ref" : "user", "$id" : ObjectId("54b23de857fe2afb0c1182bf"), "$db" : "skipe"}});
// db.usersInChat.insert({chat: {"$ref" : "chat", "$id" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "$db" : "skipe"}, user: {"$ref" : "user", "$id" : ObjectId("54b23e6ab8b3cf0211c5adf3"), "$db" : "skipe"}});
db.usersInChat.insert({chat: {_id: ObjectId("54dbac560b3ade33b8bb944b"), caption: "Bond with Mathis"}, user: {_id: ObjectId("54b23de857fe2afb0c1182bf"), sname: "James Bond"}});
db.usersInChat.insert({chat: {_id: ObjectId("54dbac560b3ade33b8bb944b"), caption: "Bond with Mathis"}, user: {_id: ObjectId("54b24e99af762f8013a30525"), sname: "Rene Mathis"}});
db.usersInChat.insert({chat: {_id: ObjectId("54dbac3e0b3ade33b8bb944a"), caption: "Bond with Vesper"}, user: {_id: ObjectId("54b24dae445e086e129d0feb"), sname: "Vesper Lynd"}});
db.usersInChat.insert({chat: {_id: ObjectId("54dbac3e0b3ade33b8bb944a"), caption: "Bond with Vesper"}, user: {_id: ObjectId("54b23de857fe2afb0c1182bf"), sname: "James Bond"}});
db.usersInChat.insert({chat: {_id: ObjectId("54b8c04bc0eceb8b5083c1cc"), caption: "Bond with M"}, user: {_id: ObjectId("54b23de857fe2afb0c1182bf"), sname: "James Bond"}});
db.usersInChat.insert({chat: {_id: ObjectId("54b8c04bc0eceb8b5083c1cc"), caption: "Bond with M"}, user: {_id: ObjectId("54b23e6ab8b3cf0211c5adf3"), sname: "M"}});
db.usersInChat.find();
db.usersInChat.count();
{ "_id" : ObjectId("54fa09bebc86060d0f1f1a81"), "chat" : { "_id" : ObjectId("54dbac560b3ade33b8bb944b"), "caption" : "Bond with Mathis" }, "user" : { "_id" : ObjectId("54b23de857fe2afb0c1182bf"), "sname" : "James Bond" } }
{ "_id" : ObjectId("54fa09bfbc86060d0f1f1a82"), "chat" : { "_id" : ObjectId("54dbac560b3ade33b8bb944b"), "caption" : "Bond with Mathis" }, "user" : { "_id" : ObjectId("54b24e99af762f8013a30525"), "sname" : "Rene Mathis" } }
{ "_id" : ObjectId("54fa09c0bc86060d0f1f1a83"), "chat" : { "_id" : ObjectId("54dbac3e0b3ade33b8bb944a"), "caption" : "Bond with Vesper" }, "user" : { "_id" : ObjectId("54b24dae445e086e129d0feb"), "sname" : "Vesper Lynd" } }
{ "_id" : ObjectId("54fa09c1bc86060d0f1f1a84"), "chat" : { "_id" : ObjectId("54dbac3e0b3ade33b8bb944a"), "caption" : "Bond with Vesper" }, "user" : { "_id" : ObjectId("54b23de857fe2afb0c1182bf"), "sname" : "James Bond" } }
{ "_id" : ObjectId("54fa09c1bc86060d0f1f1a85"), "chat" : { "_id" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "caption" : "Bond with M" }, "user" : { "_id" : ObjectId("54b23de857fe2afb0c1182bf"), "sname" : "James Bond" } }
{ "_id" : ObjectId("54fa09c2bc86060d0f1f1a86"), "chat" : { "_id" : ObjectId("54b8c04bc0eceb8b5083c1cc"), "caption" : "Bond with M" }, "user" : { "_id" : ObjectId("54b23e6ab8b3cf0211c5adf3"), "sname" : "M" } }
