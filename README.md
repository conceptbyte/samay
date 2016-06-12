# Samay
Samay is a date/time manipulation library built in GoLang. It uses a fluent API and provides functions for

- Date time Formatting
- Rounding of date and time
- Addition of date time
- Subtraction of date time
- Difference between different date time
- Formatters for date time
- Comparison of date times

# Use
Create a new instance using one of the two methods

## Using a Time object
```
samay.Create(time Time)
```

## Using a string and format
```
samay.CreateFromFormat(time.FormatANSIC, '25/12/14')
```

# Fluent API
Samay's fluent APIs can be utilised by stacking on function calls to get the desired output.

```
samay.Create(time Time).EndOfDay().AddWeeks(2).FormatANSIC()
```

```
samay.Create(time Time).AddDays(12).GreaterThan(time1 Time)
```

# Return types of functions
- Functions which return ```Samay``` object can be used to chain method calls.
- Functions such as ```FormatXYZ``` will return string and cannot be chained.
- Function such as ```GreaterThan``` will return boolean and cannot be chained.

