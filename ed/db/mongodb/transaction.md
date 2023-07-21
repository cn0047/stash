Transaction
-

#### v3 Using Two-Phase Commits in Production Applications

Mongo does not have transactions,
but have atomic operations,
or you can use next approach:

````js
db.accounts.save({name: "A", balance: 1000, pendingTransactions: []})
db.accounts.save({name: "B", balance: 1000, pendingTransactions: []})

// init
db.transactions.save({source: "A", destination: "B", value: 100, state: "initial"})

// set status pending
t = db.transactions.findOne({state: "initial"})
db.transactions.update({_id: t._id}, {$set: {state: "pending"}})

// apply transaction
db.accounts.update(
    {name: t.source, pendingTransactions: {$ne: t._id}},
    {$inc: {balance: -t.value}, $push: {pendingTransactions: t._id}}
)
db.accounts.update(
    {name: t.destination, pendingTransactions: {$ne: t._id}},
    {$inc: {balance: t.value}, $push: {pendingTransactions: t._id}}
)

// set status commited
db.transactions.update({_id: t._id}, {$set: {state: "committed"}})

// remove pending
db.accounts.update({name: t.source}, {$pull: {pendingTransactions: t._id}})
db.accounts.update({name: t.destination}, {$pull: {pendingTransactions: t._id}})

// set status done
db.transactions.update({_id: t._id}, {$set: {state: "done"}})

// rollback
// set status canceling
db.transactions.update({_id: t._id}, {$set: {state: "canceling"}})
db.accounts.update(
    {name: t.source, pendingTransactions: t._id},
    {$inc: {balance: t.value}, $pull: {pendingTransactions: t._id}}
)
db.accounts.update(
    {name: t.destination, pendingTransactions: t._id},
    {$inc: {balance: -t.value}, $pull: {pendingTransactions: t._id}}
)

db.transactions.update({_id: t._id}, {$set: {state: "canceled"}})

t = db.transactions.findAndModify({
    query: {state: "initial", application: {$exists: 0}},
    update: {$set: {state: "pending", application: "A1"}},
    new: true
})
db.transactions.find({application: "A1", state: "pending"})
````

#### v4 Added transactions

MongoDB supports multi-document MongoDB ACID transactions.
