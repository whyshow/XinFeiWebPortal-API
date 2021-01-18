package models

import (
	"XinFeiWebPortal-API/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"os"
	"strconv"
)

type Xinfei_honor struct {
	Honor_id       string `orm:"column(honor_id);pk"`
	Honor_url      string
	Honor_classify string
	Honor_year     string
	Honor_describe string
}

// 增加
func InsertHonor(honer Xinfei_honor) (bool, error) {
	o := orm.NewOrm()
	honer.Honor_id = utils.Random(5) //生成荣耀信息id
	//sql := "INSERT INTO xinfei_ry ( honor_id, honor_url,honor_year,honor_classify,honor_describe )VALUES("+ honer.Honor_id +","+ "'" + honer.Honor_url+ "'"  +","+ honer.Honor_year +","+ "'" +honer.Honor_classify + "'" +","+ "'" + honer.Honor_describe+ "'" +")"
	//fmt.Println(sql)
	//res, err := o.Raw(sql).Exec()
	//fmt.Println(res)
	_, err := o.Insert(&honer)
	if err == nil {
		return true, err
	}
	fmt.Println(err)
	return false, err
}

// 删除
func DeleteHonor(id string) (string, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Xinfei_honor{Honor_id: id}); err == nil {
		if num == 1 {
			return "删除成功!", err
		}
		return "删除失败!", err
	} else {
		return "删除失败!", err
	}
}

// 修改
// 查询一个，根据id
func SeleteHonor(id string) (Xinfei_honor, string, error) {
	o := orm.NewOrm()
	honor := Xinfei_honor{Honor_id: id}
	err := o.Read(&honor)
	if err == orm.ErrNoRows {
		return honor, "查询不到", err
	} else if err == orm.ErrMissPK {
		return honor, "找不到主键", err
	} else {
		return honor, "", nil
	}
}

// 查询分类下所有
func SeleteHonorList(classify string, n int, p int) ([]Xinfei_honor, int64, error) {
	honor := []Xinfei_honor{}
	o := orm.NewOrm()
	sql := ""
	if p == 1 {
		sql = "select * from xinfei_honor where honor_classify = " + "'" + classify + "'" + " ORDER BY honor_year DESC" + " limit 0" + "," + strconv.Itoa(n)
	} else if p > 1 {
		s := p*n - n
		sql = "select * from xinfei_honor where honor_classify = " + "'" + classify + "'" + " ORDER BY honor_year DESC " + " limit " + strconv.Itoa(s) + "," + strconv.Itoa(n)
	} else {
		sql = "select * from xinfei_honor where honor_classify = " + "'" + classify + "'" + " ORDER BY honor_year DESC" + " limit 0" + "," + strconv.Itoa(n)
	}
	// 返回行数或者错误信息
	if num, err := o.Raw(sql).QueryRows(&honor); err == nil {
		if num == 0 {
			return honor, 0, err
		} else {
			count, err := o.QueryTable("xinfei_honor").Count()
			return honor, count, err
		}
	} else {
		return honor, 0, err
	}
}

//查询所有的荣誉
func SeleteAllHonorList(n int, p int) ([]Xinfei_honor, int64, error) {
	o := orm.NewOrm()
	honor := []Xinfei_honor{}
	sql := ""
	if p == 1 {
		sql = "SELECT * FROM xinfei_honor" + " ORDER BY honor_year DESC" + " limit 0" + "," + strconv.Itoa(n)
	} else if p > 1 {
		s := p*n - n
		sql = "SELECT * FROM xinfei_honor" + " ORDER BY honor_year DESC " + " limit " + strconv.Itoa(s) + "," + strconv.Itoa(n)
	} else {
		sql = "SELECT * FROM xinfei_honor" + " ORDER BY honor_year DESC " + " limit 0" + "," + strconv.Itoa(n)
	}
	// 返回行数或者错误信息
	if num, err := o.Raw(sql).QueryRows(&honor); err == nil {
		if p == 0 {
			return honor, num, err
		} else {
			count, err := o.QueryTable("xinfei_honor").Count()
			return honor, count, err
		}
	} else {
		return honor, 0, err
	}
}

var hc = HonorConf{}

// 获取荣誉分类
func HonorClassify() (HonorConf, error) {
	if hc.Array == nil {
		var hcf = HonorConf{}
		func() {}()
		{
			f, err := os.Open("./conf/honor.json")
			defer f.Close()
			if err != nil {
				return hcf, err
			}
			bytes, _ := ioutil.ReadAll(f)
			err = json.Unmarshal(bytes, &hcf)
			if err != nil {
				return hcf, err
			}
			hc = hcf
			return hc, nil
		}
	}
	return hc, nil
}
