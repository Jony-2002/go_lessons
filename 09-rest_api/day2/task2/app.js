const nameTag = document.getElementById("name");
const surnameTag = document.getElementById("surname");
const lastnameTag = document.getElementById("lastname");
const loginTag = document.getElementById("login");
const passwordTag = document.getElementById("password");
const formTag = document.getElementById("form");

function handleSubmit(event) {
  event.preventDefault();

  // var myHeaders = new Headers();
  // myHeaders.append("Content-Type", "application/json");

  // var raw = JSON.stringify({
  //   name: nameTag.value,
  //   surname: surnameTag.value,
  //   lastname: lastnameTag.value,
  //   login: loginTag.value,
  //   password: passwordTag.value,
  // });

  // var requestOptions = {
  //   method: "POST",
  //   headers: myHeaders,
  //   body: raw,
  //   redirect: "follow",
  // };

  // fetch("http://127.0.0.1:9999/registration", requestOptions)
  //   .then((response) => response.text())
  //   .then((result) => console.log(result))
  //   .catch((error) => console.log("error", error));

  const newUser = {
    name: nameTag.value,
    surname: surnameTag.value,
    lastname: lastnameTag.value,
    login: loginTag.value,
    password: passwordTag.value,
  };

  console.log(newUser);

  fetch("http://127.0.0.1:9999/registration", {
    method: "POST",
    body: newUser,
  })
    .then((res) => {
      return res.json();
    })
    .then((data) => {
      console.log(data);
    })
    .catch((err) => {
      console.log(err);
    });
}

formTag.onsubmit = handleSubmit;
