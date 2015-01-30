define([
    '/js/account/c/chat.js',
    '/js/account/c/contact.js',
    '/js/account/c/post.js',
    'text!/js/account/t/account.layout.tpl.html'
], function (cChat, cContact, cPost, t) {
    return  Backbone.skipeView.extend({
        el: '#doc #account',
        tpl: t,
        defaultRoute: 'home',
        cChat: new cChat(),
        cContact: new cContact(),
        cPost: new cPost(),
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
