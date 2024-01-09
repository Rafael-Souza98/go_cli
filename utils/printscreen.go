package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)


func ChromeScreenShot(url string, quality int) {
	screenshotURL:= fmt.Sprintf("https://%s/", url)
	var buf []byte

	var ext string = "png"
		if quality < 100 {
			ext = "jpg"
		}

		log.Printf("Capturing the screen of site %s", screenshotURL)

		var options []chromedp.ExecAllocatorOption
		options = append(options, chromedp.WindowSize(1920, 1080))
		options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)

		actx, acancel:= chromedp.NewExecAllocator(context.Background(), options...)	
		defer acancel()

		ctx, cancel:= chromedp.NewContext(actx)
		defer cancel()

		tasks:= chromedp.Tasks{
			chromedp.Navigate(screenshotURL),
			chromedp.Sleep(time.Second * 4),
			chromedp.CaptureScreenshot(&buf),
		}
		if err:= chromedp.Run(ctx, tasks); err != nil {
			log.Fatal(err)
		}
		filename:= fmt.Sprintf("%s-%d-standard.%s", strings.Replace(url, "/", "-", -1), time.Now().UTC().Unix(), ext)
		
		if err:= os.WriteFile(filename, buf, 0644); err != nil {
			log.Fatal(err)
		}
		log.Printf("Cpature saves in %s", filename)
	}