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
	formData.Add("province", "江西省")
	formData.Add("city", "九江市")
	formData.Add("district", "共青城市")
	formData.Add("street", "青年大道79号江西农业大学南昌商学院")
	formData.Add("xszt", "0")
	formData.Add("jkzk", "0")
	formData.Add("jkzkxq", "")
	formData.Add("sfgl", "1")
	formData.Add("gldd", "")
	formData.Add("mqtw", "江西省")
	formData.Add("mqtwxq", "江西省")
	formData.Add("zddlwz", "江西省九江市共青城市青年大道79号江西农业大学南昌商学院")
	formData.Add("sddlwz", "")
	formData.Add("bprovince", "江西省")
	formData.Add("bcity", "九江市")
	formData.Add("bdistrict", "共青城市")
	formData.Add("bstreet", "青年大道79号江西农业大学南昌商学院")
	formData.Add("sprovince", "江西省")
	formData.Add("scity", "九江市")
	formData.Add("sdistrict", "共青城市")
	formData.Add("lng", "115.820966")
	formData.Add("lat", "29.167385")
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
		fmt.Println(str)
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
