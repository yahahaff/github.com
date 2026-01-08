// Package ssl
package ssl

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	sslModel "github.com/yahahaff/rapide/internal/models/ssl"
	requestsSSL "github.com/yahahaff/rapide/internal/requests/ssl"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/internal/utils"
	"github.com/yahahaff/rapide/pkg/response"
)

// SSLCertController SSL证书控制器
type SSLCertController struct {
	controllers.BaseAPIController
}

// GetSSLCertList 获取SSL证书列表
func (ctrl *SSLCertController) GetSSLCertList(c *gin.Context) {
	request := requestsSSL.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 10 // 设置默认值
	}

	// 处理页码参数，确保页码大于0
	page := request.Page
	if page <= 0 {
		page = 1
	}

	// 获取查询参数
	domain := request.Domain
	applyStatus := request.ApplyStatus

	data, total, err := service.Entrance.SSLService.SSLCertService.GetSSLCertList(page, pageSize, domain, applyStatus)
	if err != nil {
		response.Abort500(c, "获取SSL证书列表失败")
		return
	}

	result := gin.H{
		"list":     data,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}
	response.OK(c, result)
}

// CreateSSLCert 创建SSL证书
func (ctrl *SSLCertController) CreateSSLCert(c *gin.Context) {
	request := requestsSSL.SSLCertCreateRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 设置默认值
	provider := request.Provider
	if provider == "" {
		provider = "letsencrypt"
	}

	challengeType := request.ChallengeType
	if challengeType == "" {
		// 根据verifyMethod设置默认的验证方式
		if request.VerifyMethod == "auto-dns" {
			challengeType = "dns-01"
		} else {
			challengeType = "http-01"
		}
	}

	// 转换请求数据为证书模型
	cert := sslModel.SSLCert{
		Domain:           request.Domain,
		CommonName:       request.CommonName,
		Organization:     request.Organization,
		OrganizationUnit: request.OrganizationUnit,
		Country:          request.Country,
		State:            request.State,
		City:             request.City,
		Email:            request.Email,
		Type:             "DV", // Let's Encrypt 只提供 DV 证书
		Algorithm:        request.Algorithm,
		Provider:         provider,
		ChallengeType:    challengeType,
		ApplyStatus:      "pending",
		AutoRenew:        request.AutoRenew,
		RenewStatus:      "idle",
		Status:           1, // 默认为启用状态
	}

	// 调用服务层创建证书
	err := service.Entrance.SSLService.SSLCertService.CreateSSLCert(cert)
	if err != nil {
		response.Abort500(c, "创建SSL证书失败")
		return
	}

	response.OK(c, gin.H{"message": "SSL证书创建成功，正在申请中"})
}

// DownloadSSLCert 下载SSL证书，支持多种格式
// @Summary 下载SSL证书
// @Description 下载SSL证书，支持pem、pfx、pkcs12、jks、der等格式，format为all时返回包含所有格式的压缩包
// @Tags SSL证书
// @Accept json
// @Produce octet-stream
// @Param id path string true "证书ID"
// @Param format query string false "证书格式，默认为all（返回包含所有格式的压缩包）"
// @Param password query string false "证书密码，默认为changeme"
// @Success 200 {file} binary "证书文件或压缩包"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "证书不存在"
// @Failure 500 {object} response.Response "下载失败"
// @Router /api/ssl/download/{id} [get]
func (ctrl *SSLCertController) DownloadSSLCert(c *gin.Context) {
	// 1. 获取证书ID和格式
	certID := c.Param("id")
	format := c.DefaultQuery("format", "all")
	password := c.DefaultQuery("password", "changeme")

	// 2. 获取证书信息
	cert, err := service.Entrance.SSLService.SSLCertService.GetSSLCertByID(certID)
	if err != nil {
		response.Abort404(c, "证书不存在")
		return
	}

	// 3. 验证证书状态
	if cert.ApplyStatus != "success" {
		response.Abort400(c, "证书尚未申请成功，无法下载")
		return
	}

	// 4. 根据格式生成证书
	if format == "all" {
		// 生成包含所有格式的压缩包
		zipContent, err := utils.GenerateAllFormatsCertPackage(cert.Certificate, cert.PrivateKey, cert.IntermediateCert, cert.Domain, password)
		if err != nil {
			response.Abort500(c, "生成证书压缩包失败: "+err.Error())
			return
		}

		// 返回压缩包
		c.Writer.Header().Set("Content-Description", "File Transfer")
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s-certs.zip\"", cert.Domain))
		c.Writer.Header().Set("Content-Type", "application/zip")
		c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(zipContent)))
		c.Writer.WriteHeader(200)
		c.Writer.Write(zipContent)
		c.Writer.Flush()
		return
	} else {
		// 生成指定格式的证书包
		packageFiles, err := utils.GenerateCertPackage(format, cert.Certificate, cert.PrivateKey, cert.IntermediateCert, cert.Domain, password)
		if err != nil {
			response.Abort500(c, "生成证书包失败: "+err.Error())
			return
		}

		// 返回证书文件
		if len(packageFiles) == 1 {
			// 单个文件，直接返回
			for filename, content := range packageFiles {
				c.Header("Content-Description", "File Transfer")
				c.Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
				c.Header("Content-Type", "application/octet-stream")
				c.Header("Content-Length", fmt.Sprintf("%d", len(content)))
				c.Data(200, "application/octet-stream", content)
				return
			}
		} else {
			// 多个文件，返回JSON格式的文件列表
			response.OK(c, gin.H{
				"files":  packageFiles,
				"domain": cert.Domain,
				"format": format,
			})
		}
	}
}

