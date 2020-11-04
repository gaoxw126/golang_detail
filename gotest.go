package main

import (
  "fmt"
  "strings"
  "time"
  "encoding/json"
  "bytes"
  //"gopkg.in/fatih/set.v0"
  // "testing"
  //. "github.com/smartystreets/goconvey/convey"
  "strconv"
  "log"
  "path/filepath"
  "net/http"
  "unsafe"
  "encoding/binary"
  "math"
  "reflect"
  "math/rand"
)

const HttpHeaderAuthTokenKey = "authtoken"

const(
    EmptyStatus        int8 = -1
	InitStatus         int8 = 0
	HealthStatus       int8 = 1
	UnHealthStatus     int8 = 2
)

type UpdateTopo struct {
  DockerID string `json:"dockerId"`
  SpaceId  string `json:"spaceId"`
}

func main() {

    /*
    retVal := defer_return()
    fmt.Println("return value : ", retVal)

    retValAfter := defer_return_after()
    fmt.Println("return after value : ", retValAfter)
    */

    //append_multi()

    //string_compare()

    //retVal := get_command_content()
    //fmt.Println("interface return after value : ", retVal)

    //print_struct_content()
    //ticker_timer()
    //goroutine_timer()
    //string_in_quotes()
    //for_sleep()
    //map_operation()
    //test_trim()
    //set_test()
    //bit_opt()
    //SinenceToString()
    //mapRange()
    //chanCloseWrite()
    //defaultInitValue()
    //mapLenTest()
    /*
    aTest := &ClassA{}
    bTest := &ClassB{}
    aTest.bObj = bTest
    bTest.aObj = aTest
    aTest.AMethod()
    bTest.BMethod()
    */

    //splitNSubstr()

    //formatFmt()
    //switchTest()
    //splitTest()
    //mapCountor()
    //chanPanic()
    //refInit()
    //unixTimeAlign()
    //timeDuration()
    //timeDurationUnit()
    //mapCurcurrent()
    //mapConcurrentWrite()
    //selectContinue()
    //pathBaseTest()


    //go testMapLenAfterMake()
    //time.Sleep(10 * time.Second)

    //triggerAtPunctual()
    //indexSubStr()

    //subDuration()

    //interfaceAssign()
    //httpTest()

    //testNewRequest()
    //testPrintPercentSign()
    //fmt.Println("Orig spaceId : ", interceptClusterOrigSpaceName("mongo-4ix8oo2s1m_copy_1587550381_shard_1", "_copy_", "_shard_"))
    //fmt.Println("Orig spaceId : ", interceptClusterOrigSpaceName("mongo-4ix8oo2s1m_copy_1587550381_mongos_0", "_copy_", "_mongos_"))
    //fmt.Println("Orig spaceId : ", interceptClusterOrigSpaceName("mongo-4ix8oo2s1m_copy_1587550381_configsvr", "_copy_", "_configsvr"))
    //fmt.Println("Is Little Endian: ", IsLittleEndian())
    //fmt.Println("byte to int: ", BytesToInt())
    //string2Bype()
    //byte2String()
    //lengthAndCapacity()
    //sliceArrary()
    //pointToSlice()
    //checkIntType()
    //CheckTimeDefaultDura()
    //ReadFromMap()
    //printFormat()
    //testSwitch()
    //testPrint()
    //delEmptyMapKey()
    //genRandomNum()
    //replaceStrSpace()
    cmpStructEmpty()
}


type element struct {
    Name string `json"name"`
    Age  int    `json"age"`
    Born time.Time `json"born"`
}
type header struct {
    Encryption  string `json:"encryption"`
    Timestamp   int64  `json:"timestamp"`
    Key         string `json:"key"`
    Partnercode int    `json:"partnercode"`
    DateTime    time.Time `json:"datetime"`
    SliceField  []element `json:"slicefield"`
}

func cmpStructEmpty() {
    header01 := header{
		Encryption:  "sha",
		Timestamp:   1482463793,
		Key:         "2342874840784a81d4d9e335aaf76260",
		Partnercode: 10025,
	}
    jsons, err := json.Marshal(header01)
    if err != nil {
        fmt.Printf("struct %v convert to Json fail: %s\n", header01, err.Error())
        return
    }
    strJson := string(jsons)
    fmt.Printf("struct conv as JSON string:%s\n", strJson)

    header02 := header{}
    if reflect.DeepEqual(header{}, header02){
        fmt.Printf("reflect.DeepEqual Empty struct : %#v\n", header02)
        if header02.Encryption=="" {
            fmt.Printf("Empty struct string field %#v\n", header02.Encryption)
        }
        if header02.Timestamp == 0 {
            fmt.Printf("Empty struct int64 field %#v\n", header02.Timestamp)
        }
    }

    header03 := &header{}
    if reflect.DeepEqual(&header{}, header03){
        fmt.Printf("reflect.DeepEqual Empty struct : %#v\n", header03)
        if header03.Encryption=="" {
            fmt.Printf("Empty struct string field %#v\n", header03.Encryption)
        }
        if header03.Timestamp == 0 {
            fmt.Printf("Empty struct int64 field %#v\n", header03.Timestamp)
        }
    }
}
func replaceStrSpace() {
    backupPeriodOrig := "[ 1, 3 , 4  ]"
    backupPeriodNoSpace := strings.Replace(backupPeriodOrig, " ", "", len(backupPeriodOrig))
    fmt.Printf("Str orig: {%s}, After Replace space str:{%s}\n", backupPeriodOrig, backupPeriodNoSpace)
    backupPeriod := strings.TrimSuffix(strings.TrimPrefix(backupPeriodNoSpace, "["), "]")
    fmt.Printf("Str orig: {%s}, After trim [,] str:{%s}\n", backupPeriodNoSpace,  backupPeriod)
}

