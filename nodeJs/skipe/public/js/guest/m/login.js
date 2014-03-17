define([], function () {
    return  Backbone.skipeModel.extend({
        hash: '',
        defaults: {
        },
        checkOptions: {
            email: {type: 'email', key: 'email', msg: 'Invalid email.'},
            password: {type: 'password', key: 'password', msg: 'Invalid password.'},
        },
        urlRoot: function () {
            return 'guest/'+this.hash;
        },
    });
});
