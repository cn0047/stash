define(['/js/guest/m/registration.js', 'text!/js/guest/t/registration.tpl.html'], function (m, t) {
    return  Backbone.View.extend({
        events:{
            'click #signUp': 'signUp',
        },
        initialize: function () {
            this.tpl = _.template(t);
            this.model = new m();
        },
        goTo: function () {
            this.$el.show().html(this.tpl());
            app.views.app.hideLoading();
        },
        signUp: function () {
            app.views.app.showLoading();
            // console.log(this.model.toJSON())
            this.model.save();
        },
    });
});
