define(['text!/js/account/t/account.layout.tpl.html'], function (t) {
    return  Backbone.skipeView.extend({
        el: '#doc #account',
        tpl: t,
        defaultRoute: 'account/home',
        events:{
        },
        routes: {
            view_account_home: '/js/account/v/home.js',
            view_account_logOut: '/js/account/v/logOut.js',
        },
        initialize: function () {
            app.views.app.on('renderLayouts', this.renderLayouts, this);
        },
        renderLayouts: function () {
            this.$('.head').html(_.template(this.tpl));
        },
    });
});
