const { merge } = require('webpack-merge');
const { ModuleFederationPlugin } = require('webpack').container;
const moduleFederationConfig = require('./module-federation.config');

module.exports = (config) => {
  return merge(config, {
    target: [ 'web', 'es2015' ],
    output: {
      uniqueName: 'sunnystream',
      publicPath: 'auto',
    },
    optimization: { runtimeChunk: false, minimize: false },
    plugins: [
      new ModuleFederationPlugin(moduleFederationConfig),
    ]
  });
}
