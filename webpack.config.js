const path = require("path");

module.exports = {
  mode: process.env.NODE_ENV === "production" ? "production" : "development",
  entry: {
    src: path.join(__dirname, "static/src/scripts"),
  },
  output: {
    path: path.join(__dirname, "static/dist"),
    filename: "index.js",
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: [
          "style-loader",
          "css-loader",
          {
            loader: "postcss-loader",
            options: {
              postcssOptions: {
                plugins: [
                  [
                    "postcss-preset-env",
                    {
                      // options
                    },
                  ],
                ],
              },
            },
          },
        ],
      },
      {
        test: /\.png|jpe?g|gif|webp|svg|bmp|tiff/,
        use: "file-loader",
      },
      {
        test: /\.(png|jpg|jpeg|gif)/i,
        type: "asset/resource",
      },
    ],
  },
};

