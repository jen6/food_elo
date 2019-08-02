$(document).ready(function() {
  foods = JSON.parse(localStorage.foods)
  table = $('#rating_table')
  foods.sort(function(a, b){
    return b.Rating - a.Rating;
  });
  newFoods = foods.filter(food => food.Rating != 1000);
  newFoods.forEach(function(item, index, array) {
    table.append('<tr> <td>' + item.Name + '</td>' + '<td>' + item.Rating + '</td> </tr>');
  });
});
