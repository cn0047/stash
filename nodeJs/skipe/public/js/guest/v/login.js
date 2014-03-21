define(['/js/guest/m/login.js', 'text!/js/guest/t/login.tpl.html'], function (m, t) {
    return  Backbone.skipeView.extend({
        events:{
            'click #logIn': 'logIn',
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
            this.model.set({
                email: (this.$('#email').val()).trim(),
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
                console.log(200)
            }
            app.views.app.hideLoading();
        },
    });
});