func genRandomNum() {
     //r := rand.New(rand.NewSource(time.Now().UnixNano()))

     rand.Seed(time.Now().UnixNano())

    for i:=0; i<=20;i++ {
        fmt.Println("rand int : ", 2 + rand.Intn(4))
    }
}
func delEmptyMapKey() {

    var testMap map[string]int = make(map[string]int, 0)

    for i:=0;i<5;i++ {
        key := fmt.Sprintf("key_%d", i)
        testMap[key] = i
    }

    for key, value := range testMap {
        fmt.Printf("map [key:value]=[%s:%d]\n", key, value)
    }

    // Check del no-Exist key
    fmt.Println("")
    delete(testMap, "test_6")
    delete(testMap, "test_5")
    fmt.Println("")
    for key, value := range testMap {
        fmt.Printf("map [key:value]=[%s:%d]\n", key, value)
    }

}

func testPrint() {
    fmt.Println("Wanglongxiang, wanglongjian")
    fmt.Println("Time: ", time.Now())
}

func testSwitch() {
    for i:=0; i<4; i++ {
        testSwitchRun(int8(i))
    }
}

func testSwitchRun(lastStatus int8) {
        switch lastStatus {
		case HealthStatus:
            //fmt.Println("HealthStatus")
		case UnHealthStatus:
            fmt.Println("UnHealthStatus")
		case InitStatus:
            fmt.Println("Init HealthStatus")
		default:
            fmt.Println("Default HealthStatus")
					}
}

func printFormat() {
    spaceId := "mongo-hello"
    sqlStr := fmt.Sprintf("select owe_fee, status, ignore_alarm, ipv6_status, modified_date from space where space_id like \"%s%%\"", spaceId)
    fmt.Printf("sql : %s\n", sqlStr)
}

type tempStruct struct {
    temp int
}
func ReadFromMap() {

    var pool  map[string]int = make(map[string]int)
    a, OK := pool["test"]
    if OK {
        fmt.Printf("query from map: %v, %d\n", reflect.TypeOf(a), a)
    } else {
        fmt.Printf("can't query from map. default: %v-%d\n", reflect.TypeOf(a), a)
    }

    pool["test"] = 666
    a, OK = pool["test"]
    if OK {
        fmt.Printf("query from map: %v, %d\n", reflect.TypeOf(a), a)
    } else {
        fmt.Printf("can't query from map: %v, %d\n", reflect.TypeOf(a), a)
    }


    var poolMap map[string]*tempStruct = make(map[string]*tempStruct)

    temp, OK := poolMap["test"]
    if OK {
        fmt.Printf("query from struct map: %v, %v\n", reflect.TypeOf(temp), temp)
    } else {
        fmt.Printf("can't query from struct map: %v, %v\n", reflect.TypeOf(temp), temp)
    }

    poolMap["test"] = &tempStruct{
        temp: 888,
    }
    temp, OK = poolMap["test"]
    if OK {
        fmt.Printf("query from struct map: %v, %v\n", reflect.TypeOf(temp), temp)
    } else {
        fmt.Printf("can't query from struct map: %v, %#v\n", reflect.TypeOf(temp), temp)
    }

}

func CheckTimeDefaultDura() {
    fmt.Printf("default time.Duration(20) = %d \n", time.Duration(20))
    fmt.Printf("default time.Duration(20*time.Second) = %d \n", time.Duration(20 * time.Second))
}

func checkIntType() {
    fmt.Printf("math.MaxInt8  : [%v]\n", math.MaxInt8)
    fmt.Printf("math.MaxInt16 : [%v]\n", math.MaxInt16)
    fmt.Printf("math.MaxUint8 : [%v]\n", math.MaxUint8)
    fmt.Printf("math.MaxUint16: [%v]\n", math.MaxUint16)
}

type SliceTest struct {
    iV int
    sV string
}

func pointToSlice() {
    var sliceTest *[]SliceTest

//    fmt.Printf("Sliece len=%d\n", len(*sliceTest))

    appendPointToSlice(sliceTest)

 //   fmt.Printf("Sliece len=%d\n", len(*sliceTest))

}

func sliceArrary() {
    sliceTest := make([]*SliceTest, 0, 3)

    fmt.Printf("Sliece len=%d\n", len(sliceTest))

    appendSlice(sliceTest)

    fmt.Printf("Sliece len=%d\n", len(sliceTest))
}

