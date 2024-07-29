package consts

import (
	"errors"
	"fmt"
	"strings"
)

const (
	SystemName = "Pride"
)

// common err
var (
	ErrNotExists              = errors.New("数据不存在")
	ErrFindMultiRecords       = errors.New("找到多条符合条件的数据")
	ErrEmpty                  = errors.New("数据为空")
	ErrClientNotExists        = errors.New("访问的客户端不存在")
	ErrAlreadyExists          = errors.New("数据已经存在")
	ErrUnauthorized           = errors.New("未授权")
	ErrUnimplemented          = errors.New("未实现")
	ErrValidationFail         = errors.New("数据校验失败")
	ErrNilValue               = errors.New("值是 nil")
	ErrProjectOverlap         = errors.New("项目地理位置重叠")
	ErrProjectWithoutGeometry = errors.New("项目数据缺失地理围合信息")
	ErrMissingDependency      = errors.New("依赖数据缺失")
	ErrUpdateOnMergedData     = errors.New("更新已经被合并的数据")
	ErrUpdateOnNonLatest      = errors.New("更新非最新版本数据")
	ErrFieldNoDefinition      = errors.New("字段定义不存在")
	ErrSchemaValidationFail   = errors.New("数据 Schema 校验失败")
	ErrNoReadRight            = errors.New("没有读权限")
	ErrNoWriteRight           = errors.New("没有写权限")
	ErrBeyondAuthority        = errors.New("越权错误")
	ErrWrongHouseLayoutInfo   = errors.New("房屋户型信息错误")
	ErrHouseAlreadyGotLayout  = errors.New("房屋已经关联户型")
	ErrNoRight                = errors.New("没有数据权限")
	ErrMissParam              = errors.New("缺少参数")
	ErrSwitchNotSupport       = errors.New("状态不支持切换，只限`有效`和`无效`才能操作")

	CtxUserID               = "USER_ID"
	CtxStaffID              = "STAFF_ID"
	CtxClientName           = "CLIENT_NAME"
	CtxClientObject         = "CLIENT_OBJECT"
	CtxUserInfo             = "USER_INFO"
	DefaultDateFormat       = "2006-01-02"
	DefaultDateTimeFormat   = "2006-01-02 15:04:05"
	NoLinkDateFormat        = "20060102"
	DefaultNoLinkDateFormat = "20060102150405"
	CachePrefix             = "cache:"
	NoUnitPrefix            = "noUnit:"
	RequestIDKey            = "X-Request-ID" // RequestIDKey 表示请求ID的KEY
	LogGroupIDKey           = "LOG_GROUP_ID" // LogGroupIDKey 代表日志组ID
	CodePrefix              = "V"
	CodeSeq                 = "data:entity:code_seq" // 实体数据编码序列
	IssuedCodeSeq           = "data:issued:code_seq" // 异常数据编码序列
	MergeCachePrefix        = "merge:"

	DefaultPageSize         = uint64(100)
	DefaultSRID             = 4326
	ProjectIntersectionRate = 0.8
	MetaData                = "meta" // 用于表示 category/client 等元数据的权限
	TablePrefix             = strings.ToLower(fmt.Sprintf("%s_", SystemName))
	TableNameClient         = TablePrefix + "client"
	TableNameAccessRight    = TablePrefix + "access_right"
	TableNameIssued         = TablePrefix + "issued"
	TableNameHistory        = TablePrefix + "history"
	TableNameRouter         = TablePrefix + "router"
	TableNameCategory       = TablePrefix + "category"
	TableNameMergeDraft     = TablePrefix + "merge_draft"

	DirectionMap = map[string]string{
		"东":  "东",
		"南":  "南",
		"西":  "西",
		"北":  "北",
		"东南": "东南",
		"东西": "东西",
		"东北": "东北",
		"西南": "西南",
		"南北": "南北",
		"西北": "西北",
		"未知": "未知",
	}

	CityShortWordMap = map[string]string{
		"成都":  "CD",
		"昆明":  "KM",
		"贵阳":  "GY",
		"重庆":  "CQ",
		"佛山":  "FS",
		"武汉":  "WH",
		"郑州":  "ZZ",
		"西安":  "XA",
		"长春":  "CC",
		"哈尔滨": "HB",
		"沈阳":  "SY",
		"大连":  "DL",
		"广州":  "GZ",
		"长沙":  "CS",
		"苏州":  "SU",
		"无锡":  "WX",
		"泉州":  "QZ",
		"厦门":  "XM",
		"福州":  "FZ",
		"上海":  "SH",
		"青岛":  "QD",
		"济南":  "JN",
		"烟台":  "YT",
		"南京":  "NJ",
		"合肥":  "HF",
		"深圳":  "SZ",
		"海南":  "HN",
		"东莞":  "DG",
		"杭州":  "HZ",
		"南昌":  "NC",
		"宁波":  "NB",
		"北京":  "BJ",
		"天津":  "TJ",
		"太原":  "TY"}

	CurrentMergeMap = map[string]string{
		CategoryProject:  "1",
		CategoryBuilding: "2",
		CategoryUnit:     "3",
		CategoryFloor:    "4",
		CategoryHouse:    "5",
	}

	ZhantuDirectionMap = map[string]string{
		"东":  "E",
		"南":  "S",
		"西":  "W",
		"北":  "N",
		"东南": "ES",
		"东西": "EW",
		"东北": "EN",
		"西南": "WS",
		"南北": "SN",
		"西北": "WN",
		"未知": "UN",
		"其他": "OT",
	}

	TaskTypeNameMap = map[string]string{
		TaskTypeNameZhantuProject:      "战图项目导入",
		TaskTypeNameZhantuBuilding:     "战图楼栋导入",
		TaskTypeNameZhantuHouse:        "战图房屋导入",
		TaskTypeNameZhantuCarPort:      "战图车位导入",
		TaskTypeNameZhantuParkingPlace: "战图车场导入",
		TaskTypeNameZhantuFloor:        "战图楼层导入",
	}

	// CarPortTypeMap 车位类型
	CarPortTypeMap = map[string]string{
		"1": "顶层车位", "2": "地面车位", "3": "架空车位", "4": "夹层车位", "5": "地下车位", "6": "立体车位", "7": "人防车位", "8": "其它",
	}
	// CarPortCategoryMap 车位类别
	CarPortCategoryMap = map[string]string{
		"1": "小车车位", "2": "摩托车位", "3": "单车位", "4": "上落货车位", "5": "畅通易达", "0": "其他",
	}

	// ParkingPlaceFeeTypeMap 车场收费类型
	ParkingPlaceFeeTypeMap = map[string]string{
		"1": "免费", "2": "临停收费", "3": "月卡收费", "4": "混合型收费",
	}
	// ParkingPlaceLocationMap 车场位置
	ParkingPlaceLocationMap = map[string]string{
		"1": "地面车场", "2": "地下车场",
	}
	// ParkingPlaceOwnershipTypeMap 车场产权
	ParkingPlaceOwnershipTypeMap = map[string]string{
		"1": "分散产权", "2": "单一产权", "3": "专有产权", "4": "共有产权",
	}
	// ParkingPlaceTypeMap 车场类型
	ParkingPlaceTypeMap = map[string]string{
		"1": "待改造", "2": "住宅配套车场", "3": "商业车场", "4": "混合车场",
	}
	// ZXSCodeMap 直辖市
	ZXSCodeMap = map[string]string{
		"11": "北京市", "12": "天津市", "31": "上海市", "50": "重庆市",
	}
	AsynqMaxRetry = 5
)

