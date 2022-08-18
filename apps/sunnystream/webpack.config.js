const { merge } = require('webpack-merge')
const ReactRefreshPlugin = require('@pmmmwh/react-refresh-webpack-plugin')
const { ModuleFederationPlugin } = require('webpack').container
const moduleFederationConfig = require('./module-federation.config')

module.exports = (config) => {
  if (
    config.mode === 'development' &&
    config['devServer'] &&
    config['devServer'].hot
  ) {
    const babelLoader = config.module.rules.find(
      (rule) =>
        typeof rule !== 'string' &&
        rule.loader &&
        rule.loader.toString().includes('babel-loader'),
    )

    if (babelLoader && typeof babelLoader !== 'string') {
      babelLoader.options['plugins'] = [
        ...(babelLoader.options['plugins'] || []),
        [
          require.resolve('react-refresh/babel'),
          {
            skipEnvCheck: true,
          },
        ],
      ]
    }

    if (config['devServer'].client && config['devServer'].client.webSocketURL) {
      delete config['devServer'].client.webSocketURL
    }

    config.plugins.push(new ReactRefreshPlugin())
  }

  if (process.env.npm_config_argv.includes('sunnystream:serve')) {
    config['optimization'].runtimeChunk = 'single'
  } else {
    config['optimization'].runtimeChunk = false
    config.plugins.push(new ModuleFederationPlugin(moduleFederationConfig))
  }

  return merge(config, {
    target: ['web', 'es2015'],
    output: {
      uniqueName: 'sunnystream',
      publicPath: 'auto',
    },
    optimization: {
      minimize: false,
    },
  })
}
