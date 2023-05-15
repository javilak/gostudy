package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type search struct {
	Req Req `json:"req"`
}
type Req struct {
	PageNo        string `json:"page_no"`
	PageSize      string `json:"page_size"`
	Name          string `json:"name"`
	CertificateNo string `json:"certificateNo"`
	IDNum         string `json:"idCarNumber"`
}

type Ham struct {
	Code int `json:"code"`
	Res  struct {
		PrcList []struct {
			ID            string `json:"id"`
			Remarks       string `json:"remarks"`
			CreateDate    string `json:"createDate"`
			UpdateDate    string `json:"updateDate"`
			CertID        string `json:"certId"`
			IDCarNumber   string `json:"idCarNumber"`
			Name          string `json:"name"`
			Sex           string `json:"sex"`
			DateForBirth  string `json:"dateForBirth"`
			HomeCall      string `json:"homeCall"`
			Type          string `json:"type"`
			PassDate      string `json:"passDate"`
			PassAddr      string `json:"passAddr"`
			IssueDate     string `json:"issueDate"`
			Addr          string `json:"addr"`
			CertificateNo string `json:"certificateNo"`
			Province      string `json:"province"`
			City          string `json:"city"`
			Street        string `json:"street"`
			PostalCode    string `json:"postalCode"`
			Telephone     string `json:"telephone"`
			Email         string `json:"email"`
			Qq            string `json:"qq"`
			Ps            string `json:"ps"`
			FlagAddr      string `json:"flagAddr"`
			FlagHomeCall  string `json:"flagHomeCall"`
			FlagCode      string `json:"flagCode"`
			FlagTel       string `json:"flagTel"`
			FlagEmail     string `json:"flagEmail"`
			FlagQq        string `json:"flagQq"`
			FlagPs        string `json:"flagPs"`
			FileLj        string `json:"fileLj"`
			ValidityDate  string `json:"validityDate"`
			IsShow        string `json:"isShow"`
			SignUpID      string `json:"signUpId"`
			ShowImg       string `json:"showImg"`
		} `json:"prcList"`
		HkList    []interface{} `json:"hkList"`
		FgList    []interface{} `json:"fgList"`
		Count     int           `json:"count"`
		PageSize  string        `json:"page_size"`
		PageNo    string        `json:"page_no"`
		TotalPage string        `json:"total_page"`
	} `json:"res"`
	ResMeta struct {
	} `json:"res_meta"`
	Msg string `json:"msg"`
}

func main() {
	fmt.Println("输入姓名：")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("错误")
	}
	name := strings.TrimSuffix(input, "\r\n")
	fmt.Println("ID:")
	ID := bufio.NewReader(os.Stdin)
	idcard, err := ID.ReadString('\n')
	if err != nil {
		fmt.Println("错误")
	}
	num := strings.TrimSuffix(idcard, "\r\n")
	check(name, num)
}

func check(name string, ID string) {

	client := &http.Client{}
	request := search{Req: Req{PageNo: "1", PageSize: "100", Name: name, CertificateNo: "", IDNum: ID}}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "http://82.157.138.16:8091/CRAC/app/businessSupport/cracOperationCert/getOperCertByParamWeb", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Content-Type", "application/json; chartset=UTF-8")
	req.Header.Set("Cookie", "jeesite.session.id=d9b95a3ae4e447679c26724baccb8673; JSESSIONID=FBE4B00BCD66152C5F6D65E4D3A0771A")
	req.Header.Set("Origin", "http://82.157.138.16:8091")
	req.Header.Set("Proxy-Connection", "keep-alive")
	req.Header.Set("Referer", "http://82.157.138.16:8091/CRAC/crac/pages/list_cert.html?name=6a2P5paH54aZ&certificateNo=&idCarNumber=NDQwMTAyMTk5OTEyMjIwNjMw")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("mm", "null")
	req.Header.Set("qm", "null")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("error,bad request")
	}
	var ham Ham
	err = json.Unmarshal(bodyText, &ham)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("你的操作证号是：", ham.Res.PrcList[0].CertificateNo)
	//fmt.Printf("%s\n", bodyText)

}
