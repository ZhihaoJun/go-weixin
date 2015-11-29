package weixin

import (
    "fmt"
    "log"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "crypto/sha1"
)

type Weixin struct {
    AppId string
    AppSecret string
}

type AccessTokenResponse struct {
    AccessToken string `json:"access_token"`
    ExpiresIn int64 `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
    OpenId string `json:"openid"`
    Scope string `json:"scope"`
    UnionId string `json:"unionid"`
    Errcode int64 `json:"errcode"`
    Errmsg string `json:"errmsg"`
}

func (resp *AccessTokenResponse) Ok() bool {
    if resp.Errcode == 0 {
        return true
    }
    return false
}

type UserInfoResponse struct {
    OpenId string `json:"openid"`
    Nickname string `json:"nickname"`
    Sex int64 `json:"sex"`
    Province string `json:"province"`
    City string `json:"city"`
    Country string `json:"country"`
    Headimgurl string `json:"headimgurl"`
    Privilege []string `json:"privilege"`
    UnionId string `json:"unionid"`
    Errcode int64 `json:"errcode"`
    Errmsg string `json:"errmsg"`
}

func (resp *UserInfoResponse) Ok() bool {
    if resp.Errcode == 0 {
        return true
    }
    return false
}

const (
    webAuthRedirectURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
    getWebAccessToken = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
    getUserInfo = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
)

func New(appId string, appSecret string) Weixin {
    return Weixin{
        AppId: appId,
        AppSecret: appSecret,
    }
}

func (wx *Weixin) WebAuthRedirectURL(redirectURI string, scope string, state string) string {
    return fmt.Sprintf(webAuthRedirectURL, wx.AppId, redirectURI, scope, state)
}

func (wx *Weixin) GetWebAccessToken(code string) (*AccessTokenResponse, error) {
    url := fmt.Sprintf(getWebAccessToken, wx.AppId, wx.AppSecret, code)
    log.Println("get web access token request url: ", url)
    resp, err := http.Get(url)
    if err != nil {
        log.Println("failed to request url")
        return nil, err
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("failed to read response body")
        return nil, err
    }
    log.Println("response body is ", string(body))
    
    var response AccessTokenResponse
    err = json.Unmarshal(body, &response)
    if err != nil {
        log.Println("failed to parse body to json")
        return nil, err
    }
    log.Printf("body json response is %v\n", response)
    return &response, nil
}

func (wx *Weixin) GetUserInfo(accessToken string, openId string) (*UserInfoResponse, error) {
    url := fmt.Sprintf(getUserInfo, accessToken, openId)
    log.Println("get user info request url: ", url)
    resp, err := http.Get(url)
    if err != nil {
        log.Println("failed to request url")
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("failed to read response body")
        return nil, err
    }
    log.Println("response body is ", string(body))

    var response UserInfoResponse
    err = json.Unmarshal(body, &response)
    if err != nil {
        log.Println("failed to parse body to json")
        return nil, err
    }
    log.Printf("body json response is %v\n", response)
    return &response, nil
}

func (wx *Weixin) JSSDKSignature(jssdkTicket string, noncestr string, timestamp int64, url string) string {
    string1 := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", jssdkTicket, noncestr, timestamp, url)
    return fmt.Sprintf("%x", sha1.Sum([]byte(string1)))
}
