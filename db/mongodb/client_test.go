package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/hasonx/gokit/define"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CollectionFaceLog = "t_face_log"
)

// 识别记录
type FaceLog struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"` // 主键
	Version         uint32             `bson:"version" json:"version"`
	RecordId        uint32             `bson:"recordID" json:"recordID"`
	RecordNo        string             `bson:"recordNo" json:"recordNo"`       // 记录编码
	PersonScore     float32            `bson:"personScore" json:"personScore"` // 人脸评分
	ImagePath       string             `bson:"imagePath" json:"imagePath"`     // 抓拍人像
	IrImagePath     string             `bson:"irImagePath" json:"irImagePath"`
	IdCardImagePath string             `bson:"idCardImagePath" json:"idCardImagePath"`
	Status          string             `bson:"status" json:"status"`                 // 人员类型
	MaskStatus      string             `bson:"maskStatus" json:"maskStatus"`         // 口罩
	IdVerifyStatus  string             `bson:"idVerifyStatus" json:"idVerifyStatus"` // 人证比对
	IdVerifyScore   float32            `bson:"idVerifyScore" json:"idVerifyScore"`   // 人证比对分数
	Temperature     float32            `bson:"temperature" json:"temperature"`       // 温度
	DeviceId        string             `bson:"deviceID" json:"deviceID"`             // 设备ID
	MachineId       string             `bson:"machineID" json:"machineID"`
	DeviceType      string             `bson:"deviceType" json:"deviceType"`
	PassStatus      string             `bson:"passStatus" json:"passStatus"` // 通行状态
	DenyReason      string             `bson:"denyReason" json:"denyReason"`
	ProductModel    string             `bson:"productModel" json:"productModel"`
	PassCondition   *PassCondition2    `bson:"passCondition" json:"passCondition"`
	IdCardInfo      *IDCardInfo2       `bson:"idCardInfo" json:"idCardInfo"`
	PersonInfo      *PersonInfo2       `bson:"person" json:"person"`
	LocationInfo    *LocationInfo2     `bson:"locationInfo" json:"locationInfo"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"` // 创建时间
	// HealthInfo      *HealthInfo2    `bson:"health_info" json:"healthInfo"`
}

type PassCondition2 struct {
	Version                 int32   `bson:"version" json:"version,omitempty"`
	PersonStatus            bool    `bson:"personStatus" json:"personStatus"` // 人员状态
	PersonMinScore          float32 `bson:"personMinScore" json:"personMinScore"`
	Temperature             bool    `bson:"temperature" json:"temperature"`
	TemperatureMax          float32 `bson:"temperatureMax" json:"temperatureMax"`
	TemperatureMin          float32 `bson:"temperatureMin" json:"temperatureMin"`
	Mask                    bool    `bson:"mask" json:"mask"`
	MaskMinScore            float32 `bson:"maskMinScore" json:"maskMinScore"`
	Liveness                bool    `bson:"liveness" json:"liveness"`
	LivenessBGRMinScore     float32 `bson:"livenessBgrMinScore" json:"livenessBgrMinScore"`
	LivenessNIRMinScore     float32 `bson:"livenessNirMinScore" json:"livenessNirMinScore"`
	IDVerify                bool    `bson:"idVerify" json:"idVerify"`
	IDVerifyMinScore        float32 `bson:"idVerifyMinScore" json:"idVerifyMinScore"`
	NucleicAcidTest         bool    `bson:"nucleicAcidTest" json:"nucleicAcidTest"`
	NucleicAcidTestMaxHours int32   `bson:"nucleicAcidTestMaxHours" json:"nucleicAcidTestMaxHours"`
	Trip                    bool    `bson:"trip" json:"trip"`
	TripAllowUnknown        bool    `bson:"tripAllowUnknown" json:"tripAllowUnknown"`
	Vaccine                 bool    `bson:"vaccine" json:"vaccine"`
	VaccineMinDose          int32   `bson:"vaccineMinDose" json:"vaccineMinDose"`
	HealthCode              bool    `bson:"healthCode" json:"healthCode"`
	HealthCodeAllowGreen    bool    `bson:"healthCodeAllowGreen" json:"healthCodeAllowGreen"`
	HealthCodeAllowYellow   bool    `bson:"healthCodeAllowYellow" json:"healthCodeAllowYellow"`
	HealthCodeAllowRed      bool    `bson:"healthCodeAllowRed" json:"healthCodeAllowRed"`
	HealthCodeAllowGray     bool    `bson:"healthCodeAllowGray" json:"healthCodeAllowGray"`
}
type IDCardInfo2 struct {
	Id             string `bson:"id" json:"id"`
	Name           string `bson:"name" json:"name"`
	Gender         string `bson:"gender" json:"gender"`
	Address        string `bson:"address" json:"address"`
	Birthday       string `bson:"birthday" json:"birthday"`
	National       string `bson:"national" json:"national"`
	StartDate      string `bson:"startDate" json:"startDate"`
	EndDate        string `bson:"endDate" json:"endDate"`
	Maker          string `bson:"maker" json:"maker"`
	PhotoImagePath string `bson:"photoImagePath" json:"photoImagePath"`
}

