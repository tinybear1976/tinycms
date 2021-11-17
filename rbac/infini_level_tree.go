package rbac

// import (
// 	"encoding/json"
// )

// type treeNode struct {
// 	Id         int         `json:"id"`
// 	ParentId   int         `json:"-"`
// 	ParenNode  *treeNode   `json:"-"`
// 	Level      int         `json:"-"`
// 	Name       string      `json:"name"`
// 	ChildNodes []*treeNode `json:"childs,omitempty"`
// 	Data       interface{} `json:"-"`
// }

// func createRootNode(id int, name string, data interface{}) *treeNode {
// 	n := treeNode{
// 		Id:         id,
// 		ParentId:   -1,
// 		ParenNode:  nil,
// 		Level:      1,
// 		Name:       name,
// 		Data:       data,
// 		ChildNodes: make([]*treeNode, 0),
// 	}
// 	return &n
// }

// // 建议从root节点调用该方法，通过分治算法自动挂接节点
// func (n *treeNode) createChildNode(id int, name string, parent_id int, data interface{}) bool {
// 	if n.Id == parent_id {
// 		child := treeNode{
// 			Id:         id,
// 			ParentId:   parent_id,
// 			ParenNode:  n,
// 			Level:      n.Level + 1,
// 			Name:       name,
// 			Data:       data,
// 			ChildNodes: make([]*treeNode, 0),
// 		}
// 		n.ChildNodes = append(n.ChildNodes, &child)
// 		return true
// 	}
// 	ok := false
// 	for _, node := range n.ChildNodes {
// 		ok = node.createChildNode(id, name, parent_id, data)
// 		if ok {
// 			break
// 		}
// 	}
// 	return ok
// }

// func (node *treeNode) AddChildNode(id int, name string, data interface{}) {
// 	child := treeNode{
// 		Id:         id,
// 		ParentId:   node.Id,
// 		ParenNode:  node,
// 		Level:      node.Level + 1,
// 		Name:       name,
// 		Data:       data,
// 		ChildNodes: make([]*treeNode, 0),
// 	}
// 	node.ChildNodes = append(node.ChildNodes, &child)
// }

// // 节点下是否包含子节点，返回子节点数量
// func (node *treeNode) HasChilds() int {
// 	return len(node.ChildNodes)
// }

// func (node *treeNode) ToJson() (string, error) {
// 	bytes, err := json.Marshal(node)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(bytes), nil
// }

// //===== rbac =======================================================================
