$(function() {
    window.app = {
        models: {},
        collections: {},
        views: {},
        router: {},
    };

    var vent = new _.extend({}, Backbone.Events);

    app.views.sT = Backbone.View.extend({
        initialize: function(){
            vent.on('_sT', this._sT, this);
        },
        _sT: function(id){
            var m = this.collection.get('id');
            var v = app.views.sT({model:m});
            $('body').append(v.render().el);
            console.log(m);
        },
    });

    app.router = Backbone.Router.extend({
        routes: {
            '': 'index',
            'specialTasks/:id': 'sT',
        },
        index: function(){
            console.log('index');
        },
        sT: function(id){
            // console.log(id);
            vent.trigger('_sT', id);
        },
    });

    app.collections.Task = Backbone.Collection.extend({
        model: app.models.Task,
        url: '/tasks'
    });



    app.models.Task = Backbone.Model.extend({
        defaults: {
            id: 0,
            title: 'Watch Backbone',
            priority: 1,
        },
        urlRoot: 'core.php',
        url: function(){
            return this.urlRoot + "?id=" + this.id;
        },
        parse: function(r){
            // console.log();
            this.set('title', r.title);
            // console.log(this.get('title'));
        },
        events:{
            'change': 'change',
        },
        change: function(){
            console.log('change()');
        }
    });

    // new app.views.sT({collection:c});
    // new app.router();
    // Backbone.history.start();

    var t = new app.models.Task({id:2});
    // console.log(t.toJSON());
    t.fetch();
    console.log(t);
    // t.trigger('change');
    // console.log(t.attributes);
    // t.set({title:'Updated: '+t.get('title'), id:200});
    // console.log(t.save());

    // var c = new app.collections.Task();
    // c.fetch();

});