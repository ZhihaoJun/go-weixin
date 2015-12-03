package weixin

import (
    "fmt"
    "testing"
)

func TestJSSDKSignature(t *testing.T) {
    wx := Weixin{}
    result := wx.JSSDKSignature("sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg", "Wm3WZYTPz0wzccnW", 1414587457, "http://mp.weixin.qq.com?params=value")
    if result != "0f9de62fce790f9a083d5c99e95740ceb90c27ed" {
        t.Errorf("not right result: %q", result)
    }
}

func TestGetWebAccessToken(t *testing.T) {
    wx := Weixin{"",""}
    result, err := wx.GetWebAccessToken("123123")
    fmt.Println(result)
    fmt.Println(err)
}
