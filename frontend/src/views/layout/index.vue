<template>
  <el-container class="layout-container">
    <el-aside :width="isCollapse ? '64px' : '220px'" class="sidebar">
      <div class="logo">
        <el-icon v-if="isCollapse" size="30" color="#fff"><VideoCamera /></el-icon>
        <span v-else class="logo-text">动漫商城管理</span>
      </div>
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :unique-opened="true"
        router
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409EFF"
      >
        <el-menu-item index="/dashboard">
          <el-icon><HomeFilled /></el-icon>
          <template #title>首页</template>
        </el-menu-item>

        <el-sub-menu index="system">
          <template #title>
            <el-icon><Setting /></el-icon>
            <span>系统管理</span>
          </template>
          <el-menu-item index="/system/user">
            <el-icon><User /></el-icon>
            <template #title>用户管理</template>
          </el-menu-item>
          <el-menu-item index="/system/dept">
            <el-icon><OfficeBuilding /></el-icon>
            <template #title>部门管理</template>
          </el-menu-item>
          <el-menu-item index="/system/post">
            <el-icon><Postcard /></el-icon>
            <template #title>岗位管理</template>
          </el-menu-item>
          <el-menu-item index="/system/menu">
            <el-icon><Menu /></el-icon>
            <template #title>菜单管理</template>
          </el-menu-item>
          <el-menu-item index="/system/role">
            <el-icon><UserFilled /></el-icon>
            <template #title>角色管理</template>
          </el-menu-item>
          <el-menu-item index="/system/dict">
            <el-icon><Collection /></el-icon>
            <template #title>字典管理</template>
          </el-menu-item>
          <el-menu-item index="/system/config">
            <el-icon><Tools /></el-icon>
            <template #title>参数管理</template>
          </el-menu-item>
          <el-menu-item index="/system/notice">
            <el-icon><Bell /></el-icon>
            <template #title>通知公告</template>
          </el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="monitor">
          <template #title>
            <el-icon><Monitor /></el-icon>
            <span>系统监控</span>
          </template>
          <el-menu-item index="/monitor/online">
            <el-icon><Avatar /></el-icon>
            <template #title>在线用户</template>
          </el-menu-item>
          <el-menu-item index="/monitor/job">
            <el-icon><Timer /></el-icon>
            <template #title>定时任务</template>
          </el-menu-item>
          <el-menu-item index="/monitor/operlog">
            <el-icon><Document /></el-icon>
            <template #title>操作日志</template>
          </el-menu-item>
          <el-menu-item index="/monitor/logininfor">
            <el-icon><Connection /></el-icon>
            <template #title>登录日志</template>
          </el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="mall">
          <template #title>
            <el-icon><Shop /></el-icon>
            <span>商城管理</span>
          </template>
          <el-menu-item index="/mall/user">
            <el-icon><User /></el-icon>
            <template #title>商城用户</template>
          </el-menu-item>
          <el-menu-item index="/mall/product">
            <el-icon><Goods /></el-icon>
            <template #title>商品管理</template>
          </el-menu-item>
          <el-menu-item index="/mall/order">
            <el-icon><ShoppingCart /></el-icon>
            <template #title>订单管理</template>
          </el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>

    <el-container>
      <el-header class="header">
        <div class="header-left">
          <el-icon class="collapse-icon" @click="toggleCollapse">
            <Fold v-if="!isCollapse" />
            <Expand v-else />
          </el-icon>
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/dashboard' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item v-if="currentRoute.meta?.title">
              {{ currentRoute.meta.title }}
            </el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <div class="header-right">
          <el-dropdown @command="handleCommand">
            <span class="user-info">
              <el-avatar :size="32" icon="UserFilled" />
              <span class="username">{{ userStore.userInfo?.nickname || '用户' }}</span>
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">
                  <el-icon><User /></el-icon>
                  个人中心
                </el-dropdown-item>
                <el-dropdown-item command="logout" divided>
                  <el-icon><SwitchButton /></el-icon>
                  退出登录
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <el-main class="main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/store/user'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const isCollapse = ref(false)
const currentRoute = computed(() => route)
const activeMenu = computed(() => route.path)

const toggleCollapse = () => {
  isCollapse.value = !isCollapse.value
}

const handleCommand = (command) => {
  if (command === 'logout') {
    userStore.logout()
  } else if (command === 'profile') {
    // 跳转到个人中心
  }
}
</script>

<style lang="scss" scoped>
.layout-container {
  height: 100vh;
}

.sidebar {
  background-color: #304156;
  transition: width 0.3s;
  overflow: hidden;

  .logo {
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #2b3a4a;
    border-bottom: 1px solid #1f2d3d;

    .logo-text {
      color: #fff;
      font-size: 18px;
      font-weight: bold;
      white-space: nowrap;
    }
  }
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

  .header-left {
    display: flex;
    align-items: center;

    .collapse-icon {
      font-size: 20px;
      cursor: pointer;
      margin-right: 20px;
      padding: 5px;
      border-radius: 4px;
      transition: background 0.3s;

      &:hover {
        background: #f5f7fa;
      }
    }
  }

  .header-right {
    .user-info {
      display: flex;
      align-items: center;
      cursor: pointer;
      padding: 5px 10px;
      border-radius: 4px;
      transition: background 0.3s;

      &:hover {
        background: #f5f7fa;
      }

      .username {
        margin: 0 8px;
        color: #606266;
      }
    }
  }
}

.main {
  background-color: #f0f2f5;
  padding: 20px;
  overflow-y: auto;
}

.el-menu {
  border-right: none;
}
</style>
