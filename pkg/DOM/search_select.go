package DOM

type CompleteFlightFacilities struct {
	DirectFlight     FlightFacilities
	ReturnFlight     FlightFacilities
	NumberOfAdults   int
	NumberOfChildren int
	CabinClass       string
}

type FlightFacilities struct {
	Cancellation Cancellation
	Baggage      Baggage
	FlightPath   Path
}

type Cancellation struct {
	CancellationDeadlineBefore int
	CancellationPercentage     int
	Refundable                 bool
}

type Baggage struct {
	CabinAllowedWeight  int
	CabinAllowedLength  int
	CabinAllowedBreadth int
	CabinAllowedHeight  int
	HandAllowedWeight   int
	HandAllowedLength   int
	HandAllowedBreadth  int
	HandAllowedHeight   int
	FeeExtraPerKGCabin  int
	FeeExtraPerKGHand   int
	Restrictions        string
}