type V1CommonSelectorEnum struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type V1CommonDescTypeEnum struct {
	Description string `json:"description"`
	Type        int64  `json:"type"`
}

// IsAffiliatedPropertyEnum 房屋附属房产标志枚举 1.0切换接口使用，1.0那边也是写死去返回附属房产标志列表的
var IsAffiliatedPropertyEnum = []*V1CommonSelectorEnum{
	{Text: "否", Value: "0"},
	{Text: "是", Value: "1"},
}

// CarportTypeEnum 车位类型枚举 1.0切换接口使用，1.0那边也是写死去返回附属房产标志列表的，顶层车位1,地面车位2,架空车位3,夹层车位4,地下车位5,立体车位6,人防车位7,其他8
var CarportTypeEnum = []*V1CommonSelectorEnum{
	{Text: "顶层车位", Value: "1"},
	{Text: "地面车位", Value: "2"},
	{Text: "架空车位", Value: "3"},
	{Text: "夹层车位", Value: "4"},
	{Text: "地下车位", Value: "5"},
	{Text: "立体车位", Value: "6"},
	{Text: "人防车位", Value: "7"},
	{Text: "其它", Value: "8"},
}

// entity 常用的公共字段名，做成常量
const (
	FieldCategory = "category"
	FieldCode     = "code"
	FieldName     = "name"
	FieldVersion  = "version"
	FieldStatus   = "status"
	FieldGeometry = "geometry"
)

