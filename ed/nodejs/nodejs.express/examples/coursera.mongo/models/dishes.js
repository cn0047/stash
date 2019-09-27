// grab the things we need
var mongoose = require('mongoose');
var Schema = mongoose.Schema;
require('mongoose-currency').loadType(mongoose);
var Currency = mongoose.Types.Currency;

var commentSchema = new Schema({
    rating:  {type: Number, min: 1, max: 5, required: true},
    comment:  {type: String, required: true},
    author:  {type: String, required: true}
}, {
    timestamps: true
});

// create a schema
var dishSchema = new Schema({
    name: {type: String, required: true, unique: true},
    image: {type: String, required: true},
    category: {type: String, required: true},
    label: {type: String, default: ''},
    price: {type: Currency, required: true},
    description: {type: String, required: true},
    comments:[commentSchema]
}, {
    timestamps: true
});

// the schema is useless so far
// we need to create a model using it
var Dishes = mongoose.model('Dish', dishSchema);

// make this available to our Node applications
module.exports = Dishes;
