<template>
  <div class="role-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="角色名称">
          <el-input v-model="searchForm.roleName" placeholder="请输入角色名称" clearable />
        </el-form-item>
        <el-form-item label="权限字符">
          <el-input v-model="searchForm.roleKey" placeholder="请输入权限字符" clearable />
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
          <span>角色列表</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            新增角色
          </el-button>
        </div>
      </template>

      <el-table :data="roleList" stripe v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="roleName" label="角色名称" width="150" />
        <el-table-column prop="roleKey" label="权限字符" width="150" />
        <el-table-column prop="orderNum" label="显示顺序" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
            <el-button type="warning" link @click="handleAllocateMenu(row)">分配菜单</el-button>
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
      <el-form :model="roleForm" :rules="roleRules" ref="roleFormRef" label-width="100px">
        <el-form-item label="角色名称" prop="roleName">
          <el-input v-model="roleForm.roleName" placeholder="请输入角色名称" />
        </el-form-item>
        <el-form-item label="权限字符" prop="roleKey">
          <el-input v-model="roleForm.roleKey" placeholder="请输入权限字符" />
        </el-form-item>
        <el-form-item label="显示顺序" prop="orderNum">
          <el-input-number v-model="roleForm.orderNum" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="roleForm.status">
            <el-radio value="0">正常</el-radio>
            <el-radio value="1">停用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="roleForm.remark" type="textarea" :rows="3" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="menuDialogVisible" title="分配菜单权限" width="700px">
      <el-tree
        ref="menuTreeRef"
        :data="menuList"
        :props="{ label: 'menuName', children: 'children' }"
        show-checkbox
        node-key="id"
        default-expand-all
        :default-checked-keys="checkedMenuIds"
      />
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="menuDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmitMenu">确定</el-button>
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
const menuDialogVisible = ref(false)
const isEdit = ref(false)
const roleFormRef = ref(null)
const menuTreeRef = ref(null)
const currentRoleId = ref(null)

const searchForm = reactive({
  roleName: '',
  roleKey: '',
  status: ''
})

const pagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const roleList = ref([])
const menuList = ref([])
const checkedMenuIds = ref([])

const roleForm = reactive({
  id: null,
  roleName: '',
  roleKey: '',
  orderNum: 1,
  status: '0',
  remark: ''
})

const roleRules = {
  roleName: [
    { required: true, message: '请输入角色名称', trigger: 'blur' }
  ],
  roleKey: [
    { required: true, message: '请输入权限字符', trigger: 'blur' }
  ],
  orderNum: [
    { required: true, message: '请输入显示顺序', trigger: 'blur' }
  ]
}

const dialogTitle = ref('新增角色')

const fetchData = async () => {
  loading.value = true
  try {
    const res = await request.get('/system/role/list', {
      params: {
        ...searchForm,
        page_num: pagination.pageNum,
        page_size: pagination.pageSize
      }
    })
    roleList.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch (error) {
    console.error('获取角色列表失败:', error)
  } finally {
    loading.value = false
  }
}

const fetchMenuList = async () => {
  try {
    const res = await request.get('/system/menu/list')
    menuList.value = res.data || []
  } catch (error) {
    console.error('获取菜单列表失败:', error)
  }
}

const handleSearch = () => {
  pagination.pageNum = 1
  fetchData()
}

const handleReset = () => {
  searchForm.roleName = ''
  searchForm.roleKey = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增角色'
  Object.assign(roleForm, {
    id: null,
    roleName: '',
    roleKey: '',
    orderNum: 1,
    status: '0',
    remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  dialogTitle.value = '编辑角色'
  Object.assign(roleForm, {
    id: row.id,
    roleName: row.roleName,
    roleKey: row.roleKey,
    orderNum: row.orderNum || 1,
    status: row.status,
    remark: row.remark || ''
  })
  dialogVisible.value = true
}

const handleAllocateMenu = (row) => {
  currentRoleId.value = row.id
  checkedMenuIds.value = row.menuIds || []
  menuDialogVisible.value = true
}

const handleSubmit = async () => {
  if (!roleFormRef.value) return
  
  await roleFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await request.put(`/system/role/${roleForm.id}`, roleForm)
          ElMessage.success('更新成功')
        } else {
          await request.post('/system/role', roleForm)
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

const handleSubmitMenu = async () => {
  if (!menuTreeRef.value) return
  const checkedKeys = menuTreeRef.value.getCheckedKeys()
  
  try {
    await request.put(`/system/role/${currentRoleId.value}`, {
      menuIds: checkedKeys
    })
    ElMessage.success('分配成功')
    menuDialogVisible.value = false
  } catch (error) {
    console.error('分配菜单失败:', error)
  }
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该角色吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/system/role/${row.id}`)
      ElMessage.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('删除失败:', error)
    }
  }).catch(() => {})
}

onMounted(() => {
  fetchData()
  fetchMenuList()
})
</script>

<style lang="scss" scoped>
.role-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
