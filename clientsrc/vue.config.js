const path = require("path");

module.exports = {
    outputDir: path.resolve(__dirname, "../client"),
    configureWebpack: {
      module: {
        rules: [
          {
            test: /.html$/,
            loader: "vue-template-loader",
            exclude: /index.html/
          }
        ]
      }
    },
    devServer: {
        watchOptions: {
          poll: true
        }
      }
  }
