<html>
  <script>
    var eSender = {
      queue: [],
      send: function (m) {
        this.queue.push(m);
      },
      flush: function () {
        setTimeout(eSender.flush, 25);
        if (typeof(_io_track) !== 'function') {
          return;
        }
        while (eSender.queue.length > 0) {
          _io_track(eSender.queue.pop());
        }
      },
    }
    eSender.flush();
    // Previous error handler.
    var p = window.onerror;
    window.onerror = function(m, u, l) {
      if (typeof(p) === 'function') {
          p(m, u, l);
      }
      var msg = 'Message: '+m+', Url: '+u+', LineNumber: '+l;
      console.log(msg);
      return true;
    };
    window.addEventListener('error', function (e) {
      if (e && e.target && e.target.nodeName) {
        var n, u;
        n = e.target.nodeName.toLowerCase();
        if (n === 'link') {
          u = e.target.href;
        } else {
          u = e.target.src;
        }
        var msg = 'Message: error at '+n+', Url: '+u;
        console.log(msg);
      }
    }, true);
    // Another vision.
    if (window.addEventListener) {
      window.addEventListener('error', trackJavaScriptError, false);
    } else if (window.attachEvent) {
      window.attachEvent('onerror', trackJavaScriptError);
    } else {
      window.onerror = trackJavaScriptError;
    }
    // Error.
    cnsl.lg(200);
  </script>
  <body>
    <link rel="stylesheet" href="https://onthe.io/css-bdly.css2.s" type="text/css" media="all" />
    <img src="https://www.gravatar.com/avatar/b43bbab9d16d030a8744f17ca9c1da8d?s=64&d=identicon&r=PG" alt="OK">
    <img src="//cdn.sstatic.net1s/img1.png">
    <img src="//cdn.sstatic.netS2/img2.png">
    <script src="//cdn.ravenjs.com/1.1.19/jquery,native/raven.min.js">/* OK */</script>
    <script src="//cdn.sstatic.net/Js/stubs.en.js?v=911"></script>
    <link rel="stylesheet" type="text/css" href="//cdn.sstatic.net/stackoverflow/all.css.css?v=922">
  </body>
</html>
