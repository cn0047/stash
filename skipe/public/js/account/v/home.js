define(['text!/js/account/t/home.tpl.html'], function (t) {
    return  Backbone.View.extend({
        tpl: t,
        goTo: function () {
            this.$el.html(_.template(this.tpl));
            app.views.app.hideLoading();
        },
    });
});
