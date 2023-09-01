package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func BootstrapHandler(c buffalo.Context) error {

	// externalURL := "http://localhost:3001/bootstrap/chart"
	// resp, err := http.Get(externalURL)
	// if err != nil {
	// 	return c.Error(http.StatusInternalServerError, err)
	// }
	// defer resp.Body.Close()

	// htmlBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return c.Error(http.StatusInternalServerError, err)
	// }
	// htmlContent := string(htmlBytes)
	// c.Set("respBody", htmlContent)
	// c.Set("respBodys", "sss")
	// fmt.Println(htmlContent)

	c.Set("respBody", "respBody")
	c.Set("respBodys", "respBodys")

	return c.Render(http.StatusOK, r.HTML("dashboard/index.html"))
}

func ExternalWidgetHandler(c buffalo.Context) error {
	// path := c.Param("path")
	// fmt.Println(path)
	// 인증이 필요한 외부 사이트의 URL
	externalURL := "http://localhost:3001/bootstrap/chart"

	// 여기에서 토큰 또는 인증 정보를 가져오는 로직을 구현
	// token := "your_token"

	// 인증 정보를 헤더에 추가하여 외부 사이트의 내용 가져오기
	req, err := http.NewRequest("GET", externalURL, nil)
	if err != nil {
		return err
	}
	// req.Header.Add("Authorization", "Bearer "+token)

	// 외부 사이트의 내용 가져오기
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// 외부 사이트의 내용 읽기
	// ...

	// 외부 사이트의 내용을 iframe에 표시할 형태로 템플릿으로 전달
	return c.Render(http.StatusOK, r.HTML(string(bodyBytes)))
}

func BootstrapLoginFormHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("login/index.html"))
}
