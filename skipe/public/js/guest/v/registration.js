define(['/js/guest/m/registration.js', 'text!/js/guest/t/registration.tpl.html'], function (m, t) {
    return  Backbone.skipeView.extend({
        tpl: t,
        model: new m(),
        events:{
            'click #signUp': 'signUp',
        },
        initialize: function () {
            this.model.on('afterSave', this.afterSave, this);
        },
        go: function () {
            this.renderIf();
            app.views.app.hideLoading();
        },
        render: function () {
            this.$el.html(_.template(this.tpl));
        },
        signUp: function () {
            this.hideErrors();
            this.model.clear();
            this.model.set({
                email: (this.$('#email').val()).trim(),
                sname: (this.$('#sname').val()).trim(),
            });
            if (this.model.isValid()) {
                app.views.app.showLoading();
                this.model.hash = 'registration';
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
                this.$el.find('div:first').hide();
                this.$el.find('div:nth-child(2)').removeClass('hide');
            }
            app.views.app.hideLoading();
        },
    });
});
