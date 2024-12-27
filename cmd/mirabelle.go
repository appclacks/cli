package cmd

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/coder/websocket"
	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
)

func mirabelleHost() string {
	mirabelleHost := os.Getenv("MIRABELLE_API_ENDPOINT")
	if mirabelleHost == "" {
		mirabelleHost = "localhost:5558"
	}
	return mirabelleHost
}

func mirabelleSubscribeCmd() *cobra.Command {
	var query string
	var channel string
	var subscribeCmd = &cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to a Mirabelle channel",
		Run: func(cmd *cobra.Command, args []string) {
			options := websocket.DialOptions{
				HTTPHeader: http.Header{},
			}
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			host := mirabelleHost()
			values := url.Values{}
			values.Add("query", base64.StdEncoding.EncodeToString([]byte(query)))
			c, _, err := websocket.Dial(ctx, fmt.Sprintf("ws://%s/channel/%s?%s", host, channel, values.Encode()), &options)
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
					case <-wsCtx.Done():
						return
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
						fmt.Println("websocket error: " + err.Error())
						wsCancel()
						return
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
	subscribeCmd.PersistentFlags().StringVar(&query, "query", "[:true]", "Query to execute")
	subscribeCmd.PersistentFlags().StringVar(&channel, "channel", "", "Channel to subscribe to")
	err := subscribeCmd.MarkPersistentFlagRequired("channel")
	exitIfError(err)

	return subscribeCmd
}

type Event struct {
	Host        string            `json:"host,omitempty"`
	Metric      float64           `json:"metric,omitempty"`
	Service     string            `json:"service,omitempty"`
	Time        int64             `json:"time"`
	Name        string            `json:"name,omitempty"`
	State       string            `json:"state,omitempty"`
	Description string            `json:"description,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

type EventPayload struct {
	Events []Event `json:"events"`
}

func mirabelleEventCmd() *cobra.Command {
	var service string
	var metric float64
	var name string
	var attributes []string
	var state string
	var description string
	var stream string
	var tags []string

	var sendCmd = &cobra.Command{
		Use:   "send",
		Short: "Send an event to Mirabelle",
		Run: func(cmd *cobra.Command, args []string) {
			now := time.Now().UnixNano()
			hostname, err := os.Hostname()
			exitIfError(err)
			attributesMap, err := toMap(attributes)
			exitIfError(err)
			event := Event{
				Metric:      metric,
				Host:        hostname,
				Service:     service,
				Time:        now,
				Name:        name,
				State:       state,
				Description: description,
				Tags:        tags,
				Attributes:  attributesMap,
			}
			client := http.Client{}
			var reqBody io.Reader
			payload := EventPayload{
				Events: []Event{event},
			}
			json, err := json.Marshal(payload)
			exitIfError(err)
			reqBody = bytes.NewBuffer(json)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			host := mirabelleHost()
			request, err := http.NewRequestWithContext(
				ctx,
				http.MethodPut,
				fmt.Sprintf("http://%s/api/v1/stream/%s", host, stream),
				reqBody)
			exitIfError(err)
			request.Header.Add("content-type", "application/json")
			response, err := client.Do(request)
			exitIfError(err)
			b, err := io.ReadAll(response.Body)
			exitIfError(err)
			defer response.Body.Close()
			if response.StatusCode >= 300 {
				exitIfError(fmt.Errorf("Fail to send event (status %s): %s", response.Status, string(b)))
			} else {
				fmt.Println("Event successfully sent")
			}

		},
	}
	sendCmd.PersistentFlags().StringVar(&stream, "stream", "default", "The stream to which the event shold be sent")
	sendCmd.PersistentFlags().StringVar(&service, "service", "", "Service name")
	sendCmd.PersistentFlags().Float64Var(&metric, "metric", 0, "Metric value (float64)")
	sendCmd.PersistentFlags().StringVar(&name, "name", "", "Event name")
	sendCmd.PersistentFlags().StringVar(&description, "description", "", "Event description")
	sendCmd.PersistentFlags().StringVar(&state, "state", "", "Event state")
	sendCmd.PersistentFlags().StringSliceVar(&attributes, "attributes", []string{}, "key-value attributes (example: foo=bar)")
	sendCmd.PersistentFlags().StringSliceVar(&tags, "tags", []string{}, "Event tags")

	return sendCmd
}
