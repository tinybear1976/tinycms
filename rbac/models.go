package rbac

import (
	"encoding/json"
	"strings"

	"github.com/jmoiron/sqlx/types"
)

type ui_DB_Item struct {
	Role_Id     string        `db:"role_id"`
	Id          int           `db:"ui_id"`
	Key         string        `db:"ui_key"`
	UiType      string        `db:"ui_type"`
	Description string        `db:"description"`
	Parent_id   int           `db:"parent_ui_id"`
	IsAllow     types.BitBool `db:"isallow"`
}

type rbac_TreeNode struct {
	Role_Id     string           `json:"role_id"`
	Id          int              `json:"ui_id"`
	Key         string           `json:"ui_key"`
	UiType      string           `json:"ui_type"`
	Description string           `json:"label"`
	Parent_id   int              `json:"parent_ui_id"`
	IsAllow     bool             `json:"isallow"`
	ChildNodes  []*rbac_TreeNode `json:"children,omitempty"`
}

// 节点下是否包含子节点，返回子节点数量
func (node *rbac_TreeNode) hasChilds() int {
	return len(node.ChildNodes)
}

func (node *rbac_TreeNode) toJson() (string, error) {
	bytes, err := json.Marshal(node)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (node *rbac_TreeNode) toFrontJson() (string, error) {
	if node == nil {
		return "", nil
	}
	var sb strings.Builder
	sb.WriteString("{")

	each_nodes_toString(&sb, node)

	sb.WriteString("}")
	j := cleanFormatter(sb.String())
	return j, nil
}

func (node *rbac_TreeNode) unTree() (*[]ui_DB_Item, error) {
	rst_items := make([]ui_DB_Item, 0)
	if node == nil {
		return &rst_items, nil
	}
	each_nodes_toItems(&rst_items, node)
	//fmt.Println(rst_items)
	return &rst_items, nil
}

func each_nodes_toString(sb *strings.Builder, node *rbac_TreeNode) {
	sb.WriteString("\"" + node.Key + "\"")
	sb.WriteString(":")
	if node.IsAllow {
		sb.WriteString("1")
	} else {
		sb.WriteString("0")
	}
	sb.WriteString("{")
	for _, n := range node.ChildNodes {
		each_nodes_toString(sb, n)
	}
	sb.WriteString("},")
}

func each_nodes_toItems(items *[]ui_DB_Item, node *rbac_TreeNode) {
	item := ui_DB_Item{
		Role_Id:     node.Role_Id,
		Id:          node.Id,
		Key:         node.Key,
		UiType:      node.UiType,
		Description: node.Description,
		Parent_id:   node.Parent_id,
		IsAllow:     types.BitBool(node.IsAllow),
	}
	*items = append(*items, item)
	for _, n := range node.ChildNodes {
		each_nodes_toItems(items, n)
	}
}

func cleanFormatter(s string) string {
	s = strings.ReplaceAll(s, ",}", "}")
	s = strings.ReplaceAll(s, "1{}", "1")
	s = strings.ReplaceAll(s, "0{}", "0")
	s = strings.ReplaceAll(s, "1{\"", "{\"")
	s = strings.ReplaceAll(s, "0{\"", "{\"")
	return s
}

func createRootNode_rbac(pdata *ui_DB_Item) *rbac_TreeNode {
	n := rbac_TreeNode{
		Role_Id:     pdata.Role_Id,
		Id:          pdata.Id,
		Key:         pdata.Key,
		UiType:      pdata.UiType,
		Description: pdata.Description,
		Parent_id:   pdata.Parent_id,
		IsAllow:     bool(pdata.IsAllow),
		ChildNodes:  make([]*rbac_TreeNode, 0),
	}
	return &n
}

// 建议从root节点调用该方法，通过分治算法自动挂接节点
func (n *rbac_TreeNode) createChildNode_rbac(pdata *ui_DB_Item) bool {
	if n.Id == pdata.Parent_id {
		child := rbac_TreeNode{
			Role_Id:     pdata.Role_Id,
			Id:          pdata.Id,
			Key:         pdata.Key,
			UiType:      pdata.UiType,
			Description: pdata.Description,
			Parent_id:   pdata.Parent_id,
			IsAllow:     bool(pdata.IsAllow),
			ChildNodes:  make([]*rbac_TreeNode, 0),
		}
		n.ChildNodes = append(n.ChildNodes, &child)
		return true
	}
	ok := false
	for _, node := range n.ChildNodes {
		ok = node.createChildNode_rbac(pdata)
		if ok {
			break
		}
	}
	return ok
}
