# tool-kit
CLI工具箱
## 疫苗工具
```
需要去秒苗小程序抓包
获取一下数据
tk(该小程序的任何接口的请求头中都有此参数)
member_id(/seckill/linkman/findByUserId.do接口中的id的值)
id_card(/seckill/linkman/findByUserId.do接口中的idCardNo的值)
member_name(/seckill/linkman/findByUserId.do接口中的name的值)
region_code(/base/region/childRegions.do?parentCode=51接口中的value的值,注意是精确到市，不是到省，这个值只有4位，默认成都 5101)
cookie(该小程序的任何接口的请求头中都有此参数)
vaccine_id(/seckill/seckill/list.do接口中的id的值)
```