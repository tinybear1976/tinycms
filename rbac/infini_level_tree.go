package rbac

import (
	"encoding/json"
	"strings"
)

type treeNode struct {
	Id         int         `json:"id"`
	ParentId   int         `json:"-"`
	ParenNode  *treeNode   `json:"-"`
	Level      int         `json:"-"`
	Name       string      `json:"name"`
	ChildNodes []*treeNode `json:"childs,omitempty"`
	Data       interface{} `json:"-"`
}

func createRootNode(id int, name string, data interface{}) *treeNode {
	n := treeNode{
		Id:         id,
		ParentId:   -1,
		ParenNode:  nil,
		Level:      1,
		Name:       name,
		Data:       data,
		ChildNodes: make([]*treeNode, 0),
	}
	return &n
}

// 建议从root节点调用该方法，通过分治算法自动挂接节点
func (n *treeNode) createChildNode(id int, name string, parent_id int, data interface{}) bool {
	if n.Id == parent_id {
		child := treeNode{
			Id:         id,
			ParentId:   parent_id,
			ParenNode:  n,
			Level:      n.Level + 1,
			Name:       name,
			Data:       data,
			ChildNodes: make([]*treeNode, 0),
		}
		n.ChildNodes = append(n.ChildNodes, &child)
		return true
	}
	ok := false
	for _, node := range n.ChildNodes {
		ok = node.createChildNode(id, name, parent_id, data)
		if ok {
			break
		}
	}
	return ok
}

func (node *treeNode) AddChildNode(id int, name string, data interface{}) {
	child := treeNode{
		Id:         id,
		ParentId:   node.Id,
		ParenNode:  node,
		Level:      node.Level + 1,
		Name:       name,
		Data:       data,
		ChildNodes: make([]*treeNode, 0),
	}
	node.ChildNodes = append(node.ChildNodes, &child)
}

// 节点下是否包含子节点，返回子节点数量
func (node *treeNode) HasChilds() int {
	return len(node.ChildNodes)
}

func (node *treeNode) ToJson() (string, error) {
	bytes, err := json.Marshal(node)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//===== rbac =======================================================================

type rbac_TreeNode struct {
	Role_Id     string           `json:"role_id"`
	Id          int              `json:"ui_id"`
	Key         string           `json:"ui_key"`
	UiType      string           `json:"ui_type"`
	Description string           `json:"description"`
	Parent_id   int              `json:"parent_ui_id"`
	IsAllow     bool             `json:"isallow"`
	ChildNodes  []*rbac_TreeNode `json:"childs,omitempty"`
}

func createRootNode_rbac(pdata *ui_Item) *rbac_TreeNode {
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
func (n *rbac_TreeNode) createChildNode_rbac(pdata *ui_Item) bool {
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

// 节点下是否包含子节点，返回子节点数量
func (node *rbac_TreeNode) HasChilds() int {
	return len(node.ChildNodes)
}

func (node *rbac_TreeNode) ToJson() (string, error) {
	bytes, err := json.Marshal(node)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (node *rbac_TreeNode) ToFrontJson() (string, error) {
	if node == nil {
		return "", nil
	}
	var sb strings.Builder
	sb.WriteString("{")

	each_nodes(&sb, node)

	sb.WriteString("}")
	j := cleanFormatter(sb.String())
	return j, nil
}

func each_nodes(sb *strings.Builder, node *rbac_TreeNode) {
	sb.WriteString("\"" + node.Key + "\"")
	sb.WriteString(":")
	if node.IsAllow {
		sb.WriteString("1")
	} else {
		sb.WriteString("0")
	}
	sb.WriteString("{")
	for _, n := range node.ChildNodes {
		each_nodes(sb, n)
	}
	sb.WriteString("},")
}

func cleanFormatter(s string) string {
	s = strings.ReplaceAll(s, ",}", "}")
	s = strings.ReplaceAll(s, "1{}", "1")
	s = strings.ReplaceAll(s, "0{}", "0")
	s = strings.ReplaceAll(s, "1{\"", "{\"")
	s = strings.ReplaceAll(s, "0{\"", "{\"")
	return s
}
