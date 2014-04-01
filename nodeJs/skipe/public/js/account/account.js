define(['text!/js/account/t/account.layout.tpl.html'], function (t) {
    return  Backbone.View.extend({
        el: '#doc #account',
        events:{
        },
        routes: {
        },
        initialize: function () {
            console.log('init account...');
            tpl = _.template(t);
            this.$el.html(tpl());
        },
        goTo: function (route) {
        }
    });
});