type EntityStatus string

const (
	EntityStatusUnknown       EntityStatus = "UNKNOWN"
	EntityStatusInactive      EntityStatus = "INACTIVE"
	EntityStatusInit          EntityStatus = "INIT"
	EntityStatusActive        EntityStatus = "ACTIVE"
	EntityStatusFrozen        EntityStatus = "FROZEN"
	EntityStatusMerged        EntityStatus = "MERGED"
	EntityStatusException     EntityStatus = "EXCEPTION"
	EntityStatusPendingReview EntityStatus = "PENDING_REVIEW"
)

type IssuedStatus string

const (
	IssuedStatusUnknown   IssuedStatus = "UNKNOWN"
	IssuedStatusPending   IssuedStatus = "PENDING"   // 待处理
	IssuedStatusDone      IssuedStatus = "DONE"      // 已处理
	IssuedStatusDiscarded IssuedStatus = "DISCARDED" // 已丢弃
)

type TaskStatus string

const (
	TaskStatusSuccess TaskStatus = "SUCCESS"
	TaskStatusFail    TaskStatus = "FAIL"
)

const (
	SourceZhantuProject     = "战图项目同步"
	SourceUnknown           = "未知"
	SourceZhantu            = "战图"
	SourceYXS               = "翼销售"
	SourceAdmin             = "管理后台"
	SourceSelfHosted        = "自建"
	SourceHouseDictionaryV1 = "房屋字典1.0"
	SourcePulin             = "朴邻"
	SourcePuyu              = "朴寓"
	SourceZhuyingtai        = "助英台"
	SourceSelfDecoration    = "自营装修"
	SourceDaaS              = "DaaS"
	SourceMeiju             = "美居"
	SourceFuWuJia           = "服务家"
	SourcePulinDeal         = "朴邻成交报告"
	SourceVisit             = "客源带看"
	SourceZhiKe             = "知客"
	SourceMedicalReport     = "入户检测报告"
	// 接受许哥意见，对外部引用，统一用 @ 表示，后续对外部扩展查询时，做相应的适配器来统一逻辑
	CodeMappingZhantuKey            = "@zhantu"
	CodeMappingYXSKey               = "@yxs"
	CodeMappingICPKey               = "@icp"
	CodeMappingHouseDictionaryV1Key = "@house_dictionary_v1"
	CodeMappingLibraKey             = "@libra"
	CodeMappingMeijuKey             = "@meiju"
	CodeMappingPuLinDealKey         = "@house_deal_id"
	CodeMappingClientVisitKey       = "@doc_house_id"
	CodeMappingZhikeProductKey      = "@zhike_pro_code"
	CodeMappingPuLinKey             = "@pulin"
	CodeMappingPuYuKey              = "@puyu"
	CodeMappingMedicalReportKey     = "@house_code"

	AttrTypeBasicAttrs  = "basic_attrs"
	AttrTypeExtendAttrs = "extend_attrs"
	AttrTypeCodeMapping = "code_mapping"

	CategoryHouse               = "house"
	CategoryHouseRepairRecord   = "house_repair_record"
	CategoryHouseLayout         = "house_layout"
	CategoryHouseOrderRecord    = "house_order_record"
	CategoryHouseDealRecord     = "house_deal_record"
	CategoryClientVisitRecord   = "client_visit_record"
	CategoryOnewoTown           = "onewo_town"
	CategoryProject             = "project"
	CategoryFile                = "file"
	CategoryVr                  = "vrroom"
	CategoryStore               = "store"
	CategoryBuilding            = "building"
	CategoryPeriod              = "period"
	CategoryStoreContract       = "store_contract"
	CategoryStoreFinancial      = "store_financial"
	CategoryDictionary          = "dictionary"
	CategoryFloor               = "floor"
	CategoryIssued              = "issued"
	CategoryArea                = "area"
	CategoryDistrict            = "district"
	CategoryUnit                = "unit"
	CategoryParkingPlace        = "parking_place"  // 停车场
	CategoryCarPort             = "car_port"       // 车位
	CategoryTask                = "task"           // 车位
	CategoryImportHistory       = "import_history" // 导入记录
	CategoryHouseReceive        = "house_receive"  // 服务家收房记录
	CategoryBusinessDistrict    = "商圈"
	CategoryZhantuImportHistory = "zhantu_import_history"
	CategoryZhikeProduct        = "zhike_product"        // 来自知客的项目
	CategoryWorkArea            = "work_area"            // 作业范围
	CategoryMyHouse             = "my_house"             // 住这儿-我的房
	CategoryHouseMedicalReport  = "house_medical_report" // 入户检测报告

	RoomTypeApartments = "公寓" // 公寓
	RoomTypeFlat       = "平层" // 平层
	RoomTypeDuplex     = "复式" // 复式
	RoomTypeUnknown    = "未知" // 未知

	DirectionUnknown = "未知" // 未知
	DirectionEast    = "东"  // 东
	DirectionSouth   = "南"  // 南
	DirectionWest    = "西"  // 西
	DirectionNorth   = "北"  // 北

	HeaderXScopeProvince = "X-Scope-Province"
	HeaderXScopeCity     = "X-Scope-City"

	TaskTypeMergeDeal    = "merge.house.deal"    // 合并朴邻成交报告
	TaskTypeMergeRepair  = "merge.house.repair"  // 合并维修记录
	TaskTypeMergeReceive = "merge.house.receive" // 合并服务家交房记录
	TaskTypeMergeOrder   = "merge.house.order"   // 合并智装合同记录
	TaskTypeMergeVisit   = "merge.house.visit"   // 合并朴邻带看记录

	TaskTypeMergeProjectToPulin = "merge.project.pulin" // 合并项目推送到朴邻
)

