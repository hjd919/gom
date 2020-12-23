package lbs

import (
	"encoding/json"
	"fmt"

	"github.com/hjd919/gom"
	"github.com/hjd919/gom/lbs/lbsqq"
)

type QQ struct {
	Key        string
	prefixHost string
}

func (t *QQ) Init() {
	t.prefixHost = "https://apis.map.qq.com/ws"
}

// 逆地址
func (t *QQ) GeocodeRegeo(lat, lng float64) (address string, err error) {
	url := t.prefixHost + fmt.Sprintf("/geocoder/v1/?location=%f,%f&key=%s", lat, lng, t.Key)
	buf, err := gom.HTTPGet(url)
	if err != nil {
		return
	}
	entity := lbsqq.GeocodeRegeo{}
	json.Unmarshal(buf, &entity)
	if entity.Status != 0 {
		err = fmt.Errorf(entity.Message)
		return
	}
	address = entity.Result.Address
	return
}
