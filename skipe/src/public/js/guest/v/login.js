define(['/js/guest/m/login.js', 'text!/js/guest/t/login.tpl.html'], function (m, t) {
    return  Backbone.skipeView.extend({
        tpl: t,
        model: new m(),
        events:{
            'click #btnLogInDefault': 'btnLogInDefault',
            'click #btnLogIn': 'logIn',
            'click #btnFPass': 'fPass',
            'click #btnSubmitFPass': 'submitFPass',
            'change #type input[name=type]': 'switchType',
        },
        switchType: function (e) {
            this.$('#token').val('');
            this.$('#pass').val('');
            var type = this.$('input[name=type]:checked').val();
            switch (type) {
                case 'email':
                    this.model.checkOptions.token.msg = app.nls.invalidEmail;
                    this.$('#token').attr('placeholder', app.nls.EMAIL);
                    break;
                case 'sname':
                    this.model.checkOptions.token.msg = app.nls.invalidScreenName;
                    this.$('#token').attr('placeholder', app.nls.SCREEN_NAME);
                    break;
                default:
                    this.model.checkOptions.token.msg = app.nls.app.nls.invalidToken;
            }
        },
        initialize: function () {
            this.model.on('afterGetDefaultLogIn', this.afterGetDefaultLogIn, this);
            this.model.on('afterLogIn', this.afterLogIn, this);
            this.model.on('afterFPass', this.afterFPass, this);
        },
        go: function () {
            this.renderIf();
            var matches = (document.cookie).match(/guestLoginType=(email|sname)/);
            if (matches) {
                this.$el.find('#type input[name=type][value='+matches[1]+']').trigger('click');
            }
            app.views.app.hideLoading();
        },
        render: function () {
            this.$el.html(_.template(this.tpl));
        },
        fulfillAction: function (h, a) {
            this.hideErrors();
            this.model.clear();
            var type = this.$('input[name=type]:checked').val();
            this.model.checkOptions.token.type = type;
            this.model.set({
                type: type,
                token: (this.$('#token').val()).trim(),
            });
            if (this.$('#login').is(':visible')) {
                this.model.set({
                    pass: (this.$('#pass').val()).trim(),
                });
            }
            if (this.model.isValid()) {
                app.views.app.showLoading();
                this.model.hash = h;
                this.model.save({}, {
                    success: function (m) {
                        m.trigger(a);
                    }
                });
            } else {
                this.showErrors(this.model.validationError);
            }
        },
        btnLogInDefault: function () {
            app.views.app.showLoading();
            this.model.hash = 'getDefaultLogIn';
            this.model.fetch({
                success: function (m, r) {
                    m.trigger('afterGetDefaultLogIn', r);
                }
            });
        },
        afterGetDefaultLogIn: function (args) {
            this.$('input[name=type][value='+args.type+']').click();
            this.$('#token').val(args.token);
            this.$('#pass').val(args.pass);
            this.$('#btnLogIn').click();
            app.views.app.hideLoading();
        },
        logIn: function () {
            this.fulfillAction('logIn', 'afterLogIn');
        },
        fPass: function () {
            this.$('#login').hide();
            this.$('#fp').removeClass('hide');
        },
        submitFPass: function () {
            this.fulfillAction('forgotPassword', 'afterFPass');
        },
        afterSave: function (cb) {
            var e = this.model.get('errors');
            if (e) {
                this.showErrors(e);
            }
            if (this.model.get('success')) {
                cb(this);
            }
            app.views.app.hideLoading();
        },
        afterLogIn: function () {
            this.afterSave(function (t) {
                document.cookie = 'guestLoginType='+t.$('input[name=type]:checked').val();
                t.$('form').submit();
            });
        },
        afterFPass: function () {
            this.afterSave(function (t) {
                t.$el.find('div:first').hide();
                t.$el.find('div:nth-child(2)').removeClass('hide');
            });
        },
    });
});
