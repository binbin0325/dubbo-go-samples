/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"net/http"
	"strconv"
	"time"

	hessian "github.com/apache/dubbo-go-hessian2"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"

	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"github.com/apache/dubbo-go-samples/rpc/dubbo/go-client/pkg"
)

var (
	userProvider = &pkg.UserProvider{}
)

func init() {
	hessian.RegisterJavaEnum(pkg.Gender(pkg.MAN))
	hessian.RegisterJavaEnum(pkg.Gender(pkg.WOMAN))
	hessian.RegisterPOJO(&pkg.User{})
	hessian.RegisterPOJO(&pkg.LigoLastMsgInfo{})

	config.SetConsumerService(userProvider)
}

// need to setup environment variable "DUBBO_GO_CONFIG_PATH" to "conf/dubbogo.yml" before run
func main() {

	err := config.Load()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	hintCode := r.Header.Get("didi-trace.sys.didi-header-hint-code")
	if hintCode == "" {
		hintCode = "-100"
	}
	i, err := strconv.ParseInt(hintCode, 10, 32)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	test(r.Header.Get("didi-trace.sys.didi-header-rid"), int32(i))
}

func test(traceId string, hintCode int32) {
	logger.Infof("\n\n\nstart to test dubbo")
	reqUser := pkg.User{}
	reqUser.Id = "00111111"
	reqUser.Name = "demo"
	t := time.Now()
	reqUser.Time = &t
	reqUser.Params = map[string]string{"ss": "ss"}
	reqUser.TestSet = []string{"xxxxxxx"}
	reqUser.LigoLastMsgInfo = &pkg.LigoLastMsgInfo{MessageId: 1000, Text: "ligoLastMsgInfo", MessageTime: &t}
	atta := make(map[string]interface{})
	atta["didi-trace.sys.didi-header-rid"] = traceId
	if hintCode != -100 {
		atta["didi-trace.sys.didi-header-hint-code"] = hintCode
	}
	reqContext := context.WithValue(context.Background(), constant.DubboCtxKey("attachment"), atta)
	user, err := userProvider.GetUser(reqContext, &reqUser)
	if err != nil {
		panic(err)
	}
	logger.Infof("result:", user)
}
