var express = require('express');
var bodyParser = require('body-parser');

var Favorites = require('../models/favorites');
var Verify = require('./verify');

var favoriteRouter = express.Router();
favoriteRouter.use(bodyParser.json());

favoriteRouter.route('/')
    .get(Verify.verifyOrdinaryUser, function (req, res, next) {
        var cond = {postedBy: req.decoded._doc._id};
        Favorites.find(cond)
            .populate('postedBy')
            .populate('dishes')
            .exec(function (err, d) {
                if (err) throw err;
                res.json(d);
            });
    })
    .post(Verify.verifyOrdinaryUser, function (req, res, next) {
        var cond = {postedBy: req.decoded._doc._id};
        var doc = {postedBy: req.decoded._doc._id, dishes: []};
        Favorites.findOneOrCreate(cond, doc, function (err, fav) {
            if (err) throw err;
            if (fav.dishes.indexOf(req.body._id) === -1) {
                fav.dishes.push(req.body._id);
                fav.save(function (err, result) {
                    if (err) throw err;
                    console.log('Updated Comments!');
                    res.json(result);
                });
            } else {
                res.json(fav);
            }
        });
    })
    .delete(Verify.verifyOrdinaryUser, function (req, res, next) {
        var cond = {postedBy: req.decoded._doc._id};
        Favorites.find(cond).remove(function (err, resp) {
            if (err) throw err;
            res.json(resp);
        });
    });

favoriteRouter.route('/:dishId')
    .delete(Verify.verifyOrdinaryUser, function (req, res, next) {
        var cond = {postedBy: req.decoded._doc._id};
        Favorites.update(cond, {$pullAll: {dishes: [req.params.dishId]}}, function (err, resp) {
            if (err) throw err;
            res.json(resp);
        });
    });

module.exports = favoriteRouter;
