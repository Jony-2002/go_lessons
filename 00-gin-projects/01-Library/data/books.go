package data

import "gin/Library/interfaces"

var Books = []interfaces.Book{
	{
		ID:     1,
		Name:   "React",
		Author: "Niyozbekov Jahongir",
	},
	{
		ID:     2,
		Name:   "Golang",
		Author: "Qurbonmamadov Masrur",
	},
	{
		ID:     3,
		Name:   "Flutter",
		Author: "Nazarkhudoev Shuhrat",
	},
}
