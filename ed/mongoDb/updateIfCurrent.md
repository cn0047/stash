####Isolate Sequence of Operations
````js
// Update the fields of a document
// only if the fields have not changed in the collection since the query.
var myCollection = db.products;
var myDocument = myCollection.findOne({sku: 'abc123'});
if (myDocument) {
    var oldQty = myDocument.qty;
    if (myDocument.qty < 10) {
        myDocument.qty *= 4;
    } else if (myDocument.qty < 20) {
        myDocument.qty *= 3;
    } else {
        myDocument.qty *= 2;
    }
    myCollection.update(
       {
            _id: myDocument._id,
            qty: oldQty
       },
       {
            $set: { qty: myDocument.qty }
       }
    );
    var err = db.getLastErrorObj();
    if (err && err.code) {
        print("unexpected error updating document: " + tojson( err ));
    } else if (err.n == 0) {
        print("No update: no matching document for { _id: " + myDocument._id + ", qty: " + oldQty + " }")
    }
}
````
