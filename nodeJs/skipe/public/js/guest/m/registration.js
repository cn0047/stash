define([], function () {
    return  Backbone.skipeModel.extend({
        hash: '',
        defaults: {
        },
        checkOptions: {
            email: {type: 'email', key: 'email', msg: 'Invalid email.'},
            sname: {type: 'sname', key: 'sname', msg: 'Invalid screen name.'},
        },
        urlRoot: function () {
            return 'guest/'+this.hash;
        },
    });
});
