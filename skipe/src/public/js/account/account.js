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
            this.cPost.on('afterGetPosts', this.afterGetPosts, this);
            this.cChat.on('afterGetChats', this.afterGetChats, this);
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
        },
        afterGetUser: function (r) {
            this.getChats();
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
            this.cPost.hash = 'getPosts/user/'+this.user.get('_id');
            this.cPost.fetch({
                success: function (c, r) {
                    c.trigger('afterGetPosts', r);
                }
            });
        },
        afterGetPosts: function (r) {
            console.log(r);
        },
        getChats: function () {
            this.cChat.hash = 'getChats/user/'+this.user.get('_id');
            this.cChat.fetch({
                success: function (c, r) {
                    c.trigger('afterGetChats', r);
                }
            });
        },
        afterGetChats: function (r) {
            // _.template(this.tpl)
            // console.log( this.$('.chats').html() );
            el = this.$el;
        },
    });
});