type PersonInfo2 struct {
	Id             uint32     `bson:"id" json:"id"`
	TrdId          uint32     `bson:"trdID" json:"trdID"`
	Number         string     `bson:"number" json:"number"`
	NumberType     string     `bson:"numberType" json:"numberType"`
	Name           string     `bson:"name" json:"name"` // 姓名
	Gender         int32      `bson:"gender" json:"gender"`
	Age            uint32     `bson:"age" json:"age"`
	Department     string     `bson:"department" json:"department"`
	Mobile         string     `bson:"mobile" json:"mobile"`
	IdCardNo       string     `bson:"idCardNo" json:"idCardNo"`
	AccessCardType string     `bson:"accessCardType" json:"accessCardType"`
	AccessCardNo   string     `bson:"accessCardNo" json:"accessCardNo"`
	NativePlace    string     `bson:"nativePlace" json:"nativePlace"` // 民族
	Nation         string     `bson:"nation" json:"nation"`
	Address        string     `bson:"address" json:"address"`
	Note           string     `bson:"note" json:"note"`
	Status         string     `bson:"status" json:"status"`
	FaceImage      string     `bson:"faceImage" json:"faceImage"`
	CreatedAt      *time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt      *time.Time `bson:"updatedAt" json:"updatedAt"`
}
type LocationInfo2 struct {
	Latitude     float32 `bson:"latitude" json:"latitude" form:"latitude"`
	Longitude    float32 `bson:"longitude" json:"longitude" form:"longitude"`
	Accuracy     float32 `bson:"accuracy" json:"accuracy" form:"accuracy"`
	Altitude     float32 `bson:"altitude" json:"altitude" form:"altitude"`
	Bearing      float32 `bson:"bearing" json:"bearing" form:"bearing"`
	Speed        float32 `bson:"speed" json:"speed" form:"speed"`
	Country      string  `bson:"country" json:"country" form:"country"`
	Province     string  `bson:"province" json:"province" form:"province"`
	City         string  `bson:"city" json:"city" form:"city"`
	District     string  `bson:"district" json:"district" form:"district"`
	Street       string  `bson:"street" json:"street" form:"street"`
	StreetNumber string  `bson:"streetNumber" json:"streetNumber" form:"streetNumber"`
	CityCode     string  `bson:"cityCode" json:"cityCode" form:"cityCode"`
	Address      string  `bson:"address" json:"address" form:"address"`
}

type GetFaceRecordsReq struct {
	RecordNo   string   `form:"recordNo"`    // 记录编码
	PersonName string   `form:"personName"`  // 录入姓名
	DeviceId   string   `form:"deviceId"`    // 设备mac
	Status     []string `form:"status[]"`    // 人员类型
	PassStatus string   `form:"passStatus"`  // 通行状态
	MaskStatus string   `form:"maskStatus"`  // 口罩
	IsList     bool     `form:"isList"`      // 是否列表（用于筛选字段）
	DateRange  []string `form:"dateRange[]"` // 查询时间段
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

type ServiceTestSuite struct {
	suite.Suite
	faceLog *mongo.Collection
}

func (s *ServiceTestSuite) SetupSuite() {
	viper.SetConfigFile("./config-local.yml")
	err := viper.ReadInConfig()
	s.Require().Nil(err)
	mongoClient, err := NewMongo(
		viper.GetString("mongodb.uri"),
		viper.GetString("mongodb.name"),
		viper.GetDuration("mongodb.timeout"),
		viper.GetUint64("mongodb.num"),
	)
	s.Require().Nil(err)
	s.faceLog = mongoClient.Db().Collection(CollectionFaceLog)
}

func (s *ServiceTestSuite) TestService_GetAggregate() {
	ctx := context.Background()

	filter := make(bson.M)
	filter["createdAt"] = bson.M{"$gte": primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -2))}
	matchStage := bson.D{
		{Key: "$match", Value: filter},
	}
	pipeline := bson.A{
		matchStage,
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "$hour", Value: "$createdAt"},
				{Key: "$dayOfMonth", Value: "$createdAt"},
			}},
			{Key: "count", Value: bson.D{
				{Key: "$sum", Value: 1}, // 对每个小时的记录数量进行求和
			}},
		}}},
		// bson.D{{Key: "$project", Value: bson.D{
		// 	{Key: "_id", Value: 0}, // 排除_id字段
		// 	{Key: "hour", Value: "$_id"},
		// 	{Key: "count", Value: 1},
		// }}},
	}

	cursor, err := s.faceLog.Aggregate(ctx, pipeline)
	s.Require().Nil(err)
	defer func() {
		_ = cursor.Close(ctx)
	}()

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		s.Require().Nil(err)
	}

	// 打印每小时记录的数量
	for _, result := range results {
		hour := result["hour"].(int32)
		count := result["count"].(int32)
		s.T().Logf("Hour: %d, Count: %d\n", hour, count)
	}
}

