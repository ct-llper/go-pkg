package cnf

var (
	// SexMap 性别 0：未知；1：男；2：女
	SexMap map[int64]string = map[int64]string{
		0: "未知",
		1: "男",
		2: "女",
	}

	// BloodTypeMap 血型 (0:未知 1:A 2:B 3:O 4:AB)
	BloodTypeMap map[int64]string = map[int64]string{
		0: "未知",
		1: "A",
		2: "B",
		3: "O",
		4: "AB",
	}

	// PregnantStatusMap 怀孕状态 (0:未知 1:是  2:否)
	PregnantStatusMap map[int64]string = map[int64]string{
		0: "未知",
		1: "是",
		2: "否",
	}

	// MaritalStatusMap 婚姻状态（0:未知 1:已婚，2:未婚 ）
	MaritalStatusMap map[int64]string = map[int64]string{
		0: "未知",
		1: "已婚",
		2: "未婚",
		3: "丧偶",
		4: "离婚",
		9: "其他",
	}

	// FertilityStatusMap 生育状态（0:未知 1:已育  2:未育）
	FertilityStatusMap map[int64]string = map[int64]string{
		0: "未知",
		1: "已育",
		2: "未育",
	}

	// RelationShipMap 用户关系
	RelationShipMap map[int64]string = map[int64]string{
		1: "本人",
		2: "父母",
		3: "配偶",
		4: "亲戚",
		5: "朋友",
		6: "其他",
		7: "兄弟姐妹",
		8: "子女",
	}

	// IdTypeMap 证件类型
	IdTypeMap map[int64]string = map[int64]string{
		1: "居民身份证",
		2: "中国人民解放军军人证",
		3: "中国人民武装警察身份证",
		4: "港澳居民来往大陆通行证",
		5: "台湾居民来往大陆通行证",
		6: "护照",
		7: "机动车驾驶证",
		8: "医保卡",
	}
)
