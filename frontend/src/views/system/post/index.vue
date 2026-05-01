<template>
  <div class="post-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="岗位编码">
          <el-input v-model="searchForm.postCode" placeholder="请输入岗位编码" clearable />
        </el-form-item>
        <el-form-item label="岗位名称">
          <el-input v-model="searchForm.postName" placeholder="请输入岗位名称" clearable />
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
          <span>岗位列表</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            新增岗位
          </el-button>
        </div>
      </template>

      <el-table :data="postList" stripe v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="postCode" label="岗位编码" width="150" />
        <el-table-column prop="postName" label="岗位名称" width="150" />
        <el-table-column prop="orderNum" label="显示顺序" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form :model="postForm" :rules="postRules" ref="postFormRef" label-width="100px">
        <el-form-item label="岗位编码" prop="postCode">
          <el-input v-model="postForm.postCode" placeholder="请输入岗位编码" />
        </el-form-item>
        <el-form-item label="岗位名称" prop="postName">
          <el-input v-model="postForm.postName" placeholder="请输入岗位名称" />
        </el-form-item>
        <el-form-item label="显示顺序" prop="orderNum">
          <el-input-number v-model="postForm.orderNum" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="postForm.status">
            <el-radio value="0">正常</el-radio>
            <el-radio value="1">停用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="postForm.remark" type="textarea" :rows="3" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
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
const postFormRef = ref(null)

const searchForm = reactive({
  postCode: '',
  postName: '',
  status: ''
})

const pagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const postList = ref([])

const postForm = reactive({
  id: null,
  postCode: '',
  postName: '',
  orderNum: 1,
  status: '0',
  remark: ''
})

const postRules = {
  postCode: [
    { required: true, message: '请输入岗位编码', trigger: 'blur' }
  ],
  postName: [
    { required: true, message: '请输入岗位名称', trigger: 'blur' }
  ],
  orderNum: [
    { required: true, message: '请输入显示顺序', trigger: 'blur' }
  ]
}

const dialogTitle = ref('新增岗位')

const fetchData = async () => {
  loading.value = true
  try {
    const res = await request.get('/system/post/list', {
      params: {
        ...searchForm,
        page_num: pagination.pageNum,
        page_size: pagination.pageSize
      }
    })
    postList.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch (error) {
    console.error('获取岗位列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.pageNum = 1
  fetchData()
}

const handleReset = () => {
  searchForm.postCode = ''
  searchForm.postName = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增岗位'
  Object.assign(postForm, {
    id: null,
    postCode: '',
    postName: '',
    orderNum: 1,
    status: '0',
    remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  dialogTitle.value = '编辑岗位'
  Object.assign(postForm, {
    id: row.id,
    postCode: row.postCode,
    postName: row.postName,
    orderNum: row.orderNum || 1,
    status: row.status,
    remark: row.remark || ''
  })
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!postFormRef.value) return
  
  await postFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await request.put(`/system/post/${postForm.id}`, postForm)
          ElMessage.success('更新成功')
        } else {
          await request.post('/system/post', postForm)
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
  ElMessageBox.confirm('确定要删除该岗位吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/system/post/${row.id}`)
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
.post-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
