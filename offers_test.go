package gdeslon

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OffersSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *OffersSuite) SetupSuite() {
	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testOffersApiData)
	}))
}

func (s *OffersSuite) TearDownSuite() {
	if s.server != nil {
		s.server.Close()
	}
}

func (s *OffersSuite) TestOffers() {
	client := NewClient(
		s.server.URL+"/api/",
		"xxx",
	)

	type Response struct {
		Offers []struct {
			Id          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Url         string `json:"url"`
		} `json:"offers"`
	}

	r := Response{}
	params := url.Values{}
	err := client.Call("search.json", "GET", params, &r)
	s.Assert().Nil(err)

	s.Assert().Len(r.Offers, 2)
	s.Assert().EqualValues("Чайник VITEK VT-7031", r.Offers[0].Name)
	s.Assert().EqualValues(3478785306774017000, r.Offers[0].Id)
	s.Assert().EqualValues("Чайник VITEK VT-7031 Закрытая спираль Корпус из пластика", r.Offers[0].Description)
}

func TestOffersSuite(t *testing.T) {
	suite.Run(t, new(OffersSuite))
}

const testOffersApiData = `{
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
      }
   ]
}`
