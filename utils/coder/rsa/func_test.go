package rsa

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateKeyPair2File(t *testing.T) {
	os.Mkdir("keys", os.ModePerm)
	for i := 1; i <= 8; i++ {
		GenerateKeyPair2File(2048, fmt.Sprintf("keys/rsa_%d_pub", i), fmt.Sprintf("keys/rsa_%d", i))
	}

	var fileNames []string
	filepath.Walk("keys", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() && info.Name() == "keys" {
			return nil
		}
		fileNames = append(fileNames, info.Name())
		return nil
	})

	exceptFiles := []string{
		"rsa_1", "rsa_1_pub",
		"rsa_2", "rsa_2_pub",
		"rsa_3", "rsa_3_pub",
		"rsa_4", "rsa_4_pub",
		"rsa_5", "rsa_5_pub",
		"rsa_6", "rsa_6_pub",
		"rsa_7", "rsa_7_pub",
		"rsa_8", "rsa_8_pub",
	}

	Convey("files are created", t, func() {
		So(fileNames, ShouldEqual, exceptFiles)
	})
	fmt.Println(fileNames)
	os.RemoveAll("keys")
}

func TestGenerateKeyPair(t *testing.T) {
	for i := 1; i < 3; i++ {
		fmt.Printf("Pair %v ------------------------------------------------------------------ \n", i)
		_, _, err := GenerateKeyPair(2048)
		assert.Nil(t, t, err)
	}
}

func TestEncryptDecrypt(t *testing.T) {
	GenerateKeyPair2File(2048, fmt.Sprintf("keys/rsa_%d_pub", 1), fmt.Sprintf("keys/rsa_%d", 1))
	es, _ := Encrypt([]byte("你好，我是刘兴"), "keys/rsa_1_pub")
	ds, _ := Decrypt(es, "keys/rsa_1")
	assert.Equal(t, "你好，我是刘兴", string(ds))
}
