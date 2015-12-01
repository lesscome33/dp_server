package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/lesscome33/dp_server/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"log"
	"strings"
)

// Operations about object
type PriceController struct {
	beego.Controller
}

// @Title create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *PriceController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJson()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *PriceController) Get() {
	fmt.Println("test")
	for k, _ := range o.Ctx.Input.Params {
		fmt.Println("key:%s , value:%s", k, o.Ctx.Input.Params[k])
	}
	
	objectId := o.Ctx.Input.Params[":objectId"]
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJson()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *PriceController) GetAll() {
	fmt.Println("test getall");
	shape := o.GetString("shape")
	carat := o.GetString("carat")
	color := o.GetString("color")
	clarity := o.GetString("clarity")
	certificate := o.GetString("certificate")
	proportion := o.GetString("proportion")
	gotext := o.GetString("go")

	url := fmt.Sprintf("http://data.washingtondiamond.com/store/calc-netsuite.php?shape=%s&carat=%s&color=%s&clarity=%s&certificate=%s&proportion=%s&go=%s", shape, carat, color, clarity, certificate, proportion, gotext)
	result := ExampleScrape(url)

	//obs := models.GetAll()
	o.Data["json"] = result
	o.ServeJson()
}

// @Title update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *PriceController) Put() {
	objectId := o.Ctx.Input.Params[":objectId"]
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJson()
}

// @Title delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *PriceController) Delete() {
	objectId := o.Ctx.Input.Params[":objectId"]
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJson()
}

func ExampleScrape(url string) string{
	//doc, err := goquery.NewDocument("http://data.washingtondiamond.com/store/calc-netsuite.php?shape=A&carat=2.0&color=D&clarity=1&certificate=E&proportion=6&go=Calculate+Price")
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var result string;

	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		htmlText, _ := s.Html()
		htmlText = strings.TrimSpace(htmlText)
		if strings.HasPrefix(htmlText, "$") {
			fmt.Printf("Review %d: %s\n", i, htmlText)
			index := strings.Index(htmlText, " ")
			result = htmlText[:index]
			fmt.Println("result: " + result)
		}
	})
	return result;
}
