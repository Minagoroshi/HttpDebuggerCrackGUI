package main

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

func getAppVersion() string {

	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\MadeForNet\HTTPDebuggerPro`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatalln("OpenKey:", err)
	}
	defer k.Close()

	version, _, err := k.GetStringValue("AppVer")
	if err != nil {
		log.Fatal(err)
	}

	verRx := regexp.MustCompile("(\\d+.*)")
	parsedVersion := verRx.FindString(version)
	parsedVersion = strings.Replace(parsedVersion, ".", "", -1)

	return parsedVersion
}

func getSerialNumber(appVersion string) string {
	var volumeInfo uint32
	volumeName := "C:\\"

	// Call GetVolumeInformationW from the kernel32.dll to get the volume information
	r, _, err := syscall.NewLazyDLL("kernel32.dll").NewProc("GetVolumeInformationW").Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumeName))),
		0, 0, uintptr(unsafe.Pointer(&volumeInfo)), 0, 0, 0, 0,
	)
	if r == 0 {
		log.Fatal(err)
	}

	value, err := strconv.Atoi(appVersion)
	if err != nil {
		log.Fatal(err)
	}

	serialNumber := uint32(value) ^ ((^volumeInfo >> 1) + 0x2E0) ^ 0x590D4

	serialNumberStr := strconv.Itoa(int(serialNumber))

	return serialNumberStr
}

func createKey() string {
	var key string
	for len(key) != 16 {
		v1, v2, v3 := generateRandomBytes()
		key = fmt.Sprintf("%02X%02X%02X7C%02X%02X%02X%02X", v1, v2^0x7C, 0xFF^v1, v2, v3%255, v3%255^7, v1^(0xFF^(v3%255)))
	}

	return key
}

func generateRandomBytes() (byte, byte, byte) {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	return b[0], b[1], b[2]
}

func crack() {
	av := getAppVersion()
	sn := getSerialNumber(av)
	key := createKey()

	log.Println("App Version:", av)
	log.Println("Serial Number:", sn)
	log.Println("Key:", key)

	k, _, err := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\MadeForNet\HTTPDebuggerPro`, registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	err = k.SetStringValue("SN"+sn, key)
	if err != nil {
		log.Fatal(err)
	}
}
