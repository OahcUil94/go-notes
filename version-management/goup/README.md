# goup

- 记录时间: 2020-10-26
- 版本信息: v0.1.6

[github地址](https://github.com/owenthereal/goup)

## goup安装

```
$ ping raw.githubusercontent.com
Ping 请求找不到主机 raw.githubusercontent.com。请检查该名称，然后重试。
```

- 访问https://cdn.jsdelivr.net/gh/用户名/仓库名/, 最后加斜杠, 可以看到jsdelivr列出的仓库的信息
- 选择对应的版本例如@0.12.0+要下载的文件路径: https://cdn.jsdelivr.net/gh/coreos/flannel@0.12.0/Documentation/kube-flannel.yml

https://cdn.jsdelivr.net/gh/owenthereal/goup@master/install.sh

Welcome to Goup!                                                                            
                                                                                            
Goup and Go will be located at:                                                             
                                                                                            
/home/vagrant/.go
                                                                                            
The Goup command will be located at:                                                        
                                                                                            
  /home/vagrant/.go/bin                                                                     
                                                                                            
The go, gofmt and other Go commands will be located at:                                     
                                                                                            
  /home/vagrant/.go/current/bin                                                             
                                                                                            
To get started you need Goup's bin directory (/home/vagrant/.go/bin) and                    
Go's bin directory (/home/vagrant/.go/current/bin) in your PATH environment                 
variable. These two paths will be added to your PATH environment variable by                
modifying the profile files located at:                                                     
                                                                                            
  /home/vagrant/.profile                                                                    
  /home/vagrant/.zprofile                                                                   
  /home/vagrant/.bash_profile                                                               
                                                                                            
Next time you log in this will be done automatically. To configure your
current shell run source $HOME/.go/env.

Error: Getting current Go version failed: Get "https://golang.org/VERSION?m=text": dial tcp 216.239.37.1:443: i/o timeout
Usage:
  goup init [flags]

Flags:
  -h, --help           help for init
      --skip-install   Skip installing Go
      --skip-prompt    Skip confirmation prompt

Global Flags:
  -v, --verbose   Verbose

FATA[0083] Getting current Go version failed: Get "https://golang.org/VERSION?m=text": dial tcp 216.239.37.1:443: i/o timeout

https://golang.google.cn/VERSION?m=text


export GOUP_GO_HOST=golang.google.cn; curl -sSf https://raw.githubusercontent.com/owenthereal/goup/master/install.sh | sh