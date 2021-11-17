package rbac

import (
	"fmt"

	mariadb "github.com/tinybear1976/database-mariadb"
	"github.com/tinybear1976/tinycms/debugging"
	"github.com/tinybear1976/tinycms/defines"
)

// func _init_test_() {
// 	err := mariadb.New(
// 		defines.DB_MAIN,
// 		"172.16.1.250",
// 		"3306",
// 		"root",
// 		"123",
// 		"tinycms")
// 	if err != nil {
// 		logger.Log.Panic("init db for tinycms faile." + err.Error())
// 		panic("init db for tinycms faile " + err.Error())
// 	}
// }

// 按角色名称返回给前端（全量）权限，如果权限表中没有找到对应记录（异常），则返回所有模板中的数据，并全体置为拒绝状态
func getUiPermission_back(role_id string) (*rbac_TreeNode, error) {
	// _init_test_()
	conn, err := mariadb.Connect(defines.DB_MAIN)
	if err != nil {
		return nil, err
	}
	items := make([]ui_DB_Item, 0)
	sql_1 := "SELECT * FROM rbac_role_ui_permission ORDER BY parent_ui_id,ui_id;"
	sql_2 := "SELECT '" + role_id + "' AS role_id,0 AS isallow, rbac_ui_temp.* FROM rbac_ui_temp ORDER BY parent_ui_id,ui_id;"
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

func makePermissionTree(pitems *[]ui_DB_Item) *rbac_TreeNode {
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

// 前台需要返回权限，需要转换动态key的json格式{a:{a}}
func getUiPermission_front(role_id string) (*rbac_TreeNode, error) {
	// _init_test_()
	conn, err := mariadb.Connect(defines.DB_MAIN)
	if err != nil {
		return nil, err
	}
	items := make([]ui_DB_Item, 0)
	sql := "SELECT * FROM rbac_role_ui_permission WHERE isallow=1 ORDER BY parent_ui_id,ui_id;"
	debugging.Debug_ShowSql("sql", sql)
	err = conn.Select(&items, sql)
	if err != nil {
		return nil, err
	}
	var ptree *rbac_TreeNode
	if len(items) > 0 {
		// 生成数据并返回
		ptree = makePermissionTree(&items)
	}
	return ptree, nil
}

func saveUiPermission_back(pNode *rbac_TreeNode) error {
	items, err := pNode.unTree()
	if err != nil {
		return err
	}
	conn, err := mariadb.Connect(defines.DB_MAIN)
	if err != nil {
		return err
	}
	//fmt.Println(items)
	tx, err := conn.Beginx()
	if err != nil {
		return err
	}
	sql_del := "DELETE FROM rbac_role_ui_permission WHERE role_id='" + pNode.Role_Id + "';"
	debugging.Debug_ShowSql("delete role permission", sql_del)
	_, err = tx.Exec(sql_del)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, item := range *items {
		sql_ins := fmt.Sprintf("INSERT INTO rbac_role_ui_permission (role_id, ui_id, ui_key, ui_type, description, parent_ui_id, isallow) VALUES ('%s', %d, '%s', '%s', '%s', %d, %v);",
			item.Role_Id, item.Id, item.Key, item.UiType, item.Description, item.Parent_id, item.IsAllow)
		debugging.Debug_ShowSql("insert role permission", sql_ins)
		_, err = tx.Exec(sql_ins)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return err
}
