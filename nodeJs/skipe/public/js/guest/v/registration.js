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
            app.views.app.showLoading();
            this.hideErrors();
            this.model.set({
                email: (this.$('#email').val()).trim(),
                sname: (this.$('#sname').val()).trim(),
            });
            if (this.model.isValid()) {
                this.model.hash = 'registration';
                this.model.save();
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
                this.$('#'+v.param).popover({html: true, content: msg}).popover('show');
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
        },
    });
});
