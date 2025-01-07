module.exports = {
  publicPath: process.env.NODE_ENV === 'production' ? '/go-webapp/' : '/',
  outputDir: '../backend/public',
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        pathRewrite: { '^/api': '' },
      },
    },
  },
  configureWebpack: {
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src'),
      },
    },
  },
};