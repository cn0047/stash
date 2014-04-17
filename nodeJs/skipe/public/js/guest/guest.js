define(['text!/js/guest/t/guest.layout.tpl.html'], function (t) {
    return  Backbone.skipeView.extend({
        el: '#doc #guest',
        events:{
            'click .navbar-nav li a': 'clickNav',
        },
        defaultRoute: 'guest/home',
        routes: {
            view_guest_registration: '/js/guest/v/registration.js',
            view_guest_login: '/js/guest/v/login.js',
            view_guest_home: '/js/guest/v/home.js',
        },
        initialize: function () {
            console.log('init guest...');
            tpl = _.template(t);
            this.$el.html(tpl());
        },
        clickNav: function (e) {
            this.$('.navbar-nav li').removeClass('active');
            this.$(e.target).parent().addClass('active');
        },
        clickNavBrand: function (e) {
            this.$('.navbar-nav li').removeClass('active');
        },
    });
});
