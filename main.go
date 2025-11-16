package main

import (
	"io"
	"log"
	"os"

	"github.com/personalconnect/dragpass-keeper/internal/keystore"
)

// keystore에 저장되는 항목들:
// - 서버 공개키 (init시 저장)
// - 디바이스키
// - 세션코드
// - Keeper 비공개키

// (generatekeypair) 키페어 생성 요청 [Internal: 세션 코드 삭제, 기존 키페어 삭제, 새 키페어 저장]
// (savekey) 디바이스키 저장 요청
// (deletekey) 디바이스키 삭제 요청
// (getkey) 디바이스키 조회 요청
// (getsessioncode) 세션코드 조회 요청
// (savesessioncode) 세션코드 저장 요청
// (getserverpubkey) 서버 공개키 조회 요청
// (getprivatekey) Keeper 비공개키 조회 요청
// (getpublickey) Keeper 공개키 조회 요청

func main() {
	log.SetOutput(os.Stderr)
	log.Println("dragpass extension helper started")

	for {
		req, err := keystore.ReadMessage()
		if err != nil {
			if err == io.EOF {
				log.Println("lost chrome extension connection")
				break
			}
			keystore.SendResponse(keystore.ResponseMessage{Success: false, Error: "wrong message format: " + err.Error()})
			continue
		}
		keystore.LogRequest(req)
		resp := keystore.HandleRequest(req)
		keystore.SendResponse(resp)
	}
}
