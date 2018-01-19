### Engine

Fetches online content (RSS, FB, Twitter, websites, ...), filters the content
based on pre-defined rules and stores the results. 

Each `engine` processes one configuration file at a time with. Each
configuration defines the 1) sources 2) filters 3) ouputs.


## Implementation details

### Source

```
type Source interface
	[]DataUnit	Data
	[]errors 		Error

	func				Fetch()
```

Each source defined in the configuration file initializes a struct of type 
`Source` which implements the `Fetch()` method. This method will fetch all
content from the source and return an array of `DataUnit`s.

Sources are embarassingly parallel with regard to other sources. Thus, each
source will run `Fetch()` in a dedicated go routine.

### DataUnit

```
type DataUnit struct
	string 				Url
	time.Time			Timestamp
	[]byte				Payload
	[]Processor 	Pipeline	
	[]Output			Outputs
	[]byte				Results
	[]Filtered		FiltersResults

	func 					Process()
	func 					AddError(error)
	func					AddResult([]byte)
```

Data units are self contained structures with all information and data to
process, filter and output the content of a `DataUnit`.

Each `DataUnit` should be embarassibly parallel from each other. Thus, one go
routine per DataUnit will be running the logic which filters, processes and
outputs the results to the correct outputs.

The `Process()` method loops through the defined processors in the `Pipeline`
and calls `Handle()` in each of the processors. The result or error of the
processing are added to the structure. Once all the processors have ran, a
similar for loop runs for the results to be handled by the defined outputs.

### Processor and Output
```
type Processor interface
	string 	Name
	
	func Handle(DataUnit)	([]byte result, error)
```

```
type Output interface
	string 	Name
	string 	Url	

	func 		Handle(DataUnit) error
```

Processors and outputs expose the `Handle()` method for the `DataUnit` to call.
In the `Processor` case, the returning value is either the result of the
procesing or an error. The `Output` returns an error or nothing.

A processor may add or permute results or inspect results for filtering. 

### Filtered
```
type Filtered struct
	string 	Filter
	[]byte	Result
```

If a `DataUnit` is a hit for a given filter rule, the filter is added to the
`DataUnit`. The `Result` keeps track of the hit score for the filter 


