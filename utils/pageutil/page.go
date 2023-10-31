package pageutil

func PageLoop(f PageLoopFunc, page, size int) error {
	total, err := f(page, size)
	if err != nil {
		return err
	}

	step := 1
	for page*size < total {
		page += step
		_, err = f(page, size)
		if err != nil {
			return err
		}
	}
	return nil
}

type PageLoopFunc func(page, size int) (total int, err error)
