package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v3"
)

type Order struct {
	Email            string
	Description      string
	Images           []*multipart.FileHeader
	MarketingConsent bool
}

type TelegramMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type Media struct {
	Type      string `json:"type"`                 // "photo"
	Media     string `json:"media"`                // "attach://file1"
	Caption   string `json:"caption,omitempty"`    // only first item
	ParseMode string `json:"parse_mode,omitempty"` // MarkdownV2
}

func HandleForm(c *gin.Context) {

	order := formToStruct(c)

	//msg := formatOrderMessage(order)
	chatID := getChatID()

	sendOrderToTelegram(order, chatID)
	client := resend.NewClient(os.Getenv("RESEND_KEY"))

	if order.MarketingConsent {
		go func(email string) {


			params := &resend.CreateContactRequest{
				Email: email,
			}
			_, err := client.Contacts.Create(params)
			if err != nil {
				log.Printf("Failed to create contact, %v", err)
			}
		}(order.Email)
	}
	ctx := context.TODO()
	params := &resend.SendEmailRequest{
    	From:        "bogrik.com <no-reply@bogrik.com>",
    	To:          []string{order.Email},
    	Subject:     "Your order is submitted",
    	Html:        "Thank you for submitting the order. I will contact you as soon as possible with the order details",
  	}

  	sent, err := client.Emails.SendWithContext(ctx, params)

  	if err != nil {
    	panic(err)
  	}
  	fmt.Println(sent.Id)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}

func sendTelegramMessage(chatID int64, text string) error {
	token := os.Getenv("TELEGRAM_BOT_KEY")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	msg := TelegramMessage{
		ChatID: chatID,
		Text:   text,
	}
	body, _ := json.Marshal(msg)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("telegram API returned status %d", resp.StatusCode)
	}

	return nil
}

func sendOrderToTelegram(order Order, chatID int64) {
    text := formatOrderMessage(order)
    if len(order.Images) == 0 {
        // just text message
        go func() {
            err := sendTelegramMessage(chatID, text)
            if err != nil {
                log.Printf("Failed to send message: %v", err)
            }
        }()
    } else {
        // text + multiple images as media group
        go func() {
            err := sendTelegramMediaGroup(chatID, order.Images, text)
            if err != nil {
                log.Printf("Failed to send media group: %v", err)
            }
        }()
    }
}

func sendTelegramMediaGroup(chatID int64, files []*multipart.FileHeader, caption string) error {
    token := os.Getenv("TELEGRAM_BOT_TOKEN")
    url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMediaGroup", token)

    // multipart body
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    media := []Media{}

    for i, f := range files {
        filename := fmt.Sprintf("file%d", i+1)
        part, err := writer.CreateFormFile(filename, f.Filename)
        if err != nil {
            return err
        }

        file, err := f.Open()
        if err != nil {
            return err
        }

        _, err = io.Copy(part, file)
        file.Close()
        if err != nil {
            return err
        }

        m := Media{
            Type:  "photo",
            Media: "attach://" + filename,
        }
        if i == 0 {
            m.Caption = caption
            m.ParseMode = "MarkdownV2"
        }

        media = append(media, m)
    }

    mediaJSON, _ := json.Marshal(media)
    writer.WriteField("chat_id", fmt.Sprintf("%d", chatID))
    writer.WriteField("media", string(mediaJSON))

    writer.Close()

    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 300 {
        return fmt.Errorf("telegram returned status: %s", resp.Status)
    }

    return nil
}

func formToStruct(c *gin.Context) Order {
	var order Order

	order.Email = c.PostForm("email")
	order.Description = c.PostForm("description")
	order.MarketingConsent = c.PostForm("marketingConsent") != ""

	// Parse uploaded files (optional: limit to 5MB each)
	c.Request.ParseMultipartForm(10 << 20)               // 10MB max total
	order.Images = c.Request.MultipartForm.File["references"] // name="files" in form

	return order
}

func formatOrderMessage(order Order) string {
	return fmt.Sprintf(
		"📩 New Order Received\n\n"+
			"Email: %s\n"+
			"Description: %s\n",
		order.Email,
		order.Description,
	)
}

func getChatID() int64 {
	chatIDStr := os.Getenv("CHAT_ID")

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid CHAT_ID: %v", err)
	}

	return chatID
}
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, order handling will not work properly")
	}
}
