package translation_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopher-translator/handlers/history_handler"
	"gopher-translator/helpers"
	"regexp"
	"strings"
	"unicode"
)

type TranslateWordRequest struct {
	Word string `json:"english-word"`
}

type TranslateWordResponse struct {
	Word string `json:"gopher-word"`
}

type TranslateSentenceRequest struct {
	Sentence string `json:"english-sentence"`
}

type TranslateSentenceResponse struct {
	Sentence string `json:"gopher-sentence"`
}

func HandleTranslateWordRequest(ctx *gin.Context) {
	req := TranslateWordRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		helpers.GetLogger().WithFields(logrus.Fields{
			"error":    err.Error(),
			"req_type": "word",
		}).Error("Failed to bind /word request body")

		ctx.JSON(400, "Invalid request")
		return
	}

	if len(req.Word) == 0 {
		helpers.GetLogger().Error("/word request has incorrect json structure or empty value")

		ctx.JSON(400, "Invalid request")
		return
	}

	if strings.Contains(req.Word, "'") {
		helpers.GetLogger().WithFields(logrus.Fields{
			"word": req.Word,
		}).Warn("Word contains apostrophe, skipping")

		ctx.JSON(200, "Gophers don't understand shortened versions of words or apostrophes.")
		return
	}

	res := TranslateWordResponse{
		Word: translateWord(req.Word),
	}

	history_handler.UpdateHistory(req.Word, res.Word)

	helpers.GetLogger().WithFields(logrus.Fields{
		"req_type":    "word",
		"original":    req.Word,
		"translation": res.Word,
	}).Info("Success")

	ctx.JSON(200, res)

}

func HandleTranslateSentenceRequest(ctx *gin.Context) {

	req := TranslateSentenceRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		helpers.GetLogger().WithFields(logrus.Fields{
			"error":    err.Error(),
			"req_type": "sentence",
		}).Error("Failed to bind request body")

		ctx.JSON(400, "Invalid request")
		return
	}

	if !validateSentence(req.Sentence) {
		helpers.GetLogger().WithFields(logrus.Fields{
			"sentence": req.Sentence,
		}).Error("Failed to validate sentence")

		ctx.JSON(400, "Invalid request")
		return
	}

	res := TranslateSentenceResponse{}
	res.Sentence = translateSentence(req.Sentence)

	history_handler.UpdateHistory(req.Sentence, res.Sentence)

	helpers.GetLogger().WithFields(logrus.Fields{
		"req_type":    "sentence",
		"original":    req.Sentence,
		"translation": res.Sentence,
	}).Info("Success")

	ctx.JSON(200, res)

}

func translateWord(word string) string {

	translate := func(original string) string {

		original = strings.ToLower(original)

		startsWithXr := strings.HasPrefix(original, "xr")
		if startsWithXr {
			return fmt.Sprintf("ge%s", original)
		}

		startsWithVowel := strings.Contains("aeiou", string(rune(original[0])))
		if startsWithVowel {
			return fmt.Sprintf("g%s", original)
		}

		startsWithConsonantAndQU, _ := regexp.MatchString("^\\wqu\\w+$", original)
		if startsWithConsonantAndQU {
			return original[3:] + original[0:3] + "ogo"
		}

		startsWithConsonantSound, _ := regexp.MatchString("^(ph|ch|sh|th|zh|wh)\\w+$", original)
		if startsWithConsonantSound {
			return original[2:] + original[0:2] + "ogo"
		}

		return original[1:] + original[0:1] + "ogo"

	}

	translation := translate(word)

	startsWithUppercase := unicode.IsUpper(rune(word[0]))
	if startsWithUppercase {
		return string(unicode.ToUpper(rune(translation[0]))) + translation[1:]
	}
	return translation

}

func translateSentence(sentence string) string {

	var words []string

	for _, word := range strings.Fields(sentence[0 : len(sentence)-1]) {
		if strings.Contains(word, "'") {
			continue
		}

		words = append(words, translateWord(word))
	}

	return strings.Join(words, " ") + sentence[len(sentence)-1:]
}

func validateSentence(sentence string) bool {

	if len(sentence) == 0 {
		return false
	}

	if !strings.Contains("?.!", sentence[len(sentence)-1:]) {
		return false
	}

	hasMoreThanOneWord, _ := regexp.MatchString("\\s", sentence)
	if !hasMoreThanOneWord {
		return false
	}

	return true

}