func appendPointToSlice(sliceArr *[]SliceTest) {
    for i:=0; i<3; i++ {
        sObj := SliceTest{
            iV:i,
            sV:"test" + strconv.Itoa(i),
        }
        //*sliceArr = append(*sliceArr, sObj)
        *sliceArr = append(*sliceArr, sObj)
        fmt.Printf("struct obj: %v, slice Arrray len=%d, values: [%#v]\n", sObj, len(*sliceArr), sliceArr)
    }
    fmt.Printf("slice Arrray len=%d, values: [%#v]\n", len(*sliceArr), sliceArr)
}



func appendSlice(sliceArr []*SliceTest) {
    for i:=0; i<3; i++ {
        sObj := SliceTest{
            iV:i,
            sV:"test" + strconv.Itoa(i),
        }
        //*sliceArr = append(*sliceArr, sObj)
        sliceArr = append(sliceArr, &sObj)
        fmt.Printf("struct obj: %v, slice Arrray len=%d, values: [%#v]\n", sObj, len(sliceArr), sliceArr)
    }
    fmt.Printf("slice Arrray len=%d, values: [%#v]\n", len(sliceArr), sliceArr)
}


func lengthAndCapacity() {
    sliceCapTest := make([]int, 0, 3)
    //sliceLenCapTest := make([]int, 3, 3)
    var sliceLenCapTest []int

    fmt.Printf("cap slice: len[%d], cap[%d]\n", len(sliceCapTest), cap(sliceCapTest))
    fmt.Printf("len cap slice: len[%d], cap[%d]\n", len(sliceLenCapTest), cap(sliceLenCapTest))


    for i, v := range sliceCapTest {
        fmt.Printf("cap slice: index[%d], val[%d]\n", i, v)
    }

    for i, v := range sliceLenCapTest {
        fmt.Printf("len cap slice: index[%d], val[%d]\n", i, v)
    }

    for i:=0;i<3;i++ {
        sliceCapTest=append(sliceCapTest, i+5)
    }

    for i, v := range sliceCapTest {
        fmt.Printf("inited cap slice: index[%d], val[%d]\n", i, v)
    }


    // assign to un-initialized slice
    /*
    if cap(sliceLenCapTest) == 0 {
        sliceLenCapTest = make([]int, 0, 3)
        fmt.Printf("len cap slice un initailize capacity: make capacity 3\n")
        fmt.Printf("len cap slice: len[%d], cap[%d]\n", len(sliceLenCapTest), cap(sliceLenCapTest))
    }
    */
    if len(sliceLenCapTest) == 0 {
        sliceLenCapTest = make([]int, 3)
        fmt.Printf("len cap slice un initailize length and capacity: make length capacity 3\n")
        fmt.Printf("len cap slice: len[%d], cap[%d]\n", len(sliceLenCapTest), cap(sliceLenCapTest))
    }
    for i, v := range sliceCapTest {
        fmt.Printf("inited cap slice: index[%d], val[%d], to Assign un-initialized slice\n", i, v)
        sliceLenCapTest[i] = v
    }

}



func byte2String() {
    iv192 := []byte{
			0x4f, 0x02, 0x1d, 0xb2, 0x43, 0xbc, 0x63, 0x3d, 0x71, 0x78, 0x18, 0x3a, 0x9f, 0xa0, 0x71, 0xe8,
			0xb4, 0xd9, 0xad, 0xa9, 0xad, 0x7d, 0xed, 0xf4, 0xe5, 0xe7, 0x38, 0x76, 0x3f, 0x69, 0x14, 0x5a,
			0x57, 0x1b, 0x24, 0x20, 0x12, 0xfb, 0x7a, 0xe0, 0x7f, 0xa9, 0xba, 0xac, 0x3d, 0xf1, 0x02, 0xe0,
			0x08, 0xb0, 0xe2, 0x79, 0x88, 0x59, 0x88, 0x81, 0xd9, 0x20, 0xa9, 0xe6, 0x4f, 0x56, 0x15, 0xcd,
		}

    fmt.Println("IV byte: ", iv192)
    fmt.Println("IV string: ", string(iv192))
}
func string2Bype() {
    strCommonIV := "asngddgehjhijkdkddkmlkjj"

    fmt.Println("request : ", []byte(strCommonIV))
}
func testNewRequest() {
  request, err := http.NewRequest("POST", "http://test", nil)
  if err != nil {
    fmt.Println("new request failed. Error: %s", err.Error())
  }
  if request != nil {
    fmt.Println("request : ", request)
  }
}

func AddAuthToken(strToken string, r *http.Request) *http.Request {
  HttpHeaderAuthTokenKey := "authtoken"
  r.Header.Add(HttpHeaderAuthTokenKey, strToken)
  fmt.Println("[common token] add X-Auth-Token: [%s], http Request [%+v]", strToken, r)
  return r
}

func httpTest() {
  urlStr := "http://test/"
  httpBody := []byte("test body")
  bodyType := "application/json;charset=utf-8"
  ccRequest, err := http.NewRequest(urlStr, bodyType, bytes.NewBuffer(httpBody))
  if err != nil {
  }
  AddAuthToken("testtoken", ccRequest)
  fmt.Println("ccRequest: ", ccRequest)
}

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func interfaceAssign() {
    names := []string{"stanley", "david", "oscar"}
    vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}

