package go_file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadByLine(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Tidak dapat Membaca File")
	}

	defer file.Close()

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)

	for fileScan.Scan() {
		fmt.Println(fileScan.Text())
	}

}

func AppendText(fileName string,input string) {
	file, err := os.OpenFile(fileName,os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	resultOfText := fmt.Sprintf("\n%s",input)

	if _,err := file.WriteString(resultOfText); err != nil {
		fmt.Println(err)
	}
}