package projectlist

import (
	"math/rand"
)

func quitMessage() string {
	messages := types[rand.Intn(len(types))]
	return messages[rand.Intn(len(messages))]
}

var types = [][]string{
	_goodbies(),
	_jokes(),
	_quotes(),
	_knowledge_bits(),
}

func _goodbies() []string {
	return []string{
		"Catch you on the flip side!",
		"Don't let the door hit you on the way out.",
		"Toodle-oo!",
		"Smell ya later!",
		"Don't let the bed bugs bite!",
		"Later, alligator!",
		"See ya, wouldn't wanna be ya!",
		"Hasta la pasta!",
		"Keep it real!",
		"Peace, love, and tacos!",
		"Keep it funky!",
		"Gotta jet!",
		"Time to hit the road, Jack.",
		"I'm outie like a navel!",
		"See you on the dark side of the moon!",
		"Gotta bounce!",
		"Time for me to fly the coop!",
		"I'm off like a prom dress!",
		"Catch you on the rebound!",
		"Adios Amigos.",
		"Farewell and adieu.",
		"Catch you on the flip side.",
		"Until we meet again.",
		"I'm outta here.",
		"Peace out.",
		"So long, farewell.",
		"Adieu.",
		"Auf Wiedersehen.",
		"Goodbye for now.",
		"Smell ya later.",
		"It's been real.",
		"Time to fly.",
		"Keep it real.",
		"Catch you on the rebound.",
		"Toodle-loo.",
		"Hasta la vista.",
		"Over and out.",
		"Sayonara.",
		"TTFN (ta-ta for now)",
		"Gotta jet.",
		"Time to hit the road.",
		"I'm signing off.",
		"I'm out of here.",
		"See ya later alligator.",
		"Be good.",
		"Until next time.",
		"Keep it cool.",
		"Later days.",
		"In a while crocodile.",
		"Keep in touch.",
		"Take it easy.",
		"Have a good one.",
		"Time to say goodbye.",
		"Time to bid adieu.",
		"Have a good day.",
		"I'll be back.",
		"Keep the faith.",
		"See you soon-ish.",
		"See you in a bit.",
		"Catch you on the next round.",
		"Gotta run.",
		"Catch you later, folks.",
		"I'm out like a light.",
		"Catch you on the next one.",
		"Talk to you soon.",
		"Be back in a jiffy.",
		"Peace, love, and chicken grease.",
	}
}

func _jokes() []string {
	return []string{
		"What did the fish say when it hit the wall? Dam!",
		"How do you make a plumber cry? Kill his family.",
		"What do you call a bear with no teeth? A gummy bear.",
		"Why was the math book sad? Because it had too many problems.",
		"What did the zero say to the eight? Nice belt!",
		"Why don't oysters give to charity? Because they're shellfish.",
		"What did the left eye say to the right eye? Between us, something smells.",
		"Why did the chicken cross the road? To get to the other side.",
		"How many tickles does it take to make an octopus laugh? Ten tickles.",
		"What's the difference between a poorly dressed person on a unicycle\nand a well-dressed person on a bicycle? Attire.",
		"Why did the banana go out with the prune? Because it couldn't get a date.",
		"What do you call a snowman with a six-pack? An abdominal snowman.",
		"Why did the scarecrow win an award? Because he was outstanding in his field.",
		"What do you call a pile of cats? A good time.",
		"Why do elephants never use computers? They're afraid of the mouse.",
		"What do you call a group of unorganized cats? A Cat-astrophe.",
		"Why did the coffee file a police report? It got mugged.",
		"How do you catch a squirrel? Climb a tree and act like a nut.",
		"Why did the frog call his insurance company? He had a jump in his car.",
		"What do you call a belt made out of watches? A waist of time.",
		"How do you make a tissue dance? Put a little boogie in it.",
		"What do you call an alligator in a vest? An investigator.",
		"Why did the chicken cross the playground? To get to the other slide.",
		"What do you call a pile of cats? A meowtain.",
		"Why did the tomato turn red? It saw the salad dressing and blushed.",
		"What do you call an elephant that doesn't matter? An irrelephant.",
		"Why don't scientists trust atoms? Because they make up everything.",
		"How do you organize a space party? You planet.",
		"What did the beaver say when it hit the wall? Dam!",
		"Why did the cookie go to the doctor? Because it was feeling crumbly.",
	}
}

