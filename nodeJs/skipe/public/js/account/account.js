define(['text!/js/guest/t/guest.layout.tpl.html'], function (t) {
    return  Backbone.View.extend({
        events:{
        },
        routes: {
        },
        initialize: function () {
            console.log('init account...');
        },
        goTo: function (route) {
            console.log(route);
        }
    });
});
