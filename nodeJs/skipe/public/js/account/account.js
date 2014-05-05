define(['text!/js/account/t/account.layout.tpl.html'], function (t) {
    return  Backbone.skipeView.extend({
        el: '#doc #account',
        defaultRoute: 'account/home',
        model: new Backbone.skipeModel(),
        events:{
        },
        routes: {
            view_account_home: '/js/account/v/home.js',
        },
        initialize: function () {
            console.log('init account...');
            this.model.on('renderLayouts', this.renderLayouts, this);
        },
        goTo: function () {
            this.render();
        },
        render: function () {
            this.$el.html(_.template(t));
            this.model.hash = 'initLayouts';
            this.model.fetch({
                success: function (m, r) {
                    m.trigger('renderLayouts', r);
                }
            });
        },
        renderLayouts: function (args) {
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
