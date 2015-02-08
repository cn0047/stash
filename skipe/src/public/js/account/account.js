define([
    '/js/account/c/chat.js',
    '/js/account/c/contact.js',
    '/js/account/c/post.js',
    '/js/account/m/user.js',
    'text!/js/account/t/account.layout.tpl.html'
], function (cChat, cContact, cPost, user, t) {
    return  Backbone.skipeView.extend({
        el: '#doc #account',
        tpl: t,
        defaultRoute: 'home',
        cChat: new cChat(),
        cContact: new cContact(),
        cPost: new cPost(),
        user: new user(),
        events:{
        },
        routes: {
            view_account_home: '/js/account/v/home.js',
            view_account_logOut: '/js/account/v/logOut.js',
        },
        initialize: function () {
            app.views.app.on('renderLayouts', this.renderLayouts, this);
            this.user.on('afterGetUser', this.afterGetUser, this);
            this.cContact.on('afterGetContacts', this.afterGetContacts, this);
        },
        renderLayouts: function () {
            this.$('.head').html(_.template(this.tpl));
            this.activate();
        },
        activate: function () {
            this.user.hash = 'getUser';
            this.user.fetch({
                success: function (m, r) {
                    m.trigger('afterGetUser', r);
                }
            });
            this.getContacts();
            this.getPosts();
        },
        afterGetUser: function (r) {
        },
        getContacts: function (r) {
            this.cContact.hash = 'getContacts/user/'+this.user.get('_id');
            this.cContact.fetch({
                success: function (c, r) {
                    c.trigger('afterGetContacts', r);
                }
            });
        },
        afterGetContacts: function (r) {
            console.dir(r);
        },
        getPosts: function () {
            this.cContact.hash = 'getContacts/user/'+this.user.get('_id');
            this.cContact.fetch({
                success: function (c, r) {
                    c.trigger('afterGetContacts', r);
                }
            });
        },
    });
});
