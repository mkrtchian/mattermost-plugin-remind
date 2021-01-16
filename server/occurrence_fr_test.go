package main

import (
	//"encoding/json"
	//"strings"
	"testing"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
func TestCreateOccurrencesFr(t *testing.T) {

	user := &model.User{
		Email:    "-@-.-",
		Nickname: "TestUser",
		Password: model.NewId(),
		Username: "testuser",
		Roles:    model.SYSTEM_USER_ROLE_ID,
		Locale:   "fr",
	}
	testTime := time.Now().UTC().Round(time.Second)

	occurrences := []Occurrence{
		{
			Id:         model.NewId(),
			ReminderId: model.NewId(),
			Occurrence: testTime,
		},
	}

	request := &ReminderRequest{
		TeamId:   model.NewId(),
		Username: user.Username,
		Payload:  "Salut dans une minute",
		Reminder: Reminder{
			Id:          model.NewId(),
			TeamId:      model.NewId(),
			Username:    user.Username,
			Occurrences: occurrences,
			Message:     "Salue",
			Target:      "moi",
			When:        "dans one minute",
		},
	}

	stringOccurrences, _ := json.Marshal(occurrences)
	setupAPI := func() *plugintest.API {
		api := &plugintest.API{}
		api.On("LogDebug", mock.Anything, mock.Anything, mock.Anything).Maybe()
		api.On("LogError", mock.Anything, mock.Anything, mock.Anything).Maybe()
		api.On("LogInfo", mock.Anything).Maybe()
		api.On("GetUserByUsername", mock.AnythingOfType("string")).Return(user, nil)
		return api
	}

	t.Run("if creates occurrences", func(t *testing.T) {

		api := setupAPI()
		defer api.AssertExpectations(t)

		api.On("KVGet", mock.Anything).Return(stringOccurrences, nil)
		api.On("KVSet", mock.Anything, mock.Anything).Return(nil)
		p := &Plugin{}
		p.API = api

		assert.Nil(t, p.CreateOccurrences(request))

	})

	t.Run("if fails to create occurrences", func(t *testing.T) {

		api := setupAPI()
		defer api.AssertExpectations(t)

		p := &Plugin{}
		p.API = api

		request.Reminder.When = "dans foobar"
		assert.NotNil(t, p.CreateOccurrences(request))

		request.Reminder.When = "dans foo secondes"
		assert.NotNil(t, p.CreateOccurrences(request))
	})
}*/

func TestInFR(t *testing.T) {

	setupAPI := func() *plugintest.API {
		api := &plugintest.API{}
		api.On("LogDebug", mock.Anything, mock.Anything, mock.Anything).Maybe()
		api.On("LogInfo", mock.Anything).Maybe()
		return api
	}

	t.Run("if inEN locale is used", func(t *testing.T) {

		api := setupAPI()
		defer api.AssertExpectations(t)

		p := &Plugin{}
		p.API = api

		user := &model.User{
			Email:    "-@-.-",
			Nickname: "TestUser",
			Password: model.NewId(),
			Username: "testuser",
			Roles:    model.SYSTEM_USER_ROLE_ID,
			Locale:   "fr",
		}
		p.locales = make(map[string]string)
		p.locales["fr"] = "assets/i18n/fr.json"

		times, err := p.inEN("dans 1 seconde", user)
		testTimes := []time.Time{
			time.Now().Round(time.Second).Add(time.Second * time.Duration(1)).UTC(),
		}
		assert.Nil(t, err)
		assert.Equal(t, times, testTimes)
/*
		times, err = p.inEN("dans 712 minutes", user)
		testTimes = []time.Time{
			time.Now().Round(time.Second).Add(time.Minute * time.Duration(712)).UTC(),
		}
		assert.Nil(t, err)
		assert.Equal(t, times, testTimes)

		times, err = p.inEN("dans 1h", user)
		testTimes = []time.Time{
			time.Now().Round(time.Second).Add(time.Hour * time.Duration(1)).UTC(),
		}
		assert.Nil(t, err)
		assert.Equal(t, times, testTimes)

		times, err = p.inEN("dans 3 heures", user)
		testTimes = []time.Time{
			time.Now().Round(time.Second).Add(time.Hour * time.Duration(3)).UTC(),
		}
		assert.Nil(t, err)
		assert.Equal(t, times, testTimes)

		times, err = p.inEN("dans 24 heures", user)
		testTimes = []time.Time{
			time.Now().Round(time.Second).Add(time.Hour * time.Duration(24)).UTC(),
		}
		assert.Nil(t, err)
		assert.Equal(t, times, testTimes)

		times, err = p.inEN("dans 2 jours", user)
		testTimes = []time.Time{
			time.Now().Round(time.Second).Add(time.Hour * time.Duration(24) * time.Duration(2)).UTC(),
		}
		assert.Nil(t, err)
		assert.Equal(t, times, testTimes)

		times, err = p.inEN("dans 90 semaines", user)
		testTimes = []time.Time{
			time.Now().Round(time.Second).Add(time.Hour * time.Duration(24) * time.Duration(7) * time.Duration(90)).UTC(),
		}
		assert.Nil(t, err)
		assert.Equal(t, times, testTimes)

		times, err = p.inEN("dans 4 mois", user)
		testTimes = []time.Time{
			time.Now().Round(time.Second).Add(time.Hour * time.Duration(24) * time.Duration(30) * time.Duration(4)).UTC(),
		}
		assert.Nil(t, err)
		assert.Equal(t, times, testTimes)

		times, err = p.inEN("dans 1 an", user)
		testTimes = []time.Time{
			time.Now().Round(time.Second).Add(time.Hour * time.Duration(24) * time.Duration(365)).UTC(),
		}
		assert.Nil(t, err)
		assert.Equal(t, times, testTimes)*/

	})

}
/*
func TestAt(t *testing.T) {
	setupAPI := func() *plugintest.API {
		api := &plugintest.API{}
		api.On("LogDebug", mock.Anything, mock.Anything, mock.Anything).Maybe()
		api.On("LogInfo", mock.Anything).Maybe()
		api.On("LogError", mock.Anything).Maybe()
		return api
	}

	t.Run("if atEN locale is used", func(t *testing.T) {

		api := setupAPI()
		defer api.AssertExpectations(t)

		p := &Plugin{}
		p.API = api

		user := &model.User{
			Email:    "-@-.-",
			Nickname: "TestUser",
			Password: model.NewId(),
			Username: "testuser",
			Roles:    model.SYSTEM_USER_ROLE_ID,
			Locale:   "fr",
		}
		location := p.location(user)

		times, err := p.atEN("à midi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Hour(), 12)

		times, err = p.atEN("à minuit", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Hour(), 0)

		times, err = p.atEN("à 2h", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 2 || times[0].In(location).Hour() == 14)

		times, err = p.atEN("à 7h", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 7 || times[0].In(location).Hour() == 19)

		times, err = p.atEN("à 12h30", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 12 && times[0].In(location).Minute() == 30)

		times, err = p.atEN("à 19h12", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 19 && times[0].In(location).Minute() == 12)

		times, err = p.atEN("à 20h05", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 20 && times[0].In(location).Minute() == 5)

		times, err = p.atEN("à 9h52", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 9 && times[0].In(location).Minute() == 52)

		times, err = p.atEN("à 9h12", user)
		assert.Nil(t, err)
		assert.True(t, (times[0].In(location).Hour() == 9 || times[0].In(location).Hour() == 21) && times[0].In(location).Minute() == 12)

		times, err = p.atEN("à 17h15", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 17 && times[0].In(location).Minute() == 15)

		times, err = p.atEN("à 9:30", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 9 && times[0].In(location).Minute() == 30)

		times, err = p.atEN("à 12:30", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 0 && times[0].In(location).Minute() == 30)

		times, err = p.atEN("à 17h", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 17 && times[0].In(location).Minute() == 0)

		times, err = p.atEN("à 22h", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 22 && times[0].In(location).Minute() == 0)

		times, err = p.atEN("à h4", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 4 && times[0].In(location).Minute() == 0)

		times, err = p.atEN("à 14h00", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 14 && times[0].In(location).Minute() == 0)

		times, err = p.atEN("à 11h00 chaque jeudi", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, (times[0].In(location).Hour() == 11 || times[0].In(location).Hour() == 23) &&
				times[0].In(location).Weekday().String() == "jeudi")
		}

		times, err = p.atEN("à 15h chaque jour", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Hour() == 15)
		}
	})
}

func TestOn(t *testing.T) {
	setupAPI := func() *plugintest.API {
		api := &plugintest.API{}
		api.On("LogDebug", mock.Anything, mock.Anything, mock.Anything).Maybe()
		api.On("LogInfo", mock.Anything).Maybe()
		return api
	}

	t.Run("if onEN locale is used", func(t *testing.T) {

		api := setupAPI()
		defer api.AssertExpectations(t)

		p := &Plugin{}
		p.API = api

		user := &model.User{
			Email:    "-@-.-",
			Nickname: "TestUser",
			Password: model.NewId(),
			Username: "testuser",
			Roles:    model.SYSTEM_USER_ROLE_ID,
			Locale:   "fr",
		}
		location := p.location(user)

		times, err := p.onEN("le Lundi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Lundi")

		times, err = p.onEN("le mardi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Mardi")

		times, err = p.onEN("le MerCredi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Mercredi")

		times, err = p.onEN("le jeuDI", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Jeudi")

		times, err = p.onEN("le venDredi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Vendredi")

		times, err = p.onEN("le Samedi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Samedi")

		times, err = p.onEN("le dimanche", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Dimanche")

		times, err = p.onEN("les Lundis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Lundi")
		}

		times, err = p.onEN("Les Mardis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Mardi")
		}

		times, err = p.onEN("les mercredis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Mercredi")
		}

		times, err = p.onEN("les Jeudis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Jeudi")
		}

		times, err = p.onEN("les Vendredis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Vendredi")
		}

		times, err = p.onEN("les Samedis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Samedi")
		}

		times, err = p.onEN("les Dimanches", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Dimanche")
		}

		times, err = p.onEN("le lun", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Lundi")

		times, err = p.onEN("le Mer", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Mercredi")

		times, err = p.onEN("le mardi à midi", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Weekday().String() == "Mardi" && times[0].In(location).Hour() == 12)

		times, err = p.onEN("le dmanche à 3h42", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Weekday().String() == "Dimanche" && times[0].In(location).Hour() == 3 &&
			times[0].In(location).Minute() == 42)

		times, err = p.onEN("le 15 Décembre", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month().String() == "Décembre" && times[0].In(location).Day() == 15)

		times, err = p.onEN("le 12 jan", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month().String() == "Janvier" && times[0].In(location).Day() == 12)

		times, err = p.onEN("le 12 juillet", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month().String() == "Juillet" && times[0].In(location).Day() == 12)

		times, err = p.onEN("le 22 mArs", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month().String() == "Mars" && times[0].In(location).Day() == 22)

		times, err = p.onEN("le 17 mars à 17h41", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Month().String() == "Mars" && times[0].In(location).Day() == 17 &&
				times[0].In(location).Hour() == 17 && times[0].In(location).Minute() == 41)
		}

		times, err = p.onEN("le 7 septembre 2020", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month().String() == "Septembre" && times[0].In(location).Day() == 7)

		times, err = p.onEN("le 17 avril 2020", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month().String() == "Avril" && times[0].In(location).Day() == 17)

		times, err = p.onEN("le 9 avril 2020 à 11h", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Month().String() == "Avril" && times[0].In(location).Day() == 9 &&
				times[0].In(location).Hour() == 11)
		}

		times, err = p.onEN("le 10 aOût 2019", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month().String() == "Août" && times[0].In(location).Day() == 10)

		times, err = p.onEN("le 7", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Day(), 7)

		times, err = p.onEN("le 7", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Day(), 7)

		times, err = p.onEN("le sept", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Day(), 7)

		times, err = p.onEN("le 17/1/20", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Year() == 2020 && times[0].In(location).Month() == 1 &&
			times[0].In(location).Day() == 17)

		times, err = p.onEN("on 17/12/2020", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Year() == 2020 && times[0].In(location).Month() == 12 &&
			times[0].In(location).Day() == 17)

		times, err = p.onEN("le 17.1.20", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Year() == 2020 && times[0].In(location).Month() == 1 &&
			times[0].In(location).Day() == 17)

		times, err = p.onEN("le 17.12.20", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Year() == 2020 && times[0].In(location).Month() == 12 &&
			times[0].In(location).Day() == 17)

		times, err = p.onEN("le 1/12", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month() == 12 && times[0].In(location).Day() == 1)

		times, err = p.onEN("le 17-5-20", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month() == 5 && times[0].In(location).Day() == 17)

		times, err = p.onEN("le 5-12-2020", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month() == 12 && times[0].In(location).Day() == 5)

		times, err = p.onEN("le 12-12", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month() == 12 && times[0].In(location).Day() == 12)

		times, err = p.onEN("le 1-1 à minuit", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month() == 1 && times[0].In(location).Day() == 1 &&
			times[0].In(location).Hour() == 0)

		times, err = p.onEN("le 10-08-2020 13:55", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Year() == 2020 &&
			times[0].In(location).Month() == 8 && times[0].In(location).Day() == 10 &&
			times[0].In(location).Hour() == 13 && times[0].In(location).Minute() == 55)

	})
}

func TestEvery(t *testing.T) {

	setupAPI := func() *plugintest.API {
		api := &plugintest.API{}
		api.On("LogDebug", mock.Anything, mock.Anything, mock.Anything).Maybe()
		api.On("LogInfo", mock.Anything).Maybe()
		return api
	}

	t.Run("if everyEN locale is used", func(t *testing.T) {
		api := setupAPI()
		defer api.AssertExpectations(t)

		p := &Plugin{}
		p.API = api

		user := &model.User{
			Email:    "-@-.-",
			Nickname: "TestUser",
			Password: model.NewId(),
			Username: "testuser",
			Roles:    model.SYSTEM_USER_ROLE_ID,
			Locale:   "fr",
		}
		location := p.location(user)

		times, err := p.everyEN("chaque Jeudi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Jeudi")

		times, err = p.everyEN("chaque jour", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), time.Now().In(location).AddDate(0, 0, 1).Weekday().String())

		times, err = p.everyEN("chaque 18/12/2022", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month() == 12 && times[0].In(location).Year() == 2022)

		times, err = p.everyEN("chaque 25 Janvier", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month() == 1 && times[0].In(location).Day() == 25)

		times, err = p.everyEN("chaque autre mercredi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Mercredi")

		times, err = p.everyEN("chaque jour à 11h32", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Hour() == 11 && times[0].In(location).Minute() == 32)

		times, err = p.everyEN("chaque 5/5 à 7h", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month() == 5 && times[0].In(location).Day() == 5 &&
			(times[0].In(location).Hour() == 7))

		times, err = p.everyEN("chaque 20/7 à 11:00", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Month() == 7 && times[0].In(location).Day() == 20 &&
			(times[0].In(location).Hour() == 11))

		times, err = p.everyEN("chaque Lundi à 7h32", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Weekday().String() == "Lundi" && (times[0].In(location).Hour() == 7))

		times, err = p.everyEN("chaque lundi et mercredi", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Weekday().String() == "Lundi" && times[1].In(location).Weekday().String() == "Mercredi")
		}

		times, err = p.everyEN("chaque mercredi, jeudi", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Weekday().String() == "Mercredi" && times[1].In(location).Weekday().String() == "Jeudi")
		}

		times, err = p.everyEN("chaque autre vendredi et samedi", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Weekday().String() == "Vendredi" && times[1].In(location).Weekday().String() == "Samedi")
		}

		times, err = p.everyEN("chaque lundi et mercredi à 1h39", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Weekday().String() == "Lundi" &&
				times[1].In(location).Weekday().String() == "Mercredi" && times[0].In(location).Hour() == 1 && times[0].Minute() == 39)
		}

		times, err = p.everyEN("chaque lundi, mardi et dimanche à 11h00", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Weekday().String() == "Lundi" &&
				times[1].In(location).Weekday().String() == "Mardi" && times[2].In(location).Weekday().String() == "Dimanche" &&
				times[0].In(location).Hour() == 11)
		}

		times, err = p.everyEN("chaque lundi, mardi à 14h", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Weekday().String() == "Lundi" &&
				times[1].In(location).Weekday().String() == "Mardi" && times[0].In(location).Hour() == 14)
		}

		times, err = p.everyEN("chaque 30/1 et 30/9 à midi", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Month() == 1 && times[0].In(location).Day() == 30 &&
				times[1].In(location).Month() == 9 && times[1].In(location).Day() == 30 && times[0].In(location).Hour() == 12)
		}
	})
}

func TestFreeForm(t *testing.T) {

	setupAPI := func() *plugintest.API {
		api := &plugintest.API{}
		api.On("LogDebug", mock.Anything, mock.Anything, mock.Anything).Maybe()
		api.On("LogInfo", mock.Anything).Maybe()
		return api
	}

	t.Run("if freeformEN locale is used", func(t *testing.T) {
		api := setupAPI()
		defer api.AssertExpectations(t)

		p := &Plugin{}
		p.API = api

		user := &model.User{
			Email:    "-@-.-",
			Nickname: "TestUser",
			Password: model.NewId(),
			Username: "testuser",
			Roles:    model.SYSTEM_USER_ROLE_ID,
			Locale:   "fr",
		}
		location := p.location(user)

		times, err := p.freeFormEN("lundi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Lundi")

		times, err = p.freeFormEN("mardi à 21h34", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Weekday().String() == "Mardi" && times[0].In(location).Hour() == 21 &&
			times[0].In(location).Minute() == 34)

		times, err = p.freeFormEN("mercredi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Mercredi")

		times, err = p.freeFormEN("jeudi à midi", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Weekday().String() == "Jeudi" && times[0].In(location).Hour() == 12)

		times, err = p.freeFormEN("vendredi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Vendredi")

		times, err = p.freeFormEN("samedi", user)
		assert.Nil(t, err)
		assert.Equal(t, times[0].In(location).Weekday().String(), "Samedi")

		times, err = p.freeFormEN("dimanche à 16h20", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Weekday().String() == "Dimanche" && times[0].In(location).Hour() == 16 &&
			times[0].In(location).Minute() == 20)

		times, err = p.freeFormEN("aujourd'hui à 15h", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Hour(), 15)
		}

		times, err = p.freeFormEN("demain", user)
		assert.Nil(t, err)
		assert.True(t, times[0].In(location).Weekday().String() == time.Now().In(location).AddDate(0, 0, 1).Weekday().String())

		times, err = p.freeFormEN("demain à 16h", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Weekday().String() == time.Now().In(location).AddDate(0, 0, 1).Weekday().String() &&
				times[0].In(location).Hour() == 16)
		}

		times, err = p.freeFormEN("demain à 20h00", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Weekday().String() == time.Now().In(location).AddDate(0, 0, 1).Weekday().String() &&
				(times[0].In(location).Hour() == 8 || times[0].In(location).Hour() == 20))
		}

		times, err = p.freeFormEN("chaque jour", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), time.Now().In(location).AddDate(0, 0, 1).Weekday().String())
		}

		times, err = p.freeFormEN("chaque jour à 3h23", user)
		assert.Nil(t, err)
		if err == nil {
			assert.True(t, times[0].In(location).Hour() == 3 && times[0].In(location).Minute() == 23)
		}

		times, err = p.freeFormEN("lundis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Lundi")
		}

		times, err = p.freeFormEN("mardis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Mardi")
		}

		times, err = p.freeFormEN("mercredis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Mercredi")
		}

		times, err = p.freeFormEN("Jeudis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Jeudi")
		}

		times, err = p.freeFormEN("Vendredis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Vendredi")
		}

		times, err = p.freeFormEN("Samedis", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Samedi")
		}

		times, err = p.freeFormEN("Dimanches", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Weekday().String(), "Dimanche")
		}

		times, err = p.freeFormEN("17h", user)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, times[0].In(location).Hour(), 17)
		}

	})
}

func TestFormatWhen(t *testing.T) {

	user := &model.User{
		Email:    "-@-.-",
		Nickname: "TestUser",
		Password: model.NewId(),
		Username: "testuser",
		Roles:    model.SYSTEM_USER_ROLE_ID,
		Locale:   "fr",
	}

	setupAPI := func() *plugintest.API {
		api := &plugintest.API{}
		api.On("LogDebug", mock.Anything, mock.Anything, mock.Anything).Maybe()
		api.On("LogError", mock.Anything, mock.Anything, mock.Anything).Maybe()
		api.On("LogInfo", mock.Anything).Maybe()
		api.On("GetUserByUsername", mock.AnythingOfType("string")).Return(user, nil)
		return api
	}

	t.Run("if creates occurrences", func(t *testing.T) {

		api := setupAPI()
		defer api.AssertExpectations(t)

		p := &Plugin{}
		p.API = api

		timeTest := time.Now().Round(time.Second).Format(time.RFC3339)

		when := p.formatWhen(user.Username, "dans 22 secondes", timeTest, false)
		assert.True(t, strings.HasPrefix(when, "dans 22 secondes à"))

		when = p.formatWhen(user.Username, "dans 22 secondes", timeTest, true)
		assert.True(t, strings.Contains(when, "aujourd'hui"))

		when = p.formatWhen(user.Username, "à 20h", timeTest, false)
		assert.True(t, strings.HasPrefix(when, "à"))

		when = p.formatWhen(user.Username, "à 20h", timeTest, true)
		assert.True(t, strings.Contains(when, "aujourd'hui"))

		when = p.formatWhen(user.Username, "le 12/8", timeTest, false)
		assert.True(t, strings.Contains(when, "aujourd'hui"))

		when = p.formatWhen(user.Username, "le 12/8", timeTest, true)
		assert.True(t, strings.Contains(when, "aujourd'hui"))

		when = p.formatWhen(user.Username, "chaque mardi", timeTest, false)
		assert.True(t, strings.Contains(when, "chaque"))

		when = p.formatWhen(user.Username, "chaque mardi", timeTest, true)
		assert.True(t, strings.Contains(when, "chaque"))

		when = p.formatWhen(user.Username, "demain", timeTest, false)
		assert.True(t, strings.Contains(when, "aujourd'hui"))

		when = p.formatWhen(user.Username, "demain", timeTest, true)
		assert.True(t, strings.Contains(when, "aujourd'hui"))

	})
}
*/
