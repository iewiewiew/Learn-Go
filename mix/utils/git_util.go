package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

/**
 * @author       weimenghua
 * @time         2023-07-12 10:44
 * @description  Git 操作
 */

// 克隆仓库
func git_clone() {
	url := "git@gitee.com:iewiewiew/wei-demo-001.git"
	path := "/Users/menghuawei/IdeaProjects/my-project/Learn-Go/.tmp/repo/wei-demo-001"

	cmd := exec.Command("git", "clone", url, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Git repository successfully!")
}

// 推送代码
func git_push() {
	repoPath := "/Users/menghuawei/IdeaProjects/my-project/Learn-Go/.tmp/repo/wei-demo-001"

	//os.O_WRONLY: 以只写模式打开文件。只允许写入文件，不允许读取文件内容。
	//os.O_CREATE: 如果文件不存在，则创建文件。如果文件已经存在，则不执行任何操作。
	//os.O_TRUNC: 如果文件存在，并且打开模式为写入模式，则截断文件内容为零长度。
	//0644 是文件权限，它指定了新创建的文件的权限。在这个例子中，它表示所有者具有读和写的权限，而其他用户只具有读的权限。
	file, err := os.OpenFile(repoPath+"/file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(time.Now().String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file wirte suceessfully!")

	err = os.Chdir(repoPath)
	if err != nil {
		log.Fatal(err)
	}

	commitCmd := exec.Command("git", "commit", "-am", "test")
	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr
	err = commitCmd.Run()
	if err != err {
		log.Fatal(err)
	}
	fmt.Println("Git commit successfully!")

	pushCmd := exec.Command("git", "push")

	err = pushCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Git push successfully!")
}
