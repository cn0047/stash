define([
    'text!/js/account/t/home.tpl.html',
    'text!/js/account/t/mainChats.tpl.html'
], function (t, tChats) {
    return  Backbone.skipeView.extend({
        tpl: t,
        tplChats: tChats,
        events:{
        },
        initialize: function () {
        },
        go: function () {
            this.renderIf();
            app.views.app.hideLoading();
        },
        render: function () {
            this.$el.html(_.template(this.tpl));
        },
        renderChats: function (d) {
            this.$('#mainChats').html(_.template(this.tplChats)({data: d}));
        },
    });
});
