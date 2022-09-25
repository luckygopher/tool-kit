package vaccine

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/qingyunjun/tool-kit/pkg/httpclient"
	"go.uber.org/zap"
)

func (c Client) Start() error {
	var (
		vaccineID string
		startTime string
		st        string
		err       error
	)
	// 先获取指定地区的疫苗列表
	vaccines, err := c.GetVaccineList(c.cfg.RegionCode)
	if err != nil {
		c.logger.Error("get vaccine list failed", zap.Error(err))
		return errors.New(err.Error())
	}
	c.logger.Debug("get vaccine list success", zap.Any("data", vaccines))
	// 从列表结果中获取指定疫苗的信息
	for i := 0; i < len(vaccines.Data); i++ {
		if strconv.Itoa(vaccines.Data[i].ID) == c.cfg.VaccineID {
			startTime = vaccines.Data[i].StartTime
			vaccineID = strconv.Itoa(vaccines.Data[i].ID)
			break
		}
	}
	if vaccineID == "" {
		return errors.New("未获取到抢购的疫苗信息,请检查疫苗列表中是否有配置文件指定的疫苗！")
	}
	// 判断疫苗秒杀开始时间
	startDateTime, err := time.ParseInLocation(Layout, startTime, time.Local)
	if err != nil {
		c.logger.Error("parse start time failed", zap.Error(err), zap.String("startTime", startTime))
		return errors.New("解析时间错误！")
	}
	getSTTime := time.Now().Add(5 * time.Second)
	// 当前时间+5秒之后是否在疫苗秒杀开始时间之前
	if getSTTime.Before(startDateTime) {
		c.logger.Info("还未到获取st时间，等待中......")
		time.Sleep(startDateTime.Sub(getSTTime))
	}
	c.logger.Info("到达获取st时间！")
	// 循环获取st，直到成功为止
	for {
		if st, err = c.GetST(c.cfg.VaccineID); err == nil {
			break
		}
		c.logger.Error("get st failed", zap.Error(err))
	}
	// 当前时间+500毫秒
	now := time.Now().Add(500 * time.Millisecond)
	if now.Before(startDateTime) {
		c.logger.Info("获取st成功，但是还未到抢购时间，等待中......")
		time.Sleep(startDateTime.Sub(now))
	}
	c.logger.Info("秒杀时间到，开始秒杀！")
	c.Seckill(vaccineID, st)
	return nil
}

// GetST 获取加密参数st
func (c Client) GetST(vaccineID string) (string, error) {
	url := c.cfg.BaseURL + "/seckill/seckill/checkstock2.do"
	params := map[string]string{
		"id": vaccineID,
	}
	headers := c.CommonHeader()
	result := ST{}
	_, err := httpclient.HTTPClient.Get(context.TODO(), url, headers, nil, params, &result)
	if err != nil {
		c.logger.Error("get failed", zap.Error(err))
		return "", err
	}
	// 判断是否业务错误
	if result.Code != SuccessCode {
		return "", errors.New(result.Msg)
	}
	return strconv.FormatInt(result.Data.St, 10), nil
}

// Seckill 秒杀疫苗
func (c Client) Seckill(vaccineID, st string) {
	wg := sync.WaitGroup{}
	wg.Add(c.cfg.Total)
	success := false
	for i := 0; i < c.cfg.Total; i++ {
		c.logger.Info(fmt.Sprintf("当前第%d个协程正在秒杀！", i+1))
		go func(vaccineID, st string) {
			resp, err := c.Subscribe(vaccineID, c.cfg.MemberID, c.cfg.IDCard, st)
			if err != nil {
				c.logger.Info(fmt.Sprintf("当前第%d个协程正在秒杀失败！", i+1), zap.Error(err))
			} else {
				c.logger.Info(fmt.Sprintf("当前第%d个协程秒杀成功！", i+1), zap.String("resp", resp))
				success = true
			}
		}(vaccineID, st)
		c.logger.Info(fmt.Sprintf("正在休息%d毫秒，等待下一个协程秒杀", c.cfg.Step))
		time.Sleep(time.Duration(c.cfg.Step) * time.Millisecond)
		wg.Done()
	}
	wg.Wait()

	if success {
		c.logger.Info("抢购成功，请在小程序中查看！")
	} else {
		c.logger.Info("所有协程都抢购失败，再接再厉！")
	}
}

// Subscribe 秒杀订阅
func (c Client) Subscribe(vaccineID, memberID, idCard, st string) (string, error) {
	url := c.cfg.BaseURL + "/seckill/seckill/subscribe.do"
	params := map[string]string{
		"seckillId":    vaccineID,
		"vaccineIndex": "1",
		"linkmanId":    memberID,
		"idCardNo":     idCard,
	}
	headers := c.CommonHeader()
	headers["ecc-hs"] = c.EccHs(vaccineID, memberID, st)
	result := Subscribe{}
	resp, err := httpclient.HTTPClient.Get(context.TODO(), url, headers, nil, params, &result)
	if err != nil {
		c.logger.Error("get failed", zap.Error(err))
		return "", err
	}
	// 判断是否业务错误
	if result.Code != SuccessCode {
		return "", errors.New(result.Msg)
	}
	return string(resp), nil
}

// EccHs 秒杀获取header信息ecc-hs
func (c Client) EccHs(vaccineID, memberID, st string) string {
	salt := Salt
	sign1 := md5.New()
	sign1.Write([]byte(vaccineID + memberID + st))
	data := hex.EncodeToString(sign1.Sum(nil))

	sign2 := md5.New()
	sign2.Write([]byte(data + salt))
	return hex.EncodeToString(sign2.Sum(nil))
}
