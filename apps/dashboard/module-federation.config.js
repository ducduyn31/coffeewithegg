module.exports = {
  name: 'dashboard',
  remotes: ['ops', 'sunnystream'],
  shared: (libraryName, sharedConfig) => {
    // Handle some libraries that have secondary package.json file without version
    if (
      ['@apollo/client/core', '@apollo/client/link/batch'].includes(libraryName)
    ) {
      return { singleton: true, requiredVersion: '0' }
    }

    return sharedConfig
  },
}
