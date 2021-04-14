# cloudStore
毕业设计 : 一个简单的云存储系统和文件审计系统，基于golang开发
参考文献:
hyperledger fabric v1.1: https://hyperledger-fabric.readthedocs.io/en/release-1.1/getting_started.html#
native 方式部署fabric 网络： https://blog.csdn.net/u013938484/article/details/79867992
gin : https://github.com/gin-gonic/gin
gorm : https://gorm.io/zh_CN/docs/
bootstrap: https://v4.bootcss.com/

部署在云端: 
aliyun es6 供应 fabricServer
tencentyun es5 供应 webServer
fabric编译chaincode 运行setkv 程序,单peer节点，单channel
