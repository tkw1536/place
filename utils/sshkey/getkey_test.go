package sshkey

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// ssh key constants to test with
const (
	sshKeyWithoutPassPhrase = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAzgaUuhEcljVBWf2qgGFv9ulUSaENUpOEAKxARlfrcJU4GRoA
dvTwipgRaYC02j8UmzbXSYyKvj+L28CU/uKBYWPwHIH7TncxFE5kKNxMxXGtv73X
IfKYQKX/2Y6cf4EHuteYhnNWiu7fk0Jdu+oY99H8VO+aqJG0FH83CUrqOcH4mrg+
s9V6Ei0qMKojFmN6D3mDBDnmWDcJHFwd2SC0O/wQu45MsK2j13FudsIM+WbkHvpR
ug3LFOebpfY6fXa0GucJbB5mZpGG5lgfycuN6APQRT8SU8SBlUfjX4MQLKN8Nqsc
Rp35nVMRKQjHS4KEfyC/MroeblnZYUexmV8MyQIDAQABAoIBAQCs6EUF67qMLeHM
/uGLoTrwhF6i9LFTrk5Iqke/QaZs9C8Ckhn1vBfsmTdzzfr6d5p5sYr5RSRkCrz8
DyJ8z+g5rkAkDMq2zb25Bxl0WX9s7b0egNr+uLUi/K4/Djs1nzq3ip2NqVrmexfU
ZQx52zxdMDSPrA+mjbIOyb2M5Pyyvx+YCjso7IZuuIoJ7yz7ofy/ifOzvI0sImWW
8xklGiLFjzb21ZO0jQJGYHj0nMxSDAfPZgW/6a/5hIERZGOIritzBchraXk7J8HQ
dBD5tl9Cs0/N52LbraZHoZ2NW12m45LPjS8A8GR5S8qlJC0dvl59IALIrnJnsRUz
wIsFvZIJAoGBAPhBLjG/5t8T+/92Bx3p4M7QwIQyrgSGl+fKy4c2AoZ3kPHpP1Kb
Ugpo4CVOipwKQl+F92v7+gAL5NyesR3u3I1lpuWVBEI/71ALI9fCt1bQG9Nu0tZM
mCqaZk7VgcHaFrk/rdBC7yORXb0r1eHjcPKmi0q0+rMzTh4K0IWV2rofAoGBANR0
HdcPIjm4UOf+De+vfUcCtLLCF7uJvc3EK8PPxntKvmmlIW+bF+SX+CjvOGCCc7et
b3ztHffVJkGhdA+CsDZcGqzENLIzFYuTAY+FGDakPgauOzrLV+7PIWoFgAMRvvSt
eAyCbwlizGXUYEjORZYABc2EUsMWBJs/ITFT5SwXAoGBAPJiGWsjdZOqnGkQ4OP4
/LCQqtan8LWkf94lZ1BNkGufg9pdpKDP22skeGyUYcr2TVWcpDU/YRj4g+xP2Jhk
Jdy8OhZ/xxez+sEJD2bSy1Ssfe6SjrIDOLKn62nfFgCiIXufS+JB5+CvRnmzufEB
sr6HkwpO51NdrVCxuGQlKth7AoGAHBOgYeyBFGm0X4RmqRdjEgBciwc1hbZFXC0h
r4YE8ARHt8R377zqYm5nAFnk9HQpAMwt4K+hd0A3BxNkOCyIRxbS+6QOZsJzhXeP
DD2FnqsD+3QJJdL7svayrsU9TqqItuM560VNkUr6QjbX5qdD8PfdzHRBT8DYKQAl
zdQNhE8CgYADJaanTCNuiJIDfVw0xX4ZqXp2KvsUtVjWA8/zkRn28zHL1BriCAMK
yedEvZTFvMsPuv7thctGKW6vR9QyXAb6OOdS/VWs+8BiYoGBLgmhqkxXWxV1Po6e
jFrM5R1za30zIhO+77tBZi2mWdeKSn6NqvesGrfXBs8BKkIUH/0uqA==
-----END RSA PRIVATE KEY-----`
	sshKeyWithPassPhrase = `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-128-CBC,7C6DBCBC7154D4962B02F3260DAB7263

OOEQqAcVQAbonq7QF35M7W/uQI5FW9Qr4b2lVNQYoVbVurlV5xPE533omRhC68xK
mPZcSbhdVc2sHLn5Z3aPIvsEHEmvcx5prlbPbbt8rx+MKBkP+bjS7DKZsfUdZxAF
T003dWxz4q/A7TyEuFTAjT76baLs+5YiUUKSSF831novgzob64locYlC+0F6o7Ln
WJ19+HeVBymt0tGTYKpQc6MQif6aGLFeEElqRRcc1wFfqFtz/px9c+y6MGvveR2+
Ef6xrUoHijqQAxI5X+Jk51AscpQONGwRZF2NiJdCE9hCz7uu9u54GcHj03Lfxikb
ohf7+k8uK6f2YNTKsMlw15vXDjGYi5ztLpXDnqJPRGOFndStys6nLjWFtAL4E48Y
NkRtVfRwCCGrEJmiirO0OY+s2bK4S8c9jUOFYXhkt98KLfF1RC5mVXQPQHoLbcZ2
ZBMXrZ89jRiwc+KxsH2+AB5XAHWn2/GSwCoC2s8AUFUdqTPBKKBhyx0ZDpQ0dCfq
X8UsKLe4CIX8qZNiy7wEdFJlINy0TJ4BsRskSf0oSuEK3GOfK0rbWbdEO4YrZ9se
BxxGDjAfS6SG/1HY2JVtoJqVbKkEacp7zfGmMj0CFZUvB54krCCPOIcnQdVFHEcm
Vtjo0SNK9C2pDrQE1yfjS7pAmo3Pk+rvw4RbWI7yYKG+QpEY0V0ujK7PS0KBQU8Q
xscCrWEFwbPgk/pEsQNahjOsELhthwurZiwukMF4OcaZojUAYvZo57A9x+CwZ0pD
PoamxT3fsiR6xDSr0KtsgFeRpXqI5hHN90+ISCnIMcCgBbvC8TTb8QpgooRXBBeR
MyrNwQJX/45GuPiNFWzQpzFzrofz17Loi1hxMV/FGEsGGA0HVz2xTtrbcA0L1SV7
7Hs3FzTVjBU2x5bHEWscpml50ma3OFgTKp047/sMj9GZ6/MSUtTusfcTutxptWMp
oJii5BXsVkQC5r18ZLpJHIpb8iZv3PoDFpr89GfvfsI9jX9iMsmV3ing2BxfqFLI
Cy/JyEe+XV7z/xbPz+3l8iiqCQLWWE5c1Nfpg6fgUuBtCuftMdqBVSaC9F2jOhVk
Onrta9KBeOw3dzMcX7+qv/ODrCSPG6mKfRbpsWFRW7bbth+MQ6hzgjgnD4iuyYLC
i9fFKkn1THwDXoVsfjlO+JsWPJj7r09GsGIRRYA3+tn/3GxW7qrSAETa7/x/jjlm
r05F0loInAZnxM4nADvQoJTLJF8CB3LMRREopjulxmg6wZmFGNfb/didWzO+k5el
jt1kPC9rX45r3Y2tzFGyVNUP3/1JqolatwkpDnymafNNmJjoky3WpwqO+v1n4BSr
tPHD3Ipa3Vy63OSnRN3EqAi8Ecg+n1UCcQ7BkDe8xwyCOokBVXhh4EekdyVWW/1i
hLvQl0ODSlJozqbffHtUHYYNC4/ULtIPyFlRHEGANBdx7vlR6n0d/GLSOMnHg5Mc
gxk9aAnLuZoLd17/CqEqS/UX0zNnFEz+ep7vHplqCV7EwJWsLDsVV11fLpmwpsC0
nq5a3FIHcM+jpmTx1Vl8qEx+HUG9MBbpeVq98q5C83TpXzBAhRvJhh1zuoIiJMxK
-----END RSA PRIVATE KEY-----`
)

func setupTestFile(content string, t *testing.T) string {
	// create a temporary id_rsa file
	f, err := ioutil.TempFile("", "id_rsa")
	if err != nil {
		t.Errorf("Failed to create temp file: %s", err.Error())
	}
	fName := f.Name()

	// write the ssh key constant to it
	_, err = io.Copy(f, strings.NewReader(content))
	if err != nil {
		t.Errorf("Failed to write temp file: %s", err.Error())
		return ""
	}

	return fName
}

func TestSSHKeyValid(t *testing.T) {

	// setup a test file
	fName := setupTestFile(sshKeyWithoutPassPhrase, t)
	if fName == "" {
		return
	}

	defer os.Remove(fName)

	// key should be able to load
	key, err := Get(fName)
	if err != nil {
		t.Errorf("Get() failed unexpectedly: %s", err.Error())
	}

	if key == nil {
		t.Errorf("Get() failed to load key (it is nil)")
	}
}

func TestSSHKeyInvalidPassphrase(t *testing.T) {

	// setup a test file
	fName := setupTestFile(sshKeyWithPassPhrase, t)
	if fName == "" {
		return
	}
	defer os.Remove(fName)

	// key should not be able to load
	_, err := Get(fName)
	if err == nil {
		t.Errorf("Get() should have failed, but didn't")
	}
}

// TODO: Test reading a s non-existent ssh key (expected failure)
