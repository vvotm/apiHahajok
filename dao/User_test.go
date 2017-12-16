/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 13:10
 * @description: 
 */

package dao

import (
	"testing"
	"github.com/vvotm/apiHahajok/db"
	"github.com/vvotm/apiHahajok/models/request"
	"time"
)

func TestUser_UpdateUserInfo(t *testing.T) {
	dbConn := db.GetConn()

	countSql := "select id from user where openId = ? and issave = ? and updatedAt > ?"
	var count int = 0
	t.Log(time.Now().Unix())
	dbConn.QueryRow(countSql, "oln4F0elOAeB6GIM6lsAh5bTRdrE", 1, time.Now().Unix()-600).Scan(&count)
	t.Log(count)
	return



	r := &request.ReqUserInfo{
		OpenId:"oln4F0elOAeB6GIM6lsAh5bTRdrE",
		Nickname:"luowen",
		Gender:"1",
	}
	nowTime := time.Now().Unix()
	updateSql := "update user set unionId = ?, nickname = ?, avatar = ?, gender = ?, country = ?," +
		"province = ?, city = ?, lang = ?, updatedAt = ? where openId = ?"
	_, err := dbConn.Exec(updateSql, r.UnionId, r.Nickname, r.Avatar, r.Gender, r.Country, r.Province, r.City,
		r.Lang, nowTime, r.OpenId)
	t.Log(err)
}

func TestUser_GetUserInfoById(t *testing.T) {
	userDao := NewUser()
	resultSet := userDao.GetUserInfoById(1)
	t.Log(resultSet)

}
