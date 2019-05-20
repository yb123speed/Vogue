# Vogue API

web框架用的gin，日志用的zap，数据库连接用的mysql driver，配置文件读取用的是viper。

main.go作为整个api server 的入口文件，tool文件夹下面存放了日志，数据库，配置文件相关代码。app文件夹里面是api实现相关的代码。