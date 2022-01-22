package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{
		Title: "First item",
		Post:  "This is the first post",
	})

	if len(feed.Items) == 0 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{
		Title: "First item",
		Post:  "This is the first post",
	})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
