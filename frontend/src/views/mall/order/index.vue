<template>
  <div class="order-manage">
    <el-card class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="订单编号">
          <el-input v-model="searchForm.orderSn" placeholder="请输入订单编号" clearable />
        </el-form-item>
        <el-form-item label="订单状态">
          <el-select v-model="searchForm.orderStatus" placeholder="请选择订单状态" clearable>
            <el-option label="待付款" value="0" />
            <el-option label="待发货" value="1" />
            <el-option label="已发货" value="2" />
            <el-option label="已完成" value="3" />
            <el-option label="已取消" value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="下单时间">
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
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="table-container">
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>订单列表</span>
        </div>
      </template>

      <el-table :data="orderList" stripe v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="orderSn" label="订单编号" width="200" />
        <el-table-column prop="userName" label="用户名称" width="120" />
        <el-table-column prop="totalAmount" label="订单金额" width="120">
          <template #default="{ row }">
            ￥{{ row.totalAmount }}
          </template>
        </el-table-column>
        <el-table-column prop="payAmount" label="实付金额" width="120">
          <template #default="{ row }">
            ￥{{ row.payAmount }}
          </template>
        </el-table-column>
        <el-table-column prop="orderStatus" label="订单状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getOrderStatusType(row.orderStatus)">
              {{ getOrderStatusLabel(row.orderStatus) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="payType" label="支付方式" width="100">
          <template #default="{ row }">
            {{ row.payType === '1' ? '微信支付' : row.payType === '2' ? '支付宝' : '其他' }}
          </template>
        </el-table-column>
        <el-table-column prop="receiverName" label="收货人" width="100" />
        <el-table-column prop="receiverPhone" label="收货电话" width="130" />
        <el-table-column prop="createTime" label="下单时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleView(row)">详情</el-button>
            <el-button type="warning" link v-if="row.orderStatus === '1'" @click="handleShip(row)">发货</el-button>
            <el-button type="success" link v-if="row.orderStatus === '0'" @click="handlePay(row)">确认付款</el-button>
            <el-button type="danger" link v-if="row.orderStatus === '0'" @click="handleCancel(row)">取消</el-button>
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

    <el-dialog v-model="detailVisible" title="订单详情" width="800px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="订单编号">{{ detailData.orderSn }}</el-descriptions-item>
        <el-descriptions-item label="订单状态">
          <el-tag :type="getOrderStatusType(detailData.orderStatus)">
            {{ getOrderStatusLabel(detailData.orderStatus) }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="订单金额">￥{{ detailData.totalAmount }}</el-descriptions-item>
        <el-descriptions-item label="实付金额">￥{{ detailData.payAmount }}</el-descriptions-item>
        <el-descriptions-item label="支付方式">{{ detailData.payType === '1' ? '微信支付' : detailData.payType === '2' ? '支付宝' : '其他' }}</el-descriptions-item>
        <el-descriptions-item label="支付时间">{{ detailData.payTime || '-' }}</el-descriptions-item>
        <el-descriptions-item label="收货人">{{ detailData.receiverName }}</el-descriptions-item>
        <el-descriptions-item label="联系电话">{{ detailData.receiverPhone }}</el-descriptions-item>
        <el-descriptions-item label="收货地址" :span="2">{{ detailData.receiverDetail || '-' }}</el-descriptions-item>
        <el-descriptions-item label="下单时间">{{ detailData.createTime }}</el-descriptions-item>
        <el-descriptions-item label="订单备注">{{ detailData.remark || '-' }}</el-descriptions-item>
      </el-descriptions>

      <el-divider>订单商品</el-divider>
      <el-table :data="detailData.orderItems || []" stripe>
        <el-table-column prop="productName" label="商品名称" />
        <el-table-column prop="productSku" label="商品SKU" width="150" />
        <el-table-column prop="productQuantity" label="数量" width="80" />
        <el-table-column prop="productPrice" label="单价" width="120">
          <template #default="{ row }">
            ￥{{ row.productPrice }}
          </template>
        </el-table-column>
        <el-table-column prop="totalPrice" label="小计" width="120">
          <template #default="{ row }">
            ￥{{ row.totalPrice }}
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="shipVisible" title="发货" width="400px">
      <el-form :model="shipForm" :rules="shipRules" ref="shipFormRef" label-width="100px">
        <el-form-item label="物流公司" prop="logisticsCompany">
          <el-input v-model="shipForm.logisticsCompany" placeholder="请输入物流公司" />
        </el-form-item>
        <el-form-item label="物流单号" prop="logisticsSn">
          <el-input v-model="shipForm.logisticsSn" placeholder="请输入物流单号" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="shipVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmitShip">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Search, Refresh } from '@element-plus/icons-vue'
import request from '@/utils/request'

const loading = ref(false)
const detailVisible = ref(false)
const shipVisible = ref(false)
const shipFormRef = ref(null)
const dateRange = ref([])
const currentOrderId = ref(null)

const searchForm = reactive({
  orderSn: '',
  orderStatus: ''
})

const pagination = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 0
})

