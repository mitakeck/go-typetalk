package typetalk

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_TalksService_CreateTalk_should_create_a_talk(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	talkName := "test"
	b, _ := ioutil.ReadFile("../testdata/create-talk.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/talks", topicId),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testFormValues(t, r, values{
				"talkName":   talkName,
				"postIds[0]": 1,
				"postIds[1]": 2,
				"postIds[2]": 3,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Talks.CreateTalk(context.Background(), topicId, talkName, 1, 2, 3)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &CreatedTalkResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TalksService_UpdateTalk_should_update_a_talk_name(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	talkId := 1
	talkName := "test"
	b, _ := ioutil.ReadFile("../testdata/update-talk.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/talks/%d", topicId, talkId),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PUT")
			testQueryValues(t, r, values{
				"talkName": talkName,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Talks.UpdateTalk(context.Background(), topicId, talkId, talkName)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &UpdatedTalkResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TalksService_DeleteTalk_should_delete_a_talk(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	talkId := 1
	b, _ := ioutil.ReadFile("../testdata/update-talk.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/talks/%d", topicId, talkId),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Talks.DeleteTalk(context.Background(), topicId, talkId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &DeletedTalkResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TalksService_GetTalkList_should_get_talk_list(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile("../testdata/get-talk-list.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/talks", topicId),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Talks.GetTalkList(context.Background(), topicId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		Talks []*Talk `json:"talks"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Talks) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TalksService_GetMessagesInTalk_should_get_some_messages_in_talk(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	talkId := 1
	b, _ := ioutil.ReadFile("../testdata/get-messages-in-talk.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/talks/%d/posts", topicId, talkId),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testQueryValues(t, r, values{
				"count":     10,
				"from":      3,
				"direction": "forward",
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Talks.GetMessagesInTalk(context.Background(), topicId, talkId, &GetMessagesOptions{10, 3, "forward"})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &MessagesInTalk{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TalksService_AddMessageToTalk_should_add_some_messages_to_talk(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	talkId := 1
	b, _ := ioutil.ReadFile("../testdata/add-messages-to-talk.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/talks/%d/posts", topicId, talkId),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testFormValues(t, r, values{
				"postIds[0]": 1,
				"postIds[1]": 2,
				"postIds[2]": 3,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Talks.AddMessagesToTalk(context.Background(), topicId, talkId, 1, 2, 3)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &MessagesInTalk{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TalksService_RemoveMessagesFromTalk_should_remove_some_messages_from_talk(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	talkId := 1
	b, _ := ioutil.ReadFile("../testdata/remove-messages-from-talk.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/talks/%d/posts", topicId, talkId),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			testQueryValues(t, r, values{
				"postIds[0]": 1,
				"postIds[1]": 2,
				"postIds[2]": 3,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Talks.RemoveMessagesFromTalk(context.Background(), topicId, talkId, 1, 2, 3)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &RemovedMessagesResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}
