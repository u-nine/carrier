package carrier

import "time"

const (
	CategoryAll         = "*"           //所有分类
	CategoryPower       = "power"       //电量
	CategoryVoltage     = "voltage"     //电压
	CategoryElectricity = "electricity" //电流
	CategoryTemperature = "temperature" //温度
	CategoryHumidity    = "humidity"    //湿度
)

type Raw struct {
	CompanyId uint32    `json:"company_id"` //公司ID
	GatewayId uint64    `json:"gateway_id"` //网关ID
	DeviceId  uint64    `json:"device_id"`  //设备ID
	Version   time.Time `json:"Version"`    //数据版本 上报时间
	Order     uint16    `json:"order"`      //数据序列值 多数据口采集口时的位值
	Value     int64     `json:"value"`      //上报值
	Digit     uint8     `json:"digit"`      //小数据点位值
}
