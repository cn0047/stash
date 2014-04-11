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
        goTo: function (route) {
            if (_.isEmpty(route)) {
                this.goTo(this.defaultRoute);
                return;
            }
            app.views.app.showLoading();
            this.$('#content > div').hide();
            this.route = route.replace('/', '_');
            var r = 'view_'+this.route;
            if (r in this.routes) {
                if (_.isEmpty(app.views[this.route])) {
                    require([this.routes[r]], app.views[app.routers.app.scope].go);
                } else {
                    this.go();
                }
            } else {
                this.goTo(this.defaultRoute);
            }
        },
        go: function (v) {
            var scope = app.routers.app.scope;
            var route = app.views[scope].route;
            if (_.isEmpty(app.views[route])) {
                app.views[scope].$el.find('#content').append('<div id="'+route+'"></div>');
                app.views[route] = new v({el: '#doc #'+scope+' #content #'+route});
            }
            app.views[route].goTo();
        },
    });
});