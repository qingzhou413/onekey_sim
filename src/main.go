package main

import (
	"flag"
	"os"
	"net"
	"fmt"
	"strconv"
	"byteutil"
	"time"
)

var SERVER_IP = flag.String("ip", "", "服务器IP")
var PORT = 27003;

var BLE_MAC = flag.String("bleMac", "", "蓝牙网关MAC")
var BLT_MAC = flag.String("bltMac", "", "标签MAC")
var TRIGGER_TYPE = flag.String("type", "", "1 : 一键报警\r\n2 : 取消一键报警")

func main() {
	flag.Parse()

	if *SERVER_IP == "" || *BLE_MAC == "" || *BLT_MAC == "" || *TRIGGER_TYPE == "" {
		flag.Usage()
		os.Exit(-1)
	}


	fmt.Printf("server ip %s, ble %s blt %s type %s\n", *SERVER_IP, *BLE_MAC, *BLT_MAC, *TRIGGER_TYPE)

	conn, err := net.Dial("udp", *SERVER_IP+":"+strconv.Itoa(PORT))
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}

	data := buildEvent(*BLE_MAC, *BLT_MAC, *TRIGGER_TYPE)
	conn.Write(data)
	fmt.Println("send msg")

}

/*
        return data;
 */
func buildEvent(ble string, blt string, typeStr string) []byte {
	var data = make([]byte, 44)

	index := 0
	// header
	byteutil.PutShortToBuffer(data, uint16(0x8546), index)
	index += 2;
	byteutil.PutByteToBuffer(data, byte(1), index);
	index++;
	byteutil.PutByteToBuffer(data, byte(0x11), index);
	index++;
	byteutil.PutShortToBuffer(data, uint16(23), index);
	index += 2;
	bleMac, _ := strconv.ParseUint(ble, 16, 64)
	byteutil.Put6LongToBuffer(data, bleMac, index);
	index += 6;
	t := time.Now()
	byteutil.Put8LongToBuffer(data, uint64(t.Unix()*1000), index);
	index += 8;
	byteutil.PutByteToBuffer(data, byte(1), index);
	index++;

	// body
	rssi := -50
	byteutil.PutByteToBuffer(data, byte(rssi), index); // rssi
	index++;
	byteutil.PutByteToBuffer(data, byte(0), index); // txpower
	index++;
	byteutil.PutByteToBuffer(data, byte(90), index); // battery
	index++;
	event, _ := strconv.Atoi(typeStr)
	byteutil.PutByteToBuffer(data, byte(event), index); // event type 0触发
	// 1报警
	index++;
	byteutil.PutByteToBuffer(data, byte(1), index); // ble_counter
	// 累计广播次数
	index++;
	byteutil.PutByteToBuffer(data, byte(0), index); // antenna
	// num
	index++;
	byteutil.PutShortToBuffer(data, uint16(100), index); // broadcast_interval
	index += 2;
	byteutil.PutShortToBuffer(data, uint16(0), index); // trigger num
	index += 2;
	byteutil.PutIntToBuffer(data, 123, index);
	index += 4;
	byteutil.PutByteArrayToBuf(data, []byte{3, 2, 1}, index);
	index += 3;
	bltMac, _ := strconv.ParseUint(blt, 16, 64)
	byteutil.Put6LongToBuffer(data, bltMac, index);
	index += 6;
	return data
}
