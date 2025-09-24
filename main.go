package main

import (
	"crypto/md5"
	"debug/macho"
	"debug/pe"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// 检查可执行文件架构
func checkExeArchitecture(filePath string) (string, error) {
	file, err := pe.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	switch file.FileHeader.Machine {
	case pe.IMAGE_FILE_MACHINE_I386:
		return "32-bit (x86)", nil
	case pe.IMAGE_FILE_MACHINE_AMD64:
		return "64-bit (x64)", nil
	case pe.IMAGE_FILE_MACHINE_ARM64:
		return "64-bit (ARM64)", nil
	case pe.IMAGE_FILE_MACHINE_IA64:
		return "64-bit (IA64)", nil
	case pe.IMAGE_FILE_MACHINE_ARMNT:
		return "32-bit (ARM)", nil
	default:
		return fmt.Sprintf("Unknown (0x%x)", file.FileHeader.Machine), nil
	}
}

func checkMachOArchitecture(filePath string) (string, error) {
	file, err := macho.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	switch file.Cpu {
	case macho.CpuAmd64:
		return "64-bit Intel (x86_64)", nil
	case macho.CpuArm64:
		return "64-bit ARM (arm64)", nil
	case macho.Cpu386:
		return "32-bit Intel (i386)", nil
	case macho.CpuArm:
		return "32-bit ARM", nil
	default:
		return fmt.Sprintf("Unknown architecture (0x%x)", file.Cpu), nil
	}
}

func MD5FileStream(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

func main() {
	fmt.Printf("操作系统: %s\n", runtime.GOOS)
	fmt.Printf("架构: %s\n", runtime.GOARCH)

	var bit string
	var err error
	bit, err = checkExeArchitecture(filepath.Join("1.0.41", "windows", "adb.exe"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("windows adb架构:", bit)
	}
	bit, err = checkMachOArchitecture(filepath.Join("1.0.41", "darwin", "adb"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("darwin adb架构:", bit)
	}
	bit, err = checkMachOArchitecture(filepath.Join("1.0.41", "linux", "adb"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("linux adb架构:", bit)
	}

	var sum string
	sum, err = MD5FileStream(filepath.Join("1.0.41", "windows", "adb.exe"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("windows adb sum:", sum) // cdde1e5edb57c8f82627a5bde94b0591
	}
	sum, err = MD5FileStream(filepath.Join("1.0.41", "windows", "AdbWinApi.dll"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("windows AdbWinApi sum:", sum) // ed5a809dc0024d83cbab4fb9933d598d
	}
	sum, err = MD5FileStream(filepath.Join("1.0.41", "windows", "AdbWinUsbApi.dll"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("windows AdbWinUsbApi sum:", sum) // 0e24119daf1909e398fa1850b6112077
	}
	sum, err = MD5FileStream(filepath.Join("1.0.41", "darwin", "adb"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("darwin adb sum:", sum) // f40ca3a5d903b9741cabafda838abc09
	}
	sum, err = MD5FileStream(filepath.Join("1.0.41", "linux", "adb"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("linux adb sum:", sum) // 930847adb8cc12623a2f712d30a5592b
	}
}
