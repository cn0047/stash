define(['/js/guest/m/login.js', 'text!/js/guest/t/login.tpl.html'], function (m, t) {
    return  Backbone.skipeView.extend({
        events:{
            'click #logIn': 'logIn',
            'change #type input[name=type]': 'switchType',
        },
        switchType: function (e) {
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
                password: (this.$('#password').val()).trim(),
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
        afterSave: function () {
            var e = this.model.get('errors');
            if (e) {
                this.showErrors(e);
            }
            if (this.model.get('success')) {
                this.$('form').submit();
            }
            app.views.app.hideLoading();
        },
    });
});
