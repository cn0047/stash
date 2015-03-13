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
        mContact: new mContact(),
        cMyContact: new cContact(),
        cAllContact: new cContact(),
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
            this.cAllContact.on('renderAllContacts', this.renderAllContacts, this);
            this.cMyContact.on('renderMyContacts', this.renderMyContacts, this);
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
            this.getMyContacts();
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
            var userId = this.$(e.target).parent().attr('data-userId');
            var m = this.cAllContact.remove(
                this.cAllContact.where({_id: userId})
            );
            m = m[0];
            // m.hash = 'addContact/owner/'+app.views.account.userId+'/user/'+userId;
            // m.save();
            // console.log(userId);
            // console.log(m.get('sname'));
            // this.cMyContact.add([{user: {_id: userId, sname: m.get('sname')}}]);
            // this.renderAllContacts();
            // this.renderMyContacts();
        },
    });
});
