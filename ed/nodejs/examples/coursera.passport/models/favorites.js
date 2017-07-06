// grab the things we need
var mongoose = require('mongoose');
var findOneOrCreate = require('mongoose-find-one-or-create');
var Schema = mongoose.Schema;

var favoriteSchema = new Schema({
    postedBy: {type: mongoose.Schema.Types.ObjectId, ref: 'User'},
    dishes: [{type: mongoose.Schema.Types.ObjectId, ref: 'Dish'}]
}, {
    timestamps: true
});

favoriteSchema.plugin(findOneOrCreate);

// the schema is useless so far
// we need to create a model using it
var Favourites = mongoose.model('Favourite', favoriteSchema);

// make this available to our Node applications
module.exports = Favourites;
