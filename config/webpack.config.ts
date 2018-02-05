import * as path from "path"
import * as webpack from "webpack"

const DIST = path.join(__dirname, "../client/dist"),
  SRC = path.join(__dirname, "../client/src")

const config: webpack.Configuration = {
  entry: path.join(SRC, "index.tsx"),
  output: {
    filename: "bundle.js",
    path: DIST,
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".jsx"]
  },
  module: {
      loaders: [
          {
          test: /\.tsx?$/,
          loader: "awesome-typescript-loader",
          exclude: /node_modules/
          }
      ]       
  }
}

export default config