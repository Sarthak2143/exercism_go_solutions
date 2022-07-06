//Package weather does blah blah.
package weather
//CurrentCondition variable blah blah.
var CurrentCondition string 
//CurrentLocation variable does blah blah.
var CurrentLocation string 

//Forecast function does blah blah.
func Forecast(city, condition string) string {	
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
