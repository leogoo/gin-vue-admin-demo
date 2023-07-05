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
    name: 'layout',
    component: () => import('@/view/Layout.vue'),
    children: [
      {
        path: 'admin',
        meta: { title: '超级管理员'},
        children: [
          {
            path: 'authority',
            meta: { title: '角色管理'},
            component: () => import('@/view/Admin/Authority.vue'),
          },
          {
            path: 'user',
            meta: { title: '用户管理'},
            component: () => import('@/view/Admin/User.vue'),
          },
          {
            path: 'operation',
            meta: { title: '操作历史'},
            component: () => import('@/view/Admin/Operation.vue'),
          },
        ]
      },
      {
        path: 'media',
        meta: { title: '媒体库'},
        children: [
          {
            path: 'fileList',
            meta: { title: '文件列表'},
            component: () => import('@/view/Media/FileList.vue'),
          },
        ]
      },
    ]
  }
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