# faker
A package that generates fake data for GO

```

   _|_|_|            _|_|_|_|          _|                            
 _|          _|_|    _|        _|_|_|  _|  _|      _|_|    _|  _|_|  
 _|  _|_|  _|    _|  _|_|_|  _|    _|  _|_|      _|_|_|_|  _|_|      
 _|    _|  _|    _|  _|      _|    _|  _|  _|    _|        _|        
   _|_|_|    _|_|    _|        _|_|_|  _|    _|    _|_|_|  _|        


```

## Error handling

In this project I preferred panic instead of explicit error returns to keep the API clean and uniform. I think this project must be
used by humans beings not by computers directly.

It is the same reason [regexp.MustCompile](https://golang.org/pkg/regexp/#MustCompile) panic instead of return an error.
