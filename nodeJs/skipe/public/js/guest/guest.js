define(['text!/js/guest/t/guest.layout.tpl.html'], function (t) {
    return  Backbone.View.extend({
        el: '#doc #guest',
        routes: {
            view_guest_registration: '/js/guest/v/registration.js',
            view_guest_home: '/js/guest/v/registration.js',
        },
        initialize: function () {
            console.log('init guest...');
            tpl = _.template(t);
            this.$el.html(tpl());
            this.goTo('guest/home');
        },
        // goToHome: function () {
        //     app.views.app.hideLoading();
        // },
        goTo: function (route) {
            if (route) {
                app.views.app.showLoading();
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
