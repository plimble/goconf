package goconf

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

type OnChange func() error

func WatchYamlFile(path string, v interface{}, onChange OnChange) {
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			zap.L().Fatal(err.Error())
		}
		defer watcher.Close()

		configFile := filepath.Clean(path)
		configDir, _ := filepath.Split(configFile)

		done := make(chan bool)
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					if filepath.Clean(event.Name) == configFile {
						if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
							err := ReloadYaml(configFile, v)
							if err != nil {
								zap.L().Error(err.Error())
							}
							if onChange != nil {
								if err = onChange(); err != nil {
									zap.L().Error(err.Error())
								}
							}
						}
					}
				case err := <-watcher.Errors:
					log.Println("error:", err)
				}
			}
		}()

		watcher.Add(configDir)
		<-done
	}()
}
