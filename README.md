# Hebust新版联通校园网认证

# 如何使用
- 下载最新版程序 https://github.com/Olixn/Portal-Auto-Auth/releases
- 解压，上传至openwrt路由器（例如/tmp目录）
- 进入`/tmp/server`目录执行`chmod +x main`
- 配置`config.yaml`
- 添加开机启动项  
  - `/tmp/server/main /tmp/server/config.yaml`
- 添加crontab计划任务  
  - `crontab -e`
  - `5 0 * * * rm -rf /tmp/campus_run.log`
- 查看运行日志 `cat /tmp/campus_run.log`

# 二次开发
- `git clone https://github.com/Olixn/Portal-Auto-Auth.git`
- `go mod tidy`
- 执行`build.bat`生成openwrt下的linux可执行文件
- 执行`go build main.go`生成exe文件

# 完善项目
> 期待您的补充
- `fork`本仓库至你的GitHub
- `clone`你的仓库进行开发
- `Pull requests`你的代码


# 其他
使用问题欢迎issues  
程序经过压缩编译后大小有5兆左右，请确保您的路由器有足够的空间。  
本项目仅为了方便使用、学习和交流，不用于其他用途。  
二次修改发布请保留作者版权信息！  
如果您觉得此项目对您有所帮助，请给作者一个star!
