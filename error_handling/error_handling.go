package erratum

// Use implmentens simple handling of resourceopener.
func Use(o ResourceOpener, input string) (err error) {

	var resource Resource
	for {
		resource, err = o()
		if _, ok := err.(TransientError); ok {
			continue
		}
		if err != nil {
			return err
		}
		break
	}

	defer func() {
		defer resource.Close() // Close resource at the end
		if r := recover(); r != nil {
			if e, ok := r.(FrobError); ok { // Type assertion on FrobError and propogate error
				resource.Defrob(e.defrobTag)
				err = e
			} else if e, ok := r.(error); ok { // Propogate error
				err = e
			}
		}
	}()

	resource.Frob(input)
	return err
}
