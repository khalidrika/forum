
// function div1()  {
//     document.getElementById("div").style.display = "block"
//     document.getElementById("div1").style.display = "none"
// }
// function div2() {
//     document.getElementById("div").style.display = "none"
//     document.getElementById("div1").style.display = "block"
// }fae 
//ikh
 function div1()  {
    document.getElementById("div").style.display = "block"
    document.getElementById("div1").style.display = "none"
}
function div2() {
    document.getElementById("div").style.display = "none"
    document.getElementById("div1").style.display = "block"
}
function updateClock() {
    const hours = document.getElementById('hrs');
    const minutes = document.getElementById('mins');
    const seconds = document.getElementById('secs');

    const hhCircle = document.getElementById('hh');
    const mmCircle = document.getElementById('mm');
    const ssCircle = document.getElementById('ss');

    const now = new Date();
    let hrs = now.getHours();
    let mins = now.getMinutes();
    let secs = now.getSeconds();

    hrs = hrs % 12 || 12;

    hours.innerText = hrs.toString().padStart(2, '0');
    minutes.innerText = mins.toString().padStart(2, '0');
    seconds.innerText = secs.toString().padStart(2, '0');

    const circumference = 251.2;

    hhCircle.style.strokeDashoffset = circumference - (circumference * hrs) / 12;
    mmCircle.style.strokeDashoffset = circumference - (circumference * mins) / 60;
    ssCircle.style.strokeDashoffset = circumference - (circumference * secs) / 60;
}

setInterval(updateClock, 1000);
updateClock();
//nihaya dyal sa3a
