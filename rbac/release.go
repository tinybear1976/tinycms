package rbac

import "encoding/json"

// 按角色名称返回给前端（全量）权限，如果权限表中没有找到对应记录（异常），则返回所有模板中的数据，并全体置为拒绝状态
// 返回json字符串
func GetUiPermission_back(role_id string) (string, error) {
	ptree, err := getUiPermission_back(role_id)
	if err != nil {
		return "", err
	}
	return ptree.toJson()
}

func GetUiPermission_front(role_id string) (string, error) {
	ptree, err := getUiPermission_front(role_id)
	if err != nil {
		return "", err
	}
	return ptree.toFrontJson()
}

func SaveUiPermission_back(j string) error {
	var items rbac_TreeNode
	b := []byte(j)
	err := json.Unmarshal(b, &items)
	if err != nil {
		return err
	}
	err = saveUiPermission_back(&items)
	return err
}
