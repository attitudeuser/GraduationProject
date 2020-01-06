# Git学习手册

## 一、安装Git

### 1. 在linux上安装Git

### 2. 在Mac OS X上安装Git

### 3.在Windows上安装Git

​	Git官网上直接下载安装程序，安装完成后在桌面右击菜单中可以找到"Git Bash"

## 二、使用

### 1. 自报家门

```shell
git config --global user.name "Your Name" # Your Name = 作者标识
git config --global user.email "email@example.com" # 账户绑定的邮箱账号
```

### 2.创建版本库

1. 创建一个版本库（空目录）

```powershell
mkdir learngit  # 创建空目录
cd learngit  # 进入learngit目录
pwd # 显示当前目录路径
```

![1578275136314](C:\Users\HP\AppData\Roaming\Typora\typora-user-images\1578275136314.png)



2. 初始化（本地）仓库

```  shell
git init   # 把目录变成Git可以管理的仓库（见上图）
```

3. 把文件添加到版本库

```shell
git add readme.txt  # 可多次添加，一次提交
git commit -m "文档说明"  # -m参数用于提交说明信息
```

### 3. 时光机穿梭

1. 版本回退

```shell
git log  # 查看历史版本
git log --pretty=oneline  # 简洁显示
git reset --hard HEAD^ # 回退一次
git reset --hard HEAD^^ # 回退两次
git reset --hard [版本id前若干位]  # 切换到指定版本
# --hard 代表指针
# 关机后可用git reflog查看使用过的命令，查找目标id
```

2. 工作区和暂存区

   工作区就是电脑里的目录，比如learngit文件夹就是一个工作区。

   版本库是工作区里的一个隐藏目录.git，版本库中有很多东西，其中最重要的就是成为stage（或者index）的暂存区，还有Git为我们自动创建的第一个分支master，以及指向master的一个指针叫HEAD

   ![1578276909753](C:\Users\HP\AppData\Roaming\Typora\typora-user-images\1578276909753.png)

   我们把文件往Git版本库中添加的时候，是分为两步执行的：

   第一步是用git add提交更改，实际上是把文件修改添加到暂存区；

   第二步是用git commit提交更改，实际上是把暂存区的所有内容提交到当前分支。

3. 删除文件

   ```shell
   git rm test.txt  # 删除（暂存区中）文件
   git commit -m "remove test.txt"  # 提交到版本库
   git checkout -- test.txt  # 把误删的文件恢复到最新版本
   ```

### 4. 远程仓库

1. 添加远程仓库

   - 登陆码云，创建仓库

   ![1578293489595](C:\Users\HP\AppData\Roaming\Typora\typora-user-images\1578293489595.png)

   - 连接远程库

     ```shell
     git remote add origin https://gitee.com/ChenFlyU/test
     ```

     origin绑定到远程库，这是默认叫法。

   - 推送到远程库

     ```shell
     git push -u origin master
     ```

     把当前分支master推送到远程，由于远程库是空的，第一次推送master分支时，加上了-u参数，Git不但会把本地的master分支内容推送到远程新的master分支，还会把本地的master分支和远程的master分支关联起来，在以后的推送或者拉取时就可以简化命令。

     此后的提交就可以用

     ```shell
     git push origin master
     ```

     

2. 从远程库克隆

   - 先新建一个仓库（使用Readme文件初始化）

     ![1578294311537](C:\Users\HP\AppData\Roaming\Typora\typora-user-images\1578294311537.png)

   - 远程库已经准备好，利用命令克隆一个本地库

     ```shell
     git clone https://gitee.com/ChenFlyU/gitskills.git
     ```

     ![1578294576510](C:\Users\HP\AppData\Roaming\Typora\typora-user-images\1578294576510.png)

     ![1578294678463](C:\Users\HP\AppData\Roaming\Typora\typora-user-images\1578294678463.png)

