package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	search := func(nums []int, tar int) int {
		for i, v := range nums {
			if v == tar {
				return i
			}
		}

		return -1
	}

	if len(preorder) == 0 {
		return nil
	}

	// 前序数组确定左子树边界元素索引：必须通过中序数组rootVal左侧的元素个数确定
	rootVal := preorder[0]
	idxRoot := search(inorder, rootVal)

	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(preorder[1:idxRoot+1], inorder[:idxRoot])
	root.Right = buildTree(preorder[idxRoot+1:], inorder[idxRoot+1:])

	return root
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	idxRoot := search(inorder, rootVal)

	root := &TreeNode{Val: rootVal}
	root.Left = buildTree2(inorder[:idxRoot], postorder[:idxRoot])
	root.Right = buildTree2(inorder[idxRoot+1:], postorder[idxRoot:len(postorder)-1])

	return root
}

func buildTree3(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	rootVal := preorder[0]
	// 建树答案不唯一，因为这里我们直接假设了左子树根节点是存在的，但实际可能为空节点
	idxLeft := search(postorder, preorder[1])

	root := &TreeNode{Val: rootVal}
	// idxLeft+2是因为除了包含leftRoot节点外，preorder的左子树起点是从1开始的
	root.Left = buildTree3(preorder[1:idxLeft+2], postorder[:idxLeft+1])
	root.Right = buildTree3(preorder[idxLeft+2:], postorder[idxLeft+1:len(postorder)-1])

	return root
}

func search(nums []int, tar int) int {
	for i, v := range nums {
		if v == tar {
			return i
		}
	}

	return -1
}
