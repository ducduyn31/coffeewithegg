const { withModuleFederation } = require('@nx/angular/module-federation')
const config = require('./module-federation.config')
module.exports = withModuleFederation({
  ...config,
  /*
   * Remote overrides for production.
   * Each entry is a pair of an unique name and the URL where it is deployed.
   *
   * e.g.
   * remotes: [
   *   ['app1', 'https://app1.example.com'],
   *   ['app2', 'https://app2.example.com'],
   * ]
   */
  remotes: [
    ['ops', 'https://ops.coffeewithegg.com'],
    ['sunnystream', 'https://sunnystream.coffeewithegg.com/'],
  ],
})
