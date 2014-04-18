define(['text!/js/account/t/account.layout.tpl.html'], function (t) {
    return  Backbone.skipeView.extend({
        el: '#doc #account',
        defaultRoute: 'account/home',
        events:{
        },
        routes: {
            view_account_home: '/js/account/v/home.js',
        },
        initialize: function () {
            console.log('init account...');
            this.$el.html(_.template(t));
        },
    });
});
