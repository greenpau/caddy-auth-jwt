// Copyright 2020 Paul Greenberg (greenpau@outlook.com)

package jwt

import (
	"github.com/caddyserver/caddy/v2/caddytest"
	"io/ioutil"
	"testing"
	"time"
)

func TestPlugin(t *testing.T) {
	tester := caddytest.NewTester(t)
	// Define URL
	baseURL := "https://127.0.0.1:3443"

	// Load configuration file
	configFile := "assets/conf/config.json"
	configContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		t.Fatalf("Failed to load configuration file %s: %s", configFile, err)
	}
	rawConfig := string(configContent)

	tester.InitServer(rawConfig, "json")
	tester.AssertGetResponse(baseURL+"/version", 200, "1.0.0")
	//tester.AssertGetResponse(baseURL+"/alertmanager", 302, "Unauthorized User")
	//tester.AssertGetResponse(baseURL+"/alertmanager", 302, "Unauthorized User")

	//caddytest.AssertGetResponse(t, baseURL+"/health", 401, "")
	//caddytest.AssertGetResponse(t, baseURL+"/metrics", 401, "")

	//caddytest.AssertGetResponse(t, "http://localhost:2019/config/", 200, "xxx")

	time.Sleep(1 * time.Millisecond)
	// Uncomment the below line to perform manual testing
	// time.Sleep(6000 * time.Second)
}
