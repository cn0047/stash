/*
Books
    Science
        Chemistry
        Physics
            Classical Mechanics
            Quantum Mechanics
*/
db.categories.insertOne({"_id": "Quantum Mechanics", "ancestors": ["Books", "Science", "Physics"], "parent": "Physics"})
db.categories.insertOne({"_id": "Classical Mechanics", "ancestors": ["Books", "Science", "Physics"], "parent": "Physics"})
db.categories.insertOne({"_id": "Physics", "ancestors": ["Books", "Science"], "parent": "Science"})
db.categories.insertOne({"_id": "Chemistry", "ancestors": ["Books", "Science"], "parent": "Science"})
db.categories.insertOne({"_id": "Science", "ancestors": ["Books"], "parent": "Books"})
db.categories.insertOne({"_id": "Books", "ancestors": [], "parent": null})

// find() all descendant documents of a category using a single query
db.categories.find({ancestors: 'Science'});
// all descendants of "Science" are:
// "Chemistry", "Physics", "Classical Mechanics", and "Quantum Mechanics".
