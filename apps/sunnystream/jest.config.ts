/* eslint-disable */
export default {
  displayName: 'sunnystream',
  preset: '../../jest.preset.js',
  transform: {
    '^.+\\.(js|jsx|ts|tsx|css|json)?$': 'ts-jest',
    '^(?!.*\\.svg$)': '@nx/react/plugins/jest',
  },
  moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx'],
  coverageDirectory: '../../coverage/apps/sunnystream',
}
