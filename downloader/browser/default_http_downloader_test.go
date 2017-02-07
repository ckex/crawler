package browser

import (
	"testing"
	"robot/downloader/request"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

// go test -run="Test_Http_Download"
func Test_Http_Download(t *testing.T) {
	client := NewHttpDownloader()
	response, err := client.Download(&request.Request{
		Url:"http://www.tianyancha.com/search?key=%E6%B7%B1%E5%9C%B3%E5%A4%A9%E9%81%93%E8%AE%A1%E7%84%B6%E9%87%91%E8%9E%8D%E6%9C%8D%E5%8A%A1%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8&checkFrom=searchBox",
		//Url:"http://www.baidu.com",
	})
	if err != nil {
		t.Error(err)
	}
	print(response)
}

func Test_Qzone_2(t *testing.T) {
	cookies := "ptisp=ctc;pgv_info=ssid=s7095828286;pt2gguin=o0543109152;pgv_pvi=5707152384;ptcz=e14a38506c4c2cfa9fb978995962f2195cd6f8a481fa168ff9ff20976d407c4f;pgv_pvid=3892314838;p_skey=N2uVkWcuRr*5T3VoljxkhBQOWYhATUDBUIFPqCDmawo_;Loading=Yes;ptui_loginuin=543109152;pgv_si=s603205632;qz_screen=1024x768;skey=@TUeqqIoFJ;RK=kdcS5dqbep;pt4_token=8MnHnDrsDJ1REI49sIunPWo2pV4sI7DF*f5YfH2UHK8_;uin=o0543109152;QZ_FE_WEBP_SUPPORT=0;p_uin=o0543109152";
	client := http.DefaultClient
	request, _ := http.NewRequest("GET", "http://h5.qzone.qq.com/proxy/domain/ic2.qzone.qq.com/cgi-bin/feeds/feeds_html_act_all?uin=543109152&hostuin=467732755&scope=0&filter=all&flag=1&refresh=0&firstGetGroup=0&mixnocache=0&scene=0&begintime=undefined&icServerTime=&start=0&count=10&sidomain=qzonestyle.gtimg.cn&useutf8=1&outputhtmlfeed=1&refer=2&r=0.8536552901318883&g_tk=56408880", nil)
	for _, value := range strings.Split(cookies, ";") {
		c := strings.Split(value, "=")
		request.AddCookie(&http.Cookie{
			Name:c[0],
			Value:c[1],
		})
	}
	response, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}
	print(response)
}

func Test_China_Unicom(t *testing.T) {
	client := http.DefaultClient
	url := "https://uac.10010.com/portal/Service/MallLogin?callback=jQuery17205765543769533836_1484877676577&req_time=1484877685118&redirectURL=http%3A%2F%2Fwww.10010.com&userName=ckex868%40vip.qq.com&password=ckex13534117937&pwdType=01&productType=05&redirectType=03&rememberMe=1&_=1484877685119"

	resp, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	print(resp)
	u := "http://iservice.10010.com/e4/index_server.html"
	//u := "http://iservice.10010.com/e3/ToExcel.jsp?type=sound"
	//u := "http://111.111.110.98:10086/media"
	request, err := http.NewRequest("GET", u, nil)

	for _, c := range resp.Cookies() {
		request.AddCookie(c)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Login success. ")
		resp2, err2 := client.Do(request)
		if err2 != nil {
			t.Error(err2)
		}
		fmt.Println("2================>", resp2, resp2.ContentLength)
		// print(resp2)
		response, _ := ioutil.ReadAll(resp2.Body)
		ioutil.WriteFile("/Users/ckex/Desktop/unicom.html", response, 0666)
	}
	//fmt.Println("1================>", resp)
}