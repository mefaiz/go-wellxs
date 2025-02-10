package errors

type AppError struct {
    Code    int
    Message string
    Err     error
}

func (e *AppError) Error() string {
    return e.Message
} 