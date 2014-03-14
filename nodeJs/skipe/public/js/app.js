requirejs.config({
    paths: {
        text: '/js/lib/text-2.0.10',
        helpers: '/js/lib/helpers',
        view_guest_index: '/js/guest/guest',
    },
});

window.app = {
    init: {
        routers: {},
        models: {},
        views: {},
    },
    routers: {},
    helpers: {},
    views: {},
};

app.init.routers.app = Backbone.Router.extend({
    routes: {
        'account*other': 'account',
        'admin*other': 'admin',
        '*other': 'guest',
    },
    account: function (id) {
        this.goTo('account', id);
    },
    admin: function (id) {
        this.goTo('admin', id);
    },
    guest: function (id) {
        this.goTo('guest', id);
    },
    goTo: function (scope, id) {
        this.scope = scope;
        this.id = id;
        if (_.isEmpty(app.views[scope])) {
            require(['view_'+scope+'_index'], app.routers.app.go);
        } else {
            this.go();
        }
    },
    go: function (v) {
        scope = app.routers.app.scope;
        id = app.routers.app.id;
        if (_.isEmpty(app.views[scope])) {
            app.views[scope] = new v();
        }
        app.views[scope].goTo(id);
    },
});

app.init.models.app = Backbone.Model.extend({});

app.init.views.app = Backbone.View.extend({
    el: 'body',
    initialize: function () {
        this.showLoading();
        require(['helpers'], function (h) {
            app.helpers = h;
        });
        app.routers.app = new app.init.routers.app();
        Backbone.history.start();
    },
    showLoading: function () {
        this.$('footer #progress').removeClass('hide');
    },
    hideLoading: function () {
        this.$('footer #progress').addClass('hide');
    },
});

$(function () {
    app.views.app = new app.init.views.app({model: new app.init.models.app()});
});
