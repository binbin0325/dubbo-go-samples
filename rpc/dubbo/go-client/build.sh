
#!/bin/bash
# 使用go.mod时build.sh的标准模板
#
# go.mod 要求golang 版本大于等于1.11
#
# build.sh中可使用的golang版本为: 1.11.2 / 1.12.6 / 1.13.3(默认使用此版本)
#
# go.mod编译模板中默认使用go1.13.3：
export GOROOT=/usr/local/go
#
# 如果需要使用 1.12.6  请修改GOROOT：
# export GOROOT=/usr/local/go1.12.6
#
# 如果需要使用1.11.2  请修改GOROOT：
# export GOROOT=/usr/local/go1.11.2
#
export PATH=$GOROOT/bin:$PATH
#
set -e  #Exit the script if an error happens
#set -x  #执行指令后,会显示该指令及参数,可加可不加该行
workspace=$(cd $(dirname $0) && pwd -P)
#
# 打开module开关
export GO111MODULE=on
#
# 设置module proxy为公司内部公共代理地址
#
export GOPROXY=https://goproxy.io
#
# 公司内的module无法通过sum.golang.org的md5校验, 所以这里关闭了SUMDB校验
#
export GOSUMDB=off
#
# 请修改module为你期望生成的可执行文件名字
module="go-client"


if [ "$module" == "" ] ; then
    echo "===== please uncomment variable 'module' ====="
    exit 1
fi

go build -o $module ./cmd   #编译目标文件
ret=$?
if [ $ret -ne 0 ];then
    echo "===== $module build failure ====="
    exit $ret
else
    echo -n "===== $module build successfully! ====="
fi

output="output"
rm -rf $output
mkdir -p $output
mkdir -p $output/conf

# 填充output目录, output的内容即为待部署内容
    (
         # 拷贝部署脚本control.sh至output目录
        cp -r conf/* ${output}/conf/ &&
        mv ${module} ${output}/ &&        # 移动需要部署的文件到output目录下
        echo -e "===== Generate output ok ====="
    ) || { echo -e "===== Generate output failure ====="; exit 2; } # 填充output目录失败后, 退出码为 非0