const orderList = ref([])
const detailData = ref({})

const shipForm = reactive({
  logisticsCompany: '',
  logisticsSn: ''
})

const shipRules = {
  logisticsCompany: [
    { required: true, message: '请输入物流公司', trigger: 'blur' }
  ],
  logisticsSn: [
    { required: true, message: '请输入物流单号', trigger: 'blur' }
  ]
}

const orderStatusMap = {
  '0': '待付款',
  '1': '待发货',
  '2': '已发货',
  '3': '已完成',
  '4': '已取消'
}

const orderStatusTypeMap = {
  '0': 'warning',
  '1': 'primary',
  '2': 'info',
  '3': 'success',
  '4': 'danger'
}

const getOrderStatusLabel = (status) => {
  return orderStatusMap[status] || '未知'
}

const getOrderStatusType = (status) => {
  return orderStatusTypeMap[status] || 'info'
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
    const res = await request.get('/mall/order/list', { params })
    orderList.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch (error) {
    console.error('获取订单列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.pageNum = 1
  fetchData()
}

const handleReset = () => {
  searchForm.orderSn = ''
  searchForm.orderStatus = ''
  dateRange.value = []
  handleSearch()
}

const handleView = async (row) => {
  try {
    const res = await request.get(`/mall/order/${row.id}`)
    detailData.value = res.data || {}
    detailVisible.value = true
  } catch (error) {
    console.error('获取订单详情失败:', error)
  }
}

const handleShip = (row) => {
  currentOrderId.value = row.id
  shipForm.logisticsCompany = ''
  shipForm.logisticsSn = ''
  shipVisible.value = true
}

const handleSubmitShip = async () => {
  if (!shipFormRef.value) return
  
  await shipFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        await request.put(`/mall/order/${currentOrderId.value}`, {
          orderStatus: '2',
          logisticsCompany: shipForm.logisticsCompany,
          logisticsSn: shipForm.logisticsSn
        })
        ElMessage.success('发货成功')
        shipVisible.value = false
        fetchData()
      } catch (error) {
        console.error('发货失败:', error)
      }
    }
  })
}

const handlePay = (row) => {
  ElMessageBox.confirm('确定要确认付款吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.put(`/mall/order/${row.id}`, { orderStatus: '1' })
      ElMessage.success('确认付款成功')
      fetchData()
    } catch (error) {
      console.error('确认付款失败:', error)
    }
  }).catch(() => {})
}

const handleCancel = (row) => {
  ElMessageBox.confirm('确定要取消该订单吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.put(`/mall/order/${row.id}`, { orderStatus: '4' })
      ElMessage.success('取消成功')
      fetchData()
    } catch (error) {
      console.error('取消失败:', error)
    }
  }).catch(() => {})
}

onMounted(() => {
  fetchData()
})
</script>

<style lang="scss" scoped>
.order-manage {
  .search-form {
    margin-bottom: 20px;
  }

  .table-container {
    margin-bottom: 20px;
  }
}
</style>
