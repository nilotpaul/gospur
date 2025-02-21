# HTTP Error Handling

If you're already using a framework which lets you return error in handlers, then you're all set (no need to read this).

**This is for people using `http.HandlerFunc` signature for their handlers (bascially stdlib stuff).**

## Creating Custom Handler Type

```go
// Basically same as `http.HandlerFunc` but retuns an error
type HTTPHandlerFunc func(w http.ResponseWriter, r *http.Request) error

type AppError struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func NewAppError(status int, msg string) *AppError {
    return &AppError{Status: status, Msg: msg}
} 

// Implementing the error interface.
func (r *AppError) Error() string {
    return r.Msg
}
```

## Creating a Wrapper to Handle Errors

```go
func handler(h HTTPHandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var (
            status = http.StatusInternalServerError
            msg = "Internal Server Error"
        )
        
        err := h(w, r)
        if err == nil { return }

        if res, ok := err.(*AppError); ok {
            status = res.Status
            msg = res.Msg 
        }

        // Handle different cases like 404, 401, etc.
        // You can also render error pages with `templates.Render`.
        http.Error(w, msg, status)
    }
}
```

## Example

```go
// Example Handler
func handleGetSomePage(w http.ResponseWriter, r *http.Request) error {
    return templates.Render(w, http.StatusOK, "SomePage.html", nil)
} 
func handlePostSomething(w http.ResponseWriter, r *http.Request) error {
    // Returning Errors
    return NewAppError(http.StatusBadRequest, "Bad!")
    
    // If OK
    w.WriteHeader(http.StatusOK)
} 

mux := http.NewServeMux()
mux.Handle("/some-page", handler(handleGetSomePage))

http.ListenAndServe(":3000", mux)
```