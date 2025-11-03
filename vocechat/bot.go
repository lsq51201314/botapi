package vocechat

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Callback func(b *Bot, gid int64, msg string)

type Bot struct {
	url  string
	key  string
	port int
	call Callback
}

func New(url, key string, port ...int) *Bot {
	obj := Bot{
		url:  url,
		key:  key,
		port: 80,
	}
	if len(port) > 0 {
		obj.port = port[0]
	}
	return &obj
}

func (b *Bot) SetCallback(cfunc Callback) {
	b.call = cfunc
}

func (b *Bot) Run() {
	//运行服务器
	r := b.request()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", b.port),
		Handler: r,
	}
	func() {
		log.Printf("服务器正在运行(%d)...\n", b.port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("服务器正在关闭...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
	log.Println("服务器关闭成功...")
}
