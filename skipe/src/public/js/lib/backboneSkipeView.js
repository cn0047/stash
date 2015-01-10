define([], function () {
    return  Backbone.View.extend({
        hideErrors: function () {
            this.$('*').popover('destroy');
        },
        showErrors: function (e) {
            _.each (e, function (v) {
                var msg = '<small class="has-error"><small class="control-label">'+v.msg+'</small></small>';
                this.$('#'+v.param).popover({html: true, content: msg, placement: 'top'}).popover('show');
            });
        },
        /**
         * Determines controller (view) and pass event to it.
         */
        go: function (route) {
            if (_.isEmpty(route)) {
                this.go(this.defaultRoute);
                return;
            }
            app.views.app.showLoading();
            this.$('.content .page').hide();
            this.route = app.routers.app.scope+'_'+route.replace('/', '_');
            var r = 'view_'+this.route;
            if (!(r in this.routes)) {
                this.go(this.defaultRoute);
            }
            if (_.isEmpty(app.views[this.route])) {
                require([this.routes[r]], app.views[app.routers.app.scope].goTo);
            } else {
                this.goTo();
            }
        },
        goTo: function (v) {
            var scope = app.routers.app.scope;
            var route = app.views[scope].route;
            if (app.views[scope].$el.find('.content #'+route).length == 0) {
                app.views[scope].$el.find('.content').append('<div class="page" id="'+route+'"></div>');
            }
            if (_.isEmpty(app.views[route])) {
                app.views[route] = new v({el: '#doc #'+scope+' .content #'+route});
            }
            this.$('.content #'+route).show();
            app.views[route].go();
        },
        /**
         * Primitive realization of cache.
         */
        renderIf: function () {
            var r = app.views[app.routers.app.scope].route;
            if (_.isEmpty(app.cache[r])) {
                this.render();
            }
        },
    });
});