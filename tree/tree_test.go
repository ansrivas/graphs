package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TreeInsert(t *testing.T) {
	assert := assert.New(t)
	bst := NewBST()

	//Inorder on empty tree should fail
	err := bst.Inorder()
	assert.NotNil(err, "Inorder call on empty tree should NOT be nil")

	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(3)
	bst.Insert(5)

	bst.Inorder()

	assert.True(true, "Starting")
}

func Test_Diameter(t *testing.T) {
	assert := assert.New(t)
	bst := NewBST()

	//Diameter and Height on empty tree should fail.
	_, err := bst.Height()
	assert.NotNil(err, "Height call on empty tree should NOT be nil")
	_, err = bst.Diameter()
	assert.NotNil(err, "Diameter call on empty tree should NOT be nil")

	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(4)
	bst.Insert(5)
	actual, _ := bst.Height()
	assert.Equal(5, actual, "Height of the tree should be 5")
	actual, _ = bst.Diameter()
	assert.Equal(5, actual, "Diameter of the tree should be 5")

	bst = NewBST()
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(5)
	bst.Insert(9)
	bst.Insert(10)
	bst.Insert(11)
	actual, _ = bst.Height()
	assert.Equal(5, actual, "Height of the tree should be 5")
	actual, _ = bst.Diameter()
	assert.Equal(7, actual, "Diameter of the tree should be 4")
}

func Test_Size(t *testing.T) {
	bst := NewBST()
	_, err := bst.Size()
	assert.NotNil(t, err, "Size on empty tree should return nil")

	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(8)
	bst.Insert(19)
	bst.Insert(1)

	size, _ := bst.Size()
	assert.Equal(t, size, 5, "size of binary tree should be 5")

}

// func Test_RootToLeafPath(t *testing.T) {
// 	assert := assert.New(t)
// 	bst := NewBST()
//
// 	bst.Insert(4)
// 	bst.Insert(2)
// 	bst.Insert(8)
// 	bst.Insert(19)
// 	bst.Insert(1)
// 	bst.Insert(6)
//
// 	_, exists := bst.RootToLeafSum(18)
// 	assert.True(exists, "Path containing 18 exists")
//
// 	if assert.True(exists, "Path containing 18 exists") {
// 		expectedPath := []int{6, 8, 4}
// 		assert.Equal(expectedPath, path, "Path should be accurate")
// 	}
// }
