package lab

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
	"strconv"
)

type RomItem struct {
	Floor	string			`json:"floor"`
	Name 	string			`json:"name"`
	Code 	string			`json:"code"`
	Device 	string			`json:"device"`
	Vlan 	string			`json:"vlan"`
	Machine string 			`json:"machine"`
	Admin 	[]string		`json:"admin"`
}

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	roms := []model.Roms{}

	err := db.C(model.CollectionRom).Find(nil).Sort("name").All(&roms)
	if err != nil {
		c.Error(err)
	}

	list := []RomItem{}

	for _, lou := range roms {
		for _, rom := range lou.Rom {
			var Admin []string
			if rom.Admin == nil {
				Admin = []string{}
			} else {
				Admin = rom.Admin
			}
			temprom := RomItem{
				Floor: 	lou.Id.Hex(),
				Name: 	rom.Name,
				Code: 	rom.Code,
				Device: rom.Device,
				Vlan: 	rom.Vlan,
				Machine:strconv.Itoa(len(rom.Machine)) + "台机器",
				Admin: 	Admin,
			}
			list = append(list, temprom)
		}
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
		"data" : list,
	})
}
