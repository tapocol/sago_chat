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

  args := *r.Args
  if args["message"] == nil {
    res := make(map[string]interface{})
    res["message"] = "Need to include the message"
    r.Session.Send(r.Id, "fail", res)
    return
  }

  res := make(map[string]interface{})
  res["name"] = r.Session.Data["name"]
  res["message"] = args["message"]
  sago.SendAll("chat", res)
}

func Init() {
  sago.AddAction("chat", chat)
}