func _quotes() []string {
	return []string{
		"\"The best way to predict your future is to create it.\" - Abraham Lincoln",
		"\"The greatest glory in living lies not in never falling,\nbut in rising every time we fall.\" - Nelson Mandela",
		"\"Life is 10% what happens to us and 90% how we react to it.\" - Charles R. Swindoll",
		"\"Success is not final, failure is not fatal:\nIt is the courage to continue that counts.\" - Winston Churchill",
		"\"Believe in yourself and all that you are.\nKnow that there is something inside you that is greater than any obstacle.\" - Christian D. Larson",
		"\"The only way to do great work is to love what you do.\" - Steve Jobs",
		"\"Happiness is not something ready made.\nIt comes from your own actions.\" - Dalai Lama",
		"\"The only limit to our realization of tomorrow will be our doubts of today.\" - Franklin D. Roosevelt",
		"\"You only live once, but if you do it right, once is enough.\" - Mae West",
		"\"Life is like a camera. Focus on the good times,\ndevelop from the negatives, and if things don't work out, take another shot.",
		"\"The greatest mistake you can make in life is continually fearing you will make one.\" - Elbert Hubbard",
		"\"The future belongs to those who believe in the beauty of their dreams.\" - Eleanor Roosevelt",
		"\"The difference between try and triumph is just a little umph!\" - Marvin Phillips",
		"\"Be the change you wish to see in the world.\" - Mahatma Gandhi",
		"\"The best revenge is massive success.\" - Frank Sinatra",
		"\"The difference between ordinary and extraordinary is that little extra.\" - Jimmy Johnson",
		"\"The only thing to fear is fear itself.\" - Franklin D. Roosevelt",
		"\"The best way to predict the future is to create it.\" - Abraham Lincoln",
		"\"The only thing that stands between you and your dream is the will\nto try and the belief that it is actually possible.\" - Joel Brown",
		"\"The secret of getting ahead is getting started.\" - Mark Twain",
		"\"The only true wisdom is in knowing you know nothing.\" - Socrates",
	}
}

func _knowledge_bits() []string {
	return []string{
		"The capital of France is Paris.",
		"The Great Wall of China is the longest wall in the world.",
		"The human brain is made up of approximately 100 billion neurons.",
		"The Earth's atmosphere is composed mostly of nitrogen and oxygen.",
		"The first successful heart transplant was performed in 1967.",
		"The average person has a body temperature of 98.6 degrees Fahrenheit.",
		"The speed of light is approximately 299,792,458 meters per second.",
		"The largest mammal in the world is the blue whale.",
		"The tallest mountain in the world is Mount Everest.",
		"The smallest country in the world is Vatican City.",
		"The human body is made up of 60% water.",
		"The Earth's rotation causes the phenomenon of day and night.",
		"The human heart beats approximately 100,000 times a day.",
		"The longest river in the world is the Nile.",
		"The largest desert in the world is the Antarctic Desert.",
		"The most common blood type in humans is O+.",
		"The average lifespan of a human is around 78 years.",
		"The Earth's magnetic field is caused by the motion of molten iron in its core.",
		"The most widely spoken language in the world is Mandarin Chinese.",
		"The first computer was called the Electronic Numerical Integrator and Computer (ENIAC).",
		"The largest organ in the human body is the skin.",
		"The Earth's gravity is responsible for keeping us on the surface.",
		"The human body is made up of 206 bones.",
		"The largest planet in our solar system is Jupiter.",
		"The lowest point on Earth is the Challenger Deep in the Mariana Trench.",
		"The first successful flight of a powered aircraft was made by the Wright brothers in 1903.",
		"The largest mammal in the ocean is the humpback whale.",
		"The fastest land animal in the world is the cheetah.",
		"The largest volcano in the world is Mauna Loa in Hawaii.",
		"The average adult human has around 100,000 hairs on their head.",
		"The Earth's rotation causes the seasons to change.",
		"The human eye can distinguish around 10 million different colors.",
		"The largest ocean in the world is the Pacific Ocean.",
		"The most abundant element in the Earth's crust is oxygen.",
		"The first successful organ transplant was a kidney transplant performed in 1954.",
		"The human body has five senses: sight, smell, taste, touch, and hearing.",
		"The smallest planet in our solar system is Mercury.",
		"The largest bird in the world is the ostrich.",
		"The longest animal in the world is the blue whale.",
		"The largest fish in the ocean is the whale shark.",
		"The Earth's atmosphere is divided into five layers:\nthe troposphere, stratosphere, mesosphere, thermosphere, and exosphere.",
		"The human body is made up of around 70% water.",
		"The first successful human heart transplant was performed in 1967.",
		"The highest point on Earth is the summit of Mount Everest.",
		"The average human heart beats around 72 times per minute.",
		"The largest animal in the world is the blue whale.",
		"The fastest mammal in the world is the peregrine falcon.",
		"The largest reptile in the world is the saltwater crocodile.",
		"The first successful artificial heart transplant was performed in 1982.",
	}
}
