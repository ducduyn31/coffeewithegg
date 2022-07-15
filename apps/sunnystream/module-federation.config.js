module.exports = {
  name: 'sunnystream',
  filename: 'remoteEntry.js',
  exposes: {
    './Module': './src/remote-entry.tsx',
  },
  library: { type: 'var', name: 'sunnystream' },
  shared: {
    'react': { singleton: true, strictVersion: true, requiredVersion: '18.2.0' },
    'react-dom': { singleton: true, strictVersion: true, requiredVersion: '18.2.0' },
    'core-js': { singleton: true, strictVersion: true, requiredVersion: '^3.6.5' },
    'regenerator-runtime': { singleton: true, strictVersion: true, requiredVersion: '0.13.7' },
    '@nrwl/react': { singleton: true, strictVersion: true, requiredVersion: '^14.3.6' }
  },
}