func subDuration() {
  beginTime := time.Now()
  time.Sleep(1 * time.Minute)
  durSub := time.Now().Sub(beginTime)
  fmt.Println("sub duration : ", durSub)
  fmt.Println("sub duration Seconds : ", durSub.Seconds())
  fmt.Println("sub duration Minutes : ", durSub.Minutes())
}

func indexSubStr() {
  strTest2 := "mongo-qkf07fkvgl@0155275209114"
  //strTest2 := "mongo-qkf07fkvgl"
  var prefixSpaceId string = ""
  fmt.Println("spaceId= ", prefixSpaceId)
  index := strings.Index(strTest2, "2752")
  if index == -1 {
     prefixSpaceId = strTest2[0:]
     fmt.Println("spaceId= ", prefixSpaceId)
  } else {
      fmt.Println("spaceId", strTest2, " , index 2752 : ", index)
     prefixSpaceId = strTest2[0:index]
     fmt.Println("spaceId= ", prefixSpaceId)
  }
}

func triggerAtPunctual() {
  time9Clock := time.Now()
  fmt.Println("Time Now:", time9Clock)

  fmt.Println("Time Now Month: ", time9Clock.Month())
  fmt.Println("Time Now Day: ", time9Clock.Day())
  fmt.Println("Time Now hour: ", time9Clock.Hour())
  fmt.Println("Time Now Minute: ", time9Clock.Minute())
  fmt.Printf("Time Now: Month[%d], Minute[%d], Hour[%d], Minute[%d]",
  	time9Clock.Month(), time9Clock.Day(), time9Clock.Hour(), time9Clock.Minute())
}
func testMapLenAfterMake() {
  openerCh := make(chan struct{}, 100)

  for i:=1; i<=3; i++ {
    openerCh <- struct{}{}
  }

  for i:=1; i<=3; i++ {
    openerCh <- struct{}{}
  }
  i := 0
  lenMap := len(openerCh)
  fmt.Println("Len map  ", lenMap)

  for range openerCh {
    i++
    fmt.Println("Num ", i)
  }


// 后写入 Channel, 对上述 for range 无效.
  for i:=1; i<=3; i++ {
    openerCh <- struct{}{}
  }
}

func pathBaseTest() {
  fileName := "/hello/mnt/aaa.txt"
  pathBase := filepath.Base(fileName)
  fmt.Println("Base path : ",  pathBase)
}
func selectContinue() {

    go func() {
      ticker := time.NewTicker(time.Duration(3) * time.Second)
      defer ticker.Stop()
      i := 0

      for {
        select {
        case <-ticker.C:
          i++
          // 超时则认为下线，如果之前认为下线，而现在 ping 通，则让其重新活过来。
          // 如果只是单纯出现网络错误，则忽略，但会输出错误日志
          if i % 3 == 0 {

            if i % 6 == 0 {
              fmt.Println("current index then continue: ", i)
              continue
            }
            fmt.Println("current index : ", i)
          } else {
            fmt.Println("current index Not 3 * : ", i)
          }
        }
      }
    }()

    time.Sleep(time.Duration(30) * time.Second)
}

func httpWriteUnHTMLJson(res interface{}) []byte {
  // Disable HTML 格式后,会有成千上万个
  if res == nil {
    return []byte{}
  }
  bf := bytes.NewBuffer([]byte{})
  jsonEncoder := json.NewEncoder(bf)
  jsonEncoder.SetEscapeHTML(false)
  jsonEncoder.Encode(res)
  return bf.Bytes()
}

func varGlobalInit() {

}

func varLocalInit() {

}

func mapConcurrentWrite() {
  c := make(map[string]int)
  for i := 0; i < 100; i++ {
    go func() { //开100个协程并发写map
      for j := 0; j < 1000000; j++ {
        c[fmt.Sprintf("%d", j)] = j
      }
    }()
  }
  //time.Sleep(time.Second * 20) //让执行main函数的主协成等待20s,不然不会执行上面的并发操作
}

