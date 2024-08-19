package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/coder/websocket"
	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
)

func mirabelleStreamCmd() *cobra.Command {
	var query string
	var streamCmd = &cobra.Command{
		Use:   "stream",
		Short: "Stream events from Mirabelle",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			c, _, err := websocket.Dial(ctx, fmt.Sprintf("ws://localhost:5558/channel/my-channel?query=%s", query), nil)
			cancel()
			exitIfError(err)

			signals := make(chan os.Signal, 1)
			signal.Notify(
				signals,
				syscall.SIGINT,
				syscall.SIGTERM)

			var wg sync.WaitGroup
			wg.Add(1)
			wsCtx, wsCancel := context.WithCancel(context.Background())
			go func() {
				defer wg.Done()
				ticker := time.NewTicker(8 * time.Second)
				for {
					select {
					case <-wsCtx.Done():
						return
					case <-ticker.C:
						ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
						err := c.Ping(ctx)
						cancel()
						if err != nil {
							fmt.Printf("fail to send websocket heartbeat: %s", err.Error())
						}
					}

				}
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case s := <-signals:
						if s == syscall.SIGINT || s == syscall.SIGTERM {
							wsCancel()
							return
						}
					}
				}

			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				formatter := prettyjson.NewFormatter()
				for {
					_, result, err := c.Read(wsCtx)
					if err != nil {
						if errors.Is(err, context.Canceled) || strings.Contains(err.Error(), "use of closed network connection") {
							return
						}
						fmt.Println("websocket error: " + err.Error())
						continue
					}
					indented := &bytes.Buffer{}
					if err := json.Indent(indented, result, "", "  "); err != nil {
						fmt.Println("fail to parse JSON response: " + err.Error())

					}
					fmt.Println(time.Now().Format("2006-01-02 15:04:05.00000"))
					data, err := formatter.Format(result)
					if err != nil {
						fmt.Println("fail to format JSON response: " + err.Error())
					}
					fmt.Println(string(data))
				}
			}()
			wg.Wait()
			c.Close(websocket.StatusNormalClosure, "")

		},
	}
	streamCmd.PersistentFlags().StringVar(&query, "query", "true", "Query to execute")

	return streamCmd

}
