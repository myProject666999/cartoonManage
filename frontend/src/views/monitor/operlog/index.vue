<template>
  <div class="operlog-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="操作模块">
          <el-input v-model="searchForm.title" placeholder="请输入操作模块" clearable />
        </el-form-item>
        <el-form-item label="操作用户">
          <el-input v-model="searchForm.operName" placeholder="请输入操作用户" clearable />
        </el-form-item>
        <el-form-item label="操作状态">
          <el-select v-model="searchForm.status" placeholder="请选择操作状态" clearable>
            <el-option label="成功" value="0" />
            <el-option label="失败" value="1" />
          </el-select>
        </el-form-item>
        <el-form-item label="操作时间">
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
      <el-table :data="operLogList" stripe v-loading="loading">
        <el-table-column prop="id" label="日志ID" width="80" />
        <el-table-column prop="title" label="操作模块" width="150" />
        <el-table-column prop="operName" label="操作用户" width="120" />
        <el-table-column prop="businessType" label="业务类型" width="100">
          <template #default="{ row }">
            <el-tag :type="getBusinessTypeClass(row.businessType)">
              {{ getBusinessTypeLabel(row.businessType) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="method" label="请求方法" width="200" show-overflow-tooltip />
        <el-table-column prop="requestMethod" label="请求方式" width="100" />
        <el-table-column prop="operatorType" label="操作类型" width="100">
          <template #default="{ row }">
            <el-tag :type="row.operatorType === '1' ? 'primary' : 'success'">
              {{ row.operatorType === '1' ? '后台用户' : '手机端用户' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="操作状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '成功' : '失败' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="operTime" label="操作时间" width="180" />
        <el-table-column prop="costTime" label="消耗时间" width="100">
          <template #default="{ row }">
            {{ row.costTime ? row.costTime + 'ms' : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleView(row)">详情</el-button>
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

    <el-dialog v-model="detailVisible" title="操作日志详情" width="800px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="日志编号">{{ detailData.id }}</el-descriptions-item>
        <el-descriptions-item label="操作模块">{{ detailData.title }}</el-descriptions-item>
        <el-descriptions-item label="业务类型">{{ getBusinessTypeLabel(detailData.businessType) }}</el-descriptions-item>
        <el-descriptions-item label="请求方式">{{ detailData.requestMethod }}</el-descriptions-item>
        <el-descriptions-item label="请求方法">{{ detailData.method }}</el-descriptions-item>
        <el-descriptions-item label="操作方式">
          <el-tag :type="detailData.operatorType === '1' ? 'primary' : 'success'">
            {{ detailData.operatorType === '1' ? '后台用户' : '手机端用户' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="操作用户">{{ detailData.operName }}</el-descriptions-item>
        <el-descriptions-item label="部门名称">{{ detailData.deptName || '-' }}</el-descriptions-item>
        <el-descriptions-item label="请求URL" :span="2">{{ detailData.operUrl }}</el-descriptions-item>
        <el-descriptions-item label="主机地址" :span="1">{{ detailData.operIp }}</el-descriptions-item>
        <el-descriptions-item label="操作地点" :span="1">{{ detailData.operLocation || '-' }}</el-descriptions-item>
        <el-descriptions-item label="操作状态" :span="1">
          <el-tag :type="detailData.status === '0' ? 'success' : 'danger'">
            {{ detailData.status === '0' ? '成功' : '失败' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="操作时间" :span="1">{{ detailData.operTime }}</el-descriptions-item>
        <el-descriptions-item label="请求参数" :span="2">
          <pre style="white-space: pre-wrap; margin: 0; max-height: 200px; overflow: auto;">{{ detailData.operParam || '-' }}</pre>
        </el-descriptions-item>
        <el-descriptions-item label="返回参数" :span="2">
          <pre style="white-space: pre-wrap; margin: 0; max-height: 200px; overflow: auto;">{{ detailData.jsonResult || '-' }}</pre>
        </el-descriptions-item>
        <el-descriptions-item label="错误消息" :span="2" v-if="detailData.status === '1'">
          <pre style="white-space: pre-wrap; margin: 0; color: #f56c6c;">{{ detailData.errorMsg || '-' }}</pre>
        </el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="detailVisible = false">关闭</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Search, Refresh, Delete } from '@element-plus/icons-vue'
import request from '@/utils/request'

const loading = ref(false)
const detailVisible = ref(false)
const dateRange = ref([])

const searchForm = reactive({
  title: '',
  operName: '',
  status: ''
})

const pagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const operLogList = ref([])
const detailData = ref({})

const businessTypeMap = {
  '0': '其它',
  '1': '新增',
  '2': '修改',
  '3': '删除',
  '4': '授权',
  '5': '导出',
  '6': '导入',
  '7': '强退',
  '8': '生成代码',
  '9': '清空数据'
}

const businessTypeClassMap = {
  '0': 'info',
  '1': 'success',
  '2': 'primary',
  '3': 'danger',
  '4': 'warning',
  '5': 'success',
  '6': 'primary',
  '7': 'danger',
  '8': 'primary',
  '9': 'danger'
}

const getBusinessTypeLabel = (type) => {
  return businessTypeMap[type] || '其它'
}

const getBusinessTypeClass = (type) => {
  return businessTypeClassMap[type] || 'info'
}

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
    const res = await request.get('/monitor/operlog/list', { params })
    operLogList.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch (error) {
    console.error('获取操作日志列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.pageNum = 1
  fetchData()
}

const handleReset = () => {
  searchForm.title = ''
  searchForm.operName = ''
  searchForm.status = ''
  dateRange.value = []
  handleSearch()
}

const handleView = (row) => {
  detailData.value = { ...row }
  detailVisible.value = true
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该日志吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/monitor/operlog/${row.id}`)
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
      await request.delete('/monitor/operlog')
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
.operlog-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
