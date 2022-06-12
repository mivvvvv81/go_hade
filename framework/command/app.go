package command

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/framework/cobra"
	"project/framework/contract"
	"syscall"
	"time"
)

func initAppCommand() *cobra.Command {
	appCommand.AddCommand(appStartCommand)
	return appCommand
}

var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务应该控制命令",
	Long:  "业务应用控制命令，其中包含业务启动，关闭，重启，查询等功能",
	RunE: func(c *cobra.Command, args []string) error {
		c.Help()
		return nil
	},
}

var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个web服务",
	RunE: func(c *cobra.Command, args []string) error {
		//从command中获取服务容器
		container := c.GetContainer()
		//从服务容器中获取kernel的服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		//从kernel服务实例中获取引擎
		core := kernelService.HttpEngine()

		//创建一个service服务
		server := &http.Server{
			Handler: core,
			Addr:    ":8888",
		}

		go func() {
			server.ListenAndServe()
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		return nil

	},
}
