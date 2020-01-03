var mailgunTest = function() {
    var api_key = 'key-000';
    var domain = '000.mailgun.org';
    var mailgun = require('mailgun-js')({apiKey: api_key, domain: domain});

    var data = {
      from: 'FL <000@gmail.com>',
      to: '0002@gmail.com',
      subject: 'Hello',
      text: 'Testing some Mailgun awesomness!'
    };

    mailgun.messages().send(data, function (error, body) {
      console.log(body);
    });
}

var gmailTest = function() {
    var send = require('gmail-send')({
      user: 'codenamek2010@gmail.com',
      pass: 'asldfj9q8r8-qw0rjgkhv82939r0wogld2029t08u)(*)ilkkdglsdfkj',
      to:   'xxx@gmail.com',
      subject: 'test subject',
      text:    'test text'
      // html:    '<b>html text text</b>'
    })();
    console.log(send);
}

var SparkPostTest = function() {
    var SparkPost = require('sparkpost');
    var client = new SparkPost('');
    client.transmissions.send({
      content: {
        from: '000@gmail.com',
        subject: 'Hello, World!',
        html:'<html><body><p>Testing SparkPost - the world\'s most awesomest email service!</p></body></html>'
      },
      recipients: [
        {address: 'x@gmail.com'}
      ]
    })
    .then(data => {
      console.log('Woohoo! You just sent your first mailing!');
      console.log(data);
    })
    .catch(err => {
      console.log('Whoops! Something went wrong');
      console.log(err);
    });
}

SparkPostTest();
