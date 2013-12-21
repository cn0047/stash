var mongoose = require('mongoose');
var Schema = mongoose.Schema
    ,ObjectId = Schema.ObjectId;
// создание новой модели
var Widget = new Schema({
    sn : {type: String, require: true, trim: true, unique: true},
    name : {type: String, required: true, trim: true},
    desc : String,
    price : Number
});
module.exports = mongoose.model('Widget', Widget);