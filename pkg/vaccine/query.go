package vaccine

import (
	"context"
	"errors"
	"time"

	"github.com/luckygopher/tool-kit/pkg/httpclient"
	"go.uber.org/zap"
)

// GetVaccineList 获取疫苗列表
func (c Client) GetVaccineList(regionCode string) (VaccineList, error) {
	url := c.cfg.BaseURL + "/seckill/seckill/list.do"
	params := map[string]string{
		"offset":     "0",
		"limit":      "100",
		"regionCode": regionCode, // 4位，例如成都：5101
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
func (c Client) GetArea(parentCode string) (Area, error) {
	url := c.cfg.BaseURL + "/base/region/childRegions.do"
	headers := c.CommonHeader()
	param := map[string]string{
		"parentCode": parentCode,
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

// FindAreaVaccine 查询某省当前有苗的地区及该地区可秒疫苗
func (c Client) FindAreaVaccine(parentCode string) error {
	area, err := c.GetArea(parentCode)
	if err != nil {
		c.logger.Error("FindAreaVaccine:获取地区列表错误", zap.String("parentCode", parentCode))
		return err
	}
	for _, item := range area.Data {
		list, err := c.GetVaccineList(item.Value)
		if err != nil {
			c.logger.Error("FindAreaVaccine:获取疫苗列表错误", zap.String(item.Name, item.Value), zap.Error(err))
			continue
		}
		if len(list.Data) > 0 {
			c.logger.Info("疫苗情况", zap.String("地区", item.Name), zap.String("地区码", item.Value),
				zap.Any("疫苗列表", list.Data))
		}
		time.Sleep(3 * time.Second)
	}
	c.logger.Info("以上是本次查询到的结果")
	return nil
}
