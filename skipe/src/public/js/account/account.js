define(['text!/js/account/t/account.layout.tpl.html'], function (t) {
    return  Backbone.skipeView.extend({
        el: '#doc #account',
        tpl: t,
        defaultRoute: 'account/home',
        model: new Backbone.skipeModel(),
        events:{
        },
        routes: {
            view_account_home: '/js/account/v/home.js',
        },
        initialize: function () {
            console.log('init account...');
            app.views.app.on('renderLayouts', this.renderLayouts, this);
            this.model.on('afterInitLayouts', this.afterInitLayouts, this);
        },
        renderLayouts: function () {
            this.$el.html(_.template(this.tpl));
            this.model.hash = 'initLayouts';
            this.model.fetch({
                success: function (m, r) {
                    m.trigger('afterInitLayouts', r);
                }
            });
        },
        afterInitLayouts: function (args) {
            _.each(args.menu, function (v, k) {
                var html = "<li><a href='#'>"+k+" "+app.nls[v]+"</a></li>";
                if (v == 'divider') {
                    html = "<li class='divider'></li>";
                }
                this.$('.navbar #menu').append(html);
            });
        },
    });
});
