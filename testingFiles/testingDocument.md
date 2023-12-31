# 系统测试文档
|变更人|变更时间|变更内容|版本号|
|------|-------|------|-------|
|刘克典|2023-07-29|完成测试文档基本内容|v4.0|

# 1.引言
测试文档由软件设计说明所驱动。测试用于验证模块单元实现了模块设计中定义的规格。一个完整的单元测试说明应该包含白盒测试和黑盒测试。测试验证程序应该执行的工作，测试验证程序不应该执行的工作。

## 1.1编写目的
通过测试尽可能的找出项目中的错误，并加以纠正。测试不仅是最后的复审，更是保证软件质量的关键。
简单地说就是想尽一切方法尝试“破坏”它，这样才能找出失败与不足之处，最终的任务就是构造更高质量的软件产品。

## 1.2参考文献
1. IEEE标准

2. 《软件工程与计算(卷二):软件开发的技术基础》刘钦、丁二玉著

3. MSCS软件需求规格文档

# 2.测试平台
Ubuntu20.04

# 3.测试用例

## 3.1 FirstLevelServiceAdd测试用例
```
func BenchmarkFirstLevelServiceAdd(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(100)
		b := rand.Intn(100)

		info["operand_1"] = a
		info["operand_2"] = b

		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", aUrlAdd, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("发送请求失败", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read回复失败", err.Error())
			return
		}
		//对返回的信息进行处理
		str := (string)(respBytes)

		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := a + b
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("向FirstLevelCalservice/Add传递的值为 %d和%d ，预期结果为%d,得到的计算结果为 %d，结果正确，测试通过\n", a, b, a+b, int(resultNumber2))
		} else {
			fmt.Printf("向FirstLevelCalservice/Add传递的值为 %d和%d ，预期结果为%d,得到的计算结果为 %d，结果错误，测试失败\n", a, b, a+b, int(resultNumber2))
		}
		request.Body.Close()
	}
}
```
## 3.2 FirstLevelServiceSub测试用例
```
func BenchmarkFirstLevelServiceSub(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(100)
		b := rand.Intn(100)

		info["operand_1"] = a
		info["operand_2"] = b

		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", aUrlSub, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("发送请求失败", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read回复失败", err.Error())
			return
		}
		//对返回的信息进行处理
		str := (string)(respBytes)
		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := a - b
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("向FirstLevelCalservice/Sub传递的值为 %d和%d ，预期结果为%d,得到的计算结果为%d，结果正确，测试通过\n", a, b, a-b, int(resultNumber2))
		} else {
			fmt.Printf("向FirstLevelCalservice/Sub传递的值为 %d和%d ，预期结果为%d,得到的计算结果为%d，结果错误，测试失败\n", a, b, a-b, int(resultNumber2))
		}
		request.Body.Close()
	}
}
```

## 3.3 SecondLevelServiceMul测试用例
```
func BenchmarkSecondLevelServiceMul(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(10)
		b := rand.Intn(10)

		info["operand_1"] = a
		info["operand_2"] = b

		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", bUrlMul, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("发送请求失败", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read回复失败", err.Error())
			return
		}
		//对返回的信息进行处理
		str := (string)(respBytes)

		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := a * b
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("向SecondLevelCalservice/Mul传递的值为 %d和%d ，预期结果为%d,得到的计算结果为 %d，结果正确，测试通过\n", a, b, a*b, int(resultNumber2))
		} else {
			fmt.Printf("向SecondLevelCalservice/Mul传递的值为 %d和%d ，预期结果为%d,得到的计算结果为 %d，结果错误，测试失败\n", a, b, a*b, int(resultNumber2))
		}
		request.Body.Close()
	}
}

```
## 3.4 SecondLevelServiceDiv测试用例
```
func BenchmarkSecondLevelServiceDiv(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(9) + 1
		b := rand.Intn(9) + 1

		info["operand_1"] = a
		info["operand_2"] = b

		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", bUrlDiv, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("发送请求失败", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read回复失败", err.Error())
			return
		}
		//对返回的信息进行处理
		if b == 0 {
			fmt.Println("除数为 0 ")
		} else {
			str := (string)(respBytes)
			var theInd1 = strings.Index(str, "data")
			var theInd2 = strings.Index(str, "message")
			var str0 = str[theInd1+6 : theInd2-2]

			resultNumber1 := a / b
			resultNumber2, _ := strconv.Atoi(str0)

			as := Assert(resultNumber1, int(resultNumber2))
			if as {
				fmt.Printf("向SecondLevelCalservice/Div传递的值为 %d和%d ，预期结果为%d,得到的计算结果为 %d，结果正确，测试通过\n", a, b, a/b, int(resultNumber2))
			} else {
				fmt.Printf("向econdLevelCalservice/Div传递的值为 %d和%d ，预期结果为%d,得到的计算结果为 %d，结果错误，测试失败\n", a, b, a/b, int(resultNumber2))
			}
		}

		request.Body.Close()
	}
}
```
## 3.5 AdvancedLevelServiceFact测试用例
```
func BenchmarkAdvancedLevelServiceFact(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(10)
		info["operand"] = a
		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", cUrlFact, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("发送请求失败", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read回复失败", err.Error())
			return
		}
		//对返回的信息进行处理

		str := (string)(respBytes)
		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := Fact(a)
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("向AdvancedCalservice/Fact传递的值为 %d ，预期结果为%d,得到的计算结果为 %d，结果正确，测试通过\n", a, resultNumber1, int(resultNumber2))
		} else {
			fmt.Printf("向AdvancedCalservice/Fact传递的值为 %d ，预期结果为%d,得到的计算结果为 %d，结果错误，测试失败\n", a, resultNumber1, int(resultNumber2))
		}

		request.Body.Close()
	}
}
```

## 3.6 AdvancedLevelServiceFib测试用例
```
func BenchmarkAdvancedLevelServiceFib(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(5)
		info["operand"] = a
		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", cUrlFib, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("发送请求失败", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read回复失败", err.Error())
			return
		}
		//对返回的信息进行处理

		str := (string)(respBytes)
		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := Fib(a)
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("向AdvancedCalservice/Fib传递的值为 %d ，预期结果为%d,得到的计算结果为 %d，结果正确，测试通过\n", a, resultNumber1, int(resultNumber2))
		} else {
			fmt.Printf("向AdvancedCalsrvice/Fib传递的值为 %d ，预期结果为%d,得到的计算结果为 %d，结果错误，测试失败\n", a, resultNumber1, int(resultNumber2))
		}

		request.Body.Close()
	}
}
```

|测试名称|测试内容|预期结果|测试结果|
|-----|-----|------|------|
|main_test|发送http请求|验证返回信息|通过|

# 4.结论
网关的关键功能已成功实现并能通过测试，能够正常使用
