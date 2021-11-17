package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getAllUiItems_back(t *testing.T) {
	tree, err := getUiPermission_back("test")
	a := assert.New(t)
	a.Nil(err)
	a.Equal(2, tree.hasChilds())
	a.Equal(2, tree.ChildNodes[0].hasChilds())
	js, err := tree.toJson()
	a.Nil(err)
	t.Log(js)
}
