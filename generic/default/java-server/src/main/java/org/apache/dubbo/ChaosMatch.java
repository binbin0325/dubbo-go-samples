package org.apache.dubbo;

import java.util.Map;
import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicInteger;

//1:java故障注入时 下发故障到待熔断客户端，对指定的类和方法进行熔断
//2:同时对待熔断客户端入口方法，拦截使用自定义脚本注入方式,选择方法执行后注入--after=true。例如dubbo的com.alibaba.dubbo.rpc.proxy.AbstractProxyInvoker的invoker
//3:拦截内容获取参数进行匹配，看dubbo传递过来的参数 invocation.getArguments().get("放火标识")是否有放火标识
//4:如果有放火标识则开启线程（or线程池）将当前请求结果传递到放火平台接口，当前dubbo请求正常返回
//5:放火平台根据传递过来的返回结果，进行预期的匹配，得到强弱依赖。

public class ChaosMatch {

    public Object run(Map<String, Object> params) {
        org.apache.dubbo.User user = (org.apache.dubbo.User)params.get("return");

        //获取放火标识
     /*   if(1!=params.get("hintCode")){
            return user;
        }*/

        new Thread(()->{
            //call chaos api -根据框架使用的http client库编写发送请求代码
            // HttpClient client=new HttpClient("chaos.url")
            // client.request(user)
        });
        return user;
    }
}
