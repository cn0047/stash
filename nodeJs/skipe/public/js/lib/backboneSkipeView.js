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
    });
});