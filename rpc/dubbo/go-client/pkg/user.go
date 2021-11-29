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

package pkg

import (
	"context"
	"strconv"
	"time"

	hessian "github.com/apache/dubbo-go-hessian2"
)

type Gender hessian.JavaEnum

const (
	MAN hessian.JavaEnum = iota
	WOMAN
)

var genderName = map[hessian.JavaEnum]string{
	MAN:   "MAN",
	WOMAN: "WOMAN",
}

var genderValue = map[string]hessian.JavaEnum{
	"MAN":   MAN,
	"WOMAN": WOMAN,
}

func (g Gender) JavaClassName() string {
	return "org.apache.dubbo.Gender"
}

func (g Gender) String() string {
	s, ok := genderName[hessian.JavaEnum(g)]
	if ok {
		return s
	}

	return strconv.Itoa(int(g))
}

func (g Gender) EnumValue(s string) hessian.JavaEnum {
	v, ok := genderValue[s]
	if ok {
		return v
	}

	return hessian.InvalidJavaEnum
}

type User struct {
	// !!! Cannot define lowercase names of variable
	LigoLastMsgInfo *LigoLastMsgInfo
	Id              string
	Name            string
	Age             int32
	Time            *time.Time
	Params          map[string]string
	TestSet         []string
}

func (u User) JavaClassName() string {
	return "org.apache.dubbo.User"
}

type LigoLastMsgInfo struct {
	MessageId   int32
	Text        string
	MessageTime *time.Time
}

func (ligoLastMsgInfo LigoLastMsgInfo) JavaClassName() string {
	return "org.apache.dubbo.LigoLastMsgInfo"
}

type UserProvider struct {
	GetUser func(ctx context.Context, u *User) (*User, error) `dubbo:"GetUser"`
}
