var screenshot = require('url-to-image');
screenshot('http://google.com', 'google.png').done(function() {
    console.log('http://google.com screenshot saved to google.png');
});