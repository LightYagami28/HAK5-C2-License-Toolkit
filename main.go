package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)

type License struct {
	Key         string `json:"key"`
	Type        uint64 `json:"type"`
	UserLimit   uint64 `json:"user_limit"`
	DeviceLimit uint64 `json:"device_limit"`
	SiteLimit   uint64 `json:"site_limit"`
}

type Status struct {
	Hostname        string `json:"hostname"`
	Uptime          int64  `json:"uptime"`
	Version         string `json:"version"`
	UpdateAvailable bool   `json:"update_available"`
	UpdateChangelog string `json:"update_changelog"`
	UpdateVersion   string `json:"update_version"`
	UpdateDownload  string `json:"update_link"`
	Updating        bool   `json:"updating"`
	HostOS          string `json:"host_os"`
	Edition         string `json:"edition"`
	UserLimit       uint64 `json:"user_limit"`
	DeviceLimit     uint64 `json:"device_limit"`
	SiteLimit       uint64 `json:"site_limit"`
}

func main() {
	fmt.Println("[*] Hak5 C2 Licensing Toolkit by Pwn3rzs / CyberArsenal!")
	scanner := bufio.NewScanner(os.Stdin)

	db, err := bolt.Open("c2.db", 0600, nil)
	if err != nil {
		fmt.Println("[!] Error opening database:", err)
		return
	}
	defer db.Close()
	fmt.Println("[+] DB Opened!")

	for {
		fmt.Println("[*] Enter a command (generate/decode/read/crack/exit):")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "generate":
			generateHexCode()
		case "read":
			handleReadCommand(scanner, db)
		case "decode":
			handleDecodeCommand(scanner)
		case "crack":
			handleCrackCommand(db)
		case "exit":
			fmt.Println("[!] Exiting...")
			return
		default:
			fmt.Println("[!] Invalid command. Please enter 'generate', 'decode', 'read', or 'exit'.")
		}
	}
}

func handleReadCommand(scanner *bufio.Scanner, db *bolt.DB) {
	err := db.View(func(tx *bolt.Tx) error {
		fmt.Println("[*] Enter the bucket (setup/status):")
		scanner.Scan()
		bucket := scanner.Text()
		key := determineKey(bucket)

		if key == "" {
			return fmt.Errorf("[!] Invalid bucket. Please enter 'setup' or 'status'")
		}

		return readData(tx, bucket, key)
	})

	if err != nil {
		fmt.Println("[!] Error reading data:", err)
	}
}

func determineKey(bucket string) string {
	switch bucket {
	case "setup":
		return "license"
	case "status":
		return "status"
	default:
		return ""
	}
}

func readData(tx *bolt.Tx, bucket string, key string) error {
	bucketz := tx.Bucket([]byte(bucket))
	if bucketz == nil {
		return fmt.Errorf("[!] Bucket '%s' not found", bucket)
	}

	data := bucketz.Get([]byte(key))
	if data == nil {
		return fmt.Errorf("[!] Data for '%s' not found", key)
	}

	return decodeData(data, key)
}

func decodeData(data []byte, key string) error {
	var err error
	var license License
	var status Status

	decoder := gob.NewDecoder(bytes.NewReader(data))

	if key == "license" {
		err = decoder.Decode(&license)
		if err != nil {
			return err
		}
		prettyLicense, _ := json.MarshalIndent(license, "", "    ")
		fmt.Printf("[+] Decoded License struct: %+v\n", string(prettyLicense))
	} else {
		err = decoder.Decode(&status)
		if err != nil {
			return err
		}
		prettyStatus, _ := json.MarshalIndent(status, "", "    ")
		fmt.Printf("[+] Decoded Status struct: %+v\n", string(prettyStatus))
	}

	return nil
}

func handleDecodeCommand(scanner *bufio.Scanner) {
	fmt.Println("[*] Enter the struct (license/status):")
	scanner.Scan()
	mode := scanner.Text()

	fmt.Println("[*] Enter the hex string:")
	scanner.Scan()
	hexData := scanner.Text()

	if err := decodeHex(hexData, mode); err != nil {
		fmt.Println("Error decoding:", err)
	}
}

func handleCrackCommand(db *bolt.DB) {
	err := db.Update(func(tx *bolt.Tx) error {
		license := License{
			Key:         "Pwn3rzs",
			Type:        2,
			UserLimit:   10000,
			DeviceLimit: 10000,
			SiteLimit:   10000,
		}

		if err := saveData(tx, "setup", "license", license); err != nil {
			return err
		}

		status := Status{
			UserLimit:   10000,
			DeviceLimit: 10000,
			Edition:     "teams",
			SiteLimit:   10000,
		}
		return saveData(tx, "status", "status", status)
	})

	if err != nil {
		fmt.Println("[!] Error updating database:", err)
	} else {
		fmt.Println("[+] DB Values edited")
		fmt.Println("[*] Patching application")
	}
}

func saveData(tx *bolt.Tx, bucketName, key string, data interface{}) error {
	bucketz := tx.Bucket([]byte(bucketName))
	if bucketz == nil {
		return fmt.Errorf("[!] Bucket '%s' not found", bucketName)
	}

	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("[!] Encoding error: %v", err)
	}

	if err := bucketz.Put([]byte(key), buf.Bytes()); err != nil {
		return fmt.Errorf("[!] Error saving data: %v", err)
	}

	return nil
}

func generateHexCode() {
	license := License{
		Key:         "Pwn3rzs",
		Type:        1,
		UserLimit:   500,
		DeviceLimit: 500,
		SiteLimit:   500,
	}

	if hexCode, err := encodeToHex(license); err == nil {
		fmt.Println("[+] Generated Hex Code:", hexCode)
	} else {
		fmt.Println("[!] Encoding error:", err)
	}
}

func encodeToHex(data interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}

func decodeHex(hexData string, mode string) error {
	hexDecoded, err := hex.DecodeString(hexData)
	if err != nil {
		return fmt.Errorf("Error decoding hex string: %v", err)
	}

	var decoder *gob.Decoder
	if mode == "license" {
		var license License
		decoder = gob.NewDecoder(bytes.NewReader(hexDecoded))
		if err := decoder.Decode(&license); err != nil {
			return fmt.Errorf("Error decoding gob data: %v", err)
		}
		fmt.Printf("Decoded License struct: %+v\n", license)
	} else if mode == "status" {
		var status Status
		decoder = gob.NewDecoder(bytes.NewReader(hexDecoded))
		if err := decoder.Decode(&status); err != nil {
			return fmt.Errorf("Error decoding gob data: %v", err)
		}
		fmt.Printf("Decoded Status struct: %+v\n", status)
	} else {
		return fmt.Errorf("Invalid mode: %s. Please enter 'license' or 'status'.", mode)
	}

	return nil
}
