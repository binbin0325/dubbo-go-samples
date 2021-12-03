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

package org.apache.dubbo;

import java.text.SimpleDateFormat;
import java.time.LocalDate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Date;
import java.util.HashMap;
import java.util.List;

import com.alibaba.dubbo.rpc.service.EchoService;
import org.apache.dubbo.common.constants.CommonConstants;
import org.apache.dubbo.config.ApplicationConfig;
import org.apache.dubbo.config.ReferenceConfig;
import org.apache.dubbo.config.RegistryConfig;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Consumer {
    // Define a private variable (Required in Spring)
    private static UserProvider userProvider;

    public static void main(String[] args) throws Exception {
        ClassPathXmlApplicationContext context = new ClassPathXmlApplicationContext(new String[]{"META-INF/spring/dubbo.consumer.xml"});
        userProvider = (UserProvider)context.getBean("userProvider");

        start();
    }

    // Start the entry function for consumer (Specified in the configuration file)
    public static void start() throws Exception {
        System.out.println("\n\ntest");
        testGetUser();
    }

    private static void testGetUser() throws Exception {
        try {
            User user=new User();
            user.setId(400000);
            User user1 = userProvider.GetUser(user);
            System.out.println("[" + new SimpleDateFormat("HH:mm:ss").format(new Date()) + "] " +
                    " UserInfo, ID:" + user1.getId());

            System.out.println("GetUser succ");
        } catch (Throwable e) {
            System.out.println("*************exception***********");
            e.printStackTrace();
        }

    }



}
