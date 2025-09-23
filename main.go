package main

import (
	"debug/macho"
	"debug/pe"
	"fmt"
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

func main() {
	fmt.Printf("操作系统: %s\n", runtime.GOOS)
	fmt.Printf("架构: %s\n", runtime.GOARCH)

	// 判断具体架构
	switch runtime.GOARCH {
	case "amd64":
		fmt.Println("这是 64 位 Windows (x86-64)")
	case "386":
		fmt.Println("这是 32 位 Windows (x86)")
	case "arm64":
		fmt.Println("这是 ARM64 Windows")
	default:
		fmt.Printf("未知架构: %s\n", runtime.GOARCH)
	}

	var bit string
	var err error
	bit, err = checkExeArchitecture(filepath.Join("1.0.41", "windows", "adb.exe"))
	fmt.Println(bit, err)

	bit, err = checkMachOArchitecture(filepath.Join("1.0.41", "darwin", "amd64", "adb"))
	fmt.Println(bit, err)

	bit, err = checkMachOArchitecture(filepath.Join("1.0.41", "darwin", "arm64", "adb"))
	fmt.Println(bit, err)
}
