package main

import (
	"log"

	"github.com/Shark/piknik/client"
)

func main() {
	log.SetFlags(0)

	// isCopy := flag.Bool("copy", false, "store content (copy)")
	// _ = flag.Bool("paste", false, "retrieve the content (paste) - this is the default action")
	// isMove := flag.Bool("move", false, "retrieve and delete the clipboard content")
	// maxLenMb := flag.Uint64("maxlen", 0, "maximum content length to accept in Mb (0=unlimited)")
	// timeout := flag.Uint("timeout", 10, "connection timeout (seconds)")
	// dataTimeout := flag.Uint("datatimeout", 3600, "data transmission timeout (seconds)")

	var conf client.Conf
	// pskHex := tomlConf.Psk
	// psk, err := hex.DecodeString(pskHex)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// conf.Psk = psk
	// if encryptSkHex := tomlConf.EncryptSk; encryptSkHex != "" {
	// 	encryptSk, err := hex.DecodeString(encryptSkHex)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	conf.EncryptSk = encryptSk
	// }
	// if signPkHex := tomlConf.SignPk; signPkHex != "" {
	// 	signPk, err := hex.DecodeString(signPkHex)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	conf.SignPk = signPk
	// }
	// if encryptSkID := tomlConf.EncryptSkID; encryptSkID > 0 {
	// 	conf.EncryptSkID = make([]byte, 8)
	// 	binary.LittleEndian.PutUint64(conf.EncryptSkID, encryptSkID)
	// } else if len(conf.EncryptSk) > 0 {
	// 	hf, _ := blake2b.New(&blake2b.Config{
	// 		Person: []byte(DomainStr),
	// 		Size:   8,
	// 	})
	// 	hf.Write(conf.EncryptSk)
	// 	encryptSkID := hf.Sum(nil)
	// 	encryptSkID[7] &= 0x7f
	// 	conf.EncryptSkID = encryptSkID
	// }
	// conf.TTL = DefaultTTL
	// if ttl := tomlConf.TTL; ttl > 0 {
	// 	conf.TTL = time.Duration(ttl) * time.Second
	// }
	// if signSkHex := tomlConf.SignSk; signSkHex != "" {
	// 	signSk, err := hex.DecodeString(signSkHex)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	switch len(signSk) {
	// 	case 32:
	// 		if len(conf.SignPk) != 32 {
	// 			log.Fatal("Public signing key required")
	// 		}
	// 		signSk = append(signSk, conf.SignPk...)
	// 	case 64:
	// 	default:
	// 		log.Fatal("Unsupported length for the secret signing key")
	// 	}
	// 	conf.SignSk = signSk
	// }
	// conf.MaxClients = *maxClients
	// conf.Timeout = time.Duration(*timeout) * time.Second
	// conf.DataTimeout = time.Duration(*dataTimeout) * time.Second
	isCopy := true
	isMove := false
	client.RunClient(conf, isCopy, isMove)
}
