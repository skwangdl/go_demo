package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"
import "time"
import "algorithms/bubblesort"
import "algorithms/qsort"

//flag 解析命令行
var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string)(values []int, err error){
	file, err := os.Open(infile)	//os 打开文件
	if err != nil{					//判断是否有异常
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()				
	br := bufio.NewReader(file)		//bufio 文件读入内存缓存
	values = make([]int, 0)

	for{
		line, isPrefix, err1 := br.ReadLine()	//无限循环，按行读取行
		if err1 != nil{							//判断是否有异常
			if err1 != io.EOF{
				err = err1
			}
			break
		}
		if isPrefix{
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		str := string(line)					//读取的行转为字符串
		value, err1 := strconv.Atoi(str)	//字符串转为int
		if err1 != nil{
			err = err1
			return
		}
		values = append(values, value)		//对数组切片按照下标顺序添加值
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil{
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()
	for _, value := range values{		//range遍历values数组
		str := strconv.Itoa(value)		//int转string
		file.WriteString(str + "\n")	//文件写入string
	}
	return nil
}

func main(){
	flag.Parse()
	if infile != nil{
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
			case "qsort":
				qsort.QuickSort(values)
			case "bubblesort":
				bubblesort.BubbleSort(values)
			default:
				fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}