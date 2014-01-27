/*
$(function() {
	window.app = {
		models: {},
		collections: {},
		views: {},
	};

	window.tpl = function(id){
		return _.template($('#'+id).html());
	};

	app.models.Task = Backbone.Model.extend({
		validate: function(attrs){
			// console.log(attrs.title.length);
			if (attrs.t7itle.length < 2) {
				return 'Very short text.';
			}
		}
	});

	app.collections.Task = Backbone.Collection.extend({model: app.models.Task});

	app.views.Task = Backbone.View.extend({
		tagName: 'li',
		template: tpl('tTemplate'),
		initialize: function(){
			// _.bindAll(this, 'editTask', 'render');
			this.model.on('change', this.render, this);
		},
		render: function(){
			this.$el.html(this.template(this.model.toJSON()));
			return this;
		},
		events:{
			'click .edit': 'editTask',
			'click .delete': 'deleteTask',
		},
		editTask: function(){
			var newTT = prompt('New task text:', this.model.get('title'));
			if (newTT) {
				this.model.save('title', newTT);
			}
		},
		deleteTask: function(){
			
		},
	});

	app.views.Tasks = Backbone.View.extend({
		tagName: 'ul',
		render: function(){
			this.collection.each(this.addOne, this);
			return this;
		},
		addOne: function(one, t){
			var tv = new app.views.Task({model:one});
			this.$el.append(tv.render().el);
		},
	});



	window.c = new app.collections.Task([
		{title:'Watch mail',priority:1},
		{title:'Watch errors',priority:2},
		{title:'Watch graphs',priority:3},
	]);
	var v = new app.views.Tasks({collection:c});
	$('#t').html(v.render().el);
	// console.log(v..el);
});
*/

$(function() {
    window.app = {
        models: {},
        collections: {},
        views: {},
        router: {},
    };

    window.tpl = function(id){
        return _.template($('#'+id).html());
    };

    app.models.Task = Backbone.Model.extend({
        defaults: {
            title: 'no text',
            priority: 0,
        },
        // validate: functcion(attrs){
            // console.log(attrs.title.length);
            // if (attrs.title.length < 2) {
            //     return 'Very short text.';
            // }
        // }
    });

    app.collections.Task = Backbone.Collection.extend({model: app.models.Task});

    app.views.Task = Backbone.View.extend({
        tagName: 'li',
        template: tpl('tTemplate'),
        initialize: function(){
            // _.bindAll(this, 'editTask', 'render');
            this.model.on('change', this.render, this);
            this.model.on('destroy', this.remove, this);
        },
        render: function(){
            this.$el.html(this.template(this.model.toJSON()));
            return this;
        },
        events:{
            'click .edit': 'editTask',
            'click .delete': 'deleteTask',
        },
        editTask: function(){
            var newTT = prompt('New task text:', this.model.get('title'));
            if (newTT) {
                this.model.save('title', newTT);
            }
        },
        deleteTask: function(){
            // alert(200);
            this.model.destroy();
            console.log(c);
        },
        remove: function(){
            this.$el.remove();
        },
    });

    app.views.Tasks = Backbone.View.extend({
        tagName: 'ul',
        initialize: function(){
            this.collection.on('add', this.addOne, this);
        },
        render: function(){
            this.collection.each(this.addOne, this);
            return this;
        },
        addOne: function(one, t){
            var tv = new app.views.Task({model:one});
            this.$el.append(tv.render().el);
        },
    });

    app.views.addTask = Backbone.View.extend({
        el: '#addTask',
        events: {
            'submit': 'submit',
        },
        initialize: function(){
            // console.log(this.el.innerHTML);
        },
        submit: function(e){
            // alert(200);
            e.preventDefault();
            var t = $(e.currentTarget).find('input[type=text]').val();
            var T = new app.models.Task({title:t});
            this.collection.add(T);
            // console.log(T);
        },
    });

    app.router = Backbone.Router.extend({
        routes: {
            '': 'index',
            'read': 'read',
            'page/:id': 'page',
            'doc/:id/*sid': 'doc',
            '*other': 'default',
        },
        index: function(){
            console.log('index');
        },
        read: function(){
            console.log('read');
        },
        page: function(id){
            console.log('page '+id);
        },
        doc: function(id, sid){
            console.log('doc '+id+' sub id '+sid);
        },
        default: function(other){
            alert(404+' = '+other);
        },
    });


    window.c = new app.collections.Task([
        {title:'Watch mail',priority:1},
        {title:'Watch errors',priority:2},
        {title:'Watch graphs',priority:3},
    ]);
    var v = new app.views.Tasks({collection:c});
    $('#t').html(v.render().el);
    var t2 = new app.models.Task({title:'Read book',priority:5});
    var v2 = new app.views.addTask({collection:c});
    c.add(t2);

    new app.router();
    Backbone.history.start();
}); 