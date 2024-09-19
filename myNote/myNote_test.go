package main

func TestCreate(t *testing.T){

	newNote:= myNote{
		Title: "monday",
		Content: "go shopping",
		Info: Info{
			CreatedAt: time.Now(),
			UpdatedAt:time.Now(),
		},

	}

	got:= Create(newNote)
	want:= 
}