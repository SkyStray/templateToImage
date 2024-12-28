package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"templateToImage/entity"
	"templateToImage/utils"
)

func main() {
	templatePath := "template/template.docx"
	data := &entity.Data{
		Title:   "三好学生",
		Name:    "张三三",
		Content: "随着中国特色社会主义进入新时代，新时代教育服务功能也发生了新变化，教育特别是高等教育要为人民服务，为中国共产党治国理政服务，为巩固和发展中国特色社会主义制度服务，为改革开放和社会主义现代化建设服务，赋予劳动教育新的使命和内涵。2018年9月，习近平总书记在全国教育大会上指出：“要在学生中弘扬劳动精神，教育引导学生崇尚劳动、尊重劳动，懂得劳动最光荣、劳动最崇高、劳动最伟大、劳动最美丽的道理，长大后能够辛勤劳动、诚实劳动、创造性劳动”，并在阐释教育目标时首次完整提出“培养德智体美劳全面发展的社会主义建设者和接班人”，进一步突显了劳动教育在新时代教育体系中的重要地位，推动新时代劳动教育回归初心、回归育人。",
		Text:    "随着中国特色社会主义进入新时代，新时代教育服务功能也发生了新变化，教育特别是高等教育要为人民服务，为中国共产党治国理政服务，为巩固和发展中国特色社会主义制度服务，为改革开放和社会主义现代化建设服务，赋予劳动教育新的使命和内涵。2018年9月，习近平总书记在全国教育大会上指出：“要在学生中弘扬劳动精神，教育引导学生崇尚劳动、尊重劳动，懂得劳动最光荣、劳动最崇高、劳动最伟大、劳动最美丽的道理，长大后能够辛勤劳动、诚实劳动、创造性劳动”，并在阐释教育目标时首次完整提出“培养德智体美劳全面发展的社会主义建设者和接班人”，进一步突显了劳动教育在新时代教育体系中的重要地位，推动新时代劳动教育回归初心、回归育人。",
		Year:    "2024",
		Month:   "12",
		Day:     "25",
		No:      uuid.NewV1().String(),
	}

	fmt.Println("开始填充模板并生成文件!!!")
	// 填充模板并生成文件
	// 保存修改后的 DOCX 文件
	docFilePath := "wordfile/wordOutFile.docx"
	err := utils.FillTemplate(templatePath, docFilePath, data)
	if err != nil {
		log.Fatalf("填充模板时出错: %v", err)
	}
	fmt.Println("文档生成成功,docFilePath:", docFilePath)

	// 将 DOCX 转换为 PDF
	pdfFilePath := "pdffile/wordOutFile.pdf"
	if err := utils.ConvertDocxToPDF(docFilePath, pdfFilePath); err != nil {
		log.Fatalf("DOCX 转换为 PDF时出错: %v", err)
	}
	fmt.Println("将 DOCX 转换为 PDF成功,pdfFilePath:", pdfFilePath)

	// 将 PDF 转换为图像
	imageFilePath := "images/wordOutFile.jpeg"
	if err := utils.ConvertPDFToImage(pdfFilePath, imageFilePath); err != nil {
		log.Fatalf("PDF 转换为图像时出错: %v", err)
	}
	fmt.Println("PDF 转换为图像成功,imageFilePath:", imageFilePath)

	fmt.Println("文档和图像生成成功!!!")
}
