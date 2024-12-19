package main
/*
X,Y,Unique Squirrel ID,Hectare,Shift,Date,Hectare Squirrel Number,Age,Primary Fur Color,
Highlight Fur Color,Combination of Primary and Highlight Color,Color notes,Location,
Above Ground Sighter Measurement,Specific Location,Running,Chasing,Climbing,Eating,
Foraging,Other Activities,Kuks,Quaas,Moans,Tail flags,Tail twitches,Approaches,Indifferent,
Runs from,Other Interactions,Lat/Long
*/

type SquirrelData struct {
	// Longitude coordinate for squirrel sighting point
	X                                     float32 `csv:"X"`
	// Latitude coordinate for squirrel sighting point
	Y                                     float32 `csv:"Y"`
	// Identification tag for each squirrel sightings. The tag is comprised of "Hectare ID" + "Shift" + "Date" + "Hectare Squirrel Number."
	UniqueSquirrelID                      string  `csv:"Unique Squirrel ID"`
	// ID tag, which is derived from the hectare grid used to divide and count the park area. One axis that runs predominantly north-to-south is numerical (1-42), and the axis that runs predominantly east-to-west is roman characters (A-I).
	Hectare                               string  `csv:"Hectare"`
	// Value is either "AM" or "PM," to communicate whether or not the sighting session occurred in the morning or late afternoon.
	Shift                                 string  `csv:"Shift"`
	// Concatenation of the sighting session day and month.
	Date                                  string  `csv:"Date"`
	// Number within the chronological sequence of squirrel sightings for a discrete sighting session.
	HectareSquirrelNumber                 string  `csv:"Hectare Squirrel Number"`
	// Value is either "Adult" or "Juvenile."
	Age                                   string  `csv:"Age"`
	// Value is either "Gray," "Cinnamon" or "Black."
	PrimaryFurColor                       string  `csv:"Primary Fur Color"`
	// Discrete value or string values comprised of "Gray," "Cinnamon" or "Black."
	HighlightFurColor                     string  `csv:"Highlight Fur Color"`
	// A combination of the previous two columns; this column gives the total permutations of primary and highlight colors observed.
	CombinationOfPrimaryAndHighlightColor string  `csv:"Combination of Primary and Highlight Color"`
	// Sighters occasionally added commentary on the squirrel fur conditions. These notes are provided here.
	ColorNotes                            string  `csv:"Color notes"`
	// Value is either "Ground Plane" or "Above Ground." Sighters were instructed to indicate the location of where the squirrel was when first sighted.
	Location                              string  `csv:"Location"`
	// For squirrel sightings on the ground plane, fields were populated with a value of “FALSE.”
	AboveGroundSighterMeasurement         string  `csv:"Above Ground Sighter Measurement"`
	// Sighters occasionally added commentary on the squirrel location. These notes are provided here.
	SpecificLocation                      string  `csv:"Specific Location"`
	// Squirrel was seen running.
	Running                               bool    `csv:"Running"`
	// Squirrel was seen chasing another squirrel.
	Chasing                               bool    `csv:"Chasing"`
	// Squirrel was seen climbing a tree or other environmental landmark.
	Climbing                              bool    `csv:"Climbing"`
	// Squirrel was seen eating.
	Eating                                bool    `csv:"Eating"`
	// Squirrel was seen foraging for food.
	Foraging                              bool    `csv:"Foraging"`
	// Other activities that sqirrel do
	OtherActivities                       string  `csv:"Other Activities"`
	// Squirrel was heard kukking, a chirpy vocal communication used for a variety of reasons.
	Kuks                                  bool    `csv:"Kuks"`
	// Squirrel was heard quaaing, an elongated vocal communication which can indicate the presence of a ground predator such as a dog.
	Quaas                                 bool    `csv:"Quaas"`
	// Squirrel was heard moaning, a high-pitched vocal communication which can indicate the presence of an air predator such as a hawk.
	Moans                                 bool    `csv:"Moans"`
	// Squirrel was seen flagging its tail. Flagging is a whipping motion used to exaggerate squirrel's size and confuse rivals or predators. Looks as if the squirrel is scribbling with tail into the air.
	TailFlags                             bool    `csv:"Tail flags"`
	// Squirrel was seen twitching its tail. Looks like a wave running through the tail, like a breakdancer doing the arm wave. Often used to communicate interest, curiosity.
	TailTwitches                          bool    `csv:"Tail twitches"`
	// Squirrel was seen approaching human, seeking food.
	Approaches                            bool    `csv:"Approaches"`
	// Squirrel was indifferent to human presence.
	Indifferent                           bool    `csv:"Indifferent"`
	// Squirrel was seen running from humans, seeing them as a threat.
	RunsFrom                              bool    `csv:"Runs from"`
	// Sighter notes on other types of interactions between squirrels and humans.
	OtherInteractions                     string  `csv:"Other Interactions"`
	// Latitude and longitude
	LatLong                               string  `csv:"Lat/Long"`
}