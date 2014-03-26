define(['/js/guest/m/login.js', 'text!/js/guest/t/login.tpl.html'], function (m, t) {
    return  Backbone.skipeView.extend({
        events:{
            'click #btnLogIn': 'logIn',
            'click #btnFPass': 'fPass',
            'change #type input[name=type]': 'switchType',
        },
        switchType: function (e) {
            this.$('#token').val('');
            this.$('#pass').val('');
            var type = this.$('input[name=type]:checked').val();
            switch (type) {
                case 'email':
                    this.model.checkOptions.token.msg = 'Invalid email';
                    this.$('#token').attr('placeholder', 'EMAIL');
                    break;
                case 'sname':
                    this.model.checkOptions.token.msg = 'Invalid screen name';
                    this.$('#token').attr('placeholder', 'SCREEN NAME');
                    break;
                default:
                    this.model.checkOptions.token.msg = 'Invalid token.';
            }
        },
        initialize: function () {
            this.tpl = _.template(t);
            this.model = new m();
            this.model.on('afterSave', this.afterSave, this);
        },
        goTo: function () {
            this.$el.show().html(this.tpl());
            var matches = (document.cookie).match(/guestLoginType=(email|sname)/);
            this.$el.find('#type input[name=type][value='+matches[1]+']').trigger('click');
            app.views.app.hideLoading();
        },
        logIn: function () {
            this.hideErrors();
            this.model.clear();
            var type = this.$('input[name=type]:checked').val();
            this.model.checkOptions.token.type = type;
            this.model.set({
                type: type,
                token: (this.$('#token').val()).trim(),
                password: (this.$('#pass').val()).trim(),
            });
            if (this.model.isValid()) {
                app.views.app.showLoading();
                this.model.hash = 'logIn';
                this.model.save({}, {
                    success: function (m) {
                        m.trigger('afterSave');
                    }
                });
            } else {
                this.showErrors(this.model.validationError);
            }
        },
        fPass: function () {
            this.$('#pass').parent().slideUp('fast');
            // this.$('form').slideUp('fast');
            // this.$('#fPass').removeClass('hide').slideDown('fast');
        },
        afterSave: function () {
            var e = this.model.get('errors');
            if (e) {
                this.showErrors(e);
            }
            if (this.model.get('success')) {
                this.$('form').submit();
                document.cookie = 'guestLoginType='+this.$('input[name=type]:checked').val();
            }
            app.views.app.hideLoading();
        },
    });
});
