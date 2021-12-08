package models

import "time"

type Ammo struct {
	Caliber     string `json:"caliber"`
	Name        string `json:"name"`
	ShortName   string `json:"shortName"`
	Damage      int    `json:"damage"`
	Penetration int    `json:"penetration"`
	Price       int    `json:"price"`
}

// ------- tarkov-tools graphQL models -------
type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ShortName   string `json:"shortName"`
	IconLink    string `json:"iconLink"`
	Avg24hPrice int    `json:avg24hPrice`
}

type Data struct {
	ItemsByType []Item `json:"itemsByType"`
}

type GraphQLResponse struct {
	Data Data `json:"data"`
}

// ------- tarkov-market models -------
type TarkovMarketItem struct {
	UID            string    `json:"uid"`
	Name           string    `json:"name"`
	ShortName      string    `json:"shortName"`
	Price          int       `json:"price"`
	BasePrice      int       `json:"basePrice"`
	Avg24HPrice    int       `json:"avg24hPrice"`
	Avg7DaysPrice  int       `json:"avg7daysPrice"`
	TraderName     string    `json:"traderName"`
	TraderPrice    int       `json:"traderPrice"`
	TraderPriceCur string    `json:"traderPriceCur"`
	Updated        time.Time `json:"updated"`
	Slots          int       `json:"slots"`
	Diff24H        float64   `json:"diff24h"`
	Diff7Days      float64   `json:"diff7days"`
	Icon           string    `json:"icon"`
	Link           string    `json:"link"`
	WikiLink       string    `json:"wikiLink"`
	Img            string    `json:"img"`
	ImgBig         string    `json:"imgBig"`
	BsgID          string    `json:"bsgId"`
	IsFunctional   bool      `json:"isFunctional"`
	Reference      string    `json:"reference"`
}

type TarkovMarketItems []TarkovMarketItem

// ------- BSG models -------
type Prefab struct {
	Path string `json:"path"`
	Rcid string `json:"rcid"`
}

type UsePrefab struct {
	Path string `json:"path"`
	Rcid string `json:"rcid"`
}

