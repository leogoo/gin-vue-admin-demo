import vuePlugin from '@vitejs/plugin-vue'
import * as path from 'path'
import Components from 'unplugin-vue-components/vite';
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers';

export default ({
  command,
  mode
}) => {
  const config = {
    base: './', // index.html文件所在位置
    root: './', // js导入的资源路径，src
    define: {
      'process.env': {}
    },
    resolve: {
      alias: {
        '@': path.resolve(__dirname, './src'),
      }
    },
    plugins: [
      vuePlugin(),
      Components({
        dirs: ["src/components"],
        extensions: ["vue"],
        // 配置文件生成位置
        dts: 'src/components.d.ts',
        resolvers: [
          ElementPlusResolver({
            importStyle: "sass"
          })
        ]
      })
    ],
  }

  return config
}