const (
	OrientationE  = "E"
	OrientationS  = "S"
	OrientationW  = "W"
	OrientationN  = "N"
	OrientationES = "ES"
	OrientationEN = "EN"
	OrientationWS = "WS"
	OrientationWN = "WN"
	OrientationSN = "SN"
	OrientationEW = "EW"
	OrientationOT = "OT"
	OrientationUN = "UN"
)

var (
	HouseOrientationMap = map[string]string{
		OrientationE:  "东",
		OrientationS:  "南",
		OrientationW:  "西",
		OrientationN:  "北",
		OrientationES: "东南",
		OrientationEN: "东北",
		OrientationWS: "西南",
		OrientationWN: "西北",
		OrientationSN: "南北",
		OrientationEW: "东西",
		OrientationOT: "其它",
		OrientationUN: "未知",
	}
	// HouseOrientationEnum 房屋朝向枚举 1.0切换接口使用，1.0那边也是写死去返回朝向列表的
	HouseOrientationEnum = []*V1CommonSelectorEnum{
		{Text: HouseOrientationMap[OrientationE], Value: OrientationE},
		{Text: HouseOrientationMap[OrientationS], Value: OrientationS},
		{Text: HouseOrientationMap[OrientationW], Value: OrientationW},
		{Text: HouseOrientationMap[OrientationN], Value: OrientationN},
		{Text: HouseOrientationMap[OrientationES], Value: OrientationES},
		{Text: HouseOrientationMap[OrientationEN], Value: OrientationEN},
		{Text: HouseOrientationMap[OrientationWS], Value: OrientationWS},
		{Text: HouseOrientationMap[OrientationWN], Value: OrientationWN},
		{Text: HouseOrientationMap[OrientationSN], Value: OrientationSN},
		{Text: HouseOrientationMap[OrientationEW], Value: OrientationEW},
		{Text: HouseOrientationMap[OrientationOT], Value: OrientationOT},
		{Text: HouseOrientationMap[OrientationUN], Value: OrientationUN},
	}
)

const (
	PropertyRightType01 = "01"
	PropertyRightType02 = "02"
	PropertyRightType03 = "03"
	PropertyRightType04 = "04"
	PropertyRightType05 = "05"
	PropertyRightType99 = "99"
	PropertyRightType00 = "00"
)

