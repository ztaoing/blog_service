/**
* @Author:zhoutao
* @Date:2020/8/2 上午8:10
* @Desc:
 */

package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

//fatal error: all goroutines are asleep - deadlock!
func main() {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Println("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file", event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error", err)
		}
	}()
	//监听的目录或文件
	_ = watcher.Add("/Users/tao/moduleGo/blog_service/configs")
	<-done
}