type Contusion struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type ArmorDistanceDistanceDamage struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Blindness struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Props struct {
	Name                                   string                      `json:"Name"`
	ShortName                              string                      `json:"ShortName"`
	Description                            string                      `json:"Description"`
	Weight                                 float64                     `json:"Weight"`
	BackgroundColor                        string                      `json:"BackgroundColor"`
	Width                                  int                         `json:"Width"`
	Height                                 int                         `json:"Height"`
	StackMaxSize                           int                         `json:"StackMaxSize"`
	Rarity                                 string                      `json:"Rarity"`
	SpawnChance                            int                         `json:"SpawnChance"`
	CreditsPrice                           int                         `json:"CreditsPrice"`
	ItemSound                              string                      `json:"ItemSound"`
	Prefab                                 Prefab                      `json:"Prefab"`
	UsePrefab                              UsePrefab                   `json:"UsePrefab"`
	StackObjectsCount                      int                         `json:"StackObjectsCount"`
	NotShownInSlot                         bool                        `json:"NotShownInSlot"`
	ExaminedByDefault                      bool                        `json:"ExaminedByDefault"`
	ExamineTime                            int                         `json:"ExamineTime"`
	IsUndiscardable                        bool                        `json:"IsUndiscardable"`
	IsUnsaleable                           bool                        `json:"IsUnsaleable"`
	IsUnbuyable                            bool                        `json:"IsUnbuyable"`
	IsUngivable                            bool                        `json:"IsUngivable"`
	IsLockedafterEquip                     bool                        `json:"IsLockedafterEquip"`
	QuestItem                              bool                        `json:"QuestItem"`
	LootExperience                         int                         `json:"LootExperience"`
	ExamineExperience                      int                         `json:"ExamineExperience"`
	HideEntrails                           bool                        `json:"HideEntrails"`
	RepairCost                             int                         `json:"RepairCost"`
	RepairSpeed                            int                         `json:"RepairSpeed"`
	ExtraSizeLeft                          int                         `json:"ExtraSizeLeft"`
	ExtraSizeRight                         int                         `json:"ExtraSizeRight"`
	ExtraSizeUp                            int                         `json:"ExtraSizeUp"`
	ExtraSizeDown                          int                         `json:"ExtraSizeDown"`
	ExtraSizeForceAdd                      bool                        `json:"ExtraSizeForceAdd"`
	MergesWithChildren                     bool                        `json:"MergesWithChildren"`
	CanSellOnRagfair                       bool                        `json:"CanSellOnRagfair"`
	CanRequireOnRagfair                    bool                        `json:"CanRequireOnRagfair"`
	ConflictingItems                       []interface{}               `json:"ConflictingItems"`
	FixedPrice                             bool                        `json:"FixedPrice"`
	Unlootable                             bool                        `json:"Unlootable"`
	UnlootableFromSlot                     string                      `json:"UnlootableFromSlot"`
	UnlootableFromSide                     []interface{}               `json:"UnlootableFromSide"`
	ChangePriceCoef                        int                         `json:"ChangePriceCoef"`
	AllowSpawnOnLocations                  []interface{}               `json:"AllowSpawnOnLocations"`
	SendToClient                           bool                        `json:"SendToClient"`
	AnimationVariantsNumber                int                         `json:"AnimationVariantsNumber"`
	DiscardingBlock                        bool                        `json:"DiscardingBlock"`
	RagFairCommissionModifier              int                         `json:"RagFairCommissionModifier"`
	IsAlwaysAvailableForInsurance          bool                        `json:"IsAlwaysAvailableForInsurance"`
	StackMinRandom                         int                         `json:"StackMinRandom"`
	StackMaxRandom                         int                         `json:"StackMaxRandom"`
	AmmoType                               string                      `json:"ammoType"`
	Damage                                 int                         `json:"Damage"`
	AmmoAccr                               int                         `json:"ammoAccr"`
	AmmoRec                                int                         `json:"ammoRec"`
	AmmoDist                               int                         `json:"ammoDist"`
	BuckshotBullets                        int                         `json:"buckshotBullets"`
	PenetrationPower                       int                         `json:"PenetrationPower"`
	PenetrationPowerDiviation              int                         `json:"PenetrationPowerDiviation"`
	AmmoHear                               int                         `json:"ammoHear"`
	AmmoSfx                                string                      `json:"ammoSfx"`
	MisfireChance                          float64                     `json:"MisfireChance"`
	MinFragmentsCount                      int                         `json:"MinFragmentsCount"`
	MaxFragmentsCount                      int                         `json:"MaxFragmentsCount"`
	AmmoShiftChance                        int                         `json:"ammoShiftChance"`
	CasingName                             string                      `json:"casingName"`
	CasingEjectPower                       int                         `json:"casingEjectPower"`
	CasingMass                             float64                     `json:"casingMass"`
	CasingSounds                           string                      `json:"casingSounds"`
	ProjectileCount                        int                         `json:"ProjectileCount"`
	InitialSpeed                           int                         `json:"InitialSpeed"`
	PenetrationChance                      float64                     `json:"PenetrationChance"`
	RicochetChance                         float64                     `json:"RicochetChance"`
	FragmentationChance                    float64                     `json:"FragmentationChance"`
	BallisticCoeficient                    int                         `json:"BallisticCoeficient"`
	Deterioration                          int                         `json:"Deterioration"`
	SpeedRetardation                       float64                     `json:"SpeedRetardation"`
	Tracer                                 bool                        `json:"Tracer"`
	TracerColor                            string                      `json:"TracerColor"`
	TracerDistance                         int                         `json:"TracerDistance"`
	ArmorDamage                            int                         `json:"ArmorDamage"`
	Caliber                                string                      `json:"Caliber"`
	StaminaBurnPerDamage                   float64                     `json:"StaminaBurnPerDamage"`
	HeavyBleedingDelta                     int                         `json:"HeavyBleedingDelta"`
	LightBleedingDelta                     int                         `json:"LightBleedingDelta"`
	ShowBullet                             bool                        `json:"ShowBullet"`
	HasGrenaderComponent                   bool                        `json:"HasGrenaderComponent"`
	FuzeArmTimeSec                         int                         `json:"FuzeArmTimeSec"`
	ExplosionStrength                      int                         `json:"ExplosionStrength"`
	MinExplosionDistance                   int                         `json:"MinExplosionDistance"`
	MaxExplosionDistance                   int                         `json:"MaxExplosionDistance"`
	FragmentsCount                         int                         `json:"FragmentsCount"`
	FragmentType                           string                      `json:"FragmentType"`
	ShowHitEffectOnExplode                 bool                        `json:"ShowHitEffectOnExplode"`
	ExplosionType                          string                      `json:"ExplosionType"`
	AmmoLifeTimeSec                        int                         `json:"AmmoLifeTimeSec"`
	Contusion                              Contusion                   `json:"Contusion"`
	ArmorDistanceDistanceDamage            ArmorDistanceDistanceDamage `json:"ArmorDistanceDistanceDamage"`
	Blindness                              Blindness                   `json:"Blindness"`
	IsLightAndSoundShot                    bool                        `json:"IsLightAndSoundShot"`
	LightAndSoundShotAngle                 int                         `json:"LightAndSoundShotAngle"`
	LightAndSoundShotSelfContusionTime     int                         `json:"LightAndSoundShotSelfContusionTime"`
	LightAndSoundShotSelfContusionStrength int                         `json:"LightAndSoundShotSelfContusionStrength"`
}

type BSGItem struct {
	ID     string `mapstructure:"_id"`
	Name   string `mapstructure:"_name"`
	Parent string `mapstructure:"_parent"`
	Type   string `mapstructure:"_type"`
	Props  Props  `mapstructure:"_props"`
	Proto  string `mapstructure:"_proto"`
}

// ------- TarkovTracker/tarkovdata ammo models -------

type TarkovTrackerAmmo struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	ShortName       string     `json:"shortName"`
	Weight          float64    `json:"weight"`
	Caliber         string     `json:"caliber"`
	StackMaxSize    int        `json:"stackMaxSize"`
	Tracer          bool       `json:"tracer"`
	TracerColor     string     `json:"tracerColor"`
	AmmoType        string     `json:"ammoType"`
	ProjectileCount int        `json:"projectileCount"`
	Ballistics      Ballistics `json:"ballistics"`
}

type Ballistics struct {
	Damage              int     `json:"damage"`
	ArmorDamage         int     `json:"armorDamage"`
	FragmentationChance float64 `json:"fragmentationChance"`
	RicochetChance      float64 `json:"ricochetChance"`
	PenetrationChance   float64 `json:"penetrationChance"`
	PenetrationPower    int     `json:"penetrationPower"`
	Accuracy            int     `json:"accuracy"`
	Recoil              int     `json:"recoil"`
	InitialSpeed        int     `json:"initialSpeed"`
}

// Configuration to be filled by envconfig
type Config struct {
	// JSONBIN_BIN_ID  string
	// JSONBIN_API_KEY string

	MONGO_USER         string
	MONGO_PASSWORD     string
	MONGO_CLUSTER_PATH string
	MONGO_DB_NAME      string

	TM_API_KEY string

	TC_API_KEY string

	VERCEL_ENV string
}

const APIKeyHeader = "X-Tarkov-Charts-API-Key"