func mapCurcurrent() {
c := make(map[string]int)
       go func() {//开一个协程写map
            for j := 0; j < 1000000; j++ {
              c[fmt.Sprintf("%d", j)] = j
            }
       }()
       go func() {    //开一个协程读map
             for j := 0; j < 1000000; j++ {
                 fmt.Println(c[fmt.Sprintf("%d",j)])
             }
       }()

 time.Sleep(time.Second*20)
}
func timeDurationUnit() {
  time01 := time.Now()
  time.Sleep(time.Duration(3) * time.Second)
  time02 := time.Now()
  dur01 := time02.Second() - time01.Second()
  log.Println("time.Now().Second() dur : ", dur01)
  dur02 := time.Duration(time02.Second() - time01.Second())
  log.Println("time.Duration(time.Now().Second() - time.Now().Second()) dur : ", dur02)
  dur03 := time.Duration(time02.Second() - time01.Second()) * time.Second
  log.Println("time.Duration(time.Now().Second() - time.Now().Second()) dur : ", dur03)

  tm01 := time.Duration(60) * time.Second - time.Duration(30) * time.Second
  log.Println("time.Duration() * Second() : ", tm01)

  dur001 := time.Now()
  time.Sleep(time.Duration(3) * time.Second)
  dur002 := time.Now()
  durDiff := dur002.Sub(dur001)
  log.Println("Time duration : ", durDiff)
  log.Printf("Time Duration Sub : %#v", durDiff)
  log.Printf("Time Duration Sub().Seconds(): %#v", durDiff.Seconds())

  sampleInterval := time.Duration(10) * time.Second
  durReserved := time.Duration(2) * time.Second
  log.Printf("sampleInterval : %#v, durReserved : %#v", sampleInterval, durReserved)
  sampleDiff := sampleInterval - durReserved
  log.Printf("sampleInterval : %#v, durReserved : %#v , Sub : %#v", sampleInterval, durReserved, sampleDiff)
  if durDiff < sampleDiff {
    log.Printf("sampleDiff : %#v > durDiff : %#v", sampleDiff, durDiff)
    log.Printf("sampleDiff : %v > durDiff : %v", sampleDiff, durDiff)
  }
}

func timeDuration() {
  timerSche := time.NewTimer(0)
  counter := 0
  for {
    select{
    case <- timerSche.C:
      counter = counter + 1
      log.Println("Timer active: ", counter)
      timerSche.Reset(time.Duration(counter) * time.Second)
      if counter > 10 {
        timerSche.Reset(time.Duration(0) * time.Second)
      }
    }
    if counter == 20 {
      log.Println("Counter 10, break...")
      break
    }
  }
}
func unixTimeAlign() {
  curTimeNow := time.Now()
  curTime := time.Now().Unix()
  log.Print("time.Now().Unix() = ", curTime)
  curTime1 := time.Now().Unix() - int64(time.Now().Second())
  log.Print("time.Now() allignment minute = ", curTime1)
  curTime2 := curTimeNow.Unix() - int64(curTimeNow.Second())
  log.Print("Single time.Now() pre allignment minute = ", curTime2)
  curTime3 := curTimeNow.Unix() + 60 - int64(curTimeNow.Second())
  log.Print("Single time.Now() back allignment minute = ", curTime3)
}

func refInit() {
  var s1 []string
fmt.Printf("s1: %#v\n", s1) // s1: []string(nil)
s1 = append(s1, "hello")
fmt.Printf("s1: %#v\n", s1) // s1: []string{"hello"}

var s2 *[]string
fmt.Printf("s2: %#v\n", s2) // s2: (*[]string)(nil)

s3 := []string{"a", "b", "c"}
fmt.Printf("s3: %#v\n", s3) // s3: []string{"a", "b", "c"}
fmt.Printf("s3: %+v\n", s3) // s3: []string{"a", "b", "c"}

s4 := &[]string{}
fmt.Printf("s4: %#v\n", s4) // s4: &[]string{}

s5 := &s3
fmt.Printf("s5: %#v\n", s5) // s5: &[]string{"a", "b", "c"}

s6 := new([]string)
fmt.Printf("s6: %#v\n", s6) // s6: &[]string(nil)
// first argument to append must be slice; have *[]string
//s6 = append(s6, "hello") // 这是一个空引用的指针，所以报错

s7 := make([]string, 0)
fmt.Printf("s7: %#v\n", s7) // s7: []string{}

// 有毛病才用这种方式
s8 := new([]string)
*s8 = make([]string, 0)
fmt.Printf("s8: %#v\n", s8) // s8: &[]string{}

arr := [5]string{"a", "b", "c"}
s9 := arr[:]
fmt.Printf("s9: %#v\n", s9) // s9: []string{"a", "b", "c", "", ""}


}
func safeSend(ch chan int, value int) (closed bool) {
  defer func() {
    if recover() != nil {
      fmt.Println("Rewrite to closed monitor.quit channel")
      closed = true
    }
  }()
  ch <- value
  return false  // <=> closed = false; return
}

func chanPanic() {
  chInt := make(chan int, 0)
  go func() {
    for i:=1; i<=6; i++ {
      // Defer , not panic
      if closed := safeSend(chInt, i); closed {
        fmt.Println("chan closed, write failed.")
        return
      }
      // Panic
      // chInt <- i
    }
  }()

  for {
      recvInt := <- chInt
      fmt.Println("chan recvl: ", recvInt)
      if  5 == recvInt {
        close(chInt)
        fmt.Println("sleep 5 seconds , exit")
        time.Sleep(5 * time.Second)
        break
      }
  }
}

func mapCountor() {
  var azHostNum = make(map[string]int32, 0)

  for i := 1; i < 6; i++ {
    key := "az" + strconv.Itoa(i)
    fmt.Println("az key: ", key)
    if _, ok := azHostNum[key]; !ok {
      azHostNum[key] = 1
    } else {
      azHostNum[key]++
    }
    fmt.Println("az host num: ", azHostNum[key])
  }
  for i := 1; i < 6; i++ {
    key := "az" + strconv.Itoa(i)
    fmt.Println("az key: ", key)
    if _, ok := azHostNum[key]; !ok {
      azHostNum[key] = 1
    } else {
      azHostNum[key]++
    }
    fmt.Println("az host num: ", azHostNum[key])
  }
  fmt.Println("az host Num: ", azHostNum)
}

