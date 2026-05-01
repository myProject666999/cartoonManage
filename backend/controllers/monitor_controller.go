package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

type OnlineUser struct {
	Token      string    `json:"token"`
	UserID     uint64    `json:"user_id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	IPAddress  string    `json:"ip_address"`
	LoginTime  time.Time `json:"login_time"`
	LastActive time.Time `json:"last_active"`
}

var onlineUsers = make(map[string]OnlineUser)

func GetOnlineUserList(c *gin.Context) {
	users := make([]OnlineUser, 0)
	for _, user := range onlineUsers {
		users = append(users, user)
	}
	utils.Success(c, users)
}

func GetTaskList(c *gin.Context) {
	taskName := c.Query("task_name")
	taskGroup := c.Query("task_group")
	status := c.Query("status")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.Task{})

	if taskName != "" {
		query = query.Where("task_name LIKE ?", "%"+taskName+"%")
	}
	if taskGroup != "" {
		query = query.Where("task_group = ?", taskGroup)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var tasks []models.Task
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("id DESC").
		Find(&tasks)

	utils.Page(c, tasks, total, pageNum, pageSize)
}

func GetTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var task models.Task
	if err := config.DB.Where("id = ?", id).First(&task).Error; err != nil {
		utils.ErrorNotFound(c, "任务不存在")
		return
	}

	utils.Success(c, task)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	task.CreateBy = username.(string)

	if err := config.DB.Create(&task).Error; err != nil {
		utils.ErrorInternalServer(c, "创建任务失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", task)
}

func UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var task models.Task
	if err := config.DB.Where("id = ?", id).First(&task).Error; err != nil {
		utils.ErrorNotFound(c, "任务不存在")
		return
	}

	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	updateTask.UpdateBy = username.(string)

	if err := config.DB.Model(&task).Updates(&updateTask).Error; err != nil {
		utils.ErrorInternalServer(c, "更新任务失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteTask(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Delete(&models.Task{}, ids).Error; err != nil {
		utils.ErrorInternalServer(c, "删除任务失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}

func GetTaskLogList(c *gin.Context) {
	taskName := c.Query("task_name")
	taskGroup := c.Query("task_group")
	status := c.Query("status")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.TaskLog{})

	if taskName != "" {
		query = query.Where("task_name LIKE ?", "%"+taskName+"%")
	}
	if taskGroup != "" {
		query = query.Where("task_group = ?", taskGroup)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var taskLogs []models.TaskLog
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("create_time DESC").
		Find(&taskLogs)

	utils.Page(c, taskLogs, total, pageNum, pageSize)
}

func DeleteTaskLog(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Delete(&models.TaskLog{}, ids).Error; err != nil {
		utils.ErrorInternalServer(c, "删除任务日志失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}

type ServerInfo struct {
	CPU    CPUInfo    `json:"cpu"`
	Memory MemoryInfo `json:"memory"`
	Disk   DiskInfo   `json:"disk"`
	Server ServerMeta `json:"server"`
}

type CPUInfo struct {
	Cores       int     `json:"cores"`
	Usage       float64 `json:"usage"`
	ModelName   string  `json:"model_name"`
	LoadAverage float64 `json:"load_average"`
}

type MemoryInfo struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	Usage       float64 `json:"usage"`
	GoTotal     uint64  `json:"go_total"`
	GoUsed      uint64  `json:"go_used"`
}

type DiskInfo struct {
	Total uint64  `json:"total"`
	Used  uint64  `json:"used"`
	Free  uint64  `json:"free"`
	Usage float64 `json:"usage"`
}

type ServerMeta struct {
	HostName string `json:"host_name"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	GoVersion string `json:"go_version"`
}

func GetServerInfo(c *gin.Context) {
	info := ServerInfo{}

	// CPU信息
	cpuPercent, _ := cpu.Percent(time.Second, false)
	cpuInfoList, _ := cpu.Info()
	loadInfo, _ := load.Avg()

	if len(cpuInfoList) > 0 {
		info.CPU.Cores = runtime.NumCPU()
		info.CPU.ModelName = cpuInfoList[0].ModelName
	}
	if len(cpuPercent) > 0 {
		info.CPU.Usage = cpuPercent[0]
	}
	if loadInfo != nil {
		info.CPU.LoadAverage = loadInfo.Load1
	}

	// 内存信息
	memInfo, _ := mem.VirtualMemory()
	if memInfo != nil {
		info.Memory.Total = memInfo.Total
		info.Memory.Used = memInfo.Used
		info.Memory.Free = memInfo.Free
		info.Memory.Usage = memInfo.UsedPercent
	}

	// Go运行时内存
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	info.Memory.GoTotal = m.TotalAlloc
	info.Memory.GoUsed = m.Alloc

	// 磁盘信息
	diskInfo, _ := disk.Usage("/")
	if diskInfo != nil {
		info.Disk.Total = diskInfo.Total
		info.Disk.Used = diskInfo.Used
		info.Disk.Free = diskInfo.Free
		info.Disk.Usage = diskInfo.UsedPercent
	}

	// 服务器信息
	hostInfo, _ := host.Info()
	if hostInfo != nil {
		info.Server.HostName = hostInfo.Hostname
		info.Server.OS = hostInfo.OS
		info.Server.Arch = hostInfo.KernelArch
	}
	info.Server.GoVersion = runtime.Version()

	utils.Success(c, info)
}

func GetDBPoolInfo(c *gin.Context) {
	if config.DB == nil {
		utils.ErrorInternalServer(c, "数据库未连接")
		return
	}

	sqlDB := config.DB.DB()
	stats := sqlDB.Stats()

	utils.Success(c, gin.H{
		"max_open_connections": stats.MaxOpenConnections,
		"open_connections":     stats.OpenConnections,
		"in_use":               stats.InUse,
		"idle":                 stats.Idle,
		"wait_count":           stats.WaitCount,
		"wait_duration":        stats.WaitDuration.Seconds(),
		"max_idle_closed":      stats.MaxIdleClosed,
		"max_lifetime_closed":  stats.MaxLifetimeClosed,
	})
}
