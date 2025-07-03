// let age = 15;

// if(age>18)
// {
//     console.log("You can drive")
// }
// else{
//     console.log("You cannot  drive")
// }


// loop
// let a = 1;
// for (let i = 0; i < 100; i++) {
//     console.log(a = i)
// }

// it is Comment of function

// function nice(name) {
//     console.log("Hey " + name + " you are nice!")
//     console.log("Hey " + name + " you are good!")
//     console.log("Hey " + name + " your tshirt is nice!")
//     console.log("Hey " + name + " your course is good too!")
// }

// function sum(a, b, c = 3) {
//     // console.log(a + b)
//     console.log(a, b, c)
//     return a + b + c
// }


// result1 = sum(3, 2)
// result2 = sum(7, 5)
// result3 = sum(3, 13, 1)

// console.log("The sum of these numbers is: ", result1)
// console.log("The sum of these numbers is: ", result2)
// console.log("The sum of these numbers is: ", result3)


// const func1 = (x)=>{
//     console.log("I am an arrow function", x)
// }

// func1(34);
// func1(66);
// func1(84);

// let age = 15;

// if (age>18) {
//     console.log("You are eligible")
// }
// else{
//     console.log("You are not eligible")
// }

function addTask() {
  const input = document.getElementById("taskInput");
  const taskText = input.value.trim();
  if (taskText === "") return;

  const li = document.createElement("li");

  // Toggle done on click
  li.addEventListener("click", function () {
    li.classList.toggle("done");
  });

  // Task text
  li.textContent = taskText;

  // Delete button
  const deleteBtn = document.createElement("button");
  deleteBtn.textContent = "X";
  deleteBtn.classList.add("delete-btn");
  deleteBtn.addEventListener("click", function (e) {
    e.stopPropagation();
    li.remove();
  });

  li.appendChild(deleteBtn);

  document.getElementById("taskList").appendChild(li);
  input.value = "";
}

