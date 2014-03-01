mongo
-

sudo service mongodb start|stop|restart

mongo

    // current database
    db
    show dbs
    use mydb
    // insert a collection
    j = {name : "mongo"}
    db.testData.insert(j)
    show collections
    db.testData.find()
    db.testData.findOne()
    db.testData.find().limit(3)
    db.testData.find({ name : "mongo" })
    
    // iterate over the cursor with a loop
    var c = db.testData.find()
    while (c.hasNext()) printjson(c.next())
    // print 4th item in list
    printjson(c[4])