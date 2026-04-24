const fs = require('fs');
const path = require('path');
const sgMail = require('@sendgrid/mail');

const srcFile = 'someFile.pdf';
const filePath = path.join(__dirname, srcFile);
const fileAttachment = fs.readFileSync(filePath).toString('base64');
const msg = {
  to: 'x@email.com',
  from: 'y@email.com',
  subject: 'test',
  text: 'Test',
  html: '<strong>Test</strong>',
  attachments: [
    {
      content: fileAttachment,
      filename: srcFile,
      type: 'application/pdf',
      disposition: 'attachment',
    },
  ],
};

sgMail.setApiKey('API_KEY');
sgMail.send(msg)
  .then(() => console.log('Email sent with attachment'))
  .catch((error) => console.error(error.response.body));
