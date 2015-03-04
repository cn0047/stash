define([
    '/js/account/m/user.js',
    'text!/js/account/t/account.layout.tpl.html'
], function (user, t) {
    return  Backbone.skipeView.extend({
        el: '#doc #account',
        defaultRoute: 'home',
        tpl: t,
        user: new user(),
        userId: '',
        events:{
        },
        routes: {
            view_account_home: '/js/account/v/home.js',
            view_account_logOut: '/js/account/v/logOut.js',
            view_account_contacts: '/js/account/v/contacts.js',
        },
        initialize: function () {
            app.views.app.on('renderLayouts', this.renderLayouts, this);
            this.user.on('afterGetUser', this.afterGetUser, this);
        },
        renderLayouts: function () {
            this.$('.head').html(_.template(this.tpl));
            this.getUser();
        },
        getUser: function () {
            this.user.hash = 'getUser';
            this.user.fetch({
                success: function (m, r) {
                    m.trigger('afterGetUser', r);
                }
            });
        },
        afterGetUser: function (r) {
            this.userId = this.user.get('_id');
        },
    });
});
