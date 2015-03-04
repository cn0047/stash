define([
    '/js/account/c/contact.js',
    'text!/js/account/t/contacts.tpl.html',
    'text!/js/account/t/contactsMy.tpl.html'
], function (cContact, t, tMy) {
    return  Backbone.skipeView.extend({
        tpl: t,
        tMy: tMy,
        cContact: new cContact(),
        initialize: function () {
            this.cContact.on('afterGetMyContacts', this.afterGetMyContacts, this);
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
        getAllContacts: function () {
        },
    });
});
