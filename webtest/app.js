const baseURL = "https://0175-110-77-162-28.ngrok-free.io/course";
fetch(baseURL)
  .then(resp => resp.json())
  .then(data => displayData(data))
  .then(data => displayData(data));

function displayData(data) {
  document.querySelector("pre").innerHTML = JSON.stringify(data, null, 2);
}
