package cryptogram

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestAESDecrypt(t *testing.T) {
	str := "068c91aca0ba3fe687f487e3f161fe36a9af6258c1e3ae8e8840746c6a1c224cae5224b44ec4d1fd559d75c7110eb9ac34acb14a7950b1e471514c6883842a5fb02c2e0496a9f58fbedfa4fc43d4192f438e5259cde2a82445e069352281c819788402006c7cbec62f6704ed36c422b1854cefd32ac010eb482b94135fa7e53809cdc311c1a01b279453c82cc5369cbd"
	data, err := AESDecrypt(str, "ab#a123@fvwpokdi")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(data)
	str2 := "068c91aca0ba3fe687f487e3f161fe36a9af6258c1e3ae8e8840746c6a1c224cae5224b44ec4d1fd559d75c7110eb9ac34acb14a7950b1e471514c6883842a5fb02c2e0496a9f58fbedfa4fc43d4192f438e5259cde2a82445e069352281c819788402006c7cbec62f6704ed36c422b1854cefd32ac010eb482b94135fa7e5381056901042138341e19b91b13cc36f17"
	data2, err2 := AESDecrypt(str2, "ab#a123@fvwpokdi")
	if err2 != nil {
		t.Error(err2.Error())
		return
	}
	t.Log(data2)
}

func TestAesDecrypt(t *testing.T) {
	key := "ab#a123@fvwpokdi"
	val, err := AESDecrypt("1f4fada9918336643b2b1c844a1df083d616bd85d13231db6f399963a655c41affd49d99bdcf3724150c0833758bc9932d05bfc9a02f468456e4e67ca698c9a0f43eaa6db9ea9efcf0342f65f25c69b58505a07e5c508e63835950de261f5ec8cc5c572659cefcf19ff7128ceb61f06f6c23360c3c2c5436a43f189cc7de399cc92e7f711e0c81eab7aaefb281a1824e", key)
	if val == "" || err != nil {
		log.Println("fdfsd")
		return
	}
	log.Println(val)

	cookieInfo := &busCookieInfo{}
	if cookieErr := json.Unmarshal([]byte(val), &cookieInfo); cookieErr != nil {
		log.Printf("cookieErr string:%+v\n ", cookieErr)
	}
	if time.Now().Unix()-cookieInfo.Timestamp >= 1000*60*24 {
		// 超时时间24h
		log.Println("v ...interface{}")
	}
}

type busCookieInfo struct {
	UserID    int64  // 用户ID
	UnionID   string // 微信UnionID
	RealName  string // 真实姓名
	OpenID    string // 用户微信凭证
	Timestamp int64  // 时间戳，cookie有效期24h
}

func TestName(t *testing.T) {
	var key = "Rp2H9tX+jIKXYg9mAP5dMw=="
	var str = "26e22759c2a2053506a579ff26aafa8f6cd1ce5f37d090b251b8031a7dcc2387"
	data, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	t.Logf("key: %s", data)
	s, err := AESDecrypt(str, string(data))
	if err != nil {
		t.Logf("dec err: %+v", err)
		return
	}
	t.Logf("id: %s", s)
}
