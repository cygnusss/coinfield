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
            loader: "ts-loader",
            exclude: /node_modules/
          }, {
            test: /\.css$/,
            loader: 'style-loader'
          }, {
            test: /\.css$/,
            loader: 'css-loader',
            query: {
              modules: true,
              localIdentName: '[name]__[local]___[hash:base64:5]'
            }
          }
      ]       
  }
}

export default config