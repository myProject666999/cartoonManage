<template>
  <div class="menu-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="菜单名称">
          <el-input v-model="searchForm.menuName" placeholder="请输入菜单名称" clearable />
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
          <span>菜单列表</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            新增菜单
          </el-button>
        </div>
      </template>

      <el-table :data="menuList" stripe row-key="id" v-loading="loading" border default-expand-all>
        <el-table-column prop="menuName" label="菜单名称" width="200" />
        <el-table-column prop="icon" label="图标" width="120">
          <template #default="{ row }">
            <el-icon v-if="row.icon">
              <component :is="row.icon" />
            </el-icon>
          </template>
        </el-table-column>
        <el-table-column prop="menuType" label="类型" width="100">
          <template #default="{ row }">
            <el-tag :type="row.menuType === 'M' ? 'primary' : row.menuType === 'C' ? 'success' : 'warning'">
              {{ row.menuType === 'M' ? '目录' : row.menuType === 'C' ? '菜单' : '按钮' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="path" label="路由地址" width="200" />
        <el-table-column prop="perms" label="权限标识" width="200" />
        <el-table-column prop="orderNum" label="排序" width="80" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
            <el-button type="success" link @click="handleAddSub(row)">新增子菜单</el-button>
            <el-button type="danger" link @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px">
      <el-form :model="menuForm" :rules="menuRules" ref="menuFormRef" label-width="100px">
        <el-row>
          <el-col :span="12">
            <el-form-item label="上级菜单">
              <el-tree-select
                v-model="menuForm.parentId"
                :data="parentMenus"
                :props="{ label: 'menuName', value: 'id', children: 'children' }"
                placeholder="请选择上级菜单"
                clearable
                check-strictly
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="菜单类型" prop="menuType">
              <el-radio-group v-model="menuForm.menuType">
                <el-radio value="M">目录</el-radio>
                <el-radio value="C">菜单</el-radio>
                <el-radio value="F">按钮</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="菜单名称" prop="menuName">
              <el-input v-model="menuForm.menuName" placeholder="请输入菜单名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="显示排序" prop="orderNum">
              <el-input-number v-model="menuForm.orderNum" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row v-if="menuForm.menuType !== 'F'">
          <el-col :span="12">
            <el-form-item label="路由地址">
              <el-input v-model="menuForm.path" placeholder="请输入路由地址" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="组件路径">
              <el-input v-model="menuForm.component" placeholder="请输入组件路径" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="权限标识">
              <el-input v-model="menuForm.perms" placeholder="请输入权限标识" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="菜单图标">
              <el-input v-model="menuForm.icon" placeholder="请输入图标名称" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="状态">
              <el-radio-group v-model="menuForm.status">
                <el-radio value="0">正常</el-radio>
                <el-radio value="1">停用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注">
              <el-input v-model="menuForm.remark" placeholder="请输入备注" />
            </el-form-item>
          </el-col>
        </el-row>
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
const menuFormRef = ref(null)

const searchForm = reactive({
  menuName: '',
  status: ''
})

const menuList = ref([])
const parentMenus = ref([])

const menuForm = reactive({
  id: null,
  parentId: 0,
  menuName: '',
  menuType: 'C',
  path: '',
  component: '',
  perms: '',
  icon: '',
  orderNum: 1,
  status: '0',
  remark: ''
})

const menuRules = {
  menuName: [
    { required: true, message: '请输入菜单名称', trigger: 'blur' }
  ],
  menuType: [
    { required: true, message: '请选择菜单类型', trigger: 'change' }
  ],
  orderNum: [
    { required: true, message: '请输入显示顺序', trigger: 'blur' }
  ]
}

const dialogTitle = ref('新增菜单')

const fetchData = async () => {
  loading.value = true
  try {
    const res = await request.get('/system/menu/list', { params: searchForm })
    menuList.value = res.data || []
  } catch (error) {
    console.error('获取菜单列表失败:', error)
  } finally {
    loading.value = false
  }
}

const fetchParentMenus = async () => {
  try {
    const res = await request.get('/system/menu/list')
    parentMenus.value = [{ id: 0, menuName: '根目录' }, ...res.data]
  } catch (error) {
    console.error('获取上级菜单失败:', error)
  }
}

const handleSearch = () => {
  fetchData()
}

const handleReset = () => {
  searchForm.menuName = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增菜单'
  Object.assign(menuForm, {
    id: null,
    parentId: 0,
    menuName: '',
    menuType: 'C',
    path: '',
    component: '',
    perms: '',
    icon: '',
    orderNum: 1,
    status: '0',
    remark: ''
  })
  dialogVisible.value = true
}

const handleAddSub = (row) => {
  isEdit.value = false
  dialogTitle.value = '新增子菜单'
  Object.assign(menuForm, {
    id: null,
    parentId: row.id,
    menuName: '',
    menuType: 'C',
    path: '',
    component: '',
    perms: '',
    icon: '',
    orderNum: 1,
    status: '0',
    remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  dialogTitle.value = '编辑菜单'
  Object.assign(menuForm, {
    id: row.id,
    parentId: row.parentId || 0,
    menuName: row.menuName,
    menuType: row.menuType || 'C',
    path: row.path || '',
    component: row.component || '',
    perms: row.perms || '',
    icon: row.icon || '',
    orderNum: row.orderNum || 1,
    status: row.status,
    remark: row.remark || ''
  })
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!menuFormRef.value) return
  
  await menuFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await request.put(`/system/menu/${menuForm.id}`, menuForm)
          ElMessage.success('更新成功')
        } else {
          await request.post('/system/menu', menuForm)
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
  ElMessageBox.confirm('确定要删除该菜单吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`/system/menu/${row.id}`)
      ElMessage.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('删除失败:', error)
    }
  }).catch(() => {})
}

onMounted(() => {
  fetchData()
  fetchParentMenus()
})
</script>

<style lang="scss" scoped>
.menu-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