var (
	PropertyRightTypesMap = map[string]string{
		PropertyRightType01: "商品房",
		PropertyRightType02: "集资房",
		PropertyRightType03: "军产房",
		PropertyRightType04: "廉租房",
		PropertyRightType05: "农民房",
		PropertyRightType99: "其他",
		PropertyRightType00: "未知",
	}
	// PropertyRightTypesEnum 房屋产权类型枚举 1.0切换接口使用，1.0那边也是写死去返回产权类型列表的
	PropertyRightTypesEnum = []*V1CommonSelectorEnum{
		{Text: PropertyRightTypesMap[PropertyRightType01], Value: PropertyRightType01},
		{Text: PropertyRightTypesMap[PropertyRightType02], Value: PropertyRightType02},
		{Text: PropertyRightTypesMap[PropertyRightType03], Value: PropertyRightType03},
		{Text: PropertyRightTypesMap[PropertyRightType04], Value: PropertyRightType04},
		{Text: PropertyRightTypesMap[PropertyRightType05], Value: PropertyRightType05},
		{Text: PropertyRightTypesMap[PropertyRightType99], Value: PropertyRightType99},
		{Text: PropertyRightTypesMap[PropertyRightType00], Value: PropertyRightType00},
	}
)

const (
	BuildingTypeHouse   = int64(1)
	BuildingTypeCarport = int64(2)
)

var (
	BuildingTypeMap = map[int64]string{
		BuildingTypeHouse:   "住宅",
		BuildingTypeCarport: "车场",
	}
	// BuildingTypeEnum 楼栋类型列表 1.0切换接口使用，1.0那边也是写死去返回楼栋类型列表的
	BuildingTypeEnum = []*V1CommonDescTypeEnum{
		{Description: BuildingTypeMap[BuildingTypeHouse], Type: BuildingTypeHouse},
		{Description: BuildingTypeMap[BuildingTypeCarport], Type: BuildingTypeCarport},
	}
)

const (
	HouseTypeMatch    = "01" // 配套
	HouseTypeBusiness = "02" // 商业
	HouseTypeHouse    = "03" // 住宅
	HouseTypeOther    = "04" // 其他
	HouseTypeUnknown  = "05" // 未知
)

var (
	// HouseTypeMap 房屋类型枚举 1.0切换接口使用，1.0那边也是写死去返回房屋类型列表的
	HouseTypeMap = map[string]string{
		HouseTypeMatch:    "配套",
		HouseTypeBusiness: "商业",
		HouseTypeHouse:    "住宅",
		HouseTypeOther:    "其他",
		HouseTypeUnknown:  "未知",
	}

	// HouseTypeEnum 房屋类型列表 1.0切换接口使用，1.0那边也是写死去返回房屋类型列表的
	HouseTypeEnum = []*V1CommonSelectorEnum{
		{Text: HouseTypeMap[HouseTypeMatch], Value: HouseTypeMatch},
		{Text: HouseTypeMap[HouseTypeBusiness], Value: HouseTypeBusiness},
		{Text: HouseTypeMap[HouseTypeHouse], Value: HouseTypeHouse},
		{Text: HouseTypeMap[HouseTypeOther], Value: HouseTypeOther},
		{Text: HouseTypeMap[HouseTypeUnknown], Value: HouseTypeUnknown},
	}
)

const (
	ExcelUploadTypeBuilding = 1 // 楼栋
	ExcelUploadTypeProject  = 2 // 项目
)

var (
	// ExcelUploadTypeMap 房屋类型枚举 1.0切换接口使用
	ExcelUploadTypeMap = map[int64]string{
		ExcelUploadTypeBuilding: "楼栋",
		ExcelUploadTypeProject:  "项目",
	}

	// ExcelUploadTypeEnum 房屋类型列表 1.0切换接口使用，1.0那边也是写死去返回房屋类型列表的
	ExcelUploadTypeEnum = []*V1CommonDescTypeEnum{
		{Description: ExcelUploadTypeMap[ExcelUploadTypeBuilding], Type: ExcelUploadTypeBuilding},
		{Description: ExcelUploadTypeMap[ExcelUploadTypeProject], Type: ExcelUploadTypeProject},
	}
)

type AccessType string

const (
	TaskTypeNameZhantuProject      = "zhantu.project.import"
	TaskTypeNameZhantuBuilding     = "zhantu.building.import"
	TaskTypeNameZhantuFloor        = "zhantu.floor.import"
	TaskTypeNameZhantuHouse        = "zhantu.house.import"
	TaskTypeNameZhantuCarPort      = "zhantu.car_port.import"
	TaskTypeNameZhantuParkingPlace = "zhantu.parking_place.import"
)

const (
	AccessTypeUnknown AccessType = ""
	AccessTypeRead    AccessType = "R" // 读
	AccessTypeWrite   AccessType = "W" // 写
)

type ClientStatus string

const (
	ClientStatusInactive ClientStatus = "INACTIVE"
	ClientStatusActive   ClientStatus = "ACTIVE"
)

type SortOrder string // 排序方式

