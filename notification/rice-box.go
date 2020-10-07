package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "flairs1.html",
		FileModTime: time.Unix(1586630909, 0),

		Content: string("<!DOCTYPE html>\n<html>\n<head>\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n<style>\n\n\n@media only screen and (min-width: 600px) {\n   \n}\n@media only screen and (min-width: 768px) {\n\n}\nhtml {\n    /* font-family: \"Lucida Sans\", sans-serif;*/\n    font-family: Trebuchet MS;\n}\n.line {\n    background-color: #000;\n\theight: 1px;\n\tpadding: 1px;\n     line-height: 49px;\n}\n\n.header {\n    background-color: #f2f2f2;\n\theight: 600px;\n    color: purple;\n    padding: 35px;\n\tborder-radius: 10px;\n\t /*border: 1px solid;\n box-shadow: 0px 0px 5px #888888;*/\n \n  text-align: justify;\n}\n\n.head {\n    background-color: #999999;\n    color: #ffffff;\n    padding: 15px;\n\tborder-radius: 0px 0px 10px 10px;\n    font-size: 14px;\n    /*font-family: verdana, Helvetica, Arial, sans-serif;*/\n    line-height: 1.8;\n    text-align: center;\n\t\n}\n\n.footer {\n    background-color: #0099cc;\n    color: #ffffff;\n    text-align: center;\n    font-size: 19px;\n    padding: 15px;\n}\n\n.tosin {\n    background-color: #ffffff;\n    text-align: ;\n    color: #000;\n    letter-spacing: 1px;\n   /* padding: 15px;*/\n    \n    \n}\n\nul {\n  list-style: none;\n}\n\nul li::before {\n  content: \"\\2022\";\n  color: gold;\n  font-size: 24px;\n  font-weight: bold;\n  display: inline-block; \n  width: 1em;\n  margin-left: -1em;\n}\n\n\n\n</style>\n</head>\n\n<body style=\"padding-left: 200px; padding-right: 200px; padding-top: 100px; padding-bottom: 100px;\">\n\n<div class=\"tosin\">\n\n<img src=\"https://www.alphapay.live/img/welcomeMail/flairslogo.png\" alt=\"flairslogo\" style=\"width: 100%; max-width:200px; height: auto;\">\n   \n</div>\n\n<hr style=\"display: block; margin-top: 0.5em; margin-bottom: 2.5em; border-style: inset; border-width: 1px; color: #f2f2f2;\">\n\n<div class=\"header\">\n  <h2>Hello!</h2>\n  <span>Now that you have successfully created your Flairs account, you are welcome to the world of awesome financial services and limitless opportunities:</span>\n\n  <ul>\n  <li>Do all your <span style=\"font-weight: bold\">banking operations</span> all in one app,</li>\n  <li>Use your <span style=\"font-weight: bold\">Flairs VISA card</span> on all channels across the globe,</li>\n  <li>Same money on whooping <span style=\"font-weight: bold\">Flairs deals</span>,</li>\n  <li>Pay for bills and other <span style=\"font-weight: bold\">utilities</span>,</li>\n  <li>Get <span style=\"font-weight: bold\">expect finance advice</span> from Lola your Personal Finance Advisor,</li>\n  <li><span style=\"font-weight: bold\">Multiply your funds</span> with Wealth Manager & lot more!</li>\n</ul>\n\n<div class=\"ade\">\n\n<span style=\"font-weight: bold; font-size: 18px; line-height: 2.7;\">Welcome aboard!</span><br />\n<span style=\"font-weight: bold; font-size: 18px;line-height: 1.0;\">With Love, </span><br />\n\n<span style=\"font-weight: bold; font-size: 39px; color: red;padding: 30px;\">&#x2764;</span> <br />\n\n\n<div> <span><img src=\"https://www.alphapay.live/img/welcomeMail/flairslogo2.png\" alt=\"flairslogo2\" style=\"width: 100%; max-width:70px; height: auto;\"> </span> </div> \n<div style=\" padding-left: 90px; bottom:65px; text-align: justify; position: relative; font-size: 20px; font-weight: bold; font-style: italic; \" > <span>Lola, </span> <br> <span > Your Personal Finance Advisor</span> <br> <hr style=\"display: block; border-style: inset; border-width: 1px; color: #000; width: 300px; float: left; \"></div>\n\n\n</div>\n\n</div>\n\n\n\n<div class=\"head\">\n\n\n\t<div style=\"text-align: center; line-height: 10px; letter-spacing: 9px;\">  \n\t         <span><a href=\"https://bit.ly/flairsAppAndroid\"><img src=\"https://www.alphapay.live/img/welcomeMail/google.png\" alt=\"google\" style=\"width: 100%; max-width:120px; height: auto;\"></a> </span>\n             <span><a href=\"https://bit.ly/flairsAppiOS\"><img src=\"https://www.alphapay.live/img/welcomeMail/appstore.png\" alt=\"appstore\" style=\"width: 100%; max-width:120px; height: auto;\"></a> </span>\n\t</div>\n\t<div style=\"text-align: center; letter-spacing: 9px; line-height: 100px;\">  \n\t\t<span><img src=\"https://www.alphapay.live/img/welcomeMail/twitter.png\" alt=\"KingsChat\" style=\"width: 100%; max-width:30px; height: 30px;\"> </span>\n\t\t<span><img src=\"https://www.alphapay.live/img/welcomeMail/facebook.png\" alt=\"Facebook\" style=\"width: 100%; max-width:30px; height: 30px;\"> </span>\n\t\t<span><img src=\"https://www.alphapay.live/img/welcomeMail/bird.png\" alt=\"Twitter\" style=\"width: 100%; max-width:30px; height: 30px;\"> </span>\n\t\t<span><img src=\"https://www.alphapay.live/img/welcomeMail/instagram.png\" alt=\"Instagram\" style=\"width: 100%; max-width:30px; height: 30px;\"> </span>\n\t\t<span><img src=\"https://www.alphapay.live/img/welcomeMail/telegram.png\" alt=\"Telegram\" style=\"width: 100%; max-width:30px; height: 30px;\"> </span>\n\t\t<span><img src=\"https://www.alphapay.live/img/welcomeMail/chat.png\" alt=\"WhatsApp\" style=\"width: 100%; max-width:30px; height: 30px;\"> </span>\n\t\t<span><img src=\"https://www.alphapay.live/img/welcomeMail/mana.png\" alt=\"Medium\" style=\"width: 100%; max-width:30px; height: 30px;\"> </span>\n\t\t<span><img src=\"https://www.alphapay.live/img/welcomeMail/mail.png\" alt=\"Email\" style=\"width: 100%; max-width:30px; height: 30px;\"> </span>\n\t\t<span><img src=\"https://www.alphapay.live/img/welcomeMail/youtube.png\" alt=\"YouTube\" style=\"width: 100%; max-width:30px; height: 30px;\"> </span>\n\n\n\n\t</div>\n\n<span>\n  <a href=\"https://www.alphaplus.finance\" \n   style=\"text-decoration: none; color: white;\">www.alphaplus.finance</a>\n</span><br />\n\n  <span>Flairs is a product of AlphaPlus Financial technology &amp; Consulting Limited. <br>\n  Flairs is committed to protecting your privacy.For more information about \n    Flairs privacy policy, see Privacy Notice.<br>\n  2020 &#169; Copyright. All rights reserved.</span><br />\n  \n\n<span>\n  <a href=\"#\" style=\"text-decoration: none; color: gold; float: right;\">Click to unsubscribe</a></span><br>\n</div>\n\n\n\n\n</div>\n\n</body>\n</html>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1597253302, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "flairs1.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`html`, &embedded.EmbeddedBox{
		Name: `html`,
		Time: time.Unix(1597253302, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"flairs1.html": file2,
		},
	})
}
