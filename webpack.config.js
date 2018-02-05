const path = require("path"),
  DIST = path.join(__dirname, "./client/dist"),
  SRC = path.join(__dirname, "./client/src")

module.exports = {
  entry: path.join(SRC, "index.tsx"),
  output: {
    filename: "bundle.js",
    path: DIST,
  },
  module: {
    loaders: [
      { test: /\.tsx$/, loader: 'awesome-typescript-loader' },
    ],
  }
}