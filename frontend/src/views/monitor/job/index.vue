<template>
  <div class="job-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="任务名称">
          <el-input v-model="searchForm.taskName" placeholder="请输入任务名称" clearable />
        </el-form-item>
        <el-form-item label="任务组名">
          <el-input v-model="searchForm.taskGroup" placeholder="请输入任务组名" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="正常" value="0" />
            <el-option label="暂停" value="1" />
          </el-select>
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
          <span>定时任务列表</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            新增任务
          </el-button>
        </div>
      </template>

      <el-table :data="jobList" stripe v-loading="loading">
        <el-table-column prop="id" label="任务ID" width="80" />
        <el-table-column prop="taskName" label="任务名称" width="150" />
        <el-table-column prop="taskGroup" label="任务组名" width="150" />
        <el-table-column prop="invokeTarget" label="调用目标" width="250" show-overflow-tooltip />
        <el-table-column prop="cronExpression" label="cron执行表达式" width="180" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'warning'">
              {{ row.status === '0' ? '正常' : '暂停' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180" />
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
            <el-button v-if="row.status === '0'" type="warning" link @click="handlePause(row)">暂停</el-button>
            <el-button v-else type="success" link @click="handleResume(row)">恢复</el-button>
            <el-button type="info" link @click="handleRun(row)">执行</el-button>
            <el-button type="primary" link @click="handleViewLog(row)">日志</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form :model="jobForm" :rules="jobRules" ref="jobFormRef" label-width="100px">
        <el-form-item label="任务名称" prop="taskName">
          <el-input v-model="jobForm.taskName" placeholder="请输入任务名称" />
        </el-form-item>
        <el-form-item label="任务组名" prop="taskGroup">
          <el-input v-model="jobForm.taskGroup" placeholder="请输入任务组名" />
        </el-form-item>
        <el-form-item label="调用目标" prop="invokeTarget">
          <el-input v-model="jobForm.invokeTarget" placeholder="请输入调用目标" />
        </el-form-item>
        <el-form-item label="cron表达式" prop="cronExpression">
          <el-input v-model="jobForm.cronExpression" placeholder="请输入cron表达式" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="jobForm.status">
            <el-radio value="0">正常</el-radio>
            <el-radio value="1">暂停</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="jobForm.remark" type="textarea" :rows="2" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="logDialogVisible" title="执行日志列表" width="800px">
      <el-table :data="jobLogList" stripe v-loading="logLoading">
        <el-table-column prop="id" label="日志ID" width="80" />
        <el-table-column prop="taskName" label="任务名称" width="150" />
        <el-table-column prop="taskGroup" label="任务组名" width="150" />
        <el-table-column prop="invokeTarget" label="调用目标" width="200" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '成功' : '失败' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="执行时间" width="180" />
      </el-table>

      <el-pagination
        v-model:current-page="logPagination.pageNum"
        v-model:page-size="logPagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="logPagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="fetchJobLogData"
        @current-change="fetchJobLogData"
      />
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Search, Refresh, Plus } from '@element-plus/icons-vue'
import request from '@/utils/request'

const loading = ref(false)
const logLoading = ref(false)
const dialogVisible = ref(false)
const logDialogVisible = ref(false)
const isEdit = ref(false)
const jobFormRef = ref(null)
const currentJobId = ref(null)

const searchForm = reactive({
  taskName: '',
  taskGroup: '',
  status: ''
})

const pagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const logPagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const jobList = ref([])
const jobLogList = ref([])

const jobForm = reactive({
  id: null,
  taskName: '',
  taskGroup: '',
  invokeTarget: '',
  cronExpression: '',
  status: '0',
  remark: ''
})

const jobRules = {
  taskName: [
    { required: true, message: '请输入任务名称', trigger: 'blur' }
  ],
  taskGroup: [
    { required: true, message: '请输入任务组名', trigger: 'blur' }
  ],
  invokeTarget: [
    { required: true, message: '请输入调用目标', trigger: 'blur' }
  ],
  cronExpression: [
    { required: true, message: '请输入cron表达式', trigger: 'blur' }
  ]
}

const dialogTitle = ref('新增任务')

const fetchData = async () => {
  loading.value = true
  try {
    const res = await request.get('/monitor/job/list', {
      params: {
        ...searchForm,
        page_num: pagination.pageNum,
        page_size: pagination.pageSize
      }
    })
    jobList.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch (error) {
    console.error('获取定时任务列表失败:', error)
  } finally {
    loading.value = false
  }
}

const fetchJobLogData = async () => {
  logLoading.value = true
  try {
    const res = await request.get('/monitor/jobLog/list', {
      params: {
        jobId: currentJobId.value,
        page_num: logPagination.pageNum,
        page_size: logPagination.pageSize
      }
    })
    jobLogList.value = res.data?.list || []
    logPagination.total = res.data?.total || 0
  } catch (error) {
    console.error('获取任务日志列表失败:', error)
  } finally {
    logLoading.value = false
  }
}

const handleSearch = () => {
  pagination.pageNum = 1
  fetchData()
}

const handleReset = () => {
  searchForm.taskName = ''
  searchForm.taskGroup = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增任务'
  Object.assign(jobForm, {
    id: null,
    taskName: '',
    taskGroup: '',
    invokeTarget: '',
    cronExpression: '',
    status: '0',
    remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  dialogTitle.value = '编辑任务'
  Object.assign(jobForm, {
    id: row.id,
    taskName: row.taskName,
    taskGroup: row.taskGroup,
    invokeTarget: row.invokeTarget,
    cronExpression: row.cronExpression,
    status: row.status,
    remark: row.remark || ''
  })
  dialogVisible.value = true
}

const handleViewLog = (row) => {
  currentJobId.value = row.id
  logPagination.pageNum = 1
  fetchJobLogData()
  logDialogVisible.value = true
}

const handleSubmit = async () => {
  if (!jobFormRef.value) return
  
  await jobFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await request.put(`/monitor/job/${jobForm.id}`, jobForm)
          ElMessage.success('更新成功')
        } else {
          await request.post('/monitor/job', jobForm)
          ElMessage.success('创建成功')
        }
        dialogVisible.value = false
        fetchData()
      } catch (error) {
        console.error('操作失败:', error)
      }
    }
  })
}

const handlePause = (row) => {
  ElMessageBox.confirm('确定要暂停该任务吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.put(`/monitor/job/${row.id}`, { status: '1' })
      ElMessage.success('暂停成功')
      fetchData()
    } catch (error) {
      console.error('暂停失败:', error)
    }
  }).catch(() => {})
}

const handleResume = (row) => {
  ElMessageBox.confirm('确定要恢复该任务吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.put(`/monitor/job/${row.id}`, { status: '0' })
      ElMessage.success('恢复成功')
      fetchData()
    } catch (error) {
      console.error('恢复失败:', error)
    }
  }).catch(() => {})
}

const handleRun = (row) => {
  ElMessageBox.confirm('确定要立即执行该任务吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.post(`/monitor/job/run/${row.id}`)
      ElMessage.success('执行成功')
      fetchData()
    } catch (error) {
      console.error('执行失败:', error)
    }
  }).catch(() => {})
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该任务吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/monitor/job/${row.id}`)
      ElMessage.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('删除失败:', error)
    }
  }).catch(() => {})
}

onMounted(() => {
  fetchData()
})
</script>

<style lang="scss" scoped>
.job-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
