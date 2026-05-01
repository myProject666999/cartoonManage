import { defineStore } from 'pinia'
import request from '@/utils/request'
import { ElMessage } from 'element-plus'
import router from '@/router'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    userInfo: null,
    menus: []
  }),

  actions: {
    async login(loginForm) {
      const res = await request.post('/auth/login', loginForm)
      if (res.code === 200) {
        this.token = res.data.token
        this.userInfo = res.data.user
        localStorage.setItem('token', res.data.token)
        ElMessage.success('登录成功')
        router.push('/dashboard')
      }
      return res
    },

    async getUserInfo() {
      const res = await request.get('/auth/info')
      if (res.code === 200) {
        this.userInfo = res.data
      }
      return res
    },

    async logout() {
      try {
        await request.post('/auth/logout')
      } catch (e) {
        console.error('Logout error:', e)
      }
      this.token = ''
      this.userInfo = null
      localStorage.removeItem('token')
      ElMessage.success('退出登录成功')
      router.push('/login')
    }
  }
})
