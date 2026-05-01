<template>
  <div class="page-container">
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>部门管理</span>
          <el-button type="primary">
            <el-icon><Plus /></el-icon>
            新增部门
          </el-button>
        </div>
      </template>
      <el-table :data="deptList" row-key="id" border default-expand-all>
        <el-table-column prop="deptName" label="部门名称" />
        <el-table-column prop="orderNum" label="显示顺序" width="120" />
        <el-table-column prop="leader" label="负责人" width="120" />
        <el-table-column prop="phone" label="联系电话" width="150" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <el-button type="primary" link>编辑</el-button>
          <el-button type="primary" link>新增</el-button>
          <el-button type="danger" link>删除</el-button>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { Plus } from '@element-plus/icons-vue'

const deptList = ref([
  {
    id: 1,
    deptName: '总公司',
    orderNum: 1,
    leader: '张总',
    phone: '13800138001',
    status: '0',
    children: [
      {
        id: 2,
        deptName: '技术部',
        orderNum: 1,
        leader: '李经理',
        phone: '13800138002',
        status: '0'
      },
      {
        id: 3,
        deptName: '市场部',
        orderNum: 2,
        leader: '王经理',
        phone: '13800138003',
        status: '0'
      }
    ]
  }
])
</script>

<style lang="scss" scoped>
.page-container {
  padding: 20px;
}
</style>
