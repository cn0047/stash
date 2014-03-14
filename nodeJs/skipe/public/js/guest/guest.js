define(['text!/js/guest/t/guest.layout.tpl.html'], function (t) {
    return  Backbone.View.extend({
        el: '#doc #guest',
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
        goTo: function (route) {
            if (_.isEmpty(route)) {
                this.goTo('guest/home');
                return;
            }
            app.views.app.showLoading();
            this.$('#content > div').hide();
            this.route = route.replace('/', '_');
            var r = 'view_'+this.route;
            if (r in this.routes) {
                if (_.isEmpty(app.views[this.route])) {
                    require([this.routes[r]], app.views.guest.go);
                } else {
                    this.go();
                }
            } else {
                this.goTo('guest/home');
            }
        },
        go: function (v) {
            route = app.views.guest.route;
            if (_.isEmpty(app.views[route])) {
                app.views.guest.$el.find('#content').append('<div id="'+route+'"></div>');
                app.views[route] = new v({el: '#doc #guest #content #'+route});
            }
            app.views[route].goTo();
        },
    });
});
