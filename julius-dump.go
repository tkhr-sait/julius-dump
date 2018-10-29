package main

import "strings"
import "net"
import "bufio"
import "regexp"
import "encoding/xml"
//import "log"
import "fmt"

type RECOGOUT struct {
  SHYPO []struct {
    RANK    string `xml:"RANK,attr"`
    SCORE   string `xml:"SCORE,attr"`
    WHYPO []struct {
      WORD     string `xml:"WORD,attr"`
      PHONE    string `xml:"PHONE,attr"`
      CM       string `xml:"CM,attr"`
    } `xml:"WHYPO,omitempty"`
  } `xml:"SHYPO,omitempty"`
}

func main() {
  conn,err := net.Dial("tcp","localhost:10500")
  if err != nil {
    panic(err)
  }
  conn.Write([]byte("TERMINATE\n"))
  conn.Write([]byte("RESUME\n"))
  var buf = ""
  reader := bufio.NewReader(conn)
  for {
    message, err := reader.ReadString('\n')
    if err != nil {
       panic(err)
    }
    if (message != ".\n") {
      // goのパーサではclassidにタグが記載されていると
      // パースできないのでomitする
      if strings.Contains(message,"CLASSID") {
        r := regexp.MustCompile(`(.*)CLASSID="[^"]*"(.*)`)
        message = r.ReplaceAllString(message,"$1$2")
      }
      buf = buf + message
    } else {
      // log.Print(buf)
      if strings.HasPrefix(buf,"<RECOGOUT") {
        recogout := RECOGOUT{}
        xml.Unmarshal([]byte(buf),&recogout)
        // log.Print(recogout)
        word := ""
        for idx := 0; idx < len(recogout.SHYPO); idx++ {
          for idx2 := 0; idx2 < len(recogout.SHYPO[idx].WHYPO); idx2++ {
            word = word + recogout.SHYPO[idx].WHYPO[idx2].WORD
          }
        }
        fmt.Println(word)
      }
      buf = ""
    }
  }
}
