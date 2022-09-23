package vaccine

import (
	"context"
	"errors"

	"github.com/qingyunjun/tool-kit/pkg/httpclient"
	"go.uber.org/zap"
)

// GetVaccineList 获取疫苗列表
func (c Client) GetVaccineList() (VaccineList, error) {
	url := c.cfg.BaseURL + "/seckill/seckill/list.do"
	params := map[string]string{
		"offset":     "0",
		"limit":      "100",
		"regionCode": c.cfg.RegionCode, // 4位，例如成都：5101
	}
	headers := c.CommonHeader()
	result := VaccineList{}
	_, err := httpclient.HTTPClient.Get(context.TODO(), url, headers, nil, params, &result)
	if err != nil {
		c.logger.Error("get failed", zap.Error(err))
		return result, err
	}
	// 判断是否业务错误
	if result.Code != SuccessCode {
		return result, errors.New(result.Msg)
	}

	return result, nil
}

// GetArea 获取某省区域的所有市信息
func (c Client) GetArea() (Area, error) {
	url := c.cfg.BaseURL + "/base/region/childRegions.do"
	headers := c.CommonHeader()
	param := map[string]string{
		"parentCode": c.cfg.ParentCode,
	}
	result := Area{}
	_, err := httpclient.HTTPClient.Get(context.TODO(), url, headers, nil, param, &result)
	if err != nil {
		c.logger.Error("get failed", zap.Error(err))
		return result, err
	}
	// 判断是否业务错误
	if result.Code != SuccessCode {
		return result, errors.New(result.Msg)
	}
	return result, nil
}

// GetMember 获取账户信息
func (c Client) GetMember() (Member, error) {
	url := c.cfg.BaseURL + "/seckill/linkman/findByUserId.do"
	headers := c.CommonHeader()
	result := Member{}
	_, err := httpclient.HTTPClient.Get(context.TODO(), url, headers, nil, nil, &result)
	if err != nil {
		c.logger.Error("get failed", zap.Error(err))
		return result, err
	}
	// 判断是否业务错误
	if result.Code != SuccessCode {
		return result, errors.New(result.Msg)
	}
	return result, nil
}