func splitTest() {
  //strTest := "prod_bj02"
  //strTest := "prod_bj02,prod_bj02,prod_bj03"
  strTest := "prod_bj02,prod_bj02,prod_bj02"
  azArr := strings.Split(strTest, ",")
  fmt.Println(azArr)
}

func switchTest() {
  const VarSwitchValue2 = 2
  const VarSwitchValue3 = 3
  var varSwitch int = 3
  switch varSwitch {
  case 1:
    fmt.Println("varSwitch = 1")
  case VarSwitchValue2, VarSwitchValue3:
    fmt.Println("varSwitch = 2 or 3")
  default:
    fmt.Println("varSwitch other value.")
  }

  var nilVar []string
  nilVar = nil
  fmt.Println("varNil: ", nilVar)
}
func formatFmt() {
  strSub := "hello"
  strTest := fmt.Sprintf("machine %s", strSub)
  fmt.Print(strTest)
  strTest = fmt.Sprintf("machine %s", strSub)
  fmt.Print(strTest)
}

func splitNSubstr() {

  strFlavor := "T1-1C1G10G"


  //strTest := "1G10G"
  strTest := strings.Split(strFlavor, "-")[1]
  fmt.Println("- split flavor: ", strTest)

  cupstr := strings.Split(strTest, "C")[0]
  fmt.Println("str CUP : ", cupstr)
  cupInt, _ := strconv.Atoi(cupstr)
  fmt.Println("int CUP : ", cupInt)

  strMemDisk := strings.Split(strTest, "C")[1]
  fmt.Println("mem disk str: ", strMemDisk)

  subStrArr := strings.SplitN(strMemDisk, "G", 3)
  fmt.Println("subStrArr len ", len(subStrArr))
  fmt.Println("subStrArr : ", subStrArr)
  iMem, _ := strconv.Atoi(subStrArr[0])
  fmt.Println("int mem : ", iMem)
  iDisk, _ := strconv.Atoi(subStrArr[1])
  fmt.Println("int disk : ", iDisk)
}

type ClassA struct {
  a string
  bObj *ClassB
}

type ClassB struct {
  b string
  aObj *ClassA
}


func (aObj *ClassA) AMethod() {
  aObj.a = "hello A"
  fmt.Println(aObj.a)
  //aObj.bObj.BMethod()

}

func (bObj *ClassB) BMethod() {
  bObj.b = "Hello B"
  fmt.Println(bObj.b)
  bObj.aObj.AMethod()
}

func StringSliceEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }

    if (a == nil) != (b == nil) {
        return false
    }

    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}


/*
func TestStringSliceEqual(t *testing.T) {
    Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
        a := []string{"hello", "goconvey"}
        b := []string{"hello", "goconvey"}
        So(StringSliceEqual(a, b), ShouldBeTrue)
    })
}
*/

func mapLenTest() {
  testMap := make(map[int]int, 3)

  fmt.Println("map len ", len(testMap))
  for i:=0; i< 8; i++ {
    testMap[i] = 0
  }

  fmt.Println("map len ", len(testMap))
  for k, v := range testMap{
    fmt.Println("k, v: ", k, v)
  }
}
func defaultInitValue() {
  var closed bool
  fmt.Println("bool init default value: ", closed)
}
func chanCloseWrite() {
  quit:= make(chan bool)

  go func() {
    counter := 0
    counterFalse := 0
    for {
      select {
      case flag := <- quit:
        counter++
        fmt.Println("read true from quit channel.", flag, counter)
      default:
        counterFalse++
        fmt.Println("read false from quit channel.falseCounter:", counterFalse)
      }

    }
  }()


  for i := 0; i < 5; i++ {
    if 3 == i {
      close(quit)
      fmt.Println("closed chann : ", quit)
    }
    quit <- true
    fmt.Println(" chann : ", quit)

  }
  /*
  for i := 0; i < 5; i++ {
    quit <- false
  }
  */
}

func mapRange() {
  strStrMap := map[string]string {
    "1":"test1",
    "2":"test2",
    "3":"test3",
  }

  for kv := range strStrMap {
    fmt.Println("[map[string]string] element : ", kv)
    delete(strStrMap, kv)
    fmt.Println("[map[string]string] element : ", strStrMap)
  }

  fmt.Println("[map[string]string] element : ", strStrMap)
  //newkey := "testnewkey"
  //strStrMap[newkey]="lallal"
  strStrMap["testNewKey"]="lallal"
  fmt.Println("[map[string]string] element : ", strStrMap)
}

