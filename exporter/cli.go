package exporter

import (
    "encoding/json"
    "github.com/zr-hebo/sniffer-agent/model"
    "os"
    "time"
    "fmt"
)

type cliExporter struct {
    outputFile *os.File
}

func NewCliExporter() *cliExporter {
    currentTime := time.Now().Format("20060102150405")
    filename := "sniff.out" + currentTime
    outputFile, err := os.Create(filename)
    fmt.Println("sniff out file is: "+filename)
    if err != nil {
        return nil
    }

    return &cliExporter{
        outputFile: outputFile,
    }
}

func (c *cliExporter) Export(qp model.QueryPiece) error {
    // 使用encoding/json包将数据qp序列化为JSON格式
    jsonData, err := json.Marshal(qp)
    if err != nil {
        return err
    }

    // 添加换行符
    jsonData = append(jsonData, '\n')

    // 将JSON数据写入文件
    _, err = c.outputFile.Write(jsonData)
    if err != nil {
        return err
    }

    return nil
}