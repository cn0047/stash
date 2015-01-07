define(['text!/js/guest/t/home.tpl.html'], function (t) {
    return  Backbone.skipeView.extend({
        tpl: t,
        go: function () {
            this.renderIf();
            app.views.app.hideLoading();
        },
        render: function () {
            this.$el.html(_.template(this.tpl));
        },
    });
});
