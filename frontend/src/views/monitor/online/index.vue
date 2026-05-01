<template>
  <div class="online-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="登录名称">
          <el-input v-model="searchForm.userName" placeholder="请输入登录名称" clearable />
        </el-form-item>
        <el-form-item label="登录IP">
          <el-input v-model="searchForm.ipaddr" placeholder="请输入登录IP" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon>
            搜索
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="table-container">
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>在线用户列表</span>
        </div>
      </template>

      <el-table :data="onlineUserList" stripe v-loading="loading">
        <el-table-column prop="tokenId" label="会话编号" width="350" />
        <el-table-column prop="userName" label="用户名称" width="120" />
        <el-table-column prop="ipaddr" label="登录IP地址" width="150" />
        <el-table-column prop="loginLocation" label="登录地点" width="200" />
        <el-table-column prop="browser" label="浏览器" width="150" />
        <el-table-column prop="os" label="操作系统" width="150" />
        <el-table-column prop="loginTime" label="登录时间" width="180" />
        <el-table-column label="操作" width="80" fixed="right">
          <template #default="{ row }">
            <el-button type="danger" link @click="handleForceLogout(row)">强退</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="pagination.pageNum"
        v-model:page-size="pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="fetchData"
        @current-change="fetchData"
      />
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Search, Refresh } from '@element-plus/icons-vue'
import request from '@/utils/request'

const loading = ref(false)

const searchForm = reactive({
  userName: '',
  ipaddr: ''
})

const pagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const onlineUserList = ref([])

const fetchData = async () => {
  loading.value = true
  try {
    const res = await request.get('/monitor/online/list', {
      params: {
        ...searchForm,
        page_num: pagination.pageNum,
        page_size: pagination.pageSize
      }
    })
    onlineUserList.value = res.data || []
    pagination.total = res.data?.length || 0
  } catch (error) {
    console.error('获取在线用户列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.pageNum = 1
  fetchData()
}

const handleReset = () => {
  searchForm.userName = ''
  searchForm.ipaddr = ''
  handleSearch()
}

const handleForceLogout = (row) => {
  ElMessageBox.confirm('确定要强制退出该用户吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/monitor/online/${row.tokenId}`)
      ElMessage.success('强退成功')
      fetchData()
    } catch (error) {
      console.error('强退失败:', error)
    }
  }).catch(() => {})
}

onMounted(() => {
  fetchData()
})
</script>

<style lang="scss" scoped>
.online-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
