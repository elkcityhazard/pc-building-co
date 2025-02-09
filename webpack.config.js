const path = require("path");

if (process.env.NODE_ENV === "production") {
  new webpack.optimize.CommonsChunkPlugin("common.js"),
    new webpack.optimize.DedupePlugin(),
    new webpack.optimize.UglifyJsPlugin(),
    new webpack.optimize.AggressiveMergingPlugin();

  new CompressionPlugin({
    asset: "[path].gz[query]",
    algorithm: "gzip",
    test: /\.js$|\.css$|\.html$/,
    threshold: 10240,
    minRatio: 0.8,
  });
}

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
        test: /\.(png|jpe?g|gif|svg)$/i,
        use: [
          {
            loader: "url-loader",
            options: {
              limit: 8192, // Convert images < 8kb to base64 strings
              name: "[name].[hash].[ext]", // Output file names
            },
          },
          {
            loader: "image-webpack-loader",
            options: {
              mozjpeg: {
                progressive: true,
                quality: 65,
              },
              optipng: {
                enabled: false,
              },
              pngquant: {
                quality: [0.65, 0.9],
                speed: 4,
              },
              gifsicle: {
                interlaced: false,
              },
              // Disable webp option if you don't need it and prefer jpeg/png
              webp: {
                quality: 75,
              },
            },
          },
        ],
      },
    ],
  },
};

