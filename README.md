# Shadowsocks-VPN-Quick-Installer


This tool is for personal usage only. Never use it for commercial and illegal purposes. I, the author of this tool, am not responsible for any consequencies caused by using the tool, since the tool is for quick proxy configuration only. I do not suggest using it in countries that don't allow VPN; however, you are the one who make the choice. The tool is originally for myself and my colleagues only. Its purpose is to let us accessing Golang developer forum and other academic websites with ease, and now I decide to share it with people with similar needs but lack the knowledge to configure VPNs. I have done a few "improvements" so some basic configurations may be customized. 
The tool is open-sourced, and is a product of self-practicing. Please feel free to submit any bug you see. 


2018.6.25 Sispuella

此工具仅限个人使用，请不要将其用于商业、或违法用途。本人不为使用本程序可能会造成的任何风险负责。程序本身只是实现快速设置shadowsocks代理服务端，本人对使用者的使用场景没有任何了解，也不建议在禁止使用VPN的国家使用本程序进行灰色活动。工具本身无罪，更何况本工具只实现了一个脚本的功能。本人不提供任何服务器托管，工具也不支持该功能。该工具最初是为了方便作者本人与同事访问Golang开发者社区以及各类其它学术网站，现分享出来以供其他有类似需求、但没有相关知识的朋友使用。


该工具完全开源，并且是个人初学Golang练习的产物；只是每次节点被封，顶着高延迟在终端配置shadowsocks文件太过烦躁，问了问周围的高手同事们，也都有着这样的烦恼，但都没抽出时间去编写解决方案。如果发现bug或者错误，欢迎发邮件至kliest.yang@smus.ca纠正指出。

如果有需求，也许会再做一个同时支持https协议与socks5协议的代理服务端以便支持部分源（比如go get），不过那是后话了。

2018.6.25 长得萌二


使用指南：
请自行租借服务器，digitalOcean的VPS也就5美金一个月，按小时收费，速度还是非常棒的。
租借服务器时请选择CentOs系统，因为它默认安装了yum所以可以省去很多麻烦。
1. yum install git
2. git clone https://github.com/sispuella/Shadowsocks-VPN-Quick-Installer.git
3. cd Shadowsocks-VPN-Quick-Installer/Release
4. chmod 777 run
5. ./run
运行完以上指令，并且按照提示输入完毕后，如果没有出错，那么shadosocks就已经被配置完成了。（注：server port和local port 必须为数字，推荐设成8388与1080，但具体是什么数字其实无所谓，如果没有相关知识就按照8388和1080设置吧。）
输入ssserver -c /etc/shadowsocks.json -d start 启动。
此时可以通过shadowsocks客户端连接代理，具体的设置方法很简单，请自行了解。配置参数请参考自己设置的值。然后就可以上google查阅学术资料啦！
节点被封的状况，可以直接在digitalocean上销毁droplet，重新建一个并按此方法再配置一遍就好了。
