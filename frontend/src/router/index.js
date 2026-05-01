import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/store/user'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: { title: '登录' }
  },
  {
    path: '/',
    component: () => import('@/views/layout/index.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index.vue'),
        meta: { title: '首页', icon: 'HomeFilled' }
      },
      {
        path: 'system/user',
        name: 'UserManage',
        component: () => import('@/views/system/user/index.vue'),
        meta: { title: '用户管理', icon: 'User' }
      },
      {
        path: 'system/dept',
        name: 'DeptManage',
        component: () => import('@/views/system/dept/index.vue'),
        meta: { title: '部门管理', icon: 'OfficeBuilding' }
      },
      {
        path: 'system/post',
        name: 'PostManage',
        component: () => import('@/views/system/post/index.vue'),
        meta: { title: '岗位管理', icon: 'Postcard' }
      },
      {
        path: 'system/menu',
        name: 'MenuManage',
        component: () => import('@/views/system/menu/index.vue'),
        meta: { title: '菜单管理', icon: 'Menu' }
      },
      {
        path: 'system/role',
        name: 'RoleManage',
        component: () => import('@/views/system/role/index.vue'),
        meta: { title: '角色管理', icon: 'UserFilled' }
      },
      {
        path: 'system/dict',
        name: 'DictManage',
        component: () => import('@/views/system/dict/index.vue'),
        meta: { title: '字典管理', icon: 'Collection' }
      },
      {
        path: 'system/config',
        name: 'ConfigManage',
        component: () => import('@/views/system/config/index.vue'),
        meta: { title: '参数管理', icon: 'Setting' }
      },
      {
        path: 'system/notice',
        name: 'NoticeManage',
        component: () => import('@/views/system/notice/index.vue'),
        meta: { title: '通知公告', icon: 'Bell' }
      },
      {
        path: 'monitor/operlog',
        name: 'OperLog',
        component: () => import('@/views/monitor/operlog/index.vue'),
        meta: { title: '操作日志', icon: 'Document' }
      },
      {
        path: 'monitor/logininfor',
        name: 'LoginLog',
        component: () => import('@/views/monitor/logininfor/index.vue'),
        meta: { title: '登录日志', icon: 'Connection' }
      },
      {
        path: 'monitor/online',
        name: 'OnlineUser',
        component: () => import('@/views/monitor/online/index.vue'),
        meta: { title: '在线用户', icon: 'Avatar' }
      },
      {
        path: 'monitor/job',
        name: 'JobManage',
        component: () => import('@/views/monitor/job/index.vue'),
        meta: { title: '定时任务', icon: 'Timer' }
      },
      {
        path: 'mall/user',
        name: 'MallUserManage',
        component: () => import('@/views/mall/user/index.vue'),
        meta: { title: '商城用户', icon: 'User' }
      },
      {
        path: 'mall/product',
        name: 'ProductManage',
        component: () => import('@/views/mall/product/index.vue'),
        meta: { title: '商品管理', icon: 'Goods' }
      },
      {
        path: 'mall/order',
        name: 'OrderManage',
        component: () => import('@/views/mall/order/index.vue'),
        meta: { title: '订单管理', icon: 'ShoppingCart' }
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore()
  const token = localStorage.getItem('token')

  if (to.path === '/login') {
    if (token) {
      next('/dashboard')
    } else {
      next()
    }
  } else {
    if (!token) {
      next('/login')
    } else {
      if (!userStore.userInfo) {
        try {
          await userStore.getUserInfo()
        } catch (error) {
          localStorage.removeItem('token')
          next('/login')
          return
        }
      }
      next()
    }
  }
})

export default router
