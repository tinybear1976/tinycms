package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getAllUiItems(t *testing.T) {
	tree, err := getUiPermission("test")
	a := assert.New(t)
	a.Nil(err)
	a.Equal(2, tree.HasChilds())
	a.Equal(2, tree.ChildNodes[0].HasChilds())
	js, err := tree.ToJson()
	a.Nil(err)
	t.Log(js)
}
