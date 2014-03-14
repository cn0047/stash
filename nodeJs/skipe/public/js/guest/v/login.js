define(['/js/guest/m/login.js', 'text!/js/guest/t/login.tpl.html'], function (m, t) {
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
    });
});
