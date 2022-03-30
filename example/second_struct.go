package structs

import "context"

type (
	SecondStruct struct {
		String string
	}

	ThirdStruct struct {
		String string
	}
)

func (s *FirstStruct) MyOtherTestMethod(ctx context.Context, x string) (string, error) {
	return "", nil
}
