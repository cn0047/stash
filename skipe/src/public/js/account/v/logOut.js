define([], function () {
    return  Backbone.skipeView.extend({
        model: new Backbone.skipeModel(),
        initialize: function () {
            this.model.on('afterLogOut', this.afterLogOut, this);
        },
        go: function () {
            this.model.hash = 'logOut';
            this.model.save({}, {
                success: function (m, r) {
                    m.trigger('afterLogOut', r);
                }
            });
        },
        afterLogOut: function (r) {
            window.location.reload();
        },
    });
});