// RevokeSSLCert 吊销SSL证书
// @Summary 吊销SSL证书
// @Description 吊销指定ID的SSL证书
// @Tags SSL证书
// @Accept json
// @Produce json
// @Param id path string true "证书ID"
// @Success 200 {object} response.Response "吊销成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "证书不存在"
// @Failure 500 {object} response.Response "吊销失败"
// @Router /api/ssl/revoke/{id} [post]
func (ctrl *SSLCertController) RevokeSSLCert(c *gin.Context) {
	// 1. 获取证书ID
	certID := c.Param("id")

	// 2. 调用服务层吊销证书
	err := service.Entrance.SSLService.SSLCertService.RevokeSSLCert(certID)
	if err != nil {
		response.Abort500(c, "吊销证书失败: "+err.Error())
		return
	}

	// 3. 返回成功响应
	response.OK(c, gin.H{"message": "证书吊销成功"})
}

// GetSSLCertDetail 获取单个SSL证书详情
// @Summary 获取单个SSL证书详情
// @Description 获取指定ID的SSL证书详情
// @Tags SSL证书
// @Accept json
// @Produce json
// @Param id path string true "证书ID"
// @Success 200 {object} response.Response "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "证书不存在"
// @Failure 500 {object} response.Response "获取失败"
// @Router /api/ssl/detail/{id} [get]
func (ctrl *SSLCertController) GetSSLCertDetail(c *gin.Context) {
	// 1. 获取证书ID
	certID := c.Param("id")

	// 2. 调用服务层获取证书详情
	cert, err := service.Entrance.SSLService.SSLCertService.GetSSLCertByID(certID)
	if err != nil {
		response.Abort404(c, "证书不存在")
		return
	}

	// 3. 计算证书剩余天数
	now := time.Now()
	expiresInDays := int(cert.ValidityEnd.Sub(now).Hours() / 24)
	// 确保剩余天数不为负数
	if expiresInDays < 0 {
		expiresInDays = 0
	}

	// 4. 创建返回结果，包含证书信息和剩余天数
	result := gin.H{
		"id":               cert.ID,
		"domain":           cert.Domain,
		"commonName":       cert.CommonName,
		"organization":     cert.Organization,
		"organizationUnit": cert.OrganizationUnit,
		"country":          cert.Country,
		"state":            cert.State,
		"city":             cert.City,
		"email":            cert.Email,
		"type":             cert.Type,
		"algorithm":        cert.Algorithm,
		"validityStart":    cert.ValidityStart,
		"validityEnd":      cert.ValidityEnd,
		"status":           cert.Status,
		"provider":         cert.Provider,
		"challengeType":    cert.ChallengeType,
		"applyStatus":      cert.ApplyStatus,
		"errorMsg":         cert.ErrorMsg,
		"certificate":      cert.Certificate,
		"privateKey":       cert.PrivateKey,
		"intermediateCert": cert.IntermediateCert,
		"fingerprint":      cert.Fingerprint,
		"serialNumber":     cert.SerialNumber,
		"autoRenew":        cert.AutoRenew,
		"renewStatus":      cert.RenewStatus,
		"created_at":       cert.CreatedAt,
		"updated_at":       cert.UpdatedAt,
		"expiresInDays":    expiresInDays,
	}

	// 5. 返回证书详情
	response.OK(c, result)
}
