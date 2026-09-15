// postcss.config.cjs
module.exports = {
  plugins: {
    "postcss-import": {},
    "postcss-nested": {},
    "postcss-preset-env": {},
    ...(process.env.NODE_ENV === "production"
      ? {
          cssnano: {},
        }
      : {}),
  },
};