const (
	SortOrderUnknown SortOrder = ""     // 未知
	SortOrderAsc     SortOrder = "ASC"  // 正序
	SortOrderDesc    SortOrder = "DESC" // 逆序
)

type FieldType string

const (
	FieldTypeUnknown FieldType = ""
	FieldTypeString  FieldType = "string"
	FieldTypeNumber  FieldType = "number"
	FieldTypeBoolean FieldType = "boolean"
	FieldTypeArray   FieldType = "array"
	FieldTypeObject  FieldType = "object"
)

type Operator string // 查询的字段运算子

const (
	OperatorUnknown             Operator = ""
	OperatorEqual               Operator = "EQUAL"
	OperatorNotEqual            Operator = "NOT_EQUAL"
	OperatorGreater             Operator = "GREATER"
	OperatorGreaterEqual        Operator = "GREATER_EQUAL"
	OperatorLess                Operator = "LESS"
	OperatorLessEqual           Operator = "LESS_EQUAL"
	OperatorIn                  Operator = "IN"
	OperatorNotIn               Operator = "NOT_IN"
	OperatorMatch               Operator = "MATCH"
	OperatorMatchPrefix         Operator = "MATCH_PREFIX"
	OperatorIsNull              Operator = "IS_NULL"
	OperatorIsNotNull           Operator = "IS_NOT_NULL"
	OperatorGeographyIntersects Operator = "GEOGRAPHY_INTERSECTS" // 地理位置 ST_Intersects，但使用 geometry::geography 进行查询，该方法不能用上 spatial index
	OperatorGeoIntersects       Operator = "GEO_INTERSECTS"       // 地理位置 ST_Intersects，用于 geometry 字段,该方法可以用上 spatial index
	OperatorGeoIntersection     Operator = "GEO_INTERSECTION"     // 地理位置 ST_Intersection，用于 geometry 字段,该方法不能用上 spatial index
	OperatorArrayAllExist       Operator = "ARRAY_ALL_EXIST"      // 用于 JSONB 中的数组(array)类型运算符，表示数组A 必须全部包含数组B 的全部元素
	OperatorArrayAnyExist       Operator = "ARRAY_ANY_EXIST"      // 用于 JSONB 中的数组(array)类型运算符，表示数组A 只要包含数组B 的一个或多个元素
	OperatorGeoContains         Operator = "GEO_CONTAINS"         // 地理位置 ST_Contains，用于 geometry 字段,该方法可以用上 spatial index
	OperatorGeoWithin           Operator = "GEO_WITHIN"           // 地理位置 ST_Within，用于 geometry 字段,该方法可以用上 spatial index
)

type LogicRelation string

const (
	LogicRelationUnknown LogicRelation = ""
	LogicRelationAnd     LogicRelation = "AND"
	LogicRelationOr      LogicRelation = "OR"
)

type UpsertHouseLayoutError string

const (
	UpsertHouseLayoutErrorBasicInfo        = "BASIC_INFO_ERROR"
	UpsertHouseLayoutErrorHouseBinded      = "HOUSE_BINDED"
	UpsertHouseLayoutErrorImageAICheckFail = "IMAGE_AI_CHECK_FAIL"
)

const (
	ResponseBizCodeWrongHouseLayoutInfo = 500001 // 房屋已经关联其他的户型
	ResponseBizCodeHouseGotLayout       = 500002 // 房屋已经关联其他的户型
)

type BusinessStatus string

const (
	BusinessStatusUnknown  BusinessStatus = "UNKNOWN"
	BusinessStatusInactive BusinessStatus = "INACTIVE"
	BusinessStatusInit     BusinessStatus = "INIT"
	BusinessStatusActive   BusinessStatus = "ACTIVE"
)

type BusinessScope string

const (
	BusinessScopeUnknown BusinessScope = "UNKNOWN"
	BusinessScopePulin   BusinessScope = "pulin"
	BusinessScopeYanxuan BusinessScope = "yanxuan"
	BusinessScopeAll     BusinessScope = "all"
)

type MergeStatus string

const (
	MergeStatusUnknown MergeStatus = "UNKNOWN"
	MergeStatusShow    MergeStatus = "show"
	MergeStatusHidden  MergeStatus = "hidden"
	MergeStatusMerging MergeStatus = "merging"
)

const (
	MQ_TAG_PROJECT_MERAGE = "projectM"
)

const (
	DefaultDistance int64 = 1000
)
