# JD Cookie GUI

一个用于获取京东Cookie的图形界面工具，基于Python和Selenium开发。

## 功能特点

- 🖥️ 简洁直观的图形用户界面
- 🔄 自动检测Microsoft Edge浏览器
- 🍪 一键获取京东登录Cookie
- 📋 自动复制Cookie到剪贴板
- 🚀 支持打包为独立可执行文件
- ⚡ 精简优化的代码结构

## 系统要求

- Windows 10/11 64位系统
- Microsoft Edge浏览器
- EdgeDriver（已包含在项目中）

## 使用方法

### 方式一：直接运行可执行文件（推荐）

1. 下载 `dist/jd_cookie_gui.exe`
2. 双击运行程序
3. 点击"获取Cookie"按钮
4. 在弹出的浏览器中登录京东账号
5. 登录成功后，Cookie会自动显示并复制到剪贴板

### 方式二：运行Python源码

1. 安装Python 3.7+
2. 安装依赖包：
   ```bash
   pip install selenium tkinter
   ```
3. 运行程序：
   ```bash
   python jd_cookie_gui.py
   ```

## 项目结构

```
jd_cookie_gui/
├── jd_cookie_gui.py      # 主程序源码
├── jd_cookie_gui.spec    # PyInstaller打包配置
├── msedgedriver.exe      # Edge WebDriver
├── dist/                 # 打包后的可执行文件
│   └── jd_cookie_gui.exe
├── build/                # 构建临时文件
├── 使用说明.txt          # 详细使用说明
└── README.md             # 项目说明文档
```

## 技术栈

- **Python 3.x** - 主要开发语言
- **Tkinter** - GUI界面框架
- **Selenium** - 浏览器自动化
- **PyInstaller** - 打包工具
- **Microsoft Edge WebDriver** - 浏览器驱动

## 开发特色

### 代码优化
- 移除了复杂的驱动下载逻辑，简化部署
- 精简错误处理机制，提高程序稳定性
- 优化代码结构，从366行精简至228行
- 清理不必要的依赖和导入

### 用户体验
- 自动检测浏览器和驱动状态
- 友好的错误提示信息
- 一键复制功能，方便使用
- 独立可执行文件，无需安装Python环境

## 注意事项

1. 首次使用需要手动登录京东账号
2. 确保Microsoft Edge浏览器已安装
3. 如果遇到驱动问题，请检查Edge版本兼容性
4. Cookie有时效性，建议定期重新获取

## 故障排除

### 常见问题

**Q: 提示"未找到msedgedriver.exe"**
A: 确保msedgedriver.exe文件与程序在同一目录下

**Q: 浏览器无法启动**
A: 检查Microsoft Edge是否正确安装，版本是否兼容

**Q: Cookie获取失败**
A: 确保已成功登录京东账号，刷新页面后重试

## 版本信息

- **当前版本**: v1.0
- **最后更新**: 2025年1月
- **兼容性**: Windows 10/11, Microsoft Edge

## 许可证

本项目仅供学习和个人使用，请遵守相关法律法规和网站服务条款。

## 贡献

欢迎提交Issue和Pull Request来改进这个项目。

---

**免责声明**: 本工具仅用于技术学习和个人使用，使用者需自行承担使用风险，开发者不承担任何法律责任。