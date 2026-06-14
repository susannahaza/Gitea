package webhook

import (
	"net/http"
	"models/webhook"
)

func Deliver(task *webhook.HookTask, client *http.Client, req *http.Request) error {
	resp, err := client.Do(req)
	if err != nil {
		task.IsSucceed = false
		task.Error = err.Error()
		return webhook.UpdateHookTask(task)
	}
	defer resp.Body.Close()

	task.IsSucceed = true
	return webhook.UpdateHookTask(task)
}