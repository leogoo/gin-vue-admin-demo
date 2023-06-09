import vuePlugin from '@vitejs/plugin-vue'
import * as dotenv from 'dotenv'
import * as fs from 'fs'
import Components from 'unplugin-vue-components/vite';
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers';

export default ({
  command,
  mode
}) => {
  const envFile = `.env.${mode}`;
  const envConfig = dotenv.parse(fs.readFileSync(envFile, 'utf8'));
  const config = {
    base: './', // index.html文件所在位置
    root: './', // js导入的资源路径，src
    define: {
      'process.env': {}
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
