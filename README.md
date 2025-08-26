# 京东Cookie获取工具

一个简单易用的京东Cookie获取工具，基于图形化界面，支持自动化浏览器操作获取登录Cookie。

## 功能特点

- 🔐 **自动获取Cookie**: 使用Selenium WebDriver自动打开京东登录页面
- 🔒 **安全可靠**: 本地运行，不上传任何个人信息
- 🖥️ **Windows专版**: 专为Windows系统优化，使用Microsoft Edge浏览器
- 🎯 **简单易用**: 图形化界面，一键操作
- 📦 **免安装**: 提供编译好的可执行文件
- 📋 **一键复制**: 支持Cookie一键复制到剪贴板
- 💾 **数据保存**: 支持Cookie保存和导出功能

## 快速开始

### 方法一：下载可执行文件（推荐）

1. 前往 [Releases](../../releases) 页面
2. 下载Windows版本可执行文件：`jd-cookie-gui-windows.exe`
3. 双击运行即可使用

### 方法二：从源码运行

#### 环境要求

- Python 3.8+
- Microsoft Edge浏览器
- Windows 10/11 (64位)

#### 安装依赖

```bash
pip install selenium tkinter
```

#### 下载WebDriver

**EdgeDriver配置:**
1. 查看Edge浏览器版本：在地址栏输入 `edge://version/`
2. 前往 [Microsoft Edge WebDriver](https://developer.microsoft.com/en-us/microsoft-edge/tools/webdriver/) 下载对应版本
3. 将 `msedgedriver.exe` 放在项目根目录

#### 运行程序

```bash
python jd_cookie_gui.py
```

## 使用说明

1. **启动程序**: 双击运行 `jd-cookie-gui-windows.exe` 或运行Python脚本
2. **点击开始**: 点击"开始获取Cookie"按钮
3. **自动打开浏览器**: 程序会自动启动Microsoft Edge浏览器并打开京东登录页面
4. **完成登录**: 在浏览器中输入京东账号密码完成登录
5. **自动获取**: 登录成功后程序会自动获取Cookie信息
6. **复制使用**: Cookie信息显示在文本框中，可一键复制到剪贴板

## 自动编译

本项目配置了GitHub Actions自动编译，每次创建版本标签时会自动构建Windows版本的可执行文件。

### 编译流程

- **Windows版本**: 使用EdgeDriver，打包为单文件可执行程序
- **自动发布**: 编译完成后自动发布到Releases页面，包含详细的使用说明
- **版本管理**: 通过Git标签触发发布流程，支持版本控制

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
A: 请确保已安装Microsoft Edge浏览器，程序会自动检测并使用内置的EdgeDriver。

### Q: 获取Cookie失败？
A: 请确保已完全登录京东，等待页面完全加载后程序会自动获取Cookie。

### Q: 程序被杀毒软件拦截？
A: 这是正常现象，可以将程序添加到杀毒软件白名单，或临时关闭实时保护。

### Q: 首次运行需要管理员权限？
A: 某些系统配置下可能需要管理员权限来访问浏览器，请右键以管理员身份运行。

## 更新日志

### v1.0.1
- 专注Windows平台优化
- 改进用户界面体验
- 增强Cookie获取稳定性
- 添加详细使用说明

### v1.0.0
- 初始版本发布
- 图形化界面
- 自动化Cookie提取功能

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 贡献

欢迎提交Issue和Pull Request来改进这个项目！

---

**免责声明**: 本工具仅供学习交流使用，使用者需自行承担使用风险。开发者不对因使用本工具而产生的任何问题负责。