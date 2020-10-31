package models

import "testing"

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}

func Test_withCorrectParams(t *testing.T) {
	review := NewReview(4, "The Iphone X looks good")
	err := review.validate()
	if err != nil {
		t.Error("The vallidation did not pass")
		t.Fail()
	}
}

func Test_shouldFailWithWrongNumberOfStars(t *testing.T) {
	review := NewReview(8, "The Iphone X looks good")
	err := review.validate()
	if err == nil {
		t.Error("should fail with more than 5 stars")
		t.Fail()
	}
}

func Test_shouldFailWithWrongLengthOfComment(t *testing.T) {
	review := NewReview(5,
		`this comment should have more than the max length of the comments what by default is 400,
					if the constant is modificated you should verify this test and if fail modify the length of this text to pass correctly.
					Now we complete the length with lorem ipsum.
					Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce semper, nisi sed molestie consectetur, odio sapien facilisis ex,
					a tristique arcu elit eu urna. Curabitur elementum tempus enim eu sodales. Aenean accumsan scelerisque vehicula. Vivamus purus ipsum, 
					efficitur at ultrices et, scelerisque vitae nulla. Aliquam mattis ultricies nisl, nec posuere orci interdum sed. Quisque a vestibulum ipsum. 
					Cras nibh tellus, laoreet nec posuere eu, congue sit amet enim. Nulla nec lacinia quam, ac facilisis diam. Cras nec fermentum odio. 
					Donec id lectus eu ligula iaculis aliquam. In a posuere enim. Morbi sed vulputate elit. Nullam vehicula augue libero, quis consequat turpis interdum eu. 
					Suspendisse nisi dolor, finibus egestas elementum mattis`,
	)
	err := review.validate()
	if err == nil {
		t.Error("should fail with a comment very large")
		t.Fail()
	}
}
