package SocialMediaMonitor_Tgbot

import (
	"bufio"
	"encoding/json"
	"fmt"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// DCMembersCount URL: Invite Link of a Discord Server.
func DCMembersCount(URL string) (int, int) {
	id := strings.Split(URL, "invite/")[1]
	urlJson := "https://discord.com/api/v9/invites/" + id + "?with_counts=true&with_expiration=true"
	res, err := http.Get(urlJson)
	if err != nil {
		fmt.Println("err", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	var DC DCJSON
	err = json.Unmarshal(data, &DC)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	return DC.ApproximateMemberCount, DC.ApproximatePresenceCount
}

// TwFollowersCount userName: the String After @ in User Profile
func TwFollowersCount(userName string) int {
	scraper := twitterscraper.New()
	profile, err := scraper.GetProfile(userName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", profile.FollowersCount)
	return profile.FollowersCount
}

// TGChatMembersCount URL: Link of Telegram Chat Group
func TGChatMembersCount(URL string) (string, string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	resp, _ := client.Do(req)
	buf := bufio.NewScanner(resp.Body)
	if resp.StatusCode == 404 {
		return "", ""
	}
	var list []string
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		re := regexp.MustCompile(`<div class="tgme_page_extra">([\s\S]*)</div>`)
		matchStr := re.FindStringSubmatch(line)
		if len(matchStr) != 0 {
			s1 := matchStr[len(matchStr)-1]
			list = append(list, s1)
		}
	}
	newStr1 := strings.Replace(list[0], " ", "", -1)
	fmt.Println(newStr1)

	newstrList1 := strings.Split(newStr1, "members")
	newstrList2 := strings.Split(newstrList1[1], "online")
	newstrList3 := strings.Split(newstrList2[0], ",")

	return newstrList1[0], newstrList3[1]
}

// TGChannelMembersCount URL: Link of Telegram Chat Group
func TGChannelMembersCount(URL string) string {
	url := URL
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(req)
	buf := bufio.NewScanner(resp.Body)
	if resp.StatusCode == 404 {

	}
	var list []string
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		re := regexp.MustCompile(`<div class="tgme_page_extra">([\s\S]*)</div>`)
		matchStr := re.FindStringSubmatch(line)
		if len(matchStr) != 0 {
			s1 := matchStr[len(matchStr)-1]
			list = append(list, s1)
		}
	}

	newStr := strings.Replace(list[0], " ", "", -1)
	newstrList := strings.Split(newStr, "subscribers")
	fmt.Println(newstrList[0])
	return newstrList[0]
}

type DCJSON struct {
	Code      string      `json:"code"`
	Type      int         `json:"type"`
	ExpiresAt interface{} `json:"expires_at"`
	Guild     struct {
		Id                       string   `json:"id"`
		Name                     string   `json:"name"`
		Splash                   string   `json:"splash"`
		Banner                   string   `json:"banner"`
		Description              string   `json:"description"`
		Icon                     string   `json:"icon"`
		Features                 []string `json:"features"`
		VerificationLevel        int      `json:"verification_level"`
		VanityUrlCode            string   `json:"vanity_url_code"`
		PremiumSubscriptionCount int      `json:"premium_subscription_count"`
		Nsfw                     bool     `json:"nsfw"`
		NsfwLevel                int      `json:"nsfw_level"`
		WelcomeScreen            struct {
			Description     string `json:"description"`
			WelcomeChannels []struct {
				ChannelId   string      `json:"channel_id"`
				Description string      `json:"description"`
				EmojiId     interface{} `json:"emoji_id"`
				EmojiName   *string     `json:"emoji_name"`
			} `json:"welcome_channels"`
		} `json:"welcome_screen"`
	} `json:"guild"`
	Channel struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"channel"`
	ApproximateMemberCount   int `json:"approximate_member_count"`
	ApproximatePresenceCount int `json:"approximate_presence_count"`
}
