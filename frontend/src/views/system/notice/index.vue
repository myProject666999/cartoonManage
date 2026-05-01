<template>
  <div class="notice-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="公告标题">
          <el-input v-model="searchForm.noticeTitle" placeholder="请输入公告标题" clearable />
        </el-form-item>
        <el-form-item label="公告类型">
          <el-select v-model="searchForm.noticeType" placeholder="请选择公告类型" clearable>
            <el-option label="通知" value="1" />
            <el-option label="公告" value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="正常" value="0" />
            <el-option label="停用" value="1" />
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
          <span>公告列表</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            新增公告
          </el-button>
        </div>
      </template>

      <el-table :data="noticeList" stripe v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="noticeTitle" label="公告标题" width="250" />
        <el-table-column prop="noticeType" label="公告类型" width="100">
          <template #default="{ row }">
            <el-tag :type="row.noticeType === '1' ? 'primary' : 'success'">
              {{ row.noticeType === '1' ? '通知' : '公告' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createBy" label="创建者" width="120" />
        <el-table-column prop="createTime" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleView(row)">查看</el-button>
            <el-button type="success" link @click="handleEdit(row)">编辑</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px">
      <el-form :model="noticeForm" :rules="noticeRules" ref="noticeFormRef" label-width="100px">
        <el-form-item label="公告标题" prop="noticeTitle">
          <el-input v-model="noticeForm.noticeTitle" placeholder="请输入公告标题" />
        </el-form-item>
        <el-form-item label="公告类型" prop="noticeType">
          <el-radio-group v-model="noticeForm.noticeType">
            <el-radio value="1">通知</el-radio>
            <el-radio value="2">公告</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="公告内容" prop="noticeContent">
          <el-input v-model="noticeForm.noticeContent" type="textarea" :rows="8" placeholder="请输入公告内容" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="noticeForm.status">
            <el-radio value="0">正常</el-radio>
            <el-radio value="1">停用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="noticeForm.remark" type="textarea" :rows="2" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit" :disabled="isView">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Search, Refresh, Plus } from '@element-plus/icons-vue'
import request from '@/utils/request'

const loading = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const isView = ref(false)
const noticeFormRef = ref(null)

const searchForm = reactive({
  noticeTitle: '',
  noticeType: '',
  status: ''
})

const pagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const noticeList = ref([])

const noticeForm = reactive({
  id: null,
  noticeTitle: '',
  noticeType: '1',
  noticeContent: '',
  status: '0',
  remark: ''
})

const noticeRules = {
  noticeTitle: [
    { required: true, message: '请输入公告标题', trigger: 'blur' }
  ],
  noticeType: [
    { required: true, message: '请选择公告类型', trigger: 'change' }
  ],
  noticeContent: [
    { required: true, message: '请输入公告内容', trigger: 'blur' }
  ]
}

const dialogTitle = ref('新增公告')

const fetchData = async () => {
  loading.value = true
  try {
    const res = await request.get('/system/notice/list', {
      params: {
        ...searchForm,
        page_num: pagination.pageNum,
        page_size: pagination.pageSize
      }
    })
    noticeList.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch (error) {
    console.error('获取公告列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.pageNum = 1
  fetchData()
}

const handleReset = () => {
  searchForm.noticeTitle = ''
  searchForm.noticeType = ''
  searchForm.status = ''
  handleSearch()
}

const resetForm = () => {
  Object.assign(noticeForm, {
    id: null,
    noticeTitle: '',
    noticeType: '1',
    noticeContent: '',
    status: '0',
    remark: ''
  })
}

const handleAdd = () => {
  isEdit.value = false
  isView.value = false
  dialogTitle.value = '新增公告'
  resetForm()
  dialogVisible.value = true
}

const handleView = (row) => {
  isEdit.value = false
  isView.value = true
  dialogTitle.value = '查看公告'
  Object.assign(noticeForm, {
    id: row.id,
    noticeTitle: row.noticeTitle,
    noticeType: row.noticeType || '1',
    noticeContent: row.noticeContent,
    status: row.status,
    remark: row.remark || ''
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  isView.value = false
  dialogTitle.value = '编辑公告'
  Object.assign(noticeForm, {
    id: row.id,
    noticeTitle: row.noticeTitle,
    noticeType: row.noticeType || '1',
    noticeContent: row.noticeContent,
    status: row.status,
    remark: row.remark || ''
  })
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!noticeFormRef.value || isView.value) return
  
  await noticeFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await request.put(`/system/notice/${noticeForm.id}`, noticeForm)
          ElMessage.success('更新成功')
        } else {
          await request.post('/system/notice', noticeForm)
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

const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该公告吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/system/notice/${row.id}`)
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
.notice-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
