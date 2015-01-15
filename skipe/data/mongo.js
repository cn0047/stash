/**
 * @TODO: Add unique keys (email, sname) to user collection.
 */

// Users.
db.user.find();
{ "_id" : ObjectId("54b23de857fe2afb0c1182bf"), "email" : "codenamek2010+JamesBond@gmail.com",     "sname" : "James Bond",     "password" : "fee8edd194" }
{ "_id" : ObjectId("54b23e6ab8b3cf0211c5adf3"), "email" : "codenamek2010+M@gmail.com",             "sname" : "M",              "password" : "de3a50b65c" }
{ "_id" : ObjectId("54b24d47df23943812d375fc"), "email" : "codenamek2010+Q@gmail.com",             "sname" : "Q",              "password" : "91d11f0dc4" }
{ "_id" : ObjectId("54b24dae445e086e129d0feb"), "email" : "codenamek2010+VesperLynd@gmail.com",    "sname" : "Vesper Lynd",    "password" : "ae95511c7c" }
{ "_id" : ObjectId("54b24e095b9581dc127f428f"), "email" : "codenamek2010+FelixLeiter@gmail.com",   "sname" : "Felix Leiter",   "password" : "6c4e16ede8" }
{ "_id" : ObjectId("54b24e99af762f8013a30525"), "email" : "codenamek2010+ReneMathis@gmail.com",    "sname" : "Rene Mathis",    "password" : "bc9b7f38a9" }
{ "_id" : ObjectId("54b24f95061695e114e9910f"), "email" : "codenamek2010+AlexDimitrios@gmail.com", "sname" : "Alex Dimitrios", "password" : "db3815b6ca" }
{ "_id" : ObjectId("54b24fbe49b8802715620d76"), "email" : "codenamek2010+LeChiffre@gmail.com",     "sname" : "Le Chiffre",     "password" : "6464688c35" }
{ "_id" : ObjectId("54b24fe82607dbaf153f9afc"), "email" : "codenamek2010+MrWhite@gmail.com",       "sname" : "Mr White",       "password" : "83d5645f59" }

// Contacts.
db.contact.find();
{ "_id" : ObjectId("54b760acee2a81ef4a895799"), "owner" : "James Bond", "user" : { "sname" : "M" } }
{ "_id" : ObjectId("54b760acee2a81ef4a89579a"), "owner" : "James Bond", "user" : { "sname" : "Q" } }
{ "_id" : ObjectId("54b760adee2a81ef4a89579b"), "owner" : "James Bond", "user" : { "sname" : "Vesper Lynd" } }
{ "_id" : ObjectId("54b760adee2a81ef4a89579c"), "owner" : "James Bond", "user" : { "sname" : "Felix Leiter" } }
{ "_id" : ObjectId("54b760aeee2a81ef4a89579d"), "owner" : "James Bond", "user" : { "sname" : "Rene Mathis" } }
{ "_id" : ObjectId("54b76107ee2a81ef4a89579e"), "owner" : "Le Chiffre", "user" : { "sname" : "Mr White" } }
{ "_id" : ObjectId("54b76107ee2a81ef4a89579f"), "owner" : "Le Chiffre", "user" : { "sname" : "Alex Dimitrios" } }
{ "_id" : ObjectId("54b76107ee2a81ef4a8957a0"), "owner" : "Le Chiffre", "user" : { "sname" : "Rene Mathis" } }
