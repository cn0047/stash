define([], function () {
    return  Backbone.Model.extend({
        hash: '',
        urlRoot: function () {
            return 'guest/registration/'+this.hash;
        },
        defaults: {
        },
        validate: function (a, o) {
        },
    });
});
