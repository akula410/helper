package helper

import (
	"crypto/rand"
	"fmt"
	"io"
)

type uuid struct {
	separator string //separator UUID
}

var UUID uuid

func init(){
	UUID.separator = "-"
}

func (u *uuid) GetUUID(separator ...string)string{
	var s string
	if separator == nil {
		s = u.separator
	}else{
		s = separator[0]
	}

	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic(err)
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	//uuid[6] = uuid[6]&^0xf0 | 0x40


	return fmt.Sprintf("%x%s%x%s%x%s%x%s%x", uuid[0:4], s, uuid[4:6], s, uuid[6:8], s, uuid[8:10], s, uuid[10:])
}

func (u *uuid) GetUniqName(length int)string{
	uuid := make([]byte, int(length/2))
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic(err)
	}
	uuid[int(length/4)] = uuid[int(length/4)]&^0xc0 | 0x80
	return fmt.Sprintf("%x",uuid[0:])
}