<template>
  <div class="dashboard">
    <el-row :gutter="20" class="stat-row">
      <el-col :span="6" v-for="item in statCards" :key="item.title">
        <el-card class="stat-card" :body-style="{ padding: '20px' }">
          <div class="stat-content">
            <div class="stat-icon" :style="{ background: item.color }">
              <el-icon :size="32" color="#fff">
                <component :is="item.icon" />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ item.value }}</div>
              <div class="stat-title">{{ item.title }}</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="chart-row">
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>访问趋势</span>
          </template>
          <div ref="lineChartRef" class="chart-container"></div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>销售统计</span>
          </template>
          <div ref="pieChartRef" class="chart-container"></div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="table-row">
      <el-col :span="16">
        <el-card>
          <template #header>
            <span>最新订单</span>
          </template>
          <el-table :data="recentOrders" stripe>
            <el-table-column prop="orderNo" label="订单号" width="180" />
            <el-table-column prop="userName" label="用户名" width="120" />
            <el-table-column prop="totalAmount" label="订单金额" width="100">
              <template #default="{ row }">
                ¥{{ row.totalAmount }}
              </template>
            </el-table-column>
            <el-table-column prop="orderStatus" label="订单状态" width="100">
              <template #default="{ row }">
                <el-tag :type="getStatusType(row.orderStatus)">
                  {{ getStatusText(row.orderStatus) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="createTime" label="创建时间" width="180" />
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <span>系统公告</span>
          </template>
          <div class="notice-list">
            <div class="notice-item" v-for="notice in notices" :key="notice.id">
              <div class="notice-title">{{ notice.noticeTitle }}</div>
              <div class="notice-time">{{ notice.createTime }}</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import * as echarts from 'echarts'
import {
  User,
  ShoppingCart,
  Goods,
  Coin,
  ChatLineRound
} from '@element-plus/icons-vue'

const lineChartRef = ref(null)
const pieChartRef = ref(null)

const statCards = reactive([
  { title: '总用户数', value: '1,234', icon: User, color: '#409EFF' },
  { title: '今日订单', value: '56', icon: ShoppingCart, color: '#67C23A' },
  { title: '商品数量', value: '892', icon: Goods, color: '#E6A23C' },
  { title: '今日销售额', value: '¥12,345', icon: Coin, color: '#F56C6C' }
])

const recentOrders = reactive([
  { orderNo: 'OR202401010001', userName: '张三', totalAmount: 128.50, orderStatus: '1', createTime: '2024-01-01 10:30:00' },
  { orderNo: 'OR202401010002', userName: '李四', totalAmount: 256.00, orderStatus: '2', createTime: '2024-01-01 11:20:00' },
  { orderNo: 'OR202401010003', userName: '王五', totalAmount: 89.90, orderStatus: '0', createTime: '2024-01-01 14:15:00' },
  { orderNo: 'OR202401010004', userName: '赵六', totalAmount: 320.00, orderStatus: '3', createTime: '2024-01-01 15:40:00' },
  { orderNo: 'OR202401010005', userName: '钱七', totalAmount: 156.80, orderStatus: '4', createTime: '2024-01-01 16:25:00' }
])

const notices = reactive([
  { id: 1, noticeTitle: '系统将于今晚进行维护', createTime: '2024-01-01' },
  { id: 2, noticeTitle: '新功能上线：订单导出功能', createTime: '2024-01-02' },
  { id: 3, noticeTitle: '春节放假通知', createTime: '2024-01-03' }
])

const getStatusType = (status) => {
  const types = {
    '0': 'warning',
    '1': 'primary',
    '2': 'info',
    '3': 'success',
    '4': 'success',
    '5': 'danger'
  }
  return types[status] || 'info'
}

const getStatusText = (status) => {
  const texts = {
    '0': '待付款',
    '1': '待发货',
    '2': '已发货',
    '3': '已收货',
    '4': '已完成',
    '5': '已取消'
  }
  return texts[status] || '未知'
}

const initLineChart = () => {
  const chart = echarts.init(lineChartRef.value)
  const option = {
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ['访问量', '订单量']
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        name: '访问量',
        type: 'line',
        smooth: true,
        data: [120, 200, 150, 80, 70, 110, 130],
        areaStyle: {
          opacity: 0.3
        }
      },
      {
        name: '订单量',
        type: 'line',
        smooth: true,
        data: [50, 80, 60, 30, 25, 45, 55],
        areaStyle: {
          opacity: 0.3
        }
      }
    ]
  }
  chart.setOption(option)
}

const initPieChart = () => {
  const chart = echarts.init(pieChartRef.value)
  const option = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      bottom: '5%',
      left: 'center'
    },
    series: [
      {
        name: '销售统计',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 16,
            fontWeight: 'bold'
          }
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 1048, name: '动漫周边' },
          { value: 735, name: '手办模型' },
          { value: 580, name: '漫画书籍' },
          { value: 484, name: '游戏周边' },
          { value: 300, name: '其他' }
        ]
      }
    ]
  }
  chart.setOption(option)
}

onMounted(() => {
  initLineChart()
  initPieChart()
})
</script>

<style lang="scss" scoped>
.dashboard {
  .stat-row {
    margin-bottom: 20px;

    .stat-card {
      .stat-content {
        display: flex;
        align-items: center;

        .stat-icon {
          width: 60px;
          height: 60px;
          border-radius: 8px;
          display: flex;
          align-items: center;
          justify-content: center;
          margin-right: 20px;
        }

        .stat-info {
          .stat-value {
            font-size: 24px;
            font-weight: bold;
            color: #333;
          }

          .stat-title {
            font-size: 14px;
            color: #999;
            margin-top: 5px;
          }
        }
      }
    }
  }

  .chart-row {
    margin-bottom: 20px;

    .chart-container {
      height: 300px;
    }
  }

  .table-row {
    .notice-list {
      .notice-item {
        padding: 12px 0;
        border-bottom: 1px solid #eee;

        &:last-child {
          border-bottom: none;
        }

        .notice-title {
          font-size: 14px;
          color: #333;
          cursor: pointer;

          &:hover {
            color: #409EFF;
          }
        }

        .notice-time {
          font-size: 12px;
          color: #999;
          margin-top: 5px;
        }
      }
    }
  }
}
</style>
