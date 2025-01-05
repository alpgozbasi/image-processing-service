package processor

import "github.com/h2non/bimg"

func ConvertToWebP(input []byte, quality int) ([]byte, error) {
	options := bimg.Options{
		Quality: quality,
		Type:    bimg.WEBP,
	}

	// process
	output, err := bimg.NewImage(input).Process(options)
	if err != nil {
		return nil, err
	}

	return output, err
}
