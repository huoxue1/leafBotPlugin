### 插件列表

+ #### 黑名单管理
```go
import 	_ "github.com/huoxue1/leafBotPlugin/pluginBlackList"
```
使用该插件需要在config目录下配置plugin_list.json文件，该插件会自动读取文件


+ #### 每日一图
```go
import 	_ "github.com/huoxue1/leafBotPlugin/pluginDayImage"
```
基于Bing的每日一图API所编写的每日一图插件

+ #### 闪照拦截
```go
import 	_ "github.com/huoxue1/leafBotPlugin/pluginFlashImage"
```
该插件会自动拦截所有群或者私聊的闪照，并将其发送至管理员用户

