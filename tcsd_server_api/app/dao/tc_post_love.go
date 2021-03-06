// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go_api/app/dao/internal"
)

// tcPostLoveDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type tcPostLoveDao struct {
	*internal.TcPostLoveDao
}

var (
	// TcPostLove is globally public accessible object for table tc_post_love operations.
	TcPostLove tcPostLoveDao
)

func init() {
	TcPostLove = tcPostLoveDao{
		internal.NewTcPostLoveDao(),
	}
}

// Fill with you ideas below.
