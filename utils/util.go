package utils

import (
	"fmt"
	"github.com/nguyenthenguyen/docx"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
)

// ConvertDocxToPDF 转换 DOCX 文件为 PDF
func ConvertDocxToPDF(inputFile string, outputFile string) error {
	// 判断当前运行环境
	if runtime.GOOS == "windows" {
		// Windows 环境，使用 LibreOffice 命令
		// 默认安装路径
		// cmd := exec.Command(
		// 	"C:\\Program Files\\LibreOffice\\program\\soffice.exe", // LibreOffice
		// 	"--headless",          // 无界面模式
		// 	"--convert-to", "pdf", // 转换为 PDF 格式
		// 	"--outdir", filepath.Dir(outputFile), // 设置输出目录
		// 	inputFile,
		// )
		// return cmd.Run()

		// Windows 环境，使用 LibreOffice 命令
		// 自定义安装路径
		// 命令行说明 https://help.libreoffice.org/24.8/zh-CN/text/shared/guide/start_parameters.html?&DbPAR=SHARED&System=WIN
		cmd := exec.Command(
			"D:\\GoSoft\\LibreOffice 24.8.4.2\\program\\soffice.exe", // LibreOffice 安装路径
			"--headless",          // 无界面模式
			"--norestore",         // 禁止恢复上次会话
			"--convert-to", "pdf", // 转换为 PDF 格式
			"--outdir", filepath.Dir(outputFile), // 设置输出目录
			inputFile,
		)
		return cmd.Run()
	} else if runtime.GOOS == "darwin" {
		// macOS 环境，使用 LibreOffice 命令
		cmd := exec.Command(
			"/Applications/LibreOffice.app/Contents/MacOS/soffice", // LibreOffice 在 macOS 下的默认路径
			"--headless",          // 无界面模式
			"--convert-to", "pdf", // 转换为 PDF 格式
			"--outdir", filepath.Dir(outputFile), // 设置输出目录
			inputFile,
		)
		return cmd.Run()

	} else if runtime.GOOS == "linux" {
		// Linux 环境，首先尝试使用 LibreOffice（soffice），然后使用 unoconv 作为备选
		// 尝试使用 LibreOffice（soffice）
		cmd := exec.Command(
			"soffice",             // 使用 libreoffice 的 soffice 命令
			"--headless",          // 无界面模式
			"--convert-to", "pdf", // 转换为 PDF 格式
			"--outdir", filepath.Dir(outputFile), // 设置输出目录
			inputFile,
		)
		err := cmd.Run()
		if err == nil {
			return nil // 使用 LibreOffice 成功转换，直接返回
		}

		// 如果 LibreOffice 失败，尝试使用 unoconv
		cmd = exec.Command("unoconv", "-f", "pdf", "-o", outputFile, inputFile)
		return cmd.Run()
	}

	// 如果不支持的操作系统
	return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
}

// 保存修改后的 DOCX 文件
func SaveDoc(doc *docx.Docx, filePath string) error {
	if err := doc.WriteToFile(filePath); err != nil {
		return fmt.Errorf("保存 DOCX 文件失败: %v", err)
	}
	fmt.Println("成功生成word文件！")
	return nil
}

// 使用 ImageMagick 将 PDF 文件转换为图像
func ConvertPDFToImage(pdfFilePath, imageFilePath string) error {
	// 调用 ImageMagick 命令行工具进行转换
	cmd := exec.Command("magick", pdfFilePath, "-density", "600", "-resize", "3840x2160", imageFilePath)
	// 获取命令的标准输出和错误输出
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("PDF 转换为图像失败: %v. 输出: %s", err, string(cmdOutput))
	}
	fmt.Println("PDF 成功转换为图像！")
	return nil
}

// 填充模板的主逻辑
func FillTemplate[T any](templatePath, newFilePath string, data *T) error {
	// 打开并处理 DOCX 模板
	r, err := docx.ReadDocxFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to open template: %v", err)
	}
	defer r.Close()
	doc := r.Editable()
	// fmt.Println(doc) 如果模板渲染出现问题,可以看看读到的模板数据是不是正常的
	// 填充模板中的占位符
	if err := ReplacePlaceholders(doc, data); err != nil {
		return err
	}
	if err := SaveDoc(doc, newFilePath); err != nil {
		return err
	}
	return err
}

// 替换 DOCX 模板中的占位符
func ReplacePlaceholders[T any](doc *docx.Docx, data *T) error {
	// 使用反射遍历结构体字段
	val := reflect.ValueOf(data).Elem()
	typ := val.Type()

	// 替换模板中的占位符
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("json")
		if fieldName == "" {
			continue
		}
		placeholder := "{{" + fieldName + "}}"
		// 将占位符替换为字段值
		if err := doc.Replace(placeholder, field.String(), -1); err != nil {
			return fmt.Errorf("替换占位符 %s 失败: %v", placeholder, err)
		}
	}
	fmt.Println("成功替换占位符！")
	return nil
}
