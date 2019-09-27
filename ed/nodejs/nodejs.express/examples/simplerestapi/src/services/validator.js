const simpleRules = {};

/**
 * Validates `resource` parameter value.
 *
 * @param {object} req Request object.
 */
simpleRules.resource = function(req) {
  // This is simple rule which is validates that resource name is available.
  req.checkParams('resource', 'Invalid resource name.').isIn(['users', 'customers', 'countries']);
};

/**
 * Validates `id` parameter value.
 *
 * @param {object} req Request object.
 */
simpleRules.id = function(req) {
  req.checkParams('id', 'Invalid id value.').isNumeric();
};

/**
 * Performs validation request parameters by rules.
 *
 * @param {object} req Request object.
 * @param {object} res Response object.
 * @param {array} rules Array with validation rules, like: 'resource', 'id'.
 */
module.exports.validate = function(req, res, rules) {

  rules.forEach(function(ruleName) {
    simpleRules[ruleName](req);
  });

  var e = req.validationErrors();
  if (e) {
    res.json({errors: e});
    return;
  }

};
