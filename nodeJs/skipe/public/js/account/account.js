define(['text!/js/account/t/account.layout.tpl.html'], function (t) {
    return  Backbone.skipeView.extend({
        el: '#doc #account',
        events:{
        },
        defaultRoute: 'account/home',
        routes: {
            view_account_home: '/js/account/v/home.js',
        },
        initialize: function () {
            console.log('init account...');
            tpl = _.template(t);
            this.$el.html(tpl());
        },
    });
});
