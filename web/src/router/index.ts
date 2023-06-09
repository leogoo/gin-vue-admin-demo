import { createRouter, createWebHashHistory} from 'vue-router';

const routes = [
  {
    path: '/init',
    name: 'Init',
    meta: { title: '数据库初始化'},
    component: () => import('@/view/Init.vue')
  },
  {
    path: '/login',
    name: 'Login',
    meta: { title: '登录页'},
    component: () => import('@/view/Login.vue')
  },
  {
    path: '/',
    name: 'home',
    meta: { title: '主页'},
    component: () => import('@/view/Home.vue')
  },
]
// 创建路由实例并传递 `routes` 配置
const router = createRouter({
  // 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title as string;
  }
  next();
})

export { router };