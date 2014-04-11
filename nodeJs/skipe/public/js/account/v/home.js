define(['text!/js/account/t/home.tpl.html'], function (t) {
    return  Backbone.View.extend({
        initialize: function () {
            this.tpl = _.template(t);
        },
        goTo: function () {
            this.$el.show().html(this.tpl());
            app.views.app.hideLoading();
        },
    });
});