func SinenceToString() {
  sienceNumber := "1.073741824e+09"
  var newNumber  int64 //float64
  var newFloatNum float64
  iMem, err := fmt.Sscanf(sienceNumber, "%d", &newNumber)//strconv.ParseInt(strings.Trim(sienceNumber, " "), 10, 0)
  if err != nil {
    fmt.Printf("[dockersvr] ParseUint Memory=%s, Error: %v\n", sienceNumber, err.Error())
  }
  fmt.Printf("fmt.Sscanf convert %d number.", iMem)
  fmt.Printf("string 1.0737841724e+09, to Int : %d\n", newNumber)

   iMem, err = fmt.Sscanf(sienceNumber, "%e", &newFloatNum)//strconv.ParseInt(strings.Trim(sienceNumber, " "), 10, 0)
  if err != nil {
    fmt.Printf("[dockersvr] ParseUint Memory=%s, Error: %v\n", sienceNumber, err.Error())
  }
  fmt.Printf("string 1.0737841724e+09, to Float : %f\n", newFloatNum)
  fmt.Printf("string 1.0737841724e+09, to Float : %d\n", int64(newFloatNum))
}

func bit_opt() {
  bitmap := 4

  fmt.Println("Bit & result: ",bitmap & (1<<2))
  fmt.Println("Bit | result: ",bitmap | (1<<3))
}

/*
func set_test() {
  srcMap := map[int]string {1:"first", 2:"seconde", 3:"third", 4:"fourth"}
  dstSet := set.New()
  for i := 1; i < 10; i++ {
    dstSet.Add(i)
  }

  fmt.Println("map[int]string : ", srcMap)
  fmt.Printf("len set = %d, set[int] : %v\n", dstSet.Size(), dstSet)
  fmt.Printf("set[int].List() : %v\n", dstSet.List())
  //dstBakSet := dstSet.Copy()

  fmt.Println()
  index := 0
  for _, elem := range dstSet.List() {
    index++
    iElem, _ := elem.(int)
    fmt.Printf("element %d : %v \n", index, iElem)
    if _, Ok := srcMap[iElem]; !Ok {
      dstSet.Remove(iElem)
      fmt.Printf("after remove %d element not in Map, len(dstSet)=%d,  Set : %v \n", elem, dstSet.Size(), dstSet)
    }
  }

  // delete element in set and not in map.

}
*/

func test_trim() {
  //strLeft := "nova-leftdocker-id-nova-"
  strLeft := "nova-a165d23d-c5f5-45fe-a038-df80b5d624a1"
  strRight := "rightdocker-id--nova-"

  // TrimLeft/TrimRight 删除头部/尾部中匹配子串中任意字符的字符.
  // TrimPrefix/TrimSuffix 删除头部/尾部中完全匹配子串的子串.
  fmt.Println("TrimLeft and TrimPrefix:")
  fmt.Println(strings.TrimLeft(strLeft, "nova-"))
  fmt.Println(strings.TrimPrefix(strLeft, "nova-"))
  fmt.Println(strings.TrimPrefix(strLeft, "nva-"))
  fmt.Println("TrimRight and TrimSuffix:")
  fmt.Println(strings.TrimRight(strRight, "nova-"))
  fmt.Println(strings.TrimSuffix(strRight, "nova-"))
  fmt.Println(strings.TrimSuffix(strRight, "nva-"))
  fmt.Println("strRight:", strRight)
}

func for_sleep() {
  // goroutine : Update each 3 seconds
  incrVal := 1
  exitCounter := 1
  for {
    exitCounter++
    UpdateContainers(incrVal)
    fmt.Println("[container] Current incrVal: ", incrVal)
    //time.Sleep(5 * time.Nanosecond)
    time.Sleep(5 * time.Millisecond)
    //time.Sleep(5 * time.Second)
    //time.Sleep(0)
    if exitCounter == 16 {
      break
    }
  }
  return
}

func UpdateContainers(incrVar int) {
  for i := 1; i <= 8; i++ {
    incrVar = incrVar + i
    if 0 == i % 2 {
      fmt.Println("Even: ", i * 2)
    } else {
      fmt.Println("Odd: ", i)
    }
    if 0 == incrVar % 5 {
      fmt.Println("incrVar : ", incrVar)
      return
    }
  }
  return
}

func map_operation() {
  containerMap := make(map[string]string)
  containerMap["testKey"]="111"
  containerMap["testKey"]="222"
  containerMap["testKey"]="333"
  containerMap["testKey1"]="111"
  containerMap["testKey2"]="222"
  containerMap["testKey3"]="333"

  fmt.Printf("map content: %+v", containerMap)
  fmt.Println("map has elements num: ", len(containerMap))
  for key, val := range containerMap {
    fmt.Printf("map : key [%s],value [%s]", key, val)
  }

  for key, val := range containerMap {
    if 0 == strings.Compare(key, "testKey2") {
      delete(containerMap, "testKey2")
    }
    fmt.Printf("map : key [%s],value [%s]", key, val)
  }
  fmt.Printf("after delete map content: %+v", containerMap)
}
func string_in_quotes() {
    dQuoteStr := "\"what's that?\", he said."
    bQuoteStr := `"what's that?", he said.`
    fmt.Println(dQuoteStr)
    fmt.Println(bQuoteStr)
}

