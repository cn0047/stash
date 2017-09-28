const express = require('express');
const getthemall = require('getthemall');

const api = require('./../services/api');
const validator = require('./../services/validator');

const router  = express.Router();

/**
 * Most interesting route!!!
 *
 * @see https://github.com/cn007b/simplerestapi
 * @see https://github.com/cn007b/getthemall
 */
router.get('/resources', function(req, res) {
  getthemall(req, res, function (data) {
    res.json(data);
  });
});

/**
 * This route gets all documents from certain collection.
 */
router.get('/:resource', function(req, res) {
  validator.validate(req, res, ['resource']);

  var collectionName = req.params.resource;

  new Promise(function(resolve) {
    api.get(collectionName, function (d) {
      resolve(d);
    });
  }).then(function (d) {
    res.json(d);
  });

});

/**
 * This route gets particular resource by id.
 */
router.get('/:resource/:id', function(req, res) {
  validator.validate(req, res, ['resource', 'id']);

  var collectionName = req.params.resource;
  var documentId = parseInt(req.params.id);

  new Promise(function(resolve) {
    api.getById(collectionName, documentId, function (d) {
      resolve(d);
    });
  }).then(function (d) {
    res.json(d);
  });

});

module.exports = router;
