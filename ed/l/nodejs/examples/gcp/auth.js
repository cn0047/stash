const {JWT} = require('google-auth-library');
const {auth} = require('google-auth-library');

const u = `https://apigee.googleapis.com/v1/organizations/${ORG}/apis`;

const fromFile() => {
  const keys = require('./serviceAccount.json');
  const c = new JWT({
    email: keys.client_email,
    key: keys.private_key,
    scopes: ['https://www.googleapis.com/auth/cloud-platform'],
  });
  const res = await c.request({url: u});
  console.log(res.data);
}

const fromEnv() => {
  const envKeys = process.env['SA'];
  if (!envKeys) {
    throw new Error('$SA environment variable is required');
  }
  const keys = JSON.parse(envKeys);
  const c = auth.fromJSON(keys);
  c.scopes = ['https://www.googleapis.com/auth/cloud-platform'];

  const res = await c.request({url: u});
  console.log(res.data);
}
