define(['/js/guest/m/registration.js', 'text!/js/guest/t/registration.tpl.html'], function (m, t) {
    return  Backbone.View.extend({
        events:{
            'click #signUp': 'signUp',
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
        showErrors: function (e) {
            _.each (e, function (v) {
                var msg = '<small class="has-error"><small class="control-label">'+v.msg+'</small></small>';
                this.$('#'+v.param).popover({html: true, content: msg, placement: 'top'}).popover('show');
            });
        },
        hideErrors: function () {
            this.$('*').popover('destroy');
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
