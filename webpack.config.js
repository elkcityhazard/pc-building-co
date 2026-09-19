const path = require("path");
const { experiments } = require("webpack");

module.exports = {
  mode: process.env.NODE_ENV === "production" ? "production" : "development",
  entry: {
    src: path.join(__dirname, "static/src/scripts"),
  },
  output: {
    path: path.join(__dirname, "static/dist"),
    filename: "index.js",
    cssFilename: "css/custom-ps.css"
  },
  module: {
    rules: [
    ],
  },
  experiments: {
    css: true,
  }
};
