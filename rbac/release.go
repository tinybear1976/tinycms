package rbac

// 按角色名称返回给前端（全量）权限，如果权限表中没有找到对应记录（异常），则返回所有模板中的数据，并全体置为拒绝状态
// 返回json字符串
func GetUiPermission(role_id string) (string, error) {
	ptree, err := getUiPermission(role_id)
	if err != nil {
		return "", err
	}
	return ptree.ToJson()
}
