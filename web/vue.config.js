const Components = require("unplugin-vue-components/webpack");
const {ElementPlusResolver} = require("unplugin-vue-components/resolvers");

module.exports = {
    configureWebpack: {
        plugins: [
            Components({
                resolvers: [ElementPlusResolver()],
            })
        ]
    },
    pluginOptions: {
        electronBuilder: {
            nodeIntegration: true
        }
    },
    devServer: {
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:30080',
                changeOrigin: true,
                pathRewrite: {
                    '^/api': ''
                }
            }
        }
    }
}
