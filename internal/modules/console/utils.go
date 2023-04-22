package console

import (
	"calend/internal/models/access"
	"calend/internal/modules/domain/event/dto"
	dto2 "calend/internal/modules/domain/tag/dto"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"strings"
	"time"
)

func inputUuid() string {
	var uuid string
	fmt.Println("Input uuid:")
	if _, err := fmt.Scan(&uuid); err != nil {
		printError(err)
	}

	return uuid
}

/*
Пример ввода:
2006-01-02T15:04:05Z
Событие
Описание
Тип
1
*/
func inputCreateEvent() *dto.CreateEvent {
	var err error
	event := &dto.CreateEvent{}

	var timestamp string
	fmt.Println("Input event timestamp:")
	if _, err = fmt.Scan(&timestamp); err != nil {
		printError(err)
	}
	event.Timestamp, err = parseTimestamp(timestamp)
	if err != nil {
		printError(err)
	}

	var name string
	fmt.Println("Input event name:")
	if _, err = fmt.Scan(&name); err != nil {
		printError(err)
	}
	event.Name = name

	var description string
	fmt.Println("Input event description:")
	if _, err = fmt.Scan(&description); err != nil {
		printError(err)
	}
	event.Description = &description

	var typ string
	fmt.Println("Input event type:")
	if _, err = fmt.Scan(&typ); err != nil {
		printError(err)
	}
	event.Type = typ

	var isWholeDay bool
	fmt.Println("Input event is_whole_day:")
	if _, err = fmt.Scan(&isWholeDay); err != nil {
		printError(err)
	}
	event.IsWholeDay = isWholeDay

	var tagUuids string
	fmt.Println("Input event tag_uuids:")
	if _, err = fmt.Scan(&tagUuids); err != nil {
		printError(err)
	}
	event.TagUuids = strings.Fields(tagUuids)

	return event
}

func inputUpdateEvent() *dto.UpdateEvent {
	var err error
	event := &dto.UpdateEvent{}

	var timestamp string
	fmt.Println("Input event timestamp:")
	if _, err = fmt.Scan(&timestamp); err != nil {
		printError(err)
	}
	event.Timestamp, err = parseTimestamp(timestamp)
	if err != nil {
		printError(err)
	}

	var name string
	fmt.Println("Input event name:")
	if _, err = fmt.Scan(&name); err != nil {
		printError(err)
	}
	event.Name = name

	var description string
	fmt.Println("Input event description:")
	if _, err = fmt.Scan(&description); err != nil {
		printError(err)
	}
	event.Description = &description

	var typ string
	fmt.Println("Input event type:")
	if _, err = fmt.Scan(&typ); err != nil {
		printError(err)
	}
	event.Type = typ

	var isWholeDay bool
	fmt.Println("Input event is_whole_day:")
	if _, err = fmt.Scan(&isWholeDay); err != nil {
		printError(err)
	}
	event.IsWholeDay = isWholeDay

	var tagUuids string
	fmt.Println("Input event tag_uuids:")
	if _, err = fmt.Scan(&tagUuids); err != nil {
		printError(err)
	}
	event.TagUuids = strings.Fields(tagUuids)

	return event
}

func inputInvitations() dto.CreateInvitations {
	var err error

	var cnt int
	fmt.Println("Input count of invitations:")
	if _, err = fmt.Scan(&cnt); err != nil {
		printError(err)
	}
	invitations := make(dto.CreateInvitations, cnt)

	for i := 0; i < cnt; i++ {
		var userUuid string
		fmt.Println("Input user uuid:")
		if _, err = fmt.Scan(&userUuid); err != nil {
			printError(err)
		}
		invitations[i].UserUuid = userUuid

		var arCode string
		fmt.Println("Input access right code:")
		if _, err = fmt.Scan(&arCode); err != nil {
			printError(err)
		}
		invitations[i].AccessRightCode = access.Type(arCode)
	}

	return invitations
}

func inputCreateTag() *dto2.CreateTag {
	var err error
	tag := &dto2.CreateTag{}

	var name string
	fmt.Println("Input tag name:")
	if _, err = fmt.Scan(&name); err != nil {
		printError(err)
	}
	tag.Name = name

	var description string
	fmt.Println("Input tag description:")
	if _, err = fmt.Scan(&description); err != nil {
		printError(err)
	}
	tag.Description = description

	return tag
}

func inputUpdateTag() *dto2.UpdateTag {
	var err error
	tag := &dto2.UpdateTag{}

	var name string
	fmt.Println("Input tag name:")
	if _, err = fmt.Scan(&name); err != nil {
		printError(err)
	}
	tag.Name = name

	var description string
	fmt.Println("Input tag description:")
	if _, err = fmt.Scan(&description); err != nil {
		printError(err)
	}
	tag.Description = description

	return tag
}

func printError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func printEvents(events ...*dto.Event) {
	for _, event := range events {
		event, err := marshalJSON(event)
		if err != nil {
			printError(err)
		}

		fmt.Println(event)
	}
}

func printTags(tags ...*dto2.Tag) {
	for _, tag := range tags {
		tag, err := marshalJSON(tag)
		if err != nil {
			printError(err)
		}

		fmt.Println(tag)
	}
}

func marshalJSON(v any) (string, error) {
	jsonRes, err := prettyjson.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(jsonRes), nil
}

func parseTimestamp(timestamp string) (time.Time, error) {
	return time.Parse(time.RFC3339, timestamp)
}
