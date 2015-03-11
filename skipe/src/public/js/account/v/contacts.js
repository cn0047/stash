define([
    '/js/account/m/contact.js',
    '/js/account/c/contact.js',
    'text!/js/account/t/contacts.tpl.html',
    'text!/js/account/t/contactInfo.tpl.html',
    'text!/js/account/t/contactsAll.tpl.html',
    'text!/js/account/t/contactsMy.tpl.html'
], function (mContact, cContact, t, tContactInfo, tAll, tMy) {
    return  Backbone.skipeView.extend({
        tpl: t,
        tContactInfo: tContactInfo,
        tAll: tAll,
        tMy: tMy,
        mContact: mContact,
        cContact: new cContact(),
        events:{
            'click #myContacts a': 'cancelEvent',
            'click #allContacts a': 'cancelEvent',
            'click #myContacts .glyphicon-user': 'getContactInfo',
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
            this.cContact.on('afterGetAllContacts', this.afterGetAllContacts, this);
            this.cContact.on('afterGetMyContacts', this.afterGetMyContacts, this);
        },
        go: function () {
            this.renderIf();
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
            this.cContact.hash = 'getAllContacts/user/'+app.views.account.userId;
            this.cContact.fetch({
                success: function (c, r) {
                    c.trigger('afterGetAllContacts', r);
                }
            });
        },
        afterGetAllContacts: function (r) {
            this.$('#allContacts .list-group').html(
                _.template(this.tAll)({data: r})
            );
            this.getMyContacts();
        },
        getMyContacts: function () {
            this.cContact.hash = 'getContacts/user/'+app.views.account.userId;
            this.cContact.fetch({
                success: function (c, r) {
                    c.trigger('afterGetMyContacts', r);
                }
            });
        },
        afterGetMyContacts: function (r) {
            this.$('#myContacts .list-group').html(
                _.template(this.tMy)({data: r})
            );
        },
        getContactInfo: function (e) {
            app.views.app.showLoading();
            var m = new mContact();
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
            console.log(
                this.$(e.target).parent().attr('data-userId')
            );
        },
    });
});
