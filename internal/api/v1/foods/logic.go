package foods

import (
	"fmt"
	"foods-spider-server/pkg"
	"github.com/PuerkitoBio/goquery"
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
func SearchFoodLogic(foodsName string, pageNo string) ([]*FoodSearchResItem, error) {
	FoodSearchResItemList := make([]*FoodSearchResItem, 0)
	aimURL := fmt.Sprintf("%s%s", SearchFoods, pkg.URIEncoder([]string{"q",
		foodsName, "pg", pageNo}))
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

type FoodNutrientDetails struct {
	FoodName            string              `json:"food_name"`
	Brand               string              `json:"brand"`
	ServerSize          string              `json:"server_size,omitempty"`
	NutrientDetails     map[string]string   `json:"nutrient_details"`
	NutrientDetailsData []map[string]string `json:"nutrient_details_data"`
	OtherKeySlice       []string            `json:"other_key_slice"`
	OtherValueSlice     [][]string          `json:"other_value_slice"`
}

func GetFoodDetailsBySearchLinkLogic(link string) *FoodNutrientDetails {
	f := new(FoodNutrientDetails)
	newCLDColly := pkg.NewCLDColly(link)
	keyValueMap := make(map[string]string)
	newCLDColly.C.OnHTML("td.details", func(element *colly.HTMLElement) {
		element.DOM.Find("h4.separated").Each(func(i int, separatedSelection *goquery.Selection) {
			f.OtherKeySlice = append(f.OtherKeySlice, separatedSelection.Text())
		})
		element.DOM.Find("table.generic").Each(func(i int, selection *goquery.Selection) {
			otherA := selection.Find("td.borderBottom > a")
			if otherA.Length() > 0 {
				s := make([]string, 0)
				otherA.Each(func(i int, otherAItem *goquery.Selection) {
					if value, exists := otherAItem.Attr("href"); exists {
						s = append(s, otherAItem.Text())
						s = append(s, value)
					}
				})
				f.OtherValueSlice = append(f.OtherValueSlice, s)
			}
		})
	})
	newCLDColly.C.OnHTML("div.summarypanelcontent", func(element *colly.HTMLElement) {
		// 食物名称
		f.FoodName = element.DOM.Find("h1").Text()
		f.Brand = element.DOM.Find("a").Text()
	})
	newCLDColly.C.OnHTML("div.factPanel", func(element *colly.HTMLElement) {
		element.DOM.Find("div.factTitle").Each(func(i int, selection *goquery.Selection) {
			if i != 0 {
				key := selection.Text()
				value := selection.Next().Text()
				valuePure := strings.ReplaceAll(value, "克", "")
				f.NutrientDetailsData = append(f.NutrientDetailsData, map[string]string{
					"label": key,
					"value": valuePure,
				})
			}
		})
	})
	newCLDColly.C.OnHTML("div.nutrition_facts.international", func(element *colly.HTMLElement) {
		// 当前食用量
		servingSizeValue := element.DOM.Find("div.serving_size.serving_size_value")
		if servingSizeValue != nil {
			f.ServerSize = servingSizeValue.Text()
		}
		keySlice := make([]string, 0)
		// 获取营养成分
		element.ForEach("div.nutrient.left", func(i int, nutrientLeftElement *colly.HTMLElement) {
			key := nutrientLeftElement.DOM.Text()
			if len(strings.TrimSpace(key)) > 0 {
				keySlice = append(keySlice, key)
			}
		})
		valueSlice := make([]string, 0)
		element.ForEach("div.nutrient.right", func(i int, nutrientLeftElement *colly.HTMLElement) {
			attr := nutrientLeftElement.Attr("class")
			if strings.Index(attr, "per_serve") < 0 {
				value := nutrientLeftElement.DOM.Text()
				// 过滤掉焦尔值
				if len(strings.TrimSpace(value)) > 0 && i != 1 {
					valueSlice = append(valueSlice, value)
				}
			}

		})
		for i := 0; i < len(keySlice); i++ {
			keyValueMap[keySlice[i]] = valueSlice[i]
		}
	})
	// 其他可选使用量
	_ = newCLDColly.Do()
	f.NutrientDetails = keyValueMap
	return f
}
