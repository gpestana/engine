## Engine

Fetches online content (RSS, FB, Twitter, websites, ...), filters the content
based on pre-defined rules and stores the results. 

Each `engine` processes one configuration file at a time with. Each
configuration defines the 1) sources 2) processors (which includes filters) 3) 
ouputs.

### How to start it

Create configuration file:
```
{
	"sources": [{
		"type": "rss",
		"url": "https://news.ycombinator.com/rss"
	},
	{
		"type": "rss",
		"url": "http://rss.nytimes.com/services/xml/rss/nyt/World.xml"
	}],
	"filters": [{
		"type": "news",
		"filter": "a filter defined by user"
	}],
	"outputs": []
}
```


