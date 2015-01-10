define(['text!/js/guest/t/guest.layout.tpl.html'], function (t) {
    return  Backbone.skipeView.extend({
        el: '#doc #guest',
        tpl: t,
        defaultRoute: 'home',
        events:{
            'click .navbar a': 'clickNav',
        },
        routes: {
            view_guest_registration: '/js/guest/v/registration.js',
            view_guest_login: '/js/guest/v/login.js',
            view_guest_home: '/js/guest/v/home.js',
        },
        initialize: function () {
            app.views.app.on('renderLayouts', this.renderLayouts, this);
        },
        renderLayouts: function () {
            this.$('.head').html(_.template(this.tpl));
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
