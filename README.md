# elasticsearch_go

# 安装及使用Elasticsearch
- 【安装】
    - 官网下载：https://www.elastic.co/cn/downloads/elasticsearch
    - 启动：安装目录下，./bin/elasticsearch.exe
- 【使用】
    - 参考：https://www.infoq.cn/article/hvzmnkuyymckrtk-ozdp
- 【问题】
    - 提示：Elasticsearch is already running as a service and currently: Running.（进程管理器手动关闭es进程即可）

# 部署及使用Logstash
- cd bin
- logstash.bat -f  ..\config\logstash-jdbc-book.conf
- 参考：https://www.yuque.com/sourlemon/java/nk77gg
- 问题：
    - ERROR 2003 (HY000): Can't connect to MySQL server on 'localhost' (10061)，本地没有启动mysql服务

# 部署及使用Kibana
- 参考：https://www.yuque.com/sourlemon/java/buwl0q
- 启动：cd bin，cmd 下 kibana.bat
- 配置文件说明参考：https://www.lyile.cn/articles/2021/02/18/1613634074333.html
