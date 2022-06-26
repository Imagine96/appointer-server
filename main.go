package main

import (
	"log"
)

/*
	TODO

	01 Initialize server
	02 Initialize connection with db and close it
	03 Call correct procedure per req
*/

func main() {

	/* newDay := Day{
		Date:               "29-06-2022",
		confirmedSchedules: map[string]Schedule{},
		schedulesRequests:  map[string]Schedule{},
		IsFreeDay:          false,
		MaxSchedules:       6,
	} */

	client, ctx, cancel, err := connectToDB()
	if err != nil {
		log.Panic(err)
	}

	defer closeDBClient(client, ctx, cancel)

	/* if draftmanList, err := getDraftmanList(client, ctx); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(draftmanList)
	} */

	/* if d, err := getOneDraftmanById(client, ctx, "62b7b661eda2d6d42eb8052f"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*d)
	} */

	/* newDraftman := Draftman{
		/ Id:          primitive.NewObjectIDFromTimestamp(time.Now()).Hex(), /
		Id:          "62b7b661eda2d6d42eb8052f",
		ContactInfo: ContactInfo{},
		Agenda:      Agenda{},
		History:     map[string]Day{},
	} */

	/* existentDraftman := Draftman{
		Id:          "62b7b661eda2d6d42eb8052f",
		ContactInfo: ContactInfo{},
		Agenda:      Agenda{},
		History:     map[string]Day{},
	} */

	/* if err := updateDraftmanById(client, ctx, &newDraftman); err != nil {
		fmt.Println("here \n ", err)
	} */

	/* if createdId, err := insertNewDraftMan(client, ctx, newDraftman); err != nil {
		fmt.Println(fmt.Errorf("could not create draftman \n %v", err))
	} else {
		fmt.Print(createdId)
	} */

	/* if err := deleteDraftmanById(client, ctx, existentDraftman.Id); err != nil {
		fmt.Println(err)
	}

	fmt.Println("ou yeah") */

}
