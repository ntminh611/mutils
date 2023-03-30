package stored

import (
	"encoding/json"
	"os"
	"sync"
)

var saveFile sync.Mutex

func SaveFile(data interface{}, nameFile string, encrypt func([]byte) []byte) error {
	saveFile.Lock()
	defer saveFile.Unlock()
	djson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if encrypt != nil {
		djson = encrypt(djson)
	}

	return os.WriteFile(nameFile, djson, 0644)
}

// Ex Input: LoadFromFile(&result, "result", nil)
func LoadFromFile(result interface{}, nameFile string, decrypt func([]byte) []byte) error {
	data, err := os.ReadFile(nameFile)
	if err != nil {
		return err
	}
	if decrypt != nil {
		data = decrypt(data)
	}
	err = json.Unmarshal(data, &result)
	return err
}
