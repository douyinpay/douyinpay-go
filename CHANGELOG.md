## 版本变更记录

代码位置：client/option/name_option.go 

ClientVersion 每次sdk升级，需要更新版本号

v1.0.6 - 2026-08-18
--------------------
* 新增前置咨询接口（services.cashier.ApiCashierService.PrePayConsult），支持咨询抖音支付渠道、先享后付渠道的可用性与营销信息
* 补充前置咨询接口的字段注释和文档说明
* 新增直连商户分账能力：请求分账、查询分账结果、请求分账回退、查询分账回退结果、完结分账、查询剩余待分金额、添加/删除分账接收方
* 新增服务商分账能力：请求分账、查询分账结果、请求分账回退、查询分账回退结果、完结分账、查询剩余待分金额、查询特约商户分账配置、添加/删除分账接收方，以及分账结果/接收方入账通知解析
* 补充分账相关路径常量、服务商分账账单请求模型注释，以及回调资源 `sp_mchid` 字段支持

v1.0.5 - 2026-08-04
--------------------
* 直连/服务商交易模型（Transaction）新增扣款失败通知字段：ContractId、ErrCode、ErrCodeDes
* 补充直连与服务商开放接口的字段注释和文档说明

v 1.0.4 - 2026-07-10
--------------------
* 新增H5支付中签约下单接口

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
