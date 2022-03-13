package module

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const prefix = "https://fxgl.jx.edu.cn/"

type HealthyReport struct {
	StudentID string
	SchoolID  string
	Address   string
}

func NewHealthyReport(studentID, schoolID, address string) *HealthyReport {
	return &HealthyReport{
		StudentID: studentID,
		SchoolID:  schoolID,
		Address:   address,
	}
}

func (h *HealthyReport) Report() error {
	cks, err := h.login()
	if err != nil {
		return err
	}
	err = h.sign(cks)
	if err != nil {
		return err
	}
	return nil
}

func (h *HealthyReport) sign(cks []*http.Cookie) error {
	var client http.Client
	jar, _ := cookiejar.New(nil)
	u, _ := url.Parse(prefix)
	jar.SetCookies(u, cks)
	client.Jar = jar
	formData := url.Values{}
	lc, err := getLocationByAddress(h.Address)
	if err != nil {
		return err
	}
	ad := splitAddress(h.Address)
	formData.Add("province", ad.Province)
	formData.Add("city", ad.City)
	formData.Add("district", ad.District)
	formData.Add("street", ad.Street)
	formData.Add("xszt", "0")
	formData.Add("jkzk", "0")
	formData.Add("jkzkxq", "")
	formData.Add("sfgl", "1")
	formData.Add("gldd", "")
	formData.Add("mqtw", "0")
	formData.Add("mqtwxq", "")
	formData.Add("zddlwz", "江西省九江市共青城市青年大道79号江西农业大学南昌商学院")
	formData.Add("sddlwz", "")
	formData.Add("bprovince", ad.Province)
	formData.Add("bcity", ad.City)
	formData.Add("bdistrict", ad.District)
	formData.Add("bstreet", ad.Street)
	formData.Add("sprovince", ad.Province)
	formData.Add("scity", ad.City)
	formData.Add("sdistrict", ad.District)
	formData.Add("lng", fmt.Sprintf("%v", lc.Result.Location.Lng))
	formData.Add("lat", fmt.Sprintf("%v", lc.Result.Location.Lat))
	formData.Add("sfby", "1")
	resp, err := client.PostForm(prefix+h.SchoolID+"/studentQd/saveStu", formData)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	return nil
}

func (h *HealthyReport) login() ([]*http.Cookie, error) {
	p := prefix + h.SchoolID
	le := p + "/public/homeQd?loginName=" + h.StudentID + "&loginType=0"
	r, err := http.Get(le)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, err
	}
	var flag bool
	doc.Find(".link-itemFa .link-item a").Each(func(i int, s *goquery.Selection) {
		str := s.Find("div").Text()
		if strings.TrimSpace(str) == "健康签到" {
			flag = true
			return
		}
	})
	if !flag {
		return nil, errors.New("login failed")
	}
	return r.Request.Response.Cookies(), nil
}

type AddressDetail struct {
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Street   string `json:"street"`
}

func splitAddress(address string) *AddressDetail {
	var province, city, district, street, iteration string
	s := strings.Split(address, "省")
	if len(s) >= 2 {
		province = s[0] + "省"
		iteration = s[1]
	}
	s = strings.Split(iteration, "市")
	if len(s) == 2 {
		city = s[0] + "市"
		iteration = s[1]
	} else if len(s) == 3 {
		s = strings.Split(iteration, "市")
		city = s[0] + "市"
		iteration = s[1] + "市" + s[2]
	}
	// 第三级有 区/县/市（县级市）
	s = strings.Split(iteration, "市")
	if len(s) == 2 {
		district = s[0] + "市"
		iteration = s[1]
	}
	if district == "" {
		s = strings.Split(iteration, "县")
		if len(s) >= 2 {
			district = s[0] + "县"
			iteration = s[1]
		}
		if district == "" {
			s = strings.Split(iteration, "区")
			if len(s) >= 2 {
				district = s[0] + "区"
				iteration = s[1]
			}
		}
		street = iteration
	}

	return &AddressDetail{
		Province: province,
		City:     city,
		District: district,
		Street:   street,
	}
}
