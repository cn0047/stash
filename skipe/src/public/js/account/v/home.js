define([
    '/js/account/c/chat.js',
    '/js/account/c/contact.js',
    '/js/account/m/post.js',
    '/js/account/c/post.js',
    'text!/js/account/t/home.tpl.html',
    'text!/js/account/t/mainChats.tpl.html',
    'text!/js/account/t/mainPosts.tpl.html'
], function (cChat, cContact, mPost, cPost, t, tChats, tPosts) {
    return  Backbone.skipeView.extend({
        cChat: new cChat(),
        cContact: new cContact(),
        cPost: new cPost(),
        mPost: mPost,
        tpl: t,
        tplChats: tChats,
        tplPosts: tPosts,
        events:{
            'click #mainChats a': 'activateChat',
            'click #mainPosts #showUsersOfChat': 'showUsersOfChat',
            'keypress #newPost': 'newPost',
        },
        initialize: function () {
            this.cContact.on('afterGetContacts', this.afterGetContacts, this);
            this.cPost.on('afterGetPosts', this.afterGetPosts, this);
            this.cChat.on('afterGetChats', this.afterGetChats, this);
        },
        go: function () {
            this.renderIf();
            this.activate();
            app.views.app.hideLoading();
        },
        render: function () {
            this.$el.html(_.template(this.tpl));
        },
        activate: function () {
            this.getChats();
        },
        getChats: function () {
            this.cChat.hash = 'getChats/user/'+app.views.account.userId;
            this.cChat.fetch({
                success: function (c, r) {
                    c.trigger('afterGetChats', r);
                }
            });
        },
        afterGetChats: function (r) {
            this.renderChats(r);
            this.getPosts();
        },
        renderChats: function (d) {
            this.$('#mainChats').html(_.template(this.tplChats)({data: d}));
            this.$('#mainChats .list-group a:first').addClass('active');
            this.$('#mainChats .settings #chatCaption').val(
                this.$('#mainChats .list-group a:first').html()
            );
        },
        getPosts: function () {
            this.cPost.hash = 'getPosts/chat/'+this.getActiveChatId();
            this.cPost.fetch({
                success: function (c, r) {
                    c.trigger('afterGetPosts', r);
                }
            });
        },
        getActiveChatId: function () {
            return this.$('#mainChats .list-group .active').attr('data-chat');
        },
        afterGetPosts: function (r) {
            this.renderPosts(r);
        },
        renderPosts: function (d) {
            var t = this.tplPosts;
            this.$('#mainPosts .container').html('');
            _.each(d, function (v) {
                this.$('#mainPosts .container').append(_.template(t)({v: v}));
            })
            app.views.app.hideLoading();
        },
        activateChat: function (e) {
            app.views.app.showLoading();
            this.renderPosts({});
            this.$('#mainChats .list-group a').removeClass('active');
            this.$(e.currentTarget).addClass('active');
            this.$('#mainChats .settings #chatCaption').val(
                this.$(e.currentTarget).html()
            );
            this.getPosts();
        },
        newPost: function (e) {
            if (e.which === 13) {
                var m = new mPost();
                var d = {
                    chat: this.getActiveChatId(),
                    user: app.views.account.user.get('token'),
                    date: (new Date).toLocaleString(),
                    text: this.$('#newPost').val()
                };
                m.on('afterAddPost', this.afterAddPost, this);
                m.save(d, {
                    success: function (m) {
                        m.trigger('afterAddPost', d);
                    }
                });
                this.cPost.add(m);
            }
        },
        afterAddPost: function (d) {
            this.$('#mainPosts .container').append(
                _.template(this.tplPosts)({v: d})
            );
            // socet
            this.$('#newPost').val('');
        },
        showUsersOfChat: function () {
            this.$('#mainPosts #mainUsersInChat').toggleClass('hide');
            var activeChatId = this.getActiveChatId();
            var v = this;
            this.cChat.find(function (m) {
                if (m.get('chat')._id === activeChatId) {
                    m.hash = 'getUsersInChat/chat/'+activeChatId+'/user/'+app.views.account.userId;
                    console.log(m);
                    m.on('afterGetUsersInChat', v.afterGetUsersInChat, v);
                    m.fetch({
                        success: function (m, r) {
                            m.trigger('afterGetUsersInChat', r);
                        }
                    });
                }
            });
        },
        afterGetUsersInChat: function (r) {
            console.log(r);
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
    });
});
