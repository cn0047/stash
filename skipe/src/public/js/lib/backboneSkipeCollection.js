define([], function () {
    return  Backbone.Collection.extend({
        hash: '',
        url: function () {
            return app.routers.app.scope+'/'+this.hash;
        },
    });
});
