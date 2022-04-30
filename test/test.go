package main

import (
	"fmt"
	"sync"

	"github.com/tedcy/fdfs_client"
)

func TestUpload() {

	client, err := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	defer client.Destory()
	if err != nil {
		fmt.Println("NewClientWithConfig", err.Error())
		return
	}
	fileId, err := client.UploadByFilename("test.jpeg")
	if err != nil {
		fmt.Println("UploadByFilename", err.Error())
		return
	}
	fmt.Println("fileId:", fileId)
	if err := client.DownloadToFile(fileId, "tempFile", 0, 0); err != nil {
		fmt.Println("UploadByFilename:", err.Error())
		return
	}
	if buffer, err := client.DownloadToBuffer(fileId, 0, 0); err != nil {
		fmt.Println("DownloadToBuffer:", err.Error())
	} else {
		fmt.Println("DownloadToBuffer:", string(buffer))
	}

	// if err := client.DeleteFile(fileId); err != nil {
	// 	fmt.Println("DeleteFile", err.Error())
	// 	return
	// }
}

func TestUploadFile100() {
	client, err := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	defer client.Destory()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var wg sync.WaitGroup
	for i := 0; i != 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j != 10; j++ {
				if fileId, err := client.UploadByFilename("ddd.txt"); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(fileId)

					if _, err := client.DownloadToBuffer(fileId, 0, 0); err != nil {
						fmt.Println(err.Error())
					}
					if err := client.DeleteFile(fileId); err != nil {
						fmt.Println(err.Error())
					}
				}
			}
		}()
	}
	wg.Wait()
}

func TestUploadBuffer100() {
	client, err := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	defer client.Destory()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var wg sync.WaitGroup
	for i := 0; i != 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j != 1; j++ {
				if fileId, err := client.UploadByBuffer([]byte("hello world"), "go"); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(fileId)

					if _, err := client.DownloadToBuffer(fileId, 0, 0); err != nil {
						fmt.Println(err.Error())
					}
					//if err := client.DeleteFile(fileId); err != nil {
					//	fmt.Println(err.Error())
					// }
				}
			}
		}()
	}
	wg.Wait()
}

func main() {
	TestUpload()
	// TestUploadFile100()
	// TestUpload/Buffer100()
}
