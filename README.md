# 京东Cookie提取工具

一个简单易用的京东Cookie提取工具，支持自动化浏览器操作获取登录Cookie。

## 功能特点

- 🚀 **自动化操作**: 使用Selenium WebDriver自动打开京东登录页面
- 🔒 **安全可靠**: 本地运行，不上传任何个人信息
- 💻 **跨平台支持**: 支持Windows和Linux系统
- 🎯 **简单易用**: 图形化界面，一键操作
- 📦 **免安装**: 提供编译好的可执行文件

## 快速开始

### 方法一：下载可执行文件（推荐）

1. 前往 [Releases](../../releases) 页面
2. 下载适合您系统的可执行文件：
   - Windows: `jd-cookie-gui-windows.exe`
   - Linux: `jd-cookie-gui-linux`
3. 双击运行即可使用

### 方法二：从源码运行

#### 环境要求

- Python 3.8+
- Microsoft Edge浏览器（Windows）或Google Chrome浏览器（Linux）

#### 安装依赖

```bash
pip install selenium tkinter
```

#### 下载WebDriver

**Windows (EdgeDriver):**
1. 查看Edge浏览器版本：在地址栏输入 `edge://version/`
2. 前往 [Microsoft Edge WebDriver](https://developer.microsoft.com/en-us/microsoft-edge/tools/webdriver/) 下载对应版本
3. 将 `msedgedriver.exe` 放在项目根目录

**Linux (ChromeDriver):**
1. 查看Chrome版本：`google-chrome --version`
2. 前往 [ChromeDriver](https://chromedriver.chromium.org/) 下载对应版本
3. 将 `chromedriver` 放在项目根目录并添加执行权限

#### 运行程序

```bash
python jd_cookie_gui.py
```

## 使用说明

1. **启动程序**: 运行可执行文件或Python脚本
2. **初始化浏览器**: 程序会自动检测并启动浏览器
3. **登录京东**: 在打开的浏览器窗口中完成京东登录
4. **获取Cookie**: 登录成功后，点击"获取Cookie"按钮
5. **复制使用**: Cookie信息会显示在文本框中，可直接复制使用

## 自动编译

本项目配置了GitHub Actions自动编译，每次推送代码到main分支时会自动构建Windows和Linux版本的可执行文件。

### 编译流程

- **Windows版本**: 使用EdgeDriver，打包为单文件可执行程序
- **Linux版本**: 使用ChromeDriver，打包为单文件可执行程序
- **自动发布**: 编译完成的文件会作为Artifacts上传，可在Actions页面下载

## 技术栈

- **GUI框架**: Tkinter
- **浏览器自动化**: Selenium WebDriver
- **打包工具**: PyInstaller
- **CI/CD**: GitHub Actions

## 注意事项

⚠️ **重要提醒**:
- 本工具仅用于个人学习和研究目的
- 请遵守京东网站的使用条款和相关法律法规
- 不要将获取的Cookie用于任何违法违规活动
- 建议定期更新Cookie以确保有效性

## 常见问题

### Q: 程序启动后浏览器无法打开？
A: 请确保已安装对应的浏览器和WebDriver，并且版本匹配。

### Q: 获取Cookie失败？
A: 请确保已完全登录京东，并且页面加载完成后再点击获取Cookie。

### Q: Windows版本提示找不到EdgeDriver？
A: 请下载与您Edge浏览器版本匹配的EdgeDriver，并放在程序同目录下。

### Q: Linux版本无法运行？
A: 请确保已安装Chrome浏览器和对应版本的ChromeDriver，并给予执行权限。

## 更新日志

### v1.0.0
- 初始版本发布
- 支持Windows和Linux平台
- 图形化界面
- 自动化Cookie提取功能

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 贡献

欢迎提交Issue和Pull Request来改进这个项目！

---

**免责声明**: 本工具仅供学习交流使用，使用者需自行承担使用风险。开发者不对因使用本工具而产生的任何问题负责。