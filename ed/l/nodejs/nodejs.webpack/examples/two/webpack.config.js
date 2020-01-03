const path = require('path');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  context: __dirname,
  entry: [
    './src/app.js',
    './src/index.js',
  ],
  output: {
    path: path.resolve('build/'),
    filename: 'app.js'
  },

  module: {
    rules: [
      {
        test: /\.css$/,
        use: ExtractTextPlugin.extract({ fallback: 'style-loader', use: [ 'css-loader' ]})
      },
    ],
  },

  plugins: [
    new ExtractTextPlugin({ filename: '[name].css' }),
    new HtmlWebpackPlugin({ title: 'example two', hash: false, template: 'src/index.html' }),
  ]
};
