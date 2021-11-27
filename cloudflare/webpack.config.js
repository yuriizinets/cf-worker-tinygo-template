const CopyPlugin = require('copy-webpack-plugin')
const path = require('path')
const webpack = require('webpack')
const spawn = require('child_process').spawnSync

module.exports = {
  context: path.resolve(__dirname, '.'),
  devtool: 'nosources-source-map',
  entry: './index.js',
  target: 'webworker',
  plugins: [
    new webpack.IgnorePlugin({ resourceRegExp: /fs/ }),
    new webpack.IgnorePlugin({ resourceRegExp: /crypto/ }),
    new webpack.IgnorePlugin({ resourceRegExp: /util/ }),
    new webpack.IgnorePlugin({ resourceRegExp: /os/ }),
  ]
}
