<template>
  <div class="logininfor-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="用户账号">
          <el-input v-model="searchForm.userName" placeholder="请输入用户账号" clearable />
        </el-form-item>
        <el-form-item label="登录IP">
          <el-input v-model="searchForm.ipaddr" placeholder="请输入登录IP" clearable />
        </el-form-item>
        <el-form-item label="登录状态">
          <el-select v-model="searchForm.status" placeholder="请选择登录状态" clearable>
            <el-option label="成功" value="0" />
            <el-option label="失败" value="1" />
          </el-select>
        </el-form-item>
        <el-form-item label="登录时间">
          <el-date-picker
            v-model="dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
          />
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
          <el-button type="danger" @click="handleBatchDelete">
            <el-icon><Delete /></el-icon>
            清空
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="table-container">
      <el-table :data="loginLogList" stripe v-loading="loading">
        <el-table-column prop="infoId" label="访问ID" width="100" />
        <el-table-column prop="userName" label="用户账号" width="150" />
        <el-table-column prop="ipaddr" label="登录IP地址" width="150" />
        <el-table-column prop="loginLocation" label="登录地点" width="200" />
        <el-table-column prop="browser" label="浏览器" width="150" />
        <el-table-column prop="os" label="操作系统" width="150" />
        <el-table-column prop="status" label="登录状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '成功' : '失败' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="msg" label="提示消息" width="200" show-overflow-tooltip />
        <el-table-column prop="loginTime" label="登录时间" width="180" />
        <el-table-column label="操作" width="80" fixed="right">
          <template #default="{ row }">
            <el-button type="danger" link @click="handleDelete(row)">删除</el-button>
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
import { Search, Refresh, Delete } from '@element-plus/icons-vue'
import request from '@/utils/request'

const loading = ref(false)
const dateRange = ref([])

const searchForm = reactive({
  userName: '',
  ipaddr: '',
  status: ''
})

const pagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const loginLogList = ref([])

const fetchData = async () => {
  loading.value = true
  try {
    const params = {
      ...searchForm,
      page_num: pagination.pageNum,
      page_size: pagination.pageSize
    }
    if (dateRange.value && dateRange.value.length === 2) {
      params.beginTime = dateRange.value[0]
      params.endTime = dateRange.value[1]
    }
    const res = await request.get('/monitor/logininfor/list', { params })
    loginLogList.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch (error) {
    console.error('获取登录日志列表失败:', error)
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
  searchForm.status = ''
  dateRange.value = []
  handleSearch()
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该日志吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/monitor/logininfor/${row.infoId || row.id}`)
      ElMessage.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('删除失败:', error)
    }
  }).catch(() => {})
}

const handleBatchDelete = () => {
  ElMessageBox.confirm('确定要清空所有日志吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete('/monitor/logininfor')
      ElMessage.success('清空成功')
      fetchData()
    } catch (error) {
      console.error('清空失败:', error)
    }
  }).catch(() => {})
}

onMounted(() => {
  fetchData()
})
</script>

<style lang="scss" scoped>
.logininfor-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
