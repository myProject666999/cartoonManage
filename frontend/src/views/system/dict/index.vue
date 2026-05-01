<template>
  <div class="dict-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="字典名称">
          <el-input v-model="searchForm.dictName" placeholder="请输入字典名称" clearable />
        </el-form-item>
        <el-form-item label="字典类型">
          <el-input v-model="searchForm.dictType" placeholder="请输入字典类型" clearable />
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
          <span>字典类型列表</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            新增字典
          </el-button>
        </div>
      </template>

      <el-table :data="dictTypeList" stripe v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="dictName" label="字典名称" width="180" />
        <el-table-column prop="dictType" label="字典类型" width="200" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180" />
        <el-table-column prop="remark" label="备注" show-overflow-tooltip />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
            <el-button type="success" link @click="handleViewDictData(row)">字典数据</el-button>
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
      <el-form :model="dictTypeForm" :rules="dictTypeRules" ref="dictTypeFormRef" label-width="100px">
        <el-form-item label="字典名称" prop="dictName">
          <el-input v-model="dictTypeForm.dictName" placeholder="请输入字典名称" />
        </el-form-item>
        <el-form-item label="字典类型" prop="dictType">
          <el-input v-model="dictTypeForm.dictType" placeholder="请输入字典类型" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="dictTypeForm.status">
            <el-radio value="0">正常</el-radio>
            <el-radio value="1">停用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="dictTypeForm.remark" type="textarea" :rows="3" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="dictDataVisible" title="字典数据列表" width="800px">
      <div class="dict-data-toolbar">
        <el-button type="primary" @click="handleAddDictData">
          <el-icon><Plus /></el-icon>
          新增数据
        </el-button>
      </div>
      <el-table :data="dictDataList" stripe v-loading="dictDataLoading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="dictLabel" label="字典标签" width="150" />
        <el-table-column prop="dictValue" label="字典键值" width="150" />
        <el-table-column prop="dictSort" label="排序" width="100" />
        <el-table-column prop="cssClass" label="样式属性" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEditDictData(row)">编辑</el-button>
            <el-button type="danger" link @click="handleDeleteDictData(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="dictDataFormVisible" :title="dictDataFormTitle" width="500px">
      <el-form :model="dictDataForm" :rules="dictDataRules" ref="dictDataFormRef" label-width="100px">
        <el-form-item label="字典标签" prop="dictLabel">
          <el-input v-model="dictDataForm.dictLabel" placeholder="请输入字典标签" />
        </el-form-item>
        <el-form-item label="字典键值" prop="dictValue">
          <el-input v-model="dictDataForm.dictValue" placeholder="请输入字典键值" />
        </el-form-item>
        <el-form-item label="排序" prop="dictSort">
          <el-input-number v-model="dictDataForm.dictSort" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="样式属性">
          <el-input v-model="dictDataForm.cssClass" placeholder="请输入样式属性" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="dictDataForm.status">
            <el-radio value="0">正常</el-radio>
            <el-radio value="1">停用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="dictDataForm.remark" type="textarea" :rows="2" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dictDataFormVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmitDictData">确定</el-button>
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
const dictDataVisible = ref(false)
const dictDataFormVisible = ref(false)
const isEdit = ref(false)
const isDictDataEdit = ref(false)
const dictTypeFormRef = ref(null)
const dictDataFormRef = ref(null)
const dictDataLoading = ref(false)
const currentDictType = ref(null)

const searchForm = reactive({
  dictName: '',
  dictType: '',
  status: ''
})

const pagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const dictTypeList = ref([])
const dictDataList = ref([])

const dictTypeForm = reactive({
  id: null,
  dictName: '',
  dictType: '',
  status: '0',
  remark: ''
})

const dictDataForm = reactive({
  id: null,
  dictLabel: '',
  dictValue: '',
  dictSort: 1,
  cssClass: '',
  status: '0',
  remark: ''
})

const dictTypeRules = {
  dictName: [
    { required: true, message: '请输入字典名称', trigger: 'blur' }
  ],
  dictType: [
    { required: true, message: '请输入字典类型', trigger: 'blur' }
  ]
}

