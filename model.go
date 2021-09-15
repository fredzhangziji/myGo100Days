package main

type RankType int

const (
	RankOne   RankType = 1
	RankTwo   RankType = 2
	RankThree RankType = 3
	RankFour  RankType = 4
	RankFive  RankType = 5
)

// Type definition for getting all reservations
type Reservations []ReservationPayload

// Payload for ROUTE 2: POST route to make reservantion
// Payload for ROUTE 3: PUT route to update reservation
type ReservationPayload struct {
	ID            int // this is randomly generated
	Paired        bool
	PlayerOneName string
	PlayerTwoName string
	TimeSlot      string //"1530:1600"
	RankPlayerOne RankType
	RankPlayerTwo RankType
}
