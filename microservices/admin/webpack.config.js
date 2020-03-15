module.exports = {
  entry: './src/js/client.js',
  output: {
      path: __dirname + '/public',
      filename: 'bundle.js',
  },
  resolve: {
    modules: [
      'node_modules',
    ],
  },
  module: {
    rules: [
      {
        test: /\.css/,
        use: [
          'style-loader',
          {
            loader: 'css-loader',
            options: { url: false }
          }
        ]
      }
    ]
  }
};
