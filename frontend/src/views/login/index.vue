<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <h1 class="title">动漫商城管理系统</h1>
        <p class="subtitle">Cartoon Mall Management System</p>
      </div>
      <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" class="login-form">
        <el-form-item prop="username">
          <el-input v-model="loginForm.username" placeholder="请输入用户名" prefix-icon="User" size="large" />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="loginForm.password"
            type="password"
            placeholder="请输入密码"
            prefix-icon="Lock"
            size="large"
            @keyup.enter="handleLogin"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="large" :loading="loading" @click="handleLogin" style="width: 100%">
            登 录
          </el-button>
        </el-form-item>
        <div class="login-tips">
          <p>默认账号: admin</p>
          <p>默认密码: admin123</p>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/store/user'

const userStore = useUserStore()
const loginFormRef = ref(null)
const loading = ref(false)

const loginForm = reactive({
  username: 'admin',
  password: 'admin123'
})

const loginRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ]
}

const handleLogin = async () => {
  if (!loginFormRef.value) return
  
  await loginFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        await userStore.login(loginForm)
      } catch (error) {
        console.error('Login error:', error)
      } finally {
        loading.value = false
      }
    } else {
      ElMessage.error('请填写完整的登录信息')
    }
  })
}
</script>

<style lang="scss" scoped>
.login-container {
  height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-box {
  width: 400px;
  background: #fff;
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.login-header {
  text-align: center;
  margin-bottom: 30px;

  .title {
    font-size: 28px;
    font-weight: bold;
    color: #333;
    margin-bottom: 10px;
  }

  .subtitle {
    font-size: 14px;
    color: #999;
  }
}

.login-form {
  .el-form-item {
    margin-bottom: 20px;
  }
}

.login-tips {
  text-align: center;
  color: #999;
  font-size: 12px;
  margin-top: 20px;

  p {
    margin: 5px 0;
  }
}
</style>
