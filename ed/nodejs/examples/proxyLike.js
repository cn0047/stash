// Proxy object's app method doStuff

(function(doStuff) {
  app.doStuff = function() {
    console.log('app.doStuff args', arguments);
    return doStuff.apply(this, arguments);
  };
})(app.doStuff);
