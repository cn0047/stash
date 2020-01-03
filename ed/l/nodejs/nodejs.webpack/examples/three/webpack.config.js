const path = require('path');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  context: __dirname,
  entry: [
    './src/js/app.js',
    './src/js/index.js',
  ],
  output: {
    path: path.resolve('build/'),
    filename: 'bundle.js'
  },

  module: {
    rules: [
      {
        test: /\.js$/,
        use: { loader: 'babel-loader' }
      },
      {
        test: /\.css$/,
        use: ExtractTextPlugin.extract({ fallback: 'style-loader', use: [ 'css-loader' ]})
      },
    ],
  },

  plugins: [
    new ExtractTextPlugin({ filename: '[name].css' }),
    new HtmlWebpackPlugin({ title: 'example 3', hash: false, template: 'src/index.html' }),
  ]
};
