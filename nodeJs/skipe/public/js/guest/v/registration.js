define(['text!/js/guest/t/registration.tpl.html'], function (t) {
    return  Backbone.View.extend({
        initialize: function () {
            this.tpl = _.template(t);
        },
        goTo: function () {
            console.log('goToHome')
            this.$el.html(this.tpl());
            app.views.app.hideLoading();
        },
    });
});
