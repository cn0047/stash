define(['text!/js/account/t/home.tpl.html'], function (t) {
    return  Backbone.View.extend({
        tpl: t,
        go: function () {
            console.log(200);
            this.$el.html(_.template(this.tpl));
            app.views.app.hideLoading();
        },
    });
});
