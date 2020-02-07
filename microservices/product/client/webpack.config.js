module.exports = {
  entry: './client.js',
  output: {
      path: __dirname + '/dist',
      filename: 'main.js',
  },
  resolve: {
    modules: [
      'node_modules',
    ],
  },
};
