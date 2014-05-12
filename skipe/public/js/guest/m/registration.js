define([], function () {
    return  Backbone.skipeModel.extend({
        checkOptions: {
            email: {type: 'email', key: 'email', msg: app.nls.invalidEmail},
            sname: {type: 'sname', key: 'sname', msg: app.nls.invalidScreenName},
        },
    });
});
