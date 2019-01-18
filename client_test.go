package gdeslon

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *ClientSuite) SetupSuite() {
	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testApiData)
	}))
}

func (s *ClientSuite) TearDownSuite() {
	if s.server != nil {
		s.server.Close()
	}
}

func (s *ClientSuite) TestRequest() {
	client := NewClient(
		s.server.URL+"/api/",
		"xxx",
	)

	params := url.Values{}
	rep, err := client.request("search.json", "GET", params)

	s.Assert().Nil(err)
	s.Assert().NotNil(rep)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

const testApiData = `{
   "offers":[
      {
         "gsProductKey":"0fffe9cb7d5b2c449f769065e8c794d63bd7d6f7",
         "available":true,
         "merchantId":84426,
         "id":3478785306774017000,
         "article":"310726",
         "gsCategoryId":443,
         "price":1070,
         "oldprice":null,
         "charge":19.47,
         "currencyId":"RUR",
         "picture":[
            "//img.gdeslon.ru/commodities/big/5645/0523dc41439766337584079f4918.big.jpg"
         ],
         "thumbnail":[
            "//img.gdeslon.ru/commodities/small/5645/0523dc41439766337584079f4918.small.jpg"
         ],
         "name":"Чайник VITEK VT-7031",
         "description":"Чайник VITEK VT-7031 Закрытая спираль Корпус из пластика",
         "vendor":"VITEK",
         "model":"VT-7031",
         "originalPicture":[
            "http://housebt.ru/upload/iblock/319/454.jpg"
         ],
         "url":"https://af.gdeslon.ru/cc/c79d32affee0ccdbb778e9c200eddcdcd279dee6/304720ce1c1d1303",
         "destinationUrlDoNotSendTraffic":"http://housebt.ru/catalog/small_household_appliances/kitchen_appliances/electric_kettles_and_thermopots/elektroch/vitek_vt_7031_belyy.html?utm_source=gde_slon&utm_medium=partner&utm_term=310726&utm_content=310726&utm_campaign=402"
      },
      {
         "gsProductKey":"ff46cc4183bea80228c2db58244c6d1779f39947",
         "available":true,
         "merchantId":84426,
         "id":7016531937003391000,
         "article":"310732",
         "gsCategoryId":443,
         "price":1150,
         "oldprice":null,
         "charge":20.93,
         "currencyId":"RUR",
         "picture":[
            "//img.gdeslon.ru/commodities/big/fe3b/46ee07f73a7ef1defa278b5eb26e.big.jpg"
         ],
         "thumbnail":[
            "//img.gdeslon.ru/commodities/small/fe3b/46ee07f73a7ef1defa278b5eb26e.small.jpg"
         ],
         "name":"Чайник VITEK VT-7061",
         "description":"Чайник VITEK VT-7061 Закрытая спираль Корпус из пластика Установка на подставку в любом положении",
         "vendor":"VITEK",
         "model":"VT-7061",
         "originalPicture":[
            "http://housebt.ru/upload/iblock/c68/457.jpg"
         ],
         "url":"https://af.gdeslon.ru/cc/c79d32affee0ccdbb778e9c200eddcdcd279dee6/615fba9f1c45488d",
         "destinationUrlDoNotSendTraffic":"http://housebt.ru/catalog/small_household_appliances/kitchen_appliances/electric_kettles_and_thermopots/elektroch/vitek_vt_7061.html?utm_source=gde_slon&utm_medium=partner&utm_term=310732&utm_content=310732&utm_campaign=402"
      },
      {
         "gsProductKey":null,
         "available":true,
         "merchantId":84426,
         "id":12948069973561686000,
         "article":"339982",
         "gsCategoryId":443,
         "price":1420,
         "oldprice":null,
         "charge":25.84,
         "currencyId":"RUR",
         "picture":[
            "//img.gdeslon.ru/commodities/big/282b/b19b1a3ca3490bf9b34026390c9e.big.jpg"
         ],
         "thumbnail":[
            "//img.gdeslon.ru/commodities/small/282b/b19b1a3ca3490bf9b34026390c9e.small.jpg"
         ],
         "name":"Чайник VITEK VT-7044",
         "description":"Чайник VITEK VT-7044 объем 1.7 л мощность 2200 Вт корпус из стекла закрытая спираль",
         "vendor":"VITEK",
         "model":"VT-7044",
         "originalPicture":[
            "http://housebt.ru/upload/iblock/8a3/6991.jpg"
         ],
         "url":"https://af.gdeslon.ru/cc/c79d32affee0ccdbb778e9c200eddcdcd279dee6/b3b0c90a050bbb48",
         "destinationUrlDoNotSendTraffic":"http://housebt.ru/catalog/small_household_appliances/kitchen_appliances/electric_kettles_and_thermopots/elektroch/vitek_vt_7044.html?utm_source=gde_slon&utm_medium=partner&utm_term=339982&utm_content=339982&utm_campaign=402"
      },
      {
         "gsProductKey":"6c9398c7bc747a073fc66c0ebb7026c24a14090b",
         "available":true,
         "merchantId":84426,
         "id":10164713601301236000,
         "article":"310730",
         "gsCategoryId":443,
         "price":1760,
         "oldprice":null,
         "charge":32.03,
         "currencyId":"RUR",
         "picture":[
            "//img.gdeslon.ru/commodities/big/dad8/bd7b15bfcceadcdab54d6129d552.big.jpg"
         ],
         "thumbnail":[
            "//img.gdeslon.ru/commodities/small/dad8/bd7b15bfcceadcdab54d6129d552.small.jpg"
         ],
         "name":"Чайник VITEK VT-7054",
         "description":"Чайник VITEK VT-7054 Закрытая спираль Корпус из пластика Установка на подставку в любом положении",
         "vendor":"VITEK",
         "model":"VT-7054",
         "originalPicture":[
            "http://housebt.ru/upload/iblock/421/456.jpg"
         ],
         "url":"https://af.gdeslon.ru/cc/c79d32affee0ccdbb778e9c200eddcdcd279dee6/8d10512a5924492f",
         "destinationUrlDoNotSendTraffic":"http://housebt.ru/catalog/small_household_appliances/kitchen_appliances/electric_kettles_and_thermopots/elektroch/vitek_vt_7054_chernyy.html?utm_source=gde_slon&utm_medium=partner&utm_term=310730&utm_content=310730&utm_campaign=402"
      },
      {
         "gsProductKey":"2b8171dbeda3829557a57913eafcb48bec44945b",
         "available":true,
         "merchantId":84426,
         "id":7946590174837127000,
         "article":"310728",
         "gsCategoryId":443,
         "price":1390,
         "oldprice":null,
         "charge":25.3,
         "currencyId":"RUR",
         "picture":[
            "//img.gdeslon.ru/commodities/big/4a3d/fe87b03550279341c94bcf2cbf09.big.jpg"
         ],
         "thumbnail":[
            "//img.gdeslon.ru/commodities/small/4a3d/fe87b03550279341c94bcf2cbf09.small.jpg"
         ],
         "name":"Чайник VITEK VT-7039",
         "description":"Чайник VITEK VT-7039 Закрытая спираль Корпус из стали Установка на подставку в любом положении Подсветка корпуса",
         "vendor":"VITEK",
         "model":"VT-7039",
         "originalPicture":[
            "http://housebt.ru/upload/iblock/db5/455.jpg"
         ],
         "url":"https://af.gdeslon.ru/cc/c79d32affee0ccdbb778e9c200eddcdcd279dee6/6e47f5a942c6b840",
         "destinationUrlDoNotSendTraffic":"http://housebt.ru/catalog/small_household_appliances/kitchen_appliances/electric_kettles_and_thermopots/elektroch/vitek_vt_7039.html?utm_source=gde_slon&utm_medium=partner&utm_term=310728&utm_content=310728&utm_campaign=402"
      }
   ]
}`
