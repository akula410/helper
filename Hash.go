package helper

import "golang.org/x/crypto/bcrypt"

type hash struct {

}

var Hash hash

func (h *hash) GetHashBCrypt(pass string)string{
	var hashPass, err =bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return Transform.ByteToBase64(hashPass)
}

func (h *hash) AskHashBCrypt(hashPass string, pass string)bool{
	var err = bcrypt.CompareHashAndPassword(Transform.Base64ToByte(hashPass), []byte(pass))
	return err == nil
}
