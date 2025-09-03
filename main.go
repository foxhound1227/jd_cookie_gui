package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	fmt.Println("京东Cookie获取工具 (Go版本)")
	
	// 检查Chrome浏览器是否安装
	if !isChromeInstalled() {
		log.Fatal("未检测到Chrome浏览器，请先安装Chrome浏览器后再运行程序。\n\n下载地址：https://www.google.com/chrome/")
	}
	
	// 初始化浏览器
	fmt.Println("正在初始化浏览器...")
	l := launcher.New()
	l.Set("--disable-blink-features=AutomationControlled")
	l.Set("--disable-features=msEdgeTranslate")
	
	u, err := l.Launch()
	if err != nil {
		log.Fatal("启动浏览器失败: ", err)
	}
	
	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose()
	
	// 打开京东登录页面
	fmt.Println("正在打开京东登录页面...")
	page := browser.MustPage("https://plogin.m.jd.com/login/login?appid=300&returnurl=https%3A%2F%2Fm.jd.com%2F&source=wq_passport")
	
	// 等待用户登录
	fmt.Println("请在浏览器中完成京东登录...")
	fmt.Println("登录成功后，按回车键继续获取Cookie...")
	fmt.Scanln()
	
	// 获取Cookie
	fmt.Println("正在获取Cookie...")
	cookies := page.MustCookies()
	
	var ptPin, ptKey string
	for _, cookie := range cookies {
		if cookie.Name == "pt_pin" {
			ptPin = cookie.Value
		} else if cookie.Name == "pt_key" {
			ptKey = cookie.Value
		}
	}
	
	if ptPin != "" && ptKey != "" {
		cookieStr := fmt.Sprintf("pt_pin=%s;pt_key=%s;", ptPin, ptKey)
		fmt.Println("Cookie获取成功:")
		fmt.Println(cookieStr)
		
		// 尝试复制到剪贴板（简化处理，实际可能需要调用系统命令）
		fmt.Println("\n提示：请手动复制上面的Cookie信息")
	} else {
		fmt.Println("未找到所需的Cookie，请确保已登录京东")
	}
}

func isChromeInstalled() bool {
	// 检查常见的Chrome安装路径
	chromePaths := []string{
		"C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
		"C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe",
	}
	
	for _, path := range chromePaths {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}
	
	return false
}

// 简化的剪贴板复制功能（仅适用于Windows）
func copyToClipboard(text string) error {
	// 使用PowerShell命令复制到剪贴板
	cmd := exec.Command("powershell", "-command", "Set-Clipboard", "-Value", "\""+text+"\"")
	return cmd.Run()
}