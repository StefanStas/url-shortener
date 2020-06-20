package errs

type NotFoundErr struct {
    EntityType string
}

func (nf *NotFoundErr) Error() string {
        return nf.EntityType + " not found"
}

func (nf *NotFoundErr) Is(target error) bool {
    _, ok := target.(*NotFoundErr)
    return ok
}
