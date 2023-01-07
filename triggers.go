package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Replies []struct {
	pattern  string
	response string
}

func replyToChannel(channel string, message string, nick string) {
	for _, regex := range replies {
		re, err := regexp.Compile(regex.pattern)

		if err != nil {
			continue
		}

		matches := re.FindStringSubmatch(strings.TrimSpace(message))

		if len(matches) == 0 {
			continue
		}

		response := strings.Replace(regex.response, "$nick", nick, 1)
		response = strings.Replace(response, "$chan", channel, 1)

		send(fmt.Sprintf("PRIVMSG %s :%s", channel, response))
	}
}

var replies = Replies{
	{
		pattern:  "^some people are being fangoriously devoured by a gelatinous monster$",
		response: "Hillary's legs are being digested.",
	},
	{
		pattern:  "^hello fishbot$",
		response: "Hi $nick",
	},
	{
		pattern:  "^badger badger badger badger badger badger badger badger badger badger badger badger$",
		response: "mushroom mushroom!",
	},
	{
		pattern:  "^snake$",
		response: "Ah snake a snake! Snake, a snake! Ooooh, it's a snake!",
	},
	{
		pattern:  "^carrots handbags cheese\\.$",
		response: "toilets russians planets hamsters weddings poets stalin KUALA LUMPUR! pygmies budgies KUALA LUMPUR!",
	},
	{
		pattern:  "^sledgehammer$",
		response: "sledgehammers go quack!",
	},
	{
		pattern:  "^vinegar.*?aftershock.$",
		response: "Ah, a true connoisseur!",
	},
	{
		pattern:  "^moo\\?$",
		response: "To moo, or not to moo, that is the question. Whether 'tis nobler in the mind to suffer the slings and arrows of outrageous fish...",
	},
	{
		pattern:  "^herring$",
		response: "herring(n): Useful device for chopping down tall trees. Also moos (see fish).",
	},
	{
		pattern:  "^hampster$",
		response: "$nick: There is no 'p' in hamster you retard.",
	},
	{
		pattern:  "^ag$",
		response: "Ag, ag ag ag ag ag AG AG AG!",
	},
	{
		pattern:  "^fishbot owns$",
		response: "Aye, I do.",
	},
	{
		pattern:  "^vinegar$",
		response: "Nope, too sober for vinegar. Try later.",
	},
	{
		pattern:  "^martian$",
		response: "Don't run! We are your friends!",
	},
	{
		pattern:  "^just then\\, he fell into the sea$",
		response: "Ooops!",
	},
	{
		pattern:  "^aftershock$",
		response: "mmmm, Aftershock.",
	},
	{
		pattern:  "^why are you here\\?$",
		response: "Same reason. I love candy.",
	},
	{
		pattern:  "^spoon$",
		response: "There is no spoon.",
	},
	{
		pattern:  "^bounce$",
		response: "moo",
	},
	{
		pattern:  "^crack$",
		response: "Doh, there goes another bench!",
	},
	{
		pattern:  "^you can\\'t just pick people at random\\!$",
		response: "I can do anything I like, $nick, I'm eccentric! Rrarrrrrgh! Go!",
	},
	{
		pattern:  "^wertle$",
		response: "moo",
	},
	{
		pattern:  "^flibble$",
		response: "plob",
	},
	{
		pattern:  "^www\\.outwar\\.com.$",
		response: "would you please GO AWAY with that outwar rubbish!",
	},
	{
		pattern:  "^fishbot created splidge$",
		response: "omg no! Think I could show my face around here if I was responsible for THAT?",
	},
	{
		pattern:  "^now there\\'s more than one of them\\?$",
		response: "A lot more.",
	},
	{
		pattern:  "^I want everything$",
		response: "Would that include a bullet from this gun?",
	},
	{
		pattern:  "^we are getting aggravated\\.$",
		response: "Yes, we are.",
	},
	{
		pattern:  "how old are you, fishbot\\?$",
		response: "\x01ACTION is older than time itself.\x01",
	},
	{
		pattern:  "^atlantis$",
		response: "Beware the underwater headquarters of the trout and their bass henchmen. From there they plan their attacks on other continents.",
	},
	{
		pattern:  "^oh god$",
		response: "fishbot will suffice.",
	},
	{
		pattern:  "^fishbot$",
		response: "Yes?",
	},
	{
		pattern:  "^what is the matrix\\?$",
		response: "No-one can be told what the matrix is. You have to see it for yourself.",
	},
	{
		pattern:  "^what do you need\\?$",
		response: "Guns. Lots of guns.",
	},
	{
		pattern:  "^I know Kungfu$",
		response: "Show me.",
	},
	{
		pattern:  "^cake$",
		response: "fish",
	},
	{
		pattern:  "^trout go moo$",
		response: "Aye, that's cos they're fish.",
	},
	{
		pattern:  "^Kangaroo$",
		response: "The kangaroo is a four winged stinging insect.",
	},
	{
		pattern:  "^bass$",
		response: "Beware of the mutant sea bass and their laser cannons!",
	},
	{
		pattern:  "^trout$",
		response: "Trout are freshwater fish and have underwater weapons.",
	},
	{
		pattern:  "^where are we\\?$",
		response: "Last time I looked, we were in $chan.",
	},
	{
		pattern:  "^where do you want to go today\\?$",
		response: "anywhere but redmond :(.",
	},
	{
		pattern:  "^fish go moo$",
		response: "\x01ACTION notes that $nick is truly enlightened.\x01",
	},
	{
		pattern:  "^cows go moo$",
		response: "$nick: only when they are impersonating fish.",
	},
	{
		pattern:  "^fish go blubb$",
		response: "$nick LIES! Fish don't go blubb! fish go m00!",
	},
	{
		pattern:  "^you know who else moos$",
		response: "$nick: YA MUM!",
	},
	{
		pattern:  "^If there's one thing I know for sure, it\\'s that fish don\\'t m00\\.$",
		response: "$nick: HERETIC! UNBELIEVER!",
	},
	{
		pattern:  "^fishbot: Muahahaha\\. Ph33r the dark side\\. :\\)$",
		response: "$nick: You smell :P.",
	},
	{
		pattern:  "^ammuu\\?$",
		response: "$nick: fish go m00 oh yes they do!",
	},
	{
		pattern:  "^fish$",
		response: "$nick: fish go m00!",
	},
	// Actions
	{
		pattern:  "^\x01ACTION feeds fishbot hundreds and thousands\x01$",
		response: "MEDI.. er.. FISHBOT",
	},
	{
		pattern:  "^\x01ACTION strokes fishbot\x01$",
		response: "\x01ACTION m00s loudly at $nick.\x01",
	},
	{
		/*
			Unable to ignore own nickname in regex pattern, and I don't want to add
			more logic in the function, so fishbot will reply when is slapped itself.
			@See: https://stackoverflow.com/a/26792316
		*/
		pattern:  "^\x01ACTION slaps (\\w*?) around a bit with a large trout\x01$",
		response: "trouted!",
	},
	{
		pattern:  "^\x01ACTION has returned from playing counterstrike\x01$",
		response: "like we care fs :(",
	},
	{
		pattern:  "^\x01ACTION thinks happy thoughts about pretty ladies\\.\x01$",
		response: "\x01ACTION has plenty of pretty ladies. Would you like one, $nick?\x01",
	},
	{
		pattern:  "^\x01ACTION snaffles a cookie off fishbot\\.\x01$",
		response: ":(",
	},
}
