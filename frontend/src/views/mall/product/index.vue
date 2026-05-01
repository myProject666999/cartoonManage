<template>
  <div class="page-container">
    <el-card class="search-form">
      <el-form :inline="true">
        <el-form-item label="商品名称">
          <el-input placeholder="请输入商品名称" clearable />
        </el-form-item>
        <el-form-item label="商品状态">
          <el-select placeholder="请选择状态" clearable>
            <el-option label="上架" value="0" />
            <el-option label="下架" value="1" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary">
            <el-icon><Search /></el-icon>
            搜索
          </el-button>
          <el-button>
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="table-container">
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>商品列表</span>
          <el-button type="primary">
            <el-icon><Plus /></el-icon>
            新增商品
          </el-button>
        </div>
      </template>

      <el-table :data="productList" stripe>
        <el-table-column prop="productCode" label="商品编码" width="150" />
        <el-table-column prop="productName" label="商品名称" width="200" />
        <el-table-column prop="categoryName" label="分类" width="120" />
        <el-table-column prop="currentPrice" label="现价" width="100">
          <template #default="{ row }">
            ¥{{ row.currentPrice }}
          </template>
        </el-table-column>
        <el-table-column prop="originalPrice" label="原价" width="100">
          <template #default="{ row }">
            ¥{{ row.originalPrice }}
          </template>
        </el-table-column>
        <el-table-column prop="stock" label="库存" width="80" />
        <el-table-column prop="sales" label="销量" width="80" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === '0' ? 'success' : 'danger'">
              {{ row.status === '0' ? '上架' : '下架' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <el-button type="primary" link>编辑</el-button>
          <el-button type="primary" link>查看</el-button>
          <el-button type="danger" link>删除</el-button>
        </el-table-column>
      </el-table>

      <el-pagination
        :page-sizes="[10, 20, 50, 100]"
        :total="100"
        layout="total, sizes, prev, pager, next, jumper"
      />
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Search, Refresh, Plus } from '@element-plus/icons-vue'

const productList = ref([
  {
    id: 1,
    productCode: 'PRD001',
    productName: '海贼王路飞手办',
    categoryName: '手办模型',
    currentPrice: 299.00,
    originalPrice: 399.00,
    stock: 100,
    sales: 256,
    status: '0'
  },
  {
    id: 2,
    productCode: 'PRD002',
    productName: '火影忍者鸣人手办',
    categoryName: '手办模型',
    currentPrice: 289.00,
    originalPrice: 369.00,
    stock: 85,
    sales: 189,
    status: '0'
  },
  {
    id: 3,
    productCode: 'PRD003',
    productName: '龙珠悟空手办',
    categoryName: '手办模型',
    currentPrice: 329.00,
    originalPrice: 429.00,
    stock: 0,
    sales: 342,
    status: '1'
  }
])
</script>

<style lang="scss" scoped>
.page-container {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
