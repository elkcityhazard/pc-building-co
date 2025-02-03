module.exports = {
    plugins: [
      require("postcss-import"),
      require("postcss-nested"),
      require("autoprefixer"),
      require("postcss-preset-env", {features: {}}),
      ...(process.env.NODE_ENV === "production" ? [require("cssnano")] : []),
    ],
  };