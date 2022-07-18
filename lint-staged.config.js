module.exports = {
  '*.{js,json,css,scss,md,ts,html,graphql}': () => {
    return [
      `nx affected:lint --fix --parallel --uncommitted`,
      `nx format:write --uncommitted`,
    ]
  },
}
