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

	logger.Infof("\n\ntest")
	test()
}

func test() {
	logger.Infof("\n\n\nstart to test dubbo")
	reqUser := pkg.User{}
	reqUser.Id = "00111111"
	reqUser.Name = "demo"
	t := time.Now()
	reqUser.Time = &t
	reqUser.Params = map[string]string{"ss": "ss"}
	reqUser.TestSet = []string{"xxxxxxx"}
	reqUser.LigoLastMsgInfo = &pkg.LigoLastMsgInfo{MessageId: 1000, Text: "ligoLastMsgInfo", MessageTime: &t}
	user, err := userProvider.GetUser(context.TODO(), &reqUser)
	if err != nil {
		panic(err)
	}
	logger.Infof("response result: %v", user)
	logger.Infof("response result: %v", user.LigoLastMsgInfo)
}
