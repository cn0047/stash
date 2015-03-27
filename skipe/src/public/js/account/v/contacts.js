define([
    '/js/account/m/chat.js',
    '/js/account/m/contact.js',
    '/js/account/c/contact.js',
    'text!/js/account/t/contacts.tpl.html',
    'text!/js/account/t/contactInfo.tpl.html',
    'text!/js/account/t/contactsAll.tpl.html',
    'text!/js/account/t/contactsMy.tpl.html'
], function (mChat, mContact, cContact, t, tContactInfo, tAll, tMy) {
    return  Backbone.skipeView.extend({
        tpl: t,
        tContactInfo: tContactInfo,
        tAll: tAll,
        tMy: tMy,
        mChat: mChat,
        mContact: new mContact(),
        cMyContact: new cContact(),
        cAllContact: new cContact(),
        events:{
            'click #myContacts a': 'cancelEvent',
            'click #allContacts a': 'cancelEvent',
            'click #myContacts .glyphicon-user': 'getContactInfo',
            'click #myContacts .glyphicon-comment': 'startChat',
            'click #allContacts .glyphicon-plus': 'addContact',
            'mouseenter #myContacts a': 'showButtons',
            'mouseleave #myContacts a': 'hideButtons',
            'mouseenter #allContacts a': 'showButtons',
            'mouseleave #allContacts a': 'hideButtons',
        },
        showButtons: function (e) {
            this.$(e.target).find('div').removeClass('hide');
        },
        hideButtons: function (e) {
            this.$(e.target).find('div').addClass('hide');
        },
        initialize: function () {
            this.cAllContact.on('renderAllContacts', this.renderAllContacts, this);
            this.cMyContact.on('renderMyContacts', this.renderMyContacts, this);
        },
        go: function () {
            this.renderIf();
            this.getMyContacts();
            this.getAllContacts();
            app.views.app.hideLoading();
        },
        render: function () {
            this.$el.html(_.template(this.tpl));
        },
        cancelEvent: function (e) {
            e.preventDefault();
        },
        getAllContacts: function () {
            this.cAllContact.hash = 'getAllContacts/user/'+app.views.account.userId;
            this.cAllContact.fetch({
                success: function (c, r) {
                    c.trigger('renderAllContacts', r);
                }
            });
        },
        renderAllContacts: function () {
            this.$('#allContacts .list-group').html(
                _.template(this.tAll)({data: this.cAllContact.toJSON()})
            );
        },
        getMyContacts: function () {
            this.cMyContact.hash = 'getMyContacts/user/'+app.views.account.userId;
            this.cMyContact.fetch({
                success: function (c, r) {
                    c.trigger('renderMyContacts', r);
                }
            });
        },
        renderMyContacts: function () {
            this.$('#myContacts .list-group').html(
                _.template(this.tMy)({data: this.cMyContact.toJSON()})
            );
        },
        getContactInfo: function (e) {
            app.views.app.showLoading();
            var m = this.mContact;
            m.clear();
            m.hash = 'getContactInfo/user/'+this.$(e.target).parent().attr('data-userId');
            m.on('afterGetContactInfo', this.afterGetContactInfo, this);
            m.fetch({
                success: function (m, r) {
                    m.trigger('afterGetContactInfo', r);
                }
            });
        },
        afterGetContactInfo: function (d) {
            this.$('#myContacts #contactInfoModal .modal-body').html(
                _.template(this.tContactInfo)({data: d})
            );
            app.views.app.hideLoading();
        },
        addContact: function (e) {
            app.views.app.showLoading();
            var userId = this.$(e.target).parent().attr('data-userId');
            var m = this.cAllContact.remove(
                this.cAllContact.findWhere({_id: userId})
            );
            var sname = m.get('sname');
            m = this.mContact;
            m.clear();
            m.hash = 'addContact';
            var d = {
                owner: app.views.account.userId,
                userId: userId,
                sname: sname
            };
            m.on('afterAddContact', this.go, this);
            m.save(d, {
                success: function (m) {
                    m.trigger('afterAddContact');
                }
            });
        },
        startChat: function (e) {
            var $parent = this.$(e.target).parent();
            var d = {
                user: app.views.account.userId,
                userSname: app.views.account.user.get('sname'),
                withUser: $parent.attr('data-userId'),
                withUserSname: $parent.parent().find('span:first').html()
            };
            var m = new this.mChat();
            m.hash = 'startChat';
            m.on('afterStartChat', this.afterStartChat, this);
            m.save(d, {
                success: function (m, r) {
                    m.trigger('afterStartChat', r);
                }
            });
        },
        afterStartChat: function (d) {
            // console.log(d);
            // app.routers.app.go('');
            // app.views.account_home.activate();
        },
    });
});
