var foods;
var shuffledFoods = [];

var food_a;
var food_b;

function saveStatus() {
    localStorage.foods = JSON.stringify(foods);
    localStorage.shuffledFoods= JSON.stringify(shuffledFoods);
}

function shuffle(a) {
  console.log(a);
  for (let i = a.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [a[i], a[j]] = [a[j], a[i]];
  }
  return a;
}

async function loadData() {
  if(localStorage.foods){
    foods = JSON.parse(localStorage.foods);
    if(localStorage.shuffledFoods) {
      shuffledFoods = JSON.parse(localStorage.shuffledFoods);
    } else {
      shuffledFoods = JSON.parse(JSON.stringify(foods));
      shuffledFoods = shuffle(shuffledFoods);
      localStorage.shuffledFoods = JSON.stringify(shuffledFoods);
    }
    
    if(shuffledFoods.length == 0) {
      shuffledFoods = JSON.parse(JSON.stringify(foods));
      shuffledFoods = shuffle(shuffledFoods);
      localStorage.shuffledFoods = JSON.stringify(shuffledFoods);
    }
  }

	const result = await $.ajax({
      type: "GET",
      url: "/foodlist",
      dataType: "json",
	});

  if(result.error == null) {
    foods = result.foods;
    localStorage.foods = JSON.stringify(foods);


    shuffledFoods = JSON.parse(JSON.stringify(foods));
    shuffledFoods = shuffle(shuffledFoods);
    console.log(shuffledFoods);
    localStorage.shuffledFoods = JSON.stringify(shuffledFoods);
  }
}

function reshuffle() {
    shuffledFoods = JSON.parse(JSON.stringify(foods));
    shuffledFoods = shuffle(shuffledFoods);
}

function setFoods() {
  if(shuffledFoods.length < 2) {
    reshuffle()
  }
  food_a = shuffledFoods.pop()
  console.log(food_a)
  food_b = shuffledFoods.pop()
  console.log(food_b)
  $('#food_a').attr('value', food_a.Name)
  $('#food_a').attr('idx', food_a.Idx)
  $('#food_b').attr('value', food_b.Name)
  $('#food_b').attr('idx', food_b.Idx)
}


$(document).ready(function() {
  loadData();
  setFoods();
});

function action(isA) {
  if(isA) {
    newEloA = Elo.getNewRating(food_a.Rating, food_b.Rating, 1);
    newEloB = Elo.getNewRating(food_b.Rating, food_a.Rating, 0);
    foods[food_a.Idx-1].Rating = newEloA;
    foods[food_b.Idx-1].Rating = newEloB;
    console.log(foods[food_a.Idx - 1]);
    console.log(foods[food_b.Idx - 1]);
  } else {
    newEloA = Elo.getNewRating(food_a.Rating, food_b.Rating, 0);
    newEloB = Elo.getNewRating(food_b.Rating, food_a.Rating, 1);
    foods[food_a.Idx-1].Rating = newEloA;
    foods[food_b.Idx-1].Rating = newEloB;
    console.log(foods[food_a.Idx - 1]);
    console.log(foods[food_b.Idx - 1]);
  }
  setFoods();
  saveStatus();
}