func goroutine_timer() {
  tick := time.NewTimer(1e8)
  //boom := time.After(5e8)
  boom := time.After(2 * time.Second)
  funcFiledVar := "Command GetConfig execute timeout."
  go func() {
    for {
      select {
      case <- tick.C:
        fmt.Println("tike.")
      case <- boom:
        fmt.Println("BOOM!")
        fmt.Println(funcFiledVar)
        return
      default:
        //fmt.Println("     .")
        //time.Sleep(5e7  )
      }
    }
  } ()
  for i := 1; i <= 3; i++ {
    fmt.Printf("main process. Try to execute Command GetConfig %d time\n", i)
    time.Sleep(2 * time.Second)

  }

}

func ticker_timer() {
  tick := time.Tick(1e8)
  //boom := time.After(5e8)
  boom := time.After(2 * time.Second)

  for {
    select {
    case <- tick:
      fmt.Println("tike.")
    case <- boom:
      fmt.Println("BOOM!")
      return
    default:
      fmt.Println("     .")
      time.Sleep(5e7)
    }
  }
}
func print_struct_content() {
  upTopo := UpdateTopo{
    DockerID: "testDockerID",
    SpaceId: "testSpaceID",
  }

  fmt.Printf("struct %v \n", upTopo)
  fmt.Printf("struct %+v \n", upTopo)
}
func get_command_content() interface{} {
  /*
  upTopo := UpdateTopo{
    DockerID: "testDockerID",
    SpaceId: "testSpaceID",
  }
  return upTopo
  */

  strTopo := "testStringTopo"
  return strTopo

}
func string_compare() {
  emptystr := ""
  az1 := "prod_bj01"
  az2 := "prod_bj02"
  az3 := "prod_bj03"

  fmt.Printf("az1 < az2 ? %v \n", strings.Compare(az1, az2))
  fmt.Printf("az1 < az3 ? %v \n", strings.Compare(az1, az3))
  fmt.Printf("az2 < az3 ? %v \n", strings.Compare(az2, az3))
  fmt.Printf("emptystr < az1 ? %v \n", strings.Compare(emptystr, az1))
  fmt.Printf("az2 > az1 ? %v \n", strings.Compare(az2, az1))
}

func append_multi() {
  var InfluxHost [][]string
  curCluster := "[prod_bj02;172.19.29.28:8086;172.19.29.63:8086;172.19.29.42:8086"
  InfluxHost = append(InfluxHost, strings.Split(strings.TrimLeft(curCluster, "["), ";"))
  fmt.Printf("InfluxDB host: %v \n", InfluxHost)
  fmt.Printf("InfluxDB host len: %v \n", len(InfluxHost))
}
func defer_return_after() (result int) {

  defer func() {
    result++
    fmt.Println("defer after func: return value ", result)
  }()
  return 0
}
func defer_return() (result int) {

  return 0
  defer func() {
    result++
    fmt.Println("defer func: return value ", result)
  }()
  return 2
}

func testPrintPercentSign() {
    spaceId := "mongo-wrh71cm1sq"
    sqlDomainInst := fmt.Sprintf("select i.space_id, i.docker_id, i.instance_ip, i.failover_flag, i.domain, i.nlb_id from instance i, space s where s.space_id like \"%s%%\" and s.status=100 and s.cluster_type in (1,5) and i.space_id=s.space_id and length(i.domain)>0;", spaceId)
	fmt.Printf("[Dev] query domain sql: %s\n", sqlDomainInst)
	fmt.Printf("[Dev] query domain sql: %s0000", sqlDomainInst)
	//fmt.Println("[Dev] query domain sql: ", sqlDomainInst)
}

func interceptClusterOrigSpaceName(spaceId, firstTag, secondTag string) string {
	tmpSpaceId := spaceId
	if !strings.Contains(spaceId, firstTag) || !strings.Contains(spaceId, secondTag) {
		fmt.Println("String %s NOT contains [%s / %s]", spaceId, firstTag, secondTag)
	} else {
		index1 := strings.Index(spaceId, firstTag)
		index2 := strings.Index(spaceId, secondTag)
		tmpSpaceId = spaceId[0:index1] + spaceId[index2:]
	}
	return tmpSpaceId
}



func BytesToInt() int {

    bys := []byte{0x31, 0x30, 0x30}
	fmt.Printf("Need convert byte array: %#v\n", bys)
	bytebuff := bytes.NewBuffer(bys)
	var data int
	if IsLittleEndian() {
		binary.Read(bytebuff, binary.LittleEndian, &data)
		fmt.Printf("Little endian int64: %#v\n", data)
	} else {
		binary.Read(bytebuff, binary.BigEndian, &data)
		fmt.Printf("Big endian int64: %#v\n", data)
	}
	return int(data)
}


func IsLittleEndian() bool {
	//var i int32 = 0x01020304
	var i int32 = 0x313030

	// 下面这两句是为了将int32类型的指针转换为byte类型的指针
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)

	b := *pb // 取得pb位置对应的值

    fmt.Printf("int num: %#v\n", i)
    fmt.Printf("16 bit int num: %X\n", i)
    fmt.Printf("byte array: %#v\n", pb)
    fmt.Printf("b: byte array first byte: %#v\n", b)


	// 由于b是byte类型的,最多保存8位,那么只能取得开始的8位
	// 小端: 04 (03 02 01)
	// 大端: 01 (02 03 04)
	return b == 0x04
}
