package ntfyclient

import (
	"net/http"
	"strconv"
	"strings"
)

type Priority int

const (
	Min Priority = iota + 1
	Low
	Default
	High
	Max
)

func send(text string, url string, priority Priority, md bool, title string, tags []string) {
	req, _ := http.NewRequest("POST", url, strings.NewReader(text))
	req.Header.Set("Content-Type", "text/plain")
	if md {
		req.Header.Set("Markdown", "yes")
	}
	req.Header.Set("Priority", strconv.Itoa(int(priority)))
	req.Header.Set("Title", title)
	req.Header.Add("Tags", strings.Join(tags, ","))
	http.DefaultClient.Do(req)
}

type Client struct {
	url string
	tag *string
}

func NewClient(url string, tag *string) *Client {
	return &Client{url, tag}
}

func (c *Client) Send(text ...string) {
	if c.tag != nil {
		text = []string{*c.tag, strings.Join(text, " ")}
	}
	send(strings.Join(text, " "), c.url, Default, false, "", nil)
}

type Message struct {
	Text     string
	Priority Priority
	Markdown bool
	Title    string
	Tags     []string
}

func (c *Client) SendMessage(message Message) {
	send(message.Text, c.url, message.Priority, message.Markdown, message.Title, message.Tags)
}

func (c *Client) SendError(message string, stack string) {
	c.SendMessage(Message{
		Text:     message + "\n```\n" + stack + "\n```",
		Priority: Max,
		Markdown: true,
		Title:    "ERROR",
		Tags:     []string{"rotating_light"},
	})
}

func (c *Client) SendWarning(message string) {
	c.SendMessage(Message{
		Text:     message,
		Priority: High,
		Tags:     []string{"warning"},
	})
}

func (c *Client) SendDebug(message string) {
	c.SendMessage(Message{
		Text:     message,
		Priority: Min,
		Tags:     []string{"white_circle"},
	})
}
