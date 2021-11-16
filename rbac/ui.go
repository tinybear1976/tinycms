package rbac

import (
	mariadb "github.com/tinybear1976/database-mariadb"
	"github.com/tinybear1976/localsystem/logger"
	"github.com/tinybear1976/tinycms/debugging"
	"github.com/tinybear1976/tinycms/defines"
)

func _init_test_() {
	err := mariadb.New(
		defines.DB_MAIN,
		"172.16.1.250",
		"3306",
		"root",
		"123",
		"tinycms")
	if err != nil {
		logger.Log.Panic("init db for tinycms faile." + err.Error())
		panic("init db for tinycms faile " + err.Error())
	}
}

// 按角色名称返回给前端（全量）权限，如果权限表中没有找到对应记录（异常），则返回所有模板中的数据，并全体置为拒绝状态
func getUiPermission(role_id string) (*rbac_TreeNode, error) {
	_init_test_()
	conn, err := mariadb.Connect(defines.DB_MAIN)
	if err != nil {
		return nil, err
	}
	items := make([]ui_Item, 0)
	sql_1 := "select * from rbac_role_ui_permission order by parent_ui_id,ui_id;"
	sql_2 := "select '" + role_id + "' as role_id,0 as isallow, rbac_ui_temp.* from rbac_ui_temp order by parent_ui_id,ui_id;"
	debugging.Debug_ShowSql("sql_1", sql_1)
	err = conn.Select(&items, sql_1)
	if err != nil {
		return nil, err
	}
	var ptree *rbac_TreeNode
	if len(items) > 0 {
		// 生成数据并返回
		ptree = makePermissionTree(&items)
	} else {
		// 读模板，并置拒绝状态
		debugging.Debug_ShowSql("sql_2", sql_2)
		err = conn.Select(&items, sql_2)
		if err != nil {
			return nil, err
		}
		ptree = makePermissionTree(&items)
	}
	return ptree, nil
}

func makePermissionTree(pitems *[]ui_Item) *rbac_TreeNode {
	var rootNode *rbac_TreeNode
	for i := 0; i < len(*pitems); i++ {
		if (*pitems)[i].Parent_id == -1 {
			rootNode = createRootNode_rbac(&((*pitems)[i]))
		} else {
			if rootNode != nil {
				rootNode.createChildNode_rbac(&((*pitems)[i]))
			}
		}
	}
	return rootNode
}
