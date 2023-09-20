package foods

import (
	"fmt"
	"foods-spider-server/pkg"
	"github.com/gocolly/colly/v2"
	"strings"
)

type CommonFood struct {
	Cover string `json:"cover"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

// GetCommonFoodsLogic 获取常见的食物类型列表
func GetCommonFoodsLogic() ([]*CommonFood, error) {
	cldC := pkg.NewCLDColly(CommonFoodsList)
	CommonFoodsList := make([]*CommonFood, 0)
	cldC.C.OnHTML("td.borderBottom > a", func(e *colly.HTMLElement) {
		var CommonFoodsListItem = new(CommonFood)
		// 获取a标签的href属性
		href := e.Attr("href")
		// 打印结果
		// fmt.Printf("href: %s\n", href)
		CommonFoodsListItem.Link = BaseHost + href
		// 获取a标签中的img标签
		img := e.DOM.Find("img")
		if img != nil {
			// 获取img标签的src属性
			if val, exists := img.Attr("src"); exists {
				CommonFoodsListItem.Cover = val
			}
			if val, exists := img.Attr("alt"); exists {
				CommonFoodsListItem.Title = val
			}
		}
		CommonFoodsList = append(CommonFoodsList, CommonFoodsListItem)
	})
	err := cldC.Do()
	if err != nil {
		return CommonFoodsList, err
	}
	return CommonFoodsList, nil
}

type FoodSearchResItem struct {
	Title string `json:"title"`
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Desc  string `json:"desc"`
}

// SearchFoodLogic 搜索食物
func SearchFoodLogic(foodsName string) ([]*FoodSearchResItem, error) {
	FoodSearchResItemList := make([]*FoodSearchResItem, 0)
	aimURL := fmt.Sprintf("%s%s", SearchFoods, pkg.URIEncoder([]string{"q",
		foodsName, "pg", "0"}))
	cldC := pkg.NewCLDColly(aimURL)
	cldC.C.OnHTML("td.borderBottom", func(element *colly.HTMLElement) {
		var SearchItem = new(FoodSearchResItem)
		aProminent := element.DOM.Find("a.prominent")
		// 寻找a标签
		if aProminent != nil {
			SearchItem.Title = aProminent.Text()
			if val, exists := aProminent.Attr("href"); exists {
				SearchItem.Link = BaseHost + val
			}
		}
		aBrand := element.DOM.Find("a.brand")
		if aBrand != nil {
			SearchItem.Brand = aBrand.Text()
		}
		divSmallText := element.DOM.Find("div.smallText")
		if divSmallText != nil {
			smallText := strings.TrimSpace(divSmallText.Text())
			smallTextArr := strings.Split(smallText, "\n")
			smallTextArrNew := make([]string, 0)
			for i := 0; i < len(smallTextArr); i++ {
				if len(strings.TrimSpace(smallTextArr[i])) > 0 {
					smallTextArrNew = append(smallTextArrNew, strings.TrimSpace(smallTextArr[i]))
				}
			}
			//fmt.Println(smallTextArrNew[0])
			SearchItem.Desc = smallTextArrNew[0]
			FoodSearchResItemList = append(FoodSearchResItemList, SearchItem)
		}
	})
	err := cldC.Do()
	if err != nil {
		return FoodSearchResItemList, err
	}
	return FoodSearchResItemList, err
}
