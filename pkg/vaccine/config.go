package vaccine

type Config struct {
	BaseURL    string `toml:"BaseURL" yaml:"base_url"`                      // 基础URL
	TK         string `toml:"TK" yaml:"tk"`                                 // 身份验证鉴权使用，抓包从请求头中获取
	MemberID   string `toml:"MemberID" yaml:"member_id"`                    // 接种成员id
	IDCard     string `toml:"IDCard" yaml:"id_card"`                        // 接种成员身份证号码
	ParentCode string `toml:"ParentCode" yaml:"parent_code" default:"51"`   // 某省的区域代码,四川为 51
	RegionCode string `toml:"RegionCode" yaml:"region_code" default:"5101"` // 选择的地区代码,4位,成都为 5101
	Cookie     string `toml:"Cookie" yaml:"cookie"`                         // cookie
	VaccineID  string `toml:"VaccineID" yaml:"vaccine_id"`                  // 接种疫苗id
	Total      int    `toml:"Total" yaml:"total" default:"5"`               // 开启的协程数量
	Step       int    `toml:"Step" yaml:"step" default:"200"`               // 每个协程抢购时间间隔，单位毫秒
}