func (s *ServiceTestSuite) TestService_GetFaceLog() {
	req := GetFaceRecordsReq{
		RecordNo:   "",
		PersonName: "",
		DeviceId:   "",
		Status:     nil,
		PassStatus: "",
		MaskStatus: "",
		IsList:     true,
		DateRange:  nil,
	}
	opts := &options.FindOptions{}
	opts.SetSort(bson.M{"createdAt": -1})
	opts.SetSkip(0)
	opts.SetLimit(10)
	c := context.Background()
	filter := make(bson.M)

	if req.DeviceId != "" {
		filter["deviceID"] = primitive.Regex{
			Pattern: req.DeviceId,
			Options: "i",
		}
	}
	if req.PersonName != "" {
		filter["person.name"] = primitive.Regex{
			Pattern: req.PersonName,
			Options: "i",
		}
	}
	if len(req.Status) > 0 {
		filter["status"] = bson.M{"$in": req.Status}
	}
	if req.PassStatus != "" {
		filter["passStatus"] = req.PassStatus
	}

	if req.MaskStatus != "" {
		filter["maskStatus"] = req.MaskStatus
	}

	var st, et time.Time
	if len(req.DateRange) == 2 {
		start, err := time.Parse(define.TimeFormat, req.DateRange[0])
		if err == nil {
			st = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
		}
		end, err := time.Parse(define.TimeFormat, req.DateRange[1])
		if err == nil {
			et = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, end.Location())
		}
	}
	if st.IsZero() || et.IsZero() {
		// 默认只查询近三个月的
		filter["createdAt"] = bson.M{"$gte": primitive.NewDateTimeFromTime(time.Now().AddDate(0, -3, 0))}
	} else {
		filter["createdAt"] = bson.M{"$gte": primitive.NewDateTimeFromTime(st), "$lte": primitive.NewDateTimeFromTime(et)}
	}

	data := bson.A{
		bson.D{{Key: "$skip", Value: opts.Skip}},
		bson.D{{Key: "$limit", Value: opts.Limit}},
	}
	// 筛选字段
	if req.IsList {
		opts.SetProjection(bson.M{
			"_id":             0,
			"deviceID":        1,
			"recordNo":        1,
			"person.idCardNo": 1,
			"person.name":     1,
		})
		data = append(data, bson.D{{Key: "$project", Value: opts.Projection}})
	}
	// 分页查询并返回总数
	matchStage := bson.D{
		{Key: "$match", Value: filter},
	}
	sortStage := bson.D{
		{Key: "$sort", Value: opts.Sort},
	}
	count := bson.A{
		bson.D{{Key: "$count", Value: "count"}},
	}
	facetStage := bson.D{
		{Key: "$facet", Value: bson.D{
			{Key: "count", Value: count},
			{Key: "data", Value: data},
		}},
	}

	cur, err := s.faceLog.Aggregate(c, mongo.Pipeline{matchStage, sortStage, facetStage})
	s.Require().Nil(err)
	defer func() {
		_ = cur.Close(c)
	}()

	// var count2 []struct {
	// 	total int64 `bson:"total"`
	// }
	// var data2 []FaceLog

	var des []struct {
		Data  []FaceLog `bson:"data"`
		Count []struct {
			Count int64 `bson:"count"`
		} `bson:"count"`
	}
	// if cur.Next(c) {
	// 	if err = cur.Decode(&des); err != nil {
	// 		s.Require().Nil(err)
	// 	}
	// }
	if err = cur.All(c, &des); err != nil {
		s.Require().Nil(err)
	}
	s.T().Logf("res:%v", des)

	// var total int64
	// var data []FaceLog

	// var des []struct {
	// 	data  string `bson:"data"`
	// 	count string `bson:"count"`
	// }
	// if err = cur.All(c, &des); err != nil {
	// 	s.Require().Nil(err)
	// }
	// for _, de := range des {
	// 	s.T().Logf("count2:%v", de.count)
	// 	s.T().Logf("data2:%v", de.data)
	// }

	// var results []bson.D
	// if err := cur.All(c, &results); err != nil {
	// 	s.Require().Nil(err)
	// }
	// s.T().Logf("res:%v", results)
	// for _, result := range results {
	// 	m1 := result["count"]
	// 	s.T().Logf("count2:%v", m1)
	// 	c1, _ := bson.Marshal(result["count"])
	// 	err := bson.Unmarshal(c1, &count2)
	// 	s.Require().Nil(err)
	// 	s.T().Logf("count2:%v", count2)
	// }
	// for _, result := range results {
	//
	// 	s.T().Logf("data2:%v", result["data"])
	// 	c1, _ := bson.Marshal(result["count"])
	// 	d1, _ := bson.Marshal(result["data"])
	// 	s.T().Logf("c1:%s", string(c1))
	// 	s.T().Logf("d1:%s", string(d1))
	//
	// 	err := bson.Unmarshal(c1, &count2)
	// 	s.Require().Nil(err)
	// 	err = bson.Unmarshal(d1, &data2)
	// 	s.Require().Nil(err)
	// }
	// s.T().Logf("count2:%v", count2)
	// s.T().Logf("data2:%v", data2)
}
