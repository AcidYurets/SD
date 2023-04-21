package console

import (
	"calend/internal/modules/db/ent"
	ar_serv "calend/internal/modules/domain/access_right/service"
	auth_serv "calend/internal/modules/domain/auth/service"
	event_serv "calend/internal/modules/domain/event/service"
	search_serv "calend/internal/modules/domain/search/service"
	tag_serv "calend/internal/modules/domain/tag/service"
	user_serv "calend/internal/modules/domain/user/service"
	"fmt"
)

const menu = `
===== Многопользовательский календарь =====
0. Выйти
1. Вывести список всех блюд
2. Вывести список всех блюд из категории "Супы"
3. Вывести список блюд, название которых содержит слово "суп"
4. Вывести средний вес блюд
5. Вывести количество блюд в каждой категории
Что будем делать: `

func MenuLoop(
	userService *user_serv.UserService,
	tagService *tag_serv.TagService,
	authService *auth_serv.AuthService,
	arService *ar_serv.AccessRightService,
	eventService *event_serv.EventService,
	searchService *search_serv.SearchService,

	client *ent.Client,
) error {
	dbEngine := &DB{gormDb}
	linqEngine := (&Linq{}).Init(gormDb)
	jsonEngine := (&Json{}).Init(gormDb)

	for {
		fmt.Print(menu)

		var c int
		_, err := fmt.Scan(&c)
		if err != nil {
			logError(err)
			continue
		}

		switch c {
		case 1:
			if err := linqEngine.Task1(); err != nil {
				logError(err)
				continue
			}
		case 2:
			if err := linqEngine.Task2(); err != nil {
				logError(err)
				continue
			}
		case 3:
			if err := linqEngine.Task3(); err != nil {
				logError(err)
				continue
			}
		case 4:
			if err := linqEngine.Task4(); err != nil {
				logError(err)
				continue
			}
		case 5:
			if err := linqEngine.Task5(); err != nil {
				logError(err)
				continue
			}
		case 6:
			if err := jsonEngine.Task6(); err != nil {
				logError(err)
				continue
			}
		case 7:
			if err := jsonEngine.Task7(); err != nil {
				logError(err)
				continue
			}
		case 8:
			if err := jsonEngine.Task8(); err != nil {
				logError(err)
				continue
			}
		case 9:
			if err := dbEngine.Task9(); err != nil {
				logError(err)
				continue
			}
		case 10:
			if err := dbEngine.Task10(); err != nil {
				logError(err)
				continue
			}
		case 11:
			if err := dbEngine.Task11(); err != nil {
				logError(err)
				continue
			}
		case 12:
			if err := dbEngine.Task12(); err != nil {
				logError(err)
				continue
			}
		case 13:
			if err := dbEngine.Task13(); err != nil {
				logError(err)
				continue
			}
		case 14:
			if err := dbEngine.Task14(); err != nil {
				logError(err)
				continue
			}

		case 0:
			return nil

		}
	}
}
