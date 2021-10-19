package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
)

type July struct {
	Name string
	Age  int
	Sex  string
}

func ZipGameAccount(account []byte) ([]byte, error) {
	var buff bytes.Buffer
	writer := gzip.NewWriter(&buff)
	defer writer.Close()
	writer.Write(account)
	if err := writer.Flush(); err != nil {
		return nil, err
	}
	data := buff.Bytes()
	return data, nil
}

func ZipGameRead(data []byte) ([]byte, error) {
	reader, e := gzip.NewReader(bytes.NewReader(data))
	if e != nil {
		panic("读错误")
	}
	defer reader.Close()
	var content []byte
	var tmp = make([]byte, 128)

	for {
		n, err := reader.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			break
		}
		content = append(content, tmp[:n]...)
	}

	return content, nil
}

func (x July) IsStructureEmpty() bool {
	return reflect.DeepEqual(x, July{})
}

func main() {

	//str :=[]byte("hello 七月")

	ii := &July{}

	fmt.Println(ii.Name == "")
	fmt.Println(ii.IsStructureEmpty())

	j := &July{
		Name: "七月11",
		Age:  1854,
		Sex:  "男",
	}
	fmt.Println(j.IsStructureEmpty())
	data, err := json.Marshal(j)
	account, err := ZipGameAccount(data)
	accountData := base64.StdEncoding.EncodeToString(account)
	fmt.Println(accountData)
	accountDataD, err := base64.StdEncoding.DecodeString(accountData)
	fmt.Println(accountDataD)
	readByte, err := ZipGameRead(accountDataD)
	fmt.Println(readByte)
	if err != nil {
		log.Panic(err)
		panic("有毒")
	}

	jj := &July{}
	err = json.Unmarshal(readByte, jj)
	fmt.Printf("%+v", jj)
	fmt.Printf("%+v", jj.Name)

}
