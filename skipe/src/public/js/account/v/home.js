define(['text!/js/account/t/home.tpl.html'], function (t) {
    return  Backbone.skipeView.extend({
        tpl: t,
        initialize: function () {
        },
        go: function () {
            this.renderIf();
            app.views.app.hideLoading();
        },
        render: function () {
            // load contacts
            // load chats
            this.$el.html(_.template(this.tpl));
        },
    });
});
