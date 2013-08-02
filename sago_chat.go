package sago_chat

import (
  "github.com/craigjackson/sago"
)

func chat(r *sago.Request) {
  if r.Session.Data["name"] == nil {
    res := make(map[string]interface{})
    res["message"] = "Need to auth"
    r.Session.Send(r.Id, "fail", res)
    return
  }

  if r.Args["message"] == nil {
    res := make(map[string]interface{})
    res["message"] = "Need to include the message"
    r.Session.Send(r.Id, "fail", res)
    return
  }

  res := make(map[string]interface{})
  res["name"] = r.Session.Data["name"]
  res["message"] = r.Args["message"]
  sago.SendAll("chat", res)
}

func Init() {
  chat_action := sago.AddAction("chat")
  go func() {
    for v := range chat_action.Channel {
      chat(v)
    }
  }()
}