const dictDataRules = {
  dictLabel: [
    { required: true, message: '请输入字典标签', trigger: 'blur' }
  ],
  dictValue: [
    { required: true, message: '请输入字典键值', trigger: 'blur' }
  ],
  dictSort: [
    { required: true, message: '请输入排序', trigger: 'blur' }
  ]
}

const dialogTitle = ref('新增字典类型')
const dictDataFormTitle = ref('新增字典数据')

const fetchData = async () => {
  loading.value = true
  try {
    const res = await request.get('/system/dict/type/list', {
      params: {
        ...searchForm,
        page_num: pagination.pageNum,
        page_size: pagination.pageSize
      }
    })
    dictTypeList.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch (error) {
    console.error('获取字典类型列表失败:', error)
  } finally {
    loading.value = false
  }
}

const fetchDictDataList = async (dictType) => {
  dictDataLoading.value = true
  try {
    const res = await request.get('/system/dict/data/list', {
      params: {
        dictType: dictType
      }
    })
    dictDataList.value = res.data?.list || []
  } catch (error) {
    console.error('获取字典数据列表失败:', error)
  } finally {
    dictDataLoading.value = false
  }
}

const handleSearch = () => {
  pagination.pageNum = 1
  fetchData()
}

const handleReset = () => {
  searchForm.dictName = ''
  searchForm.dictType = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增字典类型'
  Object.assign(dictTypeForm, {
    id: null,
    dictName: '',
    dictType: '',
    status: '0',
    remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  dialogTitle.value = '编辑字典类型'
  Object.assign(dictTypeForm, {
    id: row.id,
    dictName: row.dictName,
    dictType: row.dictType,
    status: row.status,
    remark: row.remark || ''
  })
  dialogVisible.value = true
}

const handleViewDictData = (row) => {
  currentDictType.value = row.dictType
  fetchDictDataList(row.dictType)
  dictDataVisible.value = true
}

const handleSubmit = async () => {
  if (!dictTypeFormRef.value) return
  
  await dictTypeFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await request.put(`/system/dict/type/${dictTypeForm.id}`, dictTypeForm)
          ElMessage.success('更新成功')
        } else {
          await request.post('/system/dict/type', dictTypeForm)
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
  ElMessageBox.confirm('确定要删除该字典类型吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/system/dict/type/${row.id}`)
      ElMessage.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('删除失败:', error)
    }
  }).catch(() => {})
}

const handleAddDictData = () => {
  isDictDataEdit.value = false
  dictDataFormTitle.value = '新增字典数据'
  Object.assign(dictDataForm, {
    id: null,
    dictLabel: '',
    dictValue: '',
    dictSort: 1,
    cssClass: '',
    status: '0',
    remark: ''
  })
  dictDataFormVisible.value = true
}

const handleEditDictData = (row) => {
  isDictDataEdit.value = true
  dictDataFormTitle.value = '编辑字典数据'
  Object.assign(dictDataForm, {
    id: row.id,
    dictLabel: row.dictLabel,
    dictValue: row.dictValue,
    dictSort: row.dictSort || 1,
    cssClass: row.cssClass || '',
    status: row.status,
    remark: row.remark || ''
  })
  dictDataFormVisible.value = true
}

const handleSubmitDictData = async () => {
  if (!dictDataFormRef.value) return
  
  await dictDataFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const data = {
          ...dictDataForm,
          dictType: currentDictType.value
        }
        if (isDictDataEdit.value) {
          await request.put(`/system/dict/data/${dictDataForm.id}`, data)
          ElMessage.success('更新成功')
        } else {
          await request.post('/system/dict/data', data)
          ElMessage.success('创建成功')
        }
        dictDataFormVisible.value = false
        fetchDictDataList(currentDictType.value)
      } catch (error) {
        console.error('操作失败:', error)
      }
    }
  })
}

const handleDeleteDictData = (row) => {
  ElMessageBox.confirm('确定要删除该字典数据吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/system/dict/data/${row.id}`)
      ElMessage.success('删除成功')
      fetchDictDataList(currentDictType.value)
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
.dict-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }

  .dict-data-toolbar {
    margin-bottom: 15px;
  }
}
</style>
