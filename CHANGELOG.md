## 版本变更记录

代码位置：client/option/name_option.go 

ClientVersion 每次sdk升级，需要更新版本号

v 1.0.0 - 2026-01-26
--------------------
* 新增平台证书下载api功能
* 新增平台证书自动更新和验签功能

v 0.0.5 - 2025-09-16
--------------------
* 直联商户部分:  
  * APP纯签约(services.contract.ApiContractService.PreEntrustWeb)
  * H5 预签约(services.contract.ApiContractService.H5EntrustWeb)
  * 查询代扣签约协议(services.contract.ApiContractService.QueryContract)
  * 解约代扣签约协议(services.contract.ApiContractService.DeleteContract)          
  * 申请扣款(services.deduct.ApiDeductService.PayApply)
  * 预约扣费(services.deduct.ApiDeductService.DeductNotify)
* 服务商部分:  
  * 支付中签约下单(services.partnerpay.contract.ApiPartnerContractService.PartnerContractOrder)
  * 查询签约关系(services.partnercontract.ApiPartnerContractOrderService.PartnerQueryContract)
  * 协议解约(services.partnercontract.ApiPartnerContractOrderService.PartnerTerminateContract)
  * 申请扣款(services.deduct.parnter.ApiDeductService.PartnerPayApply)
  * 预约扣费(services.deduct.partner.ApiDeductService.PartnerContractSchedule)
  * 预约扣费结果查询(services.partnerpay.contract.ApiPartnerContractService.PartnerContractScheduleQuery)


v 0.0.4 - 2025-07-04
--------------------
* 消息通知回调（service.callback.Handler.ParseCallback），中 ENCRYPT_TYPE_AES 枚举值由“AES”调整为"AEAD-AES-256-GCM"，匹配线上文档描述

v 0.0.3 - 2025-07-03
--------------------
* 商家账单商家平台新增分账账单下载接口
* 商家账单合作伙伴平台补齐交易账单、资金账单、分账账单下载接口

v 0.0.2 - 2025-05-28
--------------------
* 商户代扣接口（com.douyinpay.api.deduct.ApiDeductService.deduct），响应体增加预支付交易会话标识(prepayId)字段
