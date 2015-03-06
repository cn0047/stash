define([
    '/js/account/c/contact.js',
    'text!/js/account/t/contacts.tpl.html',
    'text!/js/account/t/contactsAll.tpl.html',
    'text!/js/account/t/contactsMy.tpl.html'
], function (cContact, t, tAll, tMy) {
    return  Backbone.skipeView.extend({
        tpl: t,
        tAll: tAll,
        tMy: tMy,
        cContact: new cContact(),
        events:{
            'click #myContacts a': 'activateContact',
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
        activateContact: function (e) {
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
    });
});
