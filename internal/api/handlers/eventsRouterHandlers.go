package handlers

import (
	"TunaAPIGateway/config"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"slices"
	"strconv"
)

func PostEvent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	eventTypes := []string{"playback", "playlist", "user"}
	queryParams := r.URL.Query()
	userId := queryParams.Get("user_id")
	eventType := params.ByName("event_type")

	if slices.Contains(eventTypes, eventType) {
		url := fmt.Sprintf(
			"http://%s:%d/events_gateway/events/%s?user_id=%s",
			config.Config.API.EventsGatewayHost, config.Config.API.EventsGatewayPort,
			eventType, userId,
		)
		req, _ := http.NewRequest("POST", url, r.Body)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, _ := client.Do(req)
		respBody, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != 200 {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(respBody)
			return
		}
		w.Write([]byte(strconv.Itoa(resp.StatusCode) + " OK"))
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
