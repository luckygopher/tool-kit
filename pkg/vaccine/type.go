package vaccine

type Member struct {
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data []struct {
		ID           int    `json:"id"`
		UserID       int    `json:"userId"`
		Name         string `json:"name"`
		IDCardNo     string `json:"idCardNo"`
		Birthday     string `json:"birthday"`
		Sex          int    `json:"sex"`
		RegionCode   string `json:"regionCode"`
		Address      string `json:"address"`
		IsDefault    int    `json:"isDefault"`
		RelationType int    `json:"relationType"`
		CreateTime   string `json:"createTime"`
		ModifyTime   string `json:"modifyTime"`
		Yn           int    `json:"yn"`
		IDCardType   int    `json:"idCardType"`
	} `json:"data"`
	Ok    bool `json:"ok"`
	NotOk bool `json:"notOk"`
}

type Area struct {
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"data"`
	Ok    bool `json:"ok"`
	NotOk bool `json:"notOk"`
}

type ST struct {
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data struct {
		Stock int   `json:"stock"`
		St    int64 `json:"st"`
	} `json:"data"`
	Ok    bool `json:"ok"`
	NotOk bool `json:"notOk"`
}

type VaccineList struct {
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Address     string `json:"address"`
		VaccineCode string `json:"vaccineCode"`
		VaccineName string `json:"vaccineName"`
		StartTime   string `json:"startTime"`
	} `json:"data"`
	Ok    bool `json:"ok"`
	NotOk bool `json:"notOk"`
}

type Subscribe struct {
	Code  string `json:"code"`
	Msg   string `json:"msg,omitempty"`
	Ok    bool   `json:"ok"`
	NotOk bool   `json:"notOk"`
}